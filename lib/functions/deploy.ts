// import { ethers, network } from 'hardhat'

// // import { Framework } from '../../contracts/framework'
// import { getChainId } from './chain'

// export const [name, version] = ['Echo', '0.1.0']

// export async function deploy() {
// 	const chainId = await getChainId(network)

// 	const [owner, notOwner] = await ethers.getSigners()

// 	const contract = await (
// 		await ethers.getContractFactory('Echo')
// 	).deploy(name, version)

// 	const address = await contract.getAddress()

// 	const util = await new Framework(contract).init(name, version)

// 	return { chainId, contract, address, util, owner, notOwner }
// }
