import addresses from '../../addresses.json'
import { version } from 'package.json'
import {
	encodeAbiParameters,
	getContractAddress,
	Hex,
	parseAbiParameters
} from 'viem'

export const getSocketImplementation = () => {
	const versioned = addresses[version as keyof typeof addresses]
	return versioned.contracts['Plug.Socket.sol']
}

export const getSocketSalt = (nonce: bigint, admin: `0x${string}`) => {
	return encodeAbiParameters(parseAbiParameters(['uint96', 'address']), [
		nonce,
		admin
	])
}

export const getSocketAddress = (salt: Hex) => {
	const socketImplementation = getSocketImplementation()
	const address = getContractAddress({
		bytecodeHash: socketImplementation.initCodeHash as `0x${string}`,
		from: socketImplementation.deployment.address as `0x${string}`,
		opcode: 'CREATE2',
		salt
	})

	return { address }
}
