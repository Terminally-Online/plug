import packageJson from '@/package.json'
import { execSync } from 'child_process'
import { default as fs } from 'fs-extra'

import { contractsPath, etchContracts } from '@/src/lib'

const artifacts = './artifacts'
const suffix = '.initcode.json'

const directories = fs.readdirSync(artifacts)

const etcher = `${contractsPath}/libraries/Plug.Etcher.Lib.sol`
const etcherTemplate = `${contractsPath}/libraries/Plug.Etcher.Lib.Template.sol`

const etcherTest = `${contractsPath}/libraries/Plug.Etcher.Lib.t.sol`
const etcherTestTemplate = `${contractsPath}/libraries/Plug.Etcher.Lib.t.Template.sol`

const imports: string[] = []
const variables: string[] = []
const functions: string[] = []
const deployments: string[] = []
const segments: string[] = []

const tests: string[] = []

if (!fs.existsSync('src/lib/addresses.json')) {
	fs.writeFileSync('src/lib/addresses.json', '{}')
}

const addresses = JSON.parse(
	fs.readFileSync('src/lib/addresses.json').toString()
)

const currentVersion = packageJson.version
if (!addresses[currentVersion]) {
	throw new Error(`No addresses found for version ${currentVersion}`)
}

const libs = ['Plug.Factory.sol']

directories
	.filter(directory =>
		etchContracts.some(contract => directory.includes(contract.name))
	)
	.forEach(directory => {
		const fileNames = fs.readdirSync(`${artifacts}/${directory}`)
		const files = fileNames.filter(fileName => fileName.endsWith(suffix))

		if (files.length === 0) return

		files.forEach(file => {
			const data = fs.readFileSync(`${artifacts}/${directory}/${file}`)
			const json = JSON.parse(data.toString())

			const name = file.replace(suffix, '').replaceAll('.', '')
			console.log(`- Writing etcher for ${name}.`)

			const variableName = directory
				.replace('.sol', '')
				.replaceAll('.', '_')
				.replace(/([a-z])([A-Z])/g, '$1_$2')
				.toUpperCase()

			// ! Generate the contents for the Etcher script.
			{
				const functionName = name.replace(/^./, x => x.toLowerCase())

				imports.push(
					`import { ${name} } from "${
						etchContracts.find(contract =>
							directory.includes(contract.name)
						)?.relativePath
					}${directory}";`
				)

				const mined =
					addresses[currentVersion].contracts[directory]?.deployment

				if (!mined) {
					console.log('skipping with: ', currentVersion)
					console.log('directory: ', directory)
					console.log(addresses[currentVersion].contracts[directory])
				}

				variables.push(
					`bytes internal constant ${variableName}_INITCODE = hex"${json.initcode}";`
				)
				variables.push(
					`bytes32 internal constant ${variableName}_SALT = ${mined.salt};`
				)
				variables.push(
					`address internal constant ${variableName}_ADDRESS = ${mined.address};`
				)

				functions.push(`
                    /**
                     * @notice Deploy (if needed) and return the ${name} contract instance.
                     * @return $${functionName} The ${name} contract instance.
                     */
                    function ${functionName}() internal returns (${name} $${functionName}) {
						require(
							keccak256(abi.encodePacked(type(${name}).creationCode)) ==
								keccak256(abi.encodePacked(${variableName}_INITCODE)),
							"PlugEtcherLib:invalid-initcode"
						);

						if (_extcodesize(${variableName}_ADDRESS) == 0) {
                            address reality = safeCreate2(${variableName}_SALT, ${variableName}_INITCODE);
                            require(reality ==  ${variableName}_ADDRESS, "PlugEtcherLib:unexpected-address");
                        }

                        $${functionName} = ${name}(payable(${variableName}_ADDRESS));
                    }
                `)

				tests.push(`
					function test_${name}Deployment() public {
						PlugEtcherLib.${functionName}();
					}
				`)

				// ! Update Plug.Lib.sol with the statically referenced addresses.
				if (libs.includes(directory)) {
					const variableName = directory
						.replace('.sol', '_ADDRESS')
						.replaceAll('.', '_')
						.toUpperCase()

					segments.push(
						`address internal constant ${variableName} = ${mined.address};`
					)
				}
			}

			// ! Generate the contents for the deployment script.
			{
				const deployment = `
		    	    if (PlugEtcherLib.${variableName}_ADDRESS.code.length == 0) {
		    	        PlugEtcherLib.FACTORY.safeCreate2(
		    	            PlugEtcherLib.${variableName}_SALT,
		    	            PlugEtcherLib.${variableName}_INITCODE
		    	        );
		    	    }
		    	`

				fs.writeFileSync(
					`${contractsPath}/scripts/${directory.replace(
						'.sol',
						''
					)}.s.sol`,
					fs
						.readFileSync(
							`${contractsPath}/scripts/Plug.s.Template.sol`
						)
						.toString()
						.replace('/// @auto INSERT SEGMENTS', deployment)
						.replace(
							'* INSERT DOCUMENTATION',
							`
                             * @title ${directory.replace(
									'.sol',
									''
								)} Deployment
                             * @dev Deploy a ${directory.replace(
									'.sol',
									''
								)} to a new chain using the immutable
                             *      Create2 factory for constant addresses across all major EVM chains.
                             * @notice To deploy the most up to date version of ${directory.replace(
									'.sol',
									''
								)}, you can always just run
                             *         this script and everything will be deployed as configured.
                        `
						)
						.replace(
							'PlugDeployment',
							`${directory
								.replace('.sol', '')
								.replaceAll('.', '')}Deployment`
						)
				)

				deployments.push(deployment)
			}

			console.log(`✔︎ Writing etcher for ${name}.`)
		})
	})

let template = fs.readFileSync(etcherTemplate).toString()

template = template.replaceAll('Template', '')
template = template.replace('/// @notice INSERT IMPORTS', imports.join('\n'))
template = template.replace(
	'/// @auto INSERT SEGMENTS',
	variables.join('\n\n') + '\n\n' + functions.join('\n\n')
)
fs.writeFileSync(etcher, template)

fs.writeFileSync(
	etcherTest,
	fs
		.readFileSync(etcherTestTemplate)
		.toString()
		.replace('/// @auto INSERT SEGMENTS', tests.join('\n\n'))
)

// * Generate the base deployment script that has all the pieces in one.
fs.writeFileSync(
	`${contractsPath}/scripts/Plug.s.sol`,
	fs
		.readFileSync(`${contractsPath}/scripts/Plug.s.Template.sol`)
		.toString()
		.replace('/// @auto INSERT SEGMENTS', deployments.join('\n\n'))
)

// * Generate the static reference to the addresses of the contracts utilized
//   within the deployed pieces of the protocol.
fs.writeFileSync(
	`${contractsPath}/libraries/Plug.Addresses.Lib.sol`,
	fs
		.readFileSync(
			`${contractsPath}/libraries/Plug.Addresses.Lib.Template.sol`
		)
		.toString()
		.replace('/// @notice INSERT SEGMENTS', segments.join('\n\n'))
)

execSync('forge fmt')
