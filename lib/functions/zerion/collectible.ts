import { db } from "@/server/db"
import { TRPCError } from "@trpc/server"

import axios from "axios"

import { ZerionCollectibles } from "@/lib/types"

import { getZerionApiKey } from "./authentication"

export const getZerionCollectibles = async (
	socketId: string,
	socketAddress: string,
	chains: string[],
	limit = 100,
	next?: string,
	collectibles: ZerionCollectibles["data"] = []
): Promise<ZerionCollectibles["data"]> => {
	const url =
		next ??
		`https://api.zerion.io/v1/wallets/${socketAddress}/nft-positions/?filter[chain_ids]=${chains.join(",")}&currency=usd&page[size]=${limit}`
	const response = await axios.get(url, {
		headers: {
			accept: "application/json",
			authorization: getZerionApiKey()
		}
	})

	if (response.status !== 200) throw new TRPCError({ code: "INTERNAL_SERVER_ERROR" })

	const data: ZerionCollectibles = response.data

	collectibles = [...collectibles, ...data.data]

	if (data.links.next)
		return await getZerionCollectibles(socketId, socketAddress, chains, limit, data.links.next, collectibles)

	// TODO: Handle the collectible creation in the database.

	await db.$transaction(async tx => {
		await tx.collectibleCache.upsert({
			where: { socketId },
			create: {
				socketId
			},
			update: {
				updatedAt: new Date()
			}
		})

		await Promise.all(
			collectibles.map(async collectible => {
				const { attributes, relationships } = collectible
				const { nft_info, collection_info } = attributes

				const collectionFields = {
					name: collection_info.name,
					description: collection_info.description,
					iconUrl: collection_info.content.icon.url
				} as const

				await tx.collection.upsert({
					where: {
						address_chain: { address: nft_info.contract_address, chain: relationships.chain.data.id }
					},
					create: {
						address: nft_info.contract_address,
						chain: relationships.chain.data.id,
						...collectionFields
					},
					update: {
						...collectionFields
					}
				})

				// Note: We do not immediately pull the collectible metadata because it will be pulled
				// on demand when a user accesses the collectible in a collectible frame. However, we do
				// go ahead and update the fields that we already have the data for.
				const collectibleFields = {
					amount: attributes.amount ?? 0,
					name: nft_info.name ?? "",
					interface: nft_info.interface,
					isSpam: nft_info.flags.is_spam ?? false,
					previewUrl: nft_info.content?.preview?.url,
					imageUrl: nft_info.content?.detail?.url,
					videoUrl: nft_info.content?.video?.url
				} as const

				await tx.collectible.upsert({
					where: {
						cacheSocketId_tokenId_collectionAddress_collectionChain: {
							cacheSocketId: socketId,
							tokenId: nft_info.token_id,
							collectionAddress: nft_info.contract_address,
							collectionChain: relationships.chain.data.id
						}
					},
					create: {
						tokenId: nft_info.token_id,
						collectionAddress: nft_info.contract_address,
						collectionChain: relationships.chain.data.id,
						cacheSocketId: socketId,
						...collectibleFields
					},
					update: {
						...collectibleFields
					}
				})
			})
		)
	})

	return collectibles
}

const findCollectibles = async (socketId: string, search: string = "") => {
	return await db.collection.findMany({
		where: { collectibles: { some: { cacheSocketId: socketId } } },
		include: {
			collectibles: {
				where: { cacheSocketId: socketId, isSpam: false },
				orderBy: { updatedAt: "desc" }
			}
		},
		orderBy: { createdAt: "desc" }
	})
}

export const getCollectibles = async (address: string, search?: string, chains = ["ethereum"]) => {
	const socket = await db.userSocket.findFirst({
		where: { socketAddress: address }
	})

	if (socket === null) return []

	const cachedCollectibles = await db.collectibleCache.findUnique({
		where: { socketId: socket.id },
		select: { updatedAt: true }
	})

	if (cachedCollectibles && cachedCollectibles.updatedAt > new Date(Date.now() - 60 * 60 * 1000))
		return await findCollectibles(socket.id, search)

	await getZerionCollectibles(socket.id, socket.socketAddress, chains)

	return await findCollectibles(socket.id, search)
}
