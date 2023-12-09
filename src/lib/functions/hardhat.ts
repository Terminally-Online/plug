import hre, { network } from 'hardhat'
import { Network } from 'hardhat/types'

import { PlugSDK } from '@/core/sdk'

import { constants } from '@nftchance/plug-types'

export const [name, version] = ['PlugMockSocket', '0.0.0']

export default async function (
	[name, version]: [string, string] = ['PlugMockSocket', '0.0.0']
) {
	const chainId = await getChainId(network)

	const [owner, notOwner] = await hre.viem.getWalletClients()

	const contract = await hre.viem.deployContract(name, [name, version])

	const publicClient = await hre.viem.getPublicClient()

	const util = new PlugSDK(name, version, chainId, constants.types, contract)

	return {
		chainId,
		name,
		version,
		contract,
		util,
		owner,
		notOwner,
		publicClient
	}
}

export async function getChainId(network: Network) {
	return await network.provider.send('eth_chainId').then(parseInt)
}
