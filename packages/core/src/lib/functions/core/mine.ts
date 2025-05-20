import { configs } from '../bundle'
import packageJson from '@/package.json'
import { exec, execSync } from 'child_process'
import { keccak256, toUtf8Bytes, TypedDataEncoder } from 'ethers'
import { default as fs } from 'fs-extra'
import path from 'path'

import {
	constantContracts,
	DEFAULT_SCHEMA,
	etchContracts
} from '@/src/lib/constants'
import { Contract } from '@/src/lib/types'

const efficientAddressesPath = 'create2crunch/efficient_addresses.txt'

let version = packageJson.version
let last = packageJson.version
let crunchSeconds = 15
let crunchLeading = 5
let crunchTotal = 7
let crunchAddresses = 10
let factoryAddress = '0x0000000000ffe8b47b3e2130213b802212439497'
let callerAddress = '0x0000000000000000000000000000000000000000'
let quick = false
let force = false
let install = false
let match = ''

// Remove the first two elements (tsx lib/functions/mine.ts)
const args = process.argv.slice(2)
args.forEach(arg => {
	const [key, value] = arg.split('=')
	switch (key) {
		case '--version':
			version = value
			break
		case '--last':
			last = value
			break
		case '--seconds':
			crunchSeconds = parseInt(value, 10)
			break
		case '--leading':
			crunchLeading = parseInt(value, 10)
			break
		case '--total':
			crunchTotal = parseInt(value, 10)
			break
		case '--addresses':
			crunchAddresses = parseInt(value, 10)
			break
		case '--factory':
			factoryAddress = value
			break
		case '--caller':
			callerAddress = value
			break
		case '--quick':
			quick = true
			break
		case '--match':
			match = value
			break
		case '--force':
			force = true
			break
		case '--install':
			install = true
			break
	}
})

const create2 = () => {
	execSync(`rm -rf create2crunch`)

	if (install)
		execSync(
			`sudo apt install build-essential -y; curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh -s -- -y; source "$HOME/.cargo/env";`
		)

	execSync(
		`git clone https://github.com/0age/create2crunch && cd create2crunch; sed -i -e 's/0x4/0x40/g' src/lib.rs`
	)
}

const prepare = (): Record<
	string,
	{
		config: {
			seconds: number
			leading: number
			total: number
			addresses: number
			factory: string
			caller: string
			quick: boolean
		}
		contracts: Record<
			string,
			{
				initCode: string
				initCodeHash: string
				deployment: Record<'salt' | 'address' | 'rarity', string>
			}
		>
	}
> => {
	create2()

	// ? Create and read addresses.json with the addresses that have already been mined.
	if (!fs.existsSync('src/lib/addresses.json')) {
		fs.writeFileSync('src/lib/addresses.json', '{}')
	}
	const addressesJson = fs.readFileSync('src/lib/addresses.json')

	return JSON.parse(addressesJson.toString())
}

