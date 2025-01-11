import addresses from '../../addresses.json'
import {
	ByteArray,
	bytesToHex,
	encodePacked,
	getContractAddress,
	toBytes
} from 'viem'

export const getSocketSalt = (nonce: bigint, admin: `0x${string}`) => {
	const bytes = toBytes(encodePacked(['uint96', 'address'], [nonce, admin]), {
		size: 32
	})

	return { bytes, hex: bytesToHex(bytes) as `0x${string}` | string }
}

export const getSocketAddress = (salt: ByteArray) => {
	const socketImplementation = addresses['Plug.Socket.sol']
	const address = getContractAddress({
		bytecodeHash: socketImplementation.initCodeHash as `0x${string}`,
		from: addresses['Plug.Factory.sol'].deployment.address as `0x${string}`,
		opcode: 'CREATE2',
		salt
	})

	return { address, implementation: socketImplementation.deployment.address }
}
