import axios from "axios"

import { TRPCError } from "@trpc/server"

import { nativeTokenAddress } from "@/lib/constants"
import { FungibleModel, PositionModel } from "@/prisma/types"
import { db } from "@/server/db"

type PositionsResponse = {
	tokens: PositionModel[]
	defi: {
		[key: string]: {
			name: string
			icon: string
			url: string
			positions: PositionModel[]
			assets: FungibleModel[]
		}
	}
}

const MINUTE = 60 * 1000
const HOUR = 60 * MINUTE
const POSITIONS_CACHE_TIME = 12 * HOUR

const getAPIKey = () => {
	return `Basic ${process.env.ZERION_API_KEY}`
}

const findPositions = async (
	socketId: string,
	chains: string[],
	types: string[]
) => {
	const positions = await db.position.findMany({
		where: { cacheId: socketId, type: { in: types } },
		omit: {
			fungibleName: true,
			fungibleSymbol: true,
			protocolName: true,
			createdAt: true,
			updatedAt: true,
			cacheId: true
		},
		include: {
			fungible: {
				omit: { createdAt: true, updatedAt: true },
				include: {
					implementations: {
						where: { chain: { in: chains } },
						select: {
							chain: true,
							contract: true,
							decimals: true
						}
					}
				}
			},
			protocol: {
				omit: { createdAt: true, updatedAt: true }
			}
		}
	})

	return positions.reduce(
		(acc, position) => {
			const { type } = position
			const _type = type === "wallet" ? "tokens" : "defi"

			if (_type === "tokens") {
				acc[_type].push(position)
			} else if (_type === "defi") {
				const protocolName = position.protocol?.name

				if (protocolName === undefined) return acc

				if (acc.defi[protocolName] === undefined) {
					acc.defi[protocolName] = {
						name: "",
						icon: "",
						url: "",
						positions: [],
						assets: []
					}
				}

				if (
					acc.defi[protocolName].name === "" &&
					position.protocol?.name
				) {
					acc.defi[protocolName].name = position.protocol?.name
					acc.defi[protocolName].icon = position.protocol?.icon
					acc.defi[protocolName].url = position.protocol?.url
				}

				acc.defi[protocolName].positions.push(position)

				if (
					acc.defi[protocolName].assets.find(
						asset => asset.name === position.fungible.name
					) === undefined
				)
					acc.defi[protocolName].assets.push(position.fungible)
			}

			return acc
		},
		{
			tokens: [],
			defi: {}
		} as PositionsResponse
	)
}

const getFungiblePositions = async (
	socketId: string,
	socketAddress: string,
	chains: string[]
) => {
	const response = await axios.get(
		`https://api.zerion.io/v1/wallets/${socketAddress}/positions/?filter[positions]=no_filter&currency=usd&filter[chain_ids]=${chains.join(",")}&filter[trash]=only_non_trash&sort=value`,
		{
			headers: {
				accept: "application/json",
				authorization: getAPIKey()
			}
		}
	)

	if (response.status !== 200)
		throw new TRPCError({ code: "INTERNAL_SERVER_ERROR" })

	await db.protocol.createMany({
		data: response.data.data
			.map((position: any) => ({
				name: position.attributes?.application_metadata?.name ?? "",
				icon: position.attributes?.application_metadata?.icon.url ?? "",
				url: position.attributes?.application_metadata?.url ?? ""
			}))
			.filter((protocol: { name: string }) => protocol.name !== ""),
		skipDuplicates: true
	})

	const fungibleData = response.data.data.map((position: any) => {
		const implementations =
			position.attributes.fungible_info.implementations
				.filter(
					(implementation: {
						chain_id: string
						address: string
						decimals: number
					}) => chains.includes(implementation.chain_id)
				)
				.map(
					(implementation: {
						chain_id: string
						address: string
						decimals: number
					}) => ({
						chain: implementation.chain_id,
						contract: implementation.address || nativeTokenAddress,
						decimals: implementation.decimals
					})
				)

		return {
			name: position.attributes.fungible_info.name,
			symbol: position.attributes.fungible_info.symbol,
			icon: position.attributes.fungible_info.icon?.url ?? "",
			verified: position.attributes.fungible_info.flags.verified,
			implementations: implementations
		}
	})

	await db.fungible.createMany({
		data: fungibleData.map((fungible: any) => ({
			...fungible,
			implementations: undefined
		})),
		skipDuplicates: true
	})

	await db.implementation.createMany({
		data: fungibleData.map((fungible: any) => ({
			chain: fungible.implementations[0].chain,
			contract: fungible.implementations[0].contract,
			decimals: fungible.implementations[0].decimals,
			fungibleName: fungible.name,
			fungibleSymbol: fungible.symbol
		})),
		skipDuplicates: true
	})

	await db.positionCache.upsert({
		where: { socketId },
		create: {
			socketId
		},
		update: {
			updatedAt: new Date(),
			positions: {
				upsert: response.data.data.map((position: any) => ({
					where: { id: position.id },
					create: {
						id: position.id,
						chain: position.relationships.chain.data.id,
						type: position.attributes.position_type,
						balance: position.attributes.quantity.float,
						protocolName:
							position.attributes?.application_metadata?.name ??
							undefined,
						fungibleName: position.attributes.fungible_info.name,
						fungibleSymbol: position.attributes.fungible_info.symbol
					},
					update: {
						balance: position.attributes.quantity.float
					}
				})),
				deleteMany: {
					cacheId: socketId,
					id: {
						notIn: response.data.data.map(
							(position: { id: string }) => position.id
						)
					}
				}
			}
		}
	})
}

const getNonFungiblePositions = async (address: string, chains: string[]) => {
	// TODO: Retrieve from here to get the full set of NFTs held instead of the safe OpenSea list.
}

export const getPositions = async (
	address: string,
	chains = ["ethereum", "optimism", "base"],
	types = ["wallet", "deposit", "loan", "reward"]
): Promise<PositionsResponse> => {
	const socket = await db.userSocket.findFirst({
		where: { id: address }
	})

	if (!socket)
		return {
			tokens: [],
			defi: {}
		}

	const cachedPositions = await db.positionCache.findUnique({
		where: { socketId: socket.id },
		select: { updatedAt: true }
	})

	if (
		cachedPositions &&
		cachedPositions.updatedAt > new Date(Date.now() - POSITIONS_CACHE_TIME)
	)
		return await findPositions(socket.id, chains, types)

	await getFungiblePositions(socket.id, socket.socketAddress, chains)

	return await findPositions(socket.id, chains, types)
}