const mine = async (contract: Contract): Promise<void> => {
	return new Promise((resolve, reject) => {
		console.log(`⚡ Mining address for ${contract.name}`)
		console.log(`	- Seconds: ${crunchSeconds}`)
		console.log(`	- Leading: ${crunchLeading}`)
		console.log(`	- Total: ${crunchTotal}`)
		console.log(`	- Quick: ${quick}`)

		const artifactPath = `artifacts/${contract.name}/`
		const contractName = contract.name
			.replaceAll('.sol', '')
			.replaceAll('.', '')
		const initCodePath = `${artifactPath}${contractName}.initcode.json`
		const initCodeJson = JSON.parse(
			fs.readFileSync(initCodePath).toString()
		)
		const initCode = initCodeJson['initcode']
		const initCodeHash = initCodeJson['initcodeHash']

		if (force == false) {
			// Check current version first
			const hasCurrentVersion = addresses.hasOwnProperty(version)
			const hasCurrentContract =
				hasCurrentVersion &&
				addresses[version].contracts.hasOwnProperty(contract.name)
			const currentProperty = hasCurrentContract
				? addresses[version].contracts[contract.name]
				: null

			// Check last version
			const hasLastVersion = addresses.hasOwnProperty(last)
			const hasLastContract =
				hasLastVersion &&
				addresses[last].contracts.hasOwnProperty(contract.name)
			const lastProperty = hasLastContract
				? addresses[last].contracts[contract.name]
				: null

			// Skip if contract exists in current version with same hash
			if (
				currentProperty &&
				currentProperty.initCodeHash == initCodeHash
			) {
				console.log(`⏭️  Skipping ${contract.name}`)
				console.log(`	- Already exists in version ${version}`)
				console.log(`	- InitCodeHash unchanged`)
				console.log(`	- Use --force to overwrite`)
				return resolve()
			}

			// Skip if contract exists in last version with same hash (hasn't changed)
			if (lastProperty && lastProperty.initCodeHash == initCodeHash) {
				console.log(`⏭️  Skipping ${contract.name}`)
				console.log(`	- Exists in version ${last}`)
				console.log(`	- InitCodeHash unchanged`)
				console.log(`	- Use --force to overwrite`)

				// Copy the contract data to the current version
				if (!addresses[version]) {
					addresses[version] = {
						config: {
							seconds: crunchSeconds,
							leading: crunchLeading,
							total: crunchTotal,
							addresses: crunchAddresses,
							factory: factoryAddress,
							caller: callerAddress,
							quick: quick
						},
						contracts: {}
					}
				}
				addresses[version].contracts[contract.name] = {
					...lastProperty
				}

				console.log(`📋 Copied contract data to version ${version}`)
				return resolve()
			}

			if (lastProperty) {
				console.log(`🔄 Changes detected in ${contract.name}`)
				console.log(`	- Previous version: ${last}`)
				console.log(`	- Old hash: ${lastProperty.initCodeHash}`)
				console.log(`	- New hash: ${initCodeHash}`)
			}
		}

		// ? Start the heartbeat before beginning the mining so that we can check if the mining is still running.
		const interval = setInterval(() => {
			if (fs.existsSync(efficientAddressesPath) == false) {
				return
			}

			const efficientAddresses = fs
				.readFileSync(efficientAddressesPath)
				.toString()

			let running = false
			if (efficientAddresses == '') running = true
			else if (
				quick == false &&
				efficientAddresses.split('\n').length < crunchAddresses
			)
				running = true

			if (running) return

			process.kill()
			clearInterval(interval)

			let results = efficientAddresses
				.split('\n')
				.map((address: string) => address.split(' => '))
				.sort(
					(a: Array<string>, b: Array<string>) =>
						parseInt(a[1]) - parseInt(b[1])
				)

			const [salt, address, rarity] = results[0]

			const version = packageJson.version
			if (!addresses[version]) {
				addresses[version] = {
					config: {
						seconds: crunchSeconds,
						leading: crunchLeading,
						total: crunchTotal,
						addresses: crunchAddresses,
						factory: factoryAddress,
						caller: callerAddress,
						quick: quick
					},
					contracts: {}
				}
			}
			addresses[version].contracts[contract.name] = {
				initCode,
				initCodeHash: initCodeHash,
				deployment: { salt, address, rarity }
			}

			// ? Clear the efficient addresses file so that we have a clean slate for the next contract.
			fs.writeFileSync(efficientAddressesPath, '')

			console.log(`✅ Successfully mined ${contract.name}`)
			console.log(`	- Salt: ${salt}`)
			console.log(`	- Address: ${address}`)
			console.log(`	- Rarity: ${rarity}`)

			resolve()
		}, crunchSeconds * 1000)

		const process = exec(
			`cd create2crunch && export FACTORY="${factoryAddress}"; export CALLER="${callerAddress}"; export INIT_CODE_HASH="${initCodeHash}"; export LEADING=${crunchLeading}; export TOTAL=${crunchTotal}; cargo run --release $FACTORY $CALLER $INIT_CODE_HASH 0 $LEADING $TOTAL`,
			function (error, stdout, stderr) {
				if (error) {
					console.error(`exec error: ${error}`)
					return reject(error)
				}
				console.log(`stdout: ${stdout}`)
				console.error(`stderr: ${stderr}`)
			}
		)
	})
}

