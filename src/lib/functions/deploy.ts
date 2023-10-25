import hre, { network } from 'hardhat'

import { Framework } from '../../framework'
import { getChainId } from './chain'

export const [name, version] = ['FrameworkMock', '0.0.0']

export default async function () {
	const chainId = await getChainId(network)

	const [owner, notOwner] = await hre.viem.getWalletClients()

	const contract = await hre.viem.deployContract(name)
	const publicClient = await hre.viem.getPublicClient()

	const util = await new Framework(contract).init(name, version)

	return { chainId, contract, util, owner, notOwner, publicClient }
}
