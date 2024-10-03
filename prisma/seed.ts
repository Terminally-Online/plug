import { createPublicClient, http } from "viem"
import { mainnet } from "viem/chains"

import { PrismaClient } from "@prisma/client"

const prisma = new PrismaClient()

const client = createPublicClient({
	chain: mainnet,
	transport: http(process.env.ALCHEMY_API_URL)
})

const seedSockets = async () => {
	const DEFAULT_SOCKETS = [
		"0x0Bb5d848487B10F8CFBa21493c8f6D47e8a8B17E",
		"0x62180042606624f02d8a130da8a3171e9b33894d",
		"0x581BEf12967f06f2eBfcabb7504fA61f0326CD9A",
		"0xda70761A63d5D0DdE3bdE3b179126127Cccb44b3"
	]

	for (const address of DEFAULT_SOCKETS) {
		let ensName = null
		let ensAvatar = null

		try {
			ensName = await client.getEnsName({ address: address as `0x${string}` })
			if (ensName) {
				ensAvatar = await client.getEnsAvatar({ name: ensName })
			}
		} catch (error) {
			console.error(`Error fetching ENS data for ${address}:`, error)
		}

		await prisma.userSocket.upsert({
			where: { id: address },
			update: {
				socketAddress: address,
				identity: {
					upsert: {
						create: {
							ens: ensName
								? {
										create: {
											name: ensName,
											avatar: ensAvatar || undefined
										}
									}
								: undefined
						},
						update: {
							ens: ensName
								? {
										upsert: {
											create: {
												name: ensName,
												avatar: ensAvatar || undefined
											},
											update: {
												name: ensName,
												avatar: ensAvatar || undefined
											}
										}
									}
								: undefined
						}
					}
				}
			},
			create: {
				id: address,
				socketAddress: address,
				identity: {
					create: {
						ens: ensName
							? {
									create: {
										name: ensName,
										avatar: ensAvatar || undefined
									}
								}
							: undefined
					}
				}
			}
		})
	}
}

const main = async () => {
	await seedSockets()
}

main()
	.catch(e => {
		throw e
	})
	.finally(async () => {
		await prisma.$disconnect()
	})
