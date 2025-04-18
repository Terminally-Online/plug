import { default as fs, writeFileSync } from 'fs-extra'
import { keccak256 } from 'viem'

import { etchContracts } from '@/src/lib/constants'

const directory = './artifacts'

const files = fs.readdirSync(directory)

const contracts: Array<{
	name: string
	abi: any
}> = []

files
	.filter(file => file.endsWith('.sol'))
	.forEach(file => {
		const subDirectory = `${directory}/${file}`
		const subFiles = fs.readdirSync(subDirectory)

		subFiles
			.filter(file => file.endsWith('.json'))
			.filter(file => !file.includes('initcode'))
			.filter(file => !file.includes('abi'))
			.filter(file => !file.includes('Lib'))
			.forEach(subFile => {
				const json = fs.readJSONSync(`${subDirectory}/${subFile}`)
				const initcode = json.bytecode.object.slice(2)
				const fileName = subFile.replace('.json', '')
				const abi = JSON.stringify(json.abi)

				if (initcode == '') return

				const initcodeHash = keccak256(`0x${initcode}`)
				const init = JSON.stringify({
					initcode,
					initcodeHash: `${initcodeHash}`
				})

				if (
					etchContracts.some(contract => file.includes(contract.name))
				)
					contracts.push({
						name: fileName,
						abi: json.abi
					})

				writeFileSync(`${subDirectory}/${fileName}.initcode.json`, init)
				writeFileSync(`${subDirectory}/${fileName}.abi.json`, abi)
			})
	})

const contractsFile = `./src/lib/contracts.ts`
const contractsContent = `export const contracts = ${JSON.stringify(contracts, null, 4)} as const`
fs.writeFileSync(contractsFile, contractsContent)
