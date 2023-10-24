import { Network } from 'hardhat/types'

export async function getChainId(network: Network) {
	return await network.provider.send('eth_chainId').then(BigInt)
}
