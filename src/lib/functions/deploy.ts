import hre, { network } from 'hardhat'

import { Framework } from '@/framework'
import { getChainId } from '@/lib/functions/chain'
import { constants } from '@nftchance/emporium-types'

export const [name, version] = ['FrameworkMock', '0.0.0']

export default async function () {
	const chainId = await getChainId(network)

	const [owner, notOwner] = await hre.viem.getWalletClients()

	const contract = await hre.viem.deployContract(name, [name, version])

	const publicClient = await hre.viem.getPublicClient()

	const util = new Framework<typeof constants.types>(
		name,
		version,
		chainId,
		constants.types,
		contract
	)

	return { chainId, contract, util, owner, notOwner, publicClient }
}
