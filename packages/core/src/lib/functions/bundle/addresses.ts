import addresses from '../../addresses.json'
import { version } from 'package.json'
import { ByteArray, encodeAbiParameters, getContractAddress } from 'viem'

export const getSocketImplementation = () => {
	const versioned = addresses[version as keyof typeof addresses]
	return versioned.contracts['Plug.Socket.sol']
}

export const SALT_PARAMETERS = [
	{ type: 'uint96' },
	{ type: 'address' },
	{ type: 'address' },
	{ type: 'address' }
] as const

export const getSocketSalt = (
	nonce: bigint,
	admin: `0x${string}`,
	delegate: `0x${string}`,
	implementation: `0x${string}`
) => {
	return encodeAbiParameters(SALT_PARAMETERS, [
		nonce,
		admin,
		delegate,
		implementation
	])
}

export const getSocketAddress = (salt: ByteArray) => {
	const socketImplementation = getSocketImplementation()
	const address = getContractAddress({
		bytecodeHash: socketImplementation.initCodeHash as `0x${string}`,
		from: socketImplementation.deployment.address as `0x${string}`,
		opcode: 'CREATE2',
		salt
	})

	return { address }
}
