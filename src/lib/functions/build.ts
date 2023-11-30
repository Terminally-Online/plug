import { default as fse } from 'fs-extra'
import { default as path } from 'path'

const contracts = path.join(__dirname, '../../contracts')
const artifacts = path.join(__dirname, '../../../artifacts')
const dist = path.join(__dirname, '../../../dist/')

export async function main() {
	// * Copy all of the contracts.
	await fse.copy(contracts, `${dist}/contracts`)
	// * Copy all of the artifacts.
	await fse.copy(artifacts, `${dist}/artifacts`)
}

main()
