import { default as fs } from 'fs-extra'
import { exec } from 'child_process'

import { etchContracts, contractsPath } from '../constants'

// TODO: Get the initial salts and addresses for the deployed contracts.

const artifacts = './artifacts'
const suffix = '.initcode.json'
// `relativePath` is the path relative to the Etcher contract that is
// is being generated. Automatic import path solving is not implemented as it
// was just consuming more time and may introduce issues that can simply be
// avoided by providing the relative path.
//
// To get the salt and address you will need to do a little mining with your preferred
// method. The salt is the salt used to deploy the contract and the address is the address
// of the deployed contract.

const directories = fs.readdirSync(artifacts)

const etcher = `${contractsPath}/utils/Plug.Etcher.sol`
const etcherTemplate = `${contractsPath}/utils/Plug.Etcher.Template.sol`

const imports: string[] = []
const variables: string[] = []
const functions: string[] = []

const addresses = JSON.parse(fs.readFileSync('lib/addresses.json').toString())

directories
    .filter(directory => etchContracts.some(contract => directory.includes(contract.name)))
    .forEach(directory => {
        const fileNames = fs.readdirSync(`${artifacts}/${directory}`)
        const files = fileNames.filter(fileName => fileName.endsWith(suffix))

        if (files.length === 0) return

        files.forEach(file => {
            const data = fs.readFileSync(`${artifacts}/${directory}/${file}`)
            const json = JSON.parse(data.toString())

            const name = file
                .replace(suffix, '')
                .replaceAll('.', '')

            const variableName = directory
                .replace('Plug.', '')
                .replace('.sol', '')
                .replaceAll('.', '_')
                .toUpperCase()

            const functionName = name
                .replace('Plug', '')
                .replace(/^./, x => x.toLowerCase())

            imports.push(`import { ${name} } from "${etchContracts.find(contract => directory.includes(contract.name))?.relativePath}${directory}";`)

            const mined = addresses[directory]['results'][0]

            variables.push(`bytes internal constant ${variableName}_INITCODE = hex"${json.initcode}";`)
            variables.push(`bytes32 internal constant ${variableName}_SALT = ${mined[0]};`)
            variables.push(`address internal constant ${variableName}_ADDRESS = ${mined[1]};`)

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
            if(directory == 'Plug.Router.Socket.sol') { 
                const line = 'address internal constant ROUTER_SOCKET_ADDRESS'
                const receiver = fs.readFileSync(`${contractsPath}/abstracts/Plug.Receiver.sol`).toString()
                const newReceiver = receiver.replace(/address internal constant ROUTER_SOCKET_ADDRESS = 0x[0-9a-fA-F]{40};/, `${line} = ${mined[1]};`)

                fs.writeFileSync(`${contractsPath}/abstracts/Plug.Receiver.sol`, newReceiver)
            }
        })
    })

let template = fs.readFileSync(etcherTemplate).toString()

template = template.replaceAll("Template", "")
template = template
    .replace("/// @auto INSERT IMPORTS", imports.join('\n'))
template = template
    .replace("/// @auto INSERT SEGMENTS", variables.join('\n\n') + '\n\n' + functions.join('\n\n'))

exec("forge fmt")

fs.writeFileSync(etcher, template)
