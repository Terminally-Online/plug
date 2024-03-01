import { exec } from 'child_process'
import { default as fs } from 'fs-extra'

import { contractsPath, etchContracts } from '../constants'

const artifacts = './artifacts'
const suffix = '.initcode.json'

const directories = fs.readdirSync(artifacts)

const etcher = `${contractsPath}/libraries/Plug.Etcher.Lib.sol`
const etcherTemplate = `${contractsPath}/libraries/Plug.Etcher.Lib.Template.sol`

const imports: string[] = []
const variables: string[] = []
const functions: string[] = []

const addresses = JSON.parse(fs.readFileSync('lib/addresses.json').toString())

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

			const variableName = directory
				.replace('Plug.', '')
				.replace('.sol', '')
				.replaceAll('.', '_')
				.toUpperCase()

			const functionName = name
				.replace('Plug', '')
				.replace(/^./, x => x.toLowerCase())

			imports.push(
				`import { ${name} } from "${etchContracts.find(contract =>
					directory.includes(contract.name)
				)?.relativePath}${directory}";`
			)

			const mined = addresses[directory]['results'][0]

			variables.push(
				`bytes internal constant ${variableName}_INITCODE = hex"${json.initcode}";`
			)
			variables.push(
				`bytes32 internal constant ${variableName}_SALT = ${mined[0]};`
			)
			variables.push(
				`address internal constant ${variableName}_ADDRESS = ${mined[1]};`
			)

			functions.push(`
                /**
                 * @notice Deploy (if needed) and return the ${name} contract instance.
                 * @return $${functionName} The ${name} contract instance.
                 */
                function ${functionName}() internal returns (${name} $${functionName}) {
                    if (_extcodesize(${variableName}_ADDRESS) == 0) {
                        address reality = _safeCreate2(${variableName}_SALT, ${variableName}_INITCODE);
                        require(reality ==  ${variableName}_ADDRESS, "Etcher: Reality check failed");
                    }
                    $${functionName} = ${name}(payable(${variableName}_ADDRESS));
                }
            `)

			// Update the address in Plug.Receiver.sol
			if (directory == 'Plug.Router.Socket.sol') {
				const line = 'address internal constant ROUTER_SOCKET_ADDRESS'
				const receiver = fs
					.readFileSync(
						`${contractsPath}/abstracts/Plug.Receiver.sol`
					)
					.toString()
				const newReceiver = receiver.replace(
					/address internal constant ROUTER_SOCKET_ADDRESS = 0x[0-9a-fA-F]{40};/,
					`${line} = ${mined[1]};`
				)

				fs.writeFileSync(
					`${contractsPath}/abstracts/Plug.Receiver.sol`,
					newReceiver
				)
			}
		})
	})

let template = fs.readFileSync(etcherTemplate).toString()

template = template.replaceAll('Template', '')
template = template.replace('/// @auto INSERT IMPORTS', imports.join('\n'))
template = template.replace(
	'/// @auto INSERT SEGMENTS',
	variables.join('\n\n') + '\n\n' + functions.join('\n\n')
)

exec('forge fmt')

fs.writeFileSync(etcher, template)
