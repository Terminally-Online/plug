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
const deployments: string[] = []

const addresses = JSON.parse(fs.readFileSync('src/lib/addresses.json').toString())

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
                .replace('.sol', '')
                .replaceAll('.', '_')
                .replace(/([a-z])([A-Z])/g, '$1_$2')
                .toUpperCase()

            // ! Generate the contents for the Etcher script.
            {
                const functionName = name
                    .replace(/^./, x => x.toLowerCase())

                imports.push(
                    `import { ${name} } from "${etchContracts.find(contract =>
                        directory.includes(contract.name)
                    )?.relativePath}${directory}";`
                )

                const mined = addresses[directory].deployment

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
                        if (_extcodesize(${variableName}_ADDRESS) == 0) {
                            address reality = _safeCreate2(${variableName}_SALT, ${variableName}_INITCODE);
                            require(reality ==  ${variableName}_ADDRESS, "Etcher: Reality check failed");
                        }
                        $${functionName} = ${name}(payable(${variableName}_ADDRESS));
                    }
                `)

                if (directory == 'Plug.sol') {
                    const line = 'address internal constant PLUG_ADDRESS'
                    const receiver = fs
                        .readFileSync(
                            `${contractsPath}/libraries/Plug.Lib.Template.sol`
                        )
                        .toString()
                    const newReceiver = receiver.replace(
                        "/// @notice INSERT SEGMENTS",
                        `${line} = ${mined.address};`
                    )

                    fs.writeFileSync(
                        `${contractsPath}/libraries/Plug.Lib.sol`,
                        newReceiver
                    )
                }
            }

            // ! Generate the contents for the deployment script.
            {
                const deployment = `
                    PlugEtcherLib.FACTORY.safeCreate2(
                        PlugEtcherLib.${variableName}_SALT,
                        PlugEtcherLib.${variableName}_INITCODE
                    );
                `;

                fs.writeFileSync(
                    `${contractsPath}/scripts/${directory.replace(".sol", "")}.s.sol`,
                    fs.readFileSync(`${contractsPath}/scripts/Plug.s.Template.sol`)
                        .toString()
                        .replace('/// @auto INSERT SEGMENTS', deployment)
                        .replace("* INSERT DOCUMENTATION", `
                             * @title ${directory.replace('.sol', '')} Deployment
                             * @dev Deploy a ${directory.replace('.sol', '')} to a new chain using the immutable
                             *      Create2 factory for constant addresses across all major EVM chains.
                             * @notice To deploy the most up to date version of ${directory.replace(".sol", "")}, you can always just run
                             *         this script and everything will be deployed as configured.
                        `)
                        .replace("PlugDeployment", `${directory.replace(".sol", "").replaceAll(".", "")}Deployment`)
                )

                deployments.push(deployment)
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

fs.writeFileSync(etcher, template)

fs.writeFileSync(`${contractsPath}/scripts/Plug.s.sol`, fs.readFileSync(`${contractsPath}/scripts/Plug.s.Template.sol`)
    .toString()
    .replace('/// @auto INSERT SEGMENTS', deployments.join('\n\n'))
)

exec('forge fmt')
