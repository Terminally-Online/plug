import { default as fs } from 'fs-extra'
import { exec } from 'child_process'

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
type Contract = { name: string, relativePath: string, salt: string, address: string }
const contracts: Array<Contract> = [
    {
        name: 'Plug.Router.Socket.sol',
        relativePath: '../sockets/',
        salt: '0x0',
        address: 'address(0)'
    }, {
        name: 'Plug.Vault.Socket.sol',
        relativePath: '../sockets/',
        salt: '0x0',
        address: 'address(0)'
    }
]

const directories = fs.readdirSync(artifacts)

const etcher = "src/contracts/utils/Plug.Etcher.sol"
const etcherTemplate = "src/contracts/utils/Plug.Etcher.Template.sol"

const imports: string[] = []
const variables: string[] = []
const functions: string[] = []

directories
    .filter(directory => contracts.some(contract => directory.includes(contract.name)))
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

            imports.push(`import { ${name} } from "${contracts.find(contract => directory.includes(contract.name))?.relativePath}${directory}";`)

            variables.push(`bytes internal constant ${variableName}_INITCODE = hex"${json.initcode}";`)
            variables.push(`bytes32 internal constant ${variableName}_SALT = ${contracts.find(contract => directory.includes(contract.name))?.salt};`)
            variables.push(`address internal constant ${variableName}_ADDRESS = ${contracts.find(contract => directory.includes(contract.name))?.address};`)

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