const processContracts = async () => {
	let found = false

	const process = async (contracts: Readonly<Array<Contract>>) => {
		for (const contract of contracts) {
			if (match != '' && !contract.name.includes(match)) {
				continue
			}

			found = true

			await mine(contract)
		}
	}

	await process(constantContracts)
	await process(etchContracts)

	if (!found) {
		console.log(`❌ No matching contracts found`)
		console.log(`	- Search term: ${match}`)
	}
}

const timeStarted = new Date().toISOString()
const addresses = prepare()

processContracts()
	.then(() => {
		fs.writeFileSync(
			'src/lib/addresses.json',
			JSON.stringify(addresses, null, 2)
		)

		const timeEnded = new Date().toISOString()
		const duration =
			(new Date(timeEnded).getTime() - new Date(timeStarted).getTime()) /
			1000

		console.log(`🏁 Mining completed`)
		console.log(`	- Duration: ${duration} seconds`)

		generateGoFile(addresses).then(() => {
			process.exit(0)
		})
	})
	.catch(error => {
		console.error('An error occurred:', error)
		process.exit(1)
	})

async function generateGoFile(addresses: Record<string, any>) {
	const latestVersion = Object.keys(addresses).sort().pop()!
	const plugContracts = addresses[latestVersion]?.contracts || {}

	// Helper to convert to Go exported CapitalCamelCase
	function toGoExportedCamelCase(key: string) {
		if (!key) return 'Router'
		return key
			.split('_')
			.map(part => part.charAt(0).toUpperCase() + part.slice(1))
			.join('')
	}

	// Helper to convert typehash names to Go exported CapitalCamelCase
	function toGoTypehashCamelCase(typeHashName: string) {
		// Remove trailing _TYPEHASH, split, capitalize, join, then add TypeHash
		const base = typeHashName.replace(/_TYPEHASH$/, '')
		return (
			base
				.toLowerCase()
				.split('_')
				.map(part => part.charAt(0).toUpperCase() + part.slice(1))
				.join('') + 'TypeHash'
		)
	}

	const addressVars: string[] = []
	const mapEntries: string[] = []
	Object.keys(plugContracts).forEach(contractName => {
		let key = contractName
			.replaceAll('.', '')
			.replaceAll('Plug', '')
			.replaceAll('sol', '')
		key = key.toLowerCase().replace(/\s+/g, '_')
		if (!key) key = 'router'
		const varName = toGoExportedCamelCase(key)
		const address = plugContracts[contractName]?.deployment?.address || ''
		addressVars.push(`${varName} = "${address}"`)
		mapEntries.push(`\t\t"${key}": ${varName},`)
	})

	const configurations = await configs()
	const configsa = [...configurations, DEFAULT_SCHEMA.config]

	const typeHashMap: Record<string, string> = {}
	for (const config of configsa) {
		// @ts-expect-error
		const encoder = new TypedDataEncoder(config.types)
		Object.keys(config.types).forEach(typeName => {
			const typeHashName = `${typeName
				.replace(/([a-z])([A-Z])/g, '$1_$2')
				.replace(/([A-Z])([A-Z])(?=[a-z])/g, '$1_$2')
				.replace(/([0-9])([A-Z])/g, '$1_$2')
				.toUpperCase()}_TYPEHASH`
			if (!(typeHashName in typeHashMap)) {
				typeHashMap[typeHashName] = keccak256(
					toUtf8Bytes(encoder.encodeType(typeName))
				)
			}
		})
	}

	const typeHashLines = Object.entries(typeHashMap)
		.map(
			([typeHashName, typeHashValue]) =>
				`${toGoTypehashCamelCase(typeHashName)} = "${typeHashValue}"`
		)
		.join('\n')

	// Write addresses.go
	const addressesGoFile = `package common

var (
${addressVars.map(l => '\t' + l).join('\n')}

\tPlug = map[string]string{
${mapEntries.join('\n')}
\t}
)
`

	// Write typehashes.go
	const typehashesGoFile = `package common

var (
\t${typeHashLines.trim().split('\n').join('\n\t')}
)
`

	const outDir = path.join(__dirname, '../../../../../references/common')
	fs.ensureDirSync(outDir)
	fs.writeFileSync(path.join(outDir, 'addresses.go'), addressesGoFile)
	fs.writeFileSync(path.join(outDir, 'type_hashes.go'), typehashesGoFile)
}
