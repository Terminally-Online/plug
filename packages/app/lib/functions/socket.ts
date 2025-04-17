import addresses from '@terminallyonline/plug-core/addresses.json'
import { version } from '@terminallyonline/plug-core/package.json'
import { bytesToHex, encodePacked, Hex, toBytes } from 'viem'
import { base } from 'viem/chains'
import { createClient } from '../constants'

const client = createClient(base.id)

export const getSocketFactory = () => {
	const versioned = addresses[version as keyof typeof addresses]
	return versioned.contracts['Plug.Factory.sol']
}

export const getSocketImplementation = () => {
	const versioned = addresses[version as keyof typeof addresses]
	return versioned.contracts['Plug.Socket.sol']
}

export const getSocketSalt = (
	nonce: bigint,
	admin: `0x${string}`,
) => {
	const bytes = toBytes(encodePacked(['uint96', 'address'], [nonce, admin]), {
		size: 32
	})

	return { bytes, hex: bytesToHex(bytes) as `0x${string}` | string }
}

const factoryAbi = [
	{
		inputs: [{ name: "$implementation", type: "address" }, { name: "$salt", type: "bytes32" }],
		name: "getAddress",
		outputs: [{ name: "$vault", type: "address" }],
		stateMutability: "view",
		type: "function",
	},
] as const;

export const getSocketAddress = async (salt: Hex) => {
	const { deployment: { address: factory } } = getSocketFactory()
	const socketImplementation = getSocketImplementation()
	const address = await client.readContract({
		address: factory as `0x${string}`,
		abi: factoryAbi,
		functionName: 'getAddress',
		args: [socketImplementation.deployment.address as `0x${string}`, salt]

	})

	return { address }
}
