import addresses from '@terminallyonline/plug-core/addresses.json'
import { version } from '@terminallyonline/plug-core/package.json'
import { bytesToHex, encodePacked, getContractAddress, Hex, toBytes } from 'viem'

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
