import { PrismaClient } from "@prisma/client"

import { MAGIC_NONCE } from "@/server/api/routers/socket"
import { env } from "@/env"
import { encodeAbiParameters, parseAbiParameters } from "viem"
import { getSocketAddress, getSocketFactory, getSocketImplementation, getSocketSalt } from "@/lib/functions/socket"

const prisma = new PrismaClient()

const seedSockets = async () => {
	const DEFAULT_SOCKETS = [
		{
			id: "0x0Bb5d848487B10F8CFBa21493c8f6D47e8a8B17E",
			name: "0x446576.eth",
			avatar: "https://i.seadn.io/gae/Lz2sMZs95LCIDRgAf2Hlml-EjEAF6oDNcYpEOzS4BUwPYhgalIeZvaDYdLGA9hUSpe624iLK7WYPlICCTx5hpQmzD_KAt1XTzOBlEA?w=500&auto=format"
		},
		{
			id: "0x62180042606624f02d8a130da8a3171e9b33894d",
			name: "nftchance.eth",
			avatar: "https://ipfs.io/ipfs/QmbU7EHRNGH6kJ7ybmCcPfLpZkQtEtm2GD7f3bBP6SdqV1"
		},
		{
			id: "0x581BEf12967f06f2eBfcabb7504fA61f0326CD9A",
			name: "danner.eth",
			avatar: "https://ipfs.io/ipfs/QmPSkQbBdbwHQXfXmCCvJ3kERftXuDpsj1GqZvFfQo9EYu"
		},
		{
			id: "0x1ccb2945F1325e061b40Fe5b0B452f0E76fB7278",
			name: "stacker.eth",
			avatar: "https://ipfs.io/ipfs/bafkreicmlgnaxmknw7aumxfxxzysdhwqskjpemeejbijm75nnkk3nrn4xq"
		},
		{
			id: "0x50701f4f523766bFb5C195F93333107d1cB8cD90",
			name: "nftmason.eth",
			avatar: "https://ipfs.io/ipfs/bafkreicmlgnaxmknw7aumxfxxzysdhwqskjpemeejbijm75nnkk3nrn4xq"
		}
	]

	for (const socket of DEFAULT_SOCKETS) {
		const { deployment: { address: factory }} = getSocketFactory()
		const { deployment: { address: implementation } } = getSocketImplementation()

		const { hex: salt } = getSocketSalt(
			MAGIC_NONCE,
			socket.id as `0x${string}`,
		)
		const { address: socketAddress } = getSocketAddress(salt as `0x${string}`)

		const deployment = {
			deploymentFactory: factory,
			deploymentNonce: parseInt(MAGIC_NONCE.toString()),
			deploymentDelegate: env.SOLVER_DELEGATE_ADDRESS,
			deploymentImplementation: implementation,
			deploymentSalt: salt
		} as const

		const createIdentity = {
			create: {
				ens: {
					create: {
						name: socket.name,
						avatar: socket.avatar
					}
				},
				approvedAt: new Date()
			}
		} as const
		const upsertIndentity = {
			upsert: {
				create: {
					ens: {
						create: {
							name: socket.name,
							avatar: socket.avatar
						}
					},
					approvedAt: new Date()
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
					},
					approvedAt: new Date()
				}
			}
		} as const

		await prisma.socket.upsert({
			where: { id: socket.id },
			update: {
				admin: true,
				socketAddress,
				...deployment,
				identity: upsertIndentity
			},
			create: {
				id: socket.id,
				admin: true,
				socketAddress,
				...deployment,
				identity: createIdentity
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
