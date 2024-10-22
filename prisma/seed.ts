import { TEMPORARY_ADDRESS } from "@/server/api/routers/socket"

import { PrismaClient } from "@prisma/client"

const prisma = new PrismaClient()

const seedSockets = async () => {
	const DEFAULT_SOCKETS = [
		{
			address: "0x0Bb5d848487B10F8CFBa21493c8f6D47e8a8B17E",
			name: "0x446576.eth",
			avatar: "https://i.seadn.io/gae/Lz2sMZs95LCIDRgAf2Hlml-EjEAF6oDNcYpEOzS4BUwPYhgalIeZvaDYdLGA9hUSpe624iLK7WYPlICCTx5hpQmzD_KAt1XTzOBlEA?w=500&auto=format"
		},
		{
			address: "0x62180042606624f02d8a130da8a3171e9b33894d",
			name: "nftchance.eth",
			avatar: "https://ipfs.io/ipfs/QmbU7EHRNGH6kJ7ybmCcPfLpZkQtEtm2GD7f3bBP6SdqV1"
		},
		{
			address: "0x581BEf12967f06f2eBfcabb7504fA61f0326CD9A",
			name: "danner.eth",
			avatar: "https://ipfs.io/ipfs/QmPSkQbBdbwHQXfXmCCvJ3kERftXuDpsj1GqZvFfQo9EYu"
		},
		{
			address: "0xda70761A63d5D0DdE3bdE3b179126127Cccb44b3",
			name: "reka.eth",
			avatar: "https://cryptocoven.s3.amazonaws.com/7bd11d2655904bd6ae89e3fd6f16ab46.png"
		}
	]

	const socketAddress = TEMPORARY_ADDRESS

	for (const socket of DEFAULT_SOCKETS) {
		await prisma.userSocket.upsert({
			where: { id: socket.address },
			update: {
				socketAddress,
				identity: {
					upsert: {
						create: {
							ens: {
								create: {
									name: socket.name,
									avatar: socket.avatar
								}
							}
						},
						update: {
							ens: {
								upsert: {
									create: {
										name: socket.name,
										avatar: socket.avatar
									},
									update: {
										name: socket.name,
										avatar: socket.avatar
									}
								}
							}
						}
					}
				}
			},
			create: {
				id: socket.address,
				socketAddress,
				identity: {
					create: {
						ens: {
							create: {
								name: socket.name,
								avatar: socket.avatar
							}
						}
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
