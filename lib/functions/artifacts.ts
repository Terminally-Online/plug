import { default as fs, writeFileSync } from 'fs-extra'
import { keccak256 } from 'viem'

const directory = './artifacts'
const files = fs.readdirSync(directory)

files
	.filter(file => file.endsWith('.sol'))
	.forEach(file => {
		const subDirectory = `${directory}/${file}`
		const subFiles = fs.readdirSync(subDirectory)

		subFiles
			.filter(file => file.endsWith('.json'))
			.filter(file => !file.includes('initcode'))
			.forEach(subFile => {
				const json = fs.readJSONSync(`${subDirectory}/${subFile}`)
				const initcode = json.bytecode.object.slice(2)
				const fileName = subFile.replace('.json', '')
				const abi = JSON.stringify(json.abi)

				if (initcode == '') return

				const initcodeHash = keccak256(`0x${initcode}`)

				const init = JSON.stringify({
					initcode,
					initcodeHash: `0x${initcodeHash}`
				})

				writeFileSync(`${subDirectory}/${fileName}.initcode.json`, init)
				writeFileSync(`${subDirectory}/${fileName}.abi.json`, abi)
			})
	})
