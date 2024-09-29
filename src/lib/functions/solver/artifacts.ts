import { default as fse } from "fs-extra"
import path from "pathe"

import { Network } from "@/src/lib"

async function getJsonFilesInDirectory(directory: string) {
	let jsonFiles: Array<string> = []

	if ((await fse.pathExists(directory)) === false) return jsonFiles

	for (let item of await fse.readdir(directory)) {
		const fullPath = path.join(directory, item)
		const stats = await fse.stat(fullPath)

		if (stats.isDirectory()) {
			jsonFiles = jsonFiles.concat(
				await getJsonFilesInDirectory(fullPath)
			)
		} else if (path.extname(fullPath) === ".json") {
			jsonFiles.push(fullPath)
		}
	}

	return jsonFiles
}

export async function getArtifacts(network: Network) {
	if (network.artifacts === undefined || network.artifacts === "") return []

	return (await getJsonFilesInDirectory(network.artifacts))
		.filter(file => file.endsWith(".json"))
		.map(file => JSON.parse(fse.readFileSync(file).toString()))
		.filter(file => file.contractName)
		.map(file => {
			let { contractName: name, abi, bytecode, deployedBytecode } = file

			if (bytecode === "0x") bytecode = undefined
			if (deployedBytecode === "0x") deployedBytecode = undefined

			return {
				key: name.replace(".sol", ""),
				network,
				name,
				abi: JSON.stringify(abi),
				bytecode,
				deployedBytecode
			}
		})
}
