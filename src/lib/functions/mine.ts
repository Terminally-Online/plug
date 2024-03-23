import { constantContracts, etchContracts } from '../constants'
import { Contract } from '../types'
import { exec, execSync } from 'child_process'
import dedent from 'dedent'
import { default as fs } from 'fs-extra'

const efficientAddressesPath = 'create2crunch/efficient_addresses.txt'

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
		initCodeHash: string
		deployment: Record<'salt' | 'address' | 'rarity', string>
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
		console.log(dedent`
            ⏹︎ Generating efficient address:
                Contract Name: ${contract.name}
                Seconds: ${crunchSeconds}
                Leading: ${crunchLeading}
                Total: ${crunchTotal}
                Quick: ${quick}
        `)

		const artifactPath = `artifacts/${contract.name}/`
		const contractName = contract.name
			.replaceAll('.sol', '')
			.replaceAll('.', '')
		const initCodePath = `${artifactPath}${contractName}.initcode.json`
		const initCodeJson = JSON.parse(
			fs.readFileSync(initCodePath).toString()
		)
		const initCodeHash = initCodeJson['initcodeHash']

		// ? Check if the initCodeHash is already in the addresses.json file and has a matching address.
		if (force == false) {
			const hasOwnProperty = addresses.hasOwnProperty(contract.name)
			const property = addresses[contract.name]
			if (hasOwnProperty && property.initCodeHash == initCodeHash) {
				console.log(dedent`
                    ⏹︎ Skipping ${contract.name} as it already exists in the addresses.json file and has not changed.
                        ⏹︎ Use --force to overwrite.`)
				return resolve()
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

			addresses[contract.name] = {
				initCodeHash: initCodeHash,
				deployment: { salt, address, rarity }
			}

			// ? Clear the efficient addresses file so that we have a clean slate for the next contract.
			fs.writeFileSync(efficientAddressesPath, '')

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

	const process = async (contracts: Contract[]) => {
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
		console.log(`⏹︎ No contracts found with match: ${match}`)
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

		console.log(`✔︎ All contract addresses mined in ${duration} seconds.`)

		process.exit(0)
	})
	.catch(error => {
		console.error('An error occurred:', error)
		process.exit(1)
	})
