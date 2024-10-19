import { db } from "@/server/db"
import { TRPCError } from "@trpc/server"

import axios from "axios"

import { ZerionCollectibles } from "@/lib/types"

import { getZerionApiKey } from "./authentication"

export const getZerionCollectibles = async (
	chains: string[],
	socketId: string,
	socketAddress?: string,
	limit = 100,
	next?: string,
	collectibles: ZerionCollectibles["data"] = []
): Promise<ZerionCollectibles["data"]> => {
	const url =
		next ??
		`https://api.zerion.io/v1/wallets/${socketAddress ?? socketId}/nft-positions/?filter[chain_ids]=${chains.join(",")}&currency=usd&page[size]=${limit}`
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
		return await getZerionCollectibles(chains, socketId, socketAddress, limit, data.links.next, collectibles)

	await db.$transaction(async tx => {
		const id = `${socketId}-${socketAddress}`
		await tx.collectibleCache.upsert({
			where: { id },
			create: {
				id,
				socketId
			},
			update: {
				updatedAt: new Date()
			}
		})

		await Promise.all(
			collectibles.map(async collectible => {
				const { attributes, relationships } = collectible
				const { nft_info, collection_info, changed_at } = attributes

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

				// NOTE: We do not immediately pull the collectible metadata because it will be pulled
				// on demand when a user accesses the collectible in a collectible frame. However, we do
				// go ahead and update the fields that we already have the data for.
				const collectibleFields = {
					amount: attributes.amount ?? 0,
					name: nft_info.name ?? "",
					interface: nft_info.interface,
					isSpam: nft_info.flags.is_spam ?? false,
					previewUrl: nft_info.content?.preview?.url,
					imageUrl: nft_info.content?.detail?.url,
					videoUrl: nft_info.content?.video?.url,
					createdAt: new Date(changed_at)
				} as const

				await tx.collectible.upsert({
					where: {
						cacheId_tokenId_collectionAddress_collectionChain: {
							cacheId: id,
							tokenId: nft_info.token_id,
							collectionAddress: nft_info.contract_address,
							collectionChain: relationships.chain.data.id
						}
					},
					create: {
						cacheId: id,
						tokenId: nft_info.token_id,
						collectionAddress: nft_info.contract_address,
						collectionChain: relationships.chain.data.id,
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

const findCollectibles = async (cacheId?: string) => {
	const collectibles = await db.collection.findMany({
		where: { collectibles: { some: { cacheId } } },
		include: {
			collectibles: {
				where: { cacheId, isSpam: false },
				orderBy: { updatedAt: "desc" }
			}
		},
		orderBy: { createdAt: "desc" }
	})
	return collectibles
}

/**
 * Retrieves collectibles for a given address or socket address.
 * @throws {TRPCError} Throws a NOT_FOUND error if the socket is not found.
 * @throws {TRPCError} Throws a FORBIDDEN error if the socket address isn't the address of the wallet owned socket.
 */
export const getCollectibles = async (
	address: string,
	socketAddress?: string,
	chains: string[] = ["ethereum"]
): Promise<Awaited<ReturnType<typeof findCollectibles>>> => {
	const socket = await db.userSocket.findFirst({
		where: { id: address }
	})

	if (socket === null) throw new TRPCError({ code: "NOT_FOUND" })
	if (socket.socketAddress !== socketAddress) throw new TRPCError({ code: "FORBIDDEN" })

	// NOTE: The user can retrieve collectibles for their own address as well as the
	// address of their socket. To power this, we store both caches relative to the user
	// socket that was created once the user is authenticated.
	// ...
	// Wallet: `0x612...49d-`.
	// Socket: `0x612...49d-0x524...c3b`.
	// ...
	// This method while a bit less readable, it confirms that we only ever enable users
	// to retrieve collectibles for their own address as well as the address of their socket.
	const id = `${socket.id}-${socketAddress}`
	const cachedCollectibles = await db.collectibleCache.findUnique({
		where: { id },
		select: { updatedAt: true }
	})

	if (!cachedCollectibles || cachedCollectibles.updatedAt > new Date(Date.now() - 60 * 60 * 1000))
		await getZerionCollectibles(chains, socket.id, socketAddress)

	return await findCollectibles(id)
}
