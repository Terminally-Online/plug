import axios from "axios"

import { TRPCError } from "@trpc/server"

import { nativeTokenAddress } from "@/lib/constants"
import { tokens } from "@/lib/constants"
import { FungibleModel, PositionModel } from "@/prisma/types"
import { db } from "@/server/db"

import { getPrices } from "../llama"

const prohibitedNameInclusions = [
	".com",
	".io",
	".org",
	".net",
	".app",
	".gg",
	".xyz",
	".claims",
	".finance",
	".tech",
	".exchange",
	".wallet",
	".capital",
	".fund",
	".capital",
	".su",
	".cloud",
	".events"
]

const prohibitedSymbolInclusions = [
	...prohibitedNameInclusions,
	"claim",
	"airdrop",
	"visit"
]

type ZerionPositionsResponse = {
	links: {
		self: string
	}
	data: Array<{
		type: string
		id: string
		attributes: {
			parent: null
			protocol: string | null
			name: string
			position_type: string
			quantity: {
				int: string
				decimals: number
				float: number
				numeric: string
			}
			value: number
			price: number
			changes: {
				absolute_1d: number
				percent_1d: number
			}
			fungible_info: {
				name: string
				symbol: string
				icon: {
					url: string
				}
				flags: {
					verified: boolean
				}
				implementations: Array<{
					chain_id: string
					address: string
					decimals: number
				}>
			}
			flags: {
				displayable: boolean
				is_trash: boolean
			}
			updated_at: string
			updated_at_block: number | null
			application_metadata?: {
				name: string
				icon: {
					url: string
				}
				url: string
			}
		}
		relationships: {
			chain: {
				links: {
					related: string
				}
				data: {
					type: string
					id: string
				}
			}
			fungible: {
				links: {
					related: string
				}
				data: {
					type: string
					id: string
				}
			}
			dapp?: {
				data: {
					type: string
					id: string
				}
			}
		}
	}>
}

type PositionsResponse = {
	tokens: FungibleModel[]
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

const getIsExcluded = (token: AlchemyTokenBalance) => {
	const IsMissingName = token.name === undefined

	const isProhibitedName = prohibitedNameInclusions.some(inclusion =>
		token.name?.toLowerCase().includes(inclusion)
	)

	const isProhibitedSymbol = prohibitedSymbolInclusions.some(inclusion =>
		token.symbol?.toLowerCase().includes(inclusion)
	)

	return !(IsMissingName || isProhibitedName || isProhibitedSymbol)
}

const findPositions = async (socketId: string) => {
	const tokens = await db.fungible.findMany({
		where: {
			AND: [
				{
					implementations: {
						some: { balances: { some: { socketId } } }
					}
				},
				...prohibitedNameInclusions.map(inclusion => ({
					NOT: {
						name: {
							contains: inclusion
						}
					}
				})),
				...prohibitedSymbolInclusions.map(inclusion => ({
					NOT: {
						symbol: {
							contains: inclusion
						}
					}
				}))
			]
		},
		select: {
			name: true,
			symbol: true,
			icon: true,
			verified: true,
			implementations: {
				where: { balances: { some: { balance: { gt: 0 } } } },
				omit: {
					createdAt: true,
					updatedAt: true,
					fungibleName: true,
					fungibleSymbol: true
				},
				include: {
					balances: {
						where: { socketId, balance: { gt: 0 } },
						select: {
							balance: true
						}
					}
				}
			}
		}
	})

	const prices = await getPrices(
		tokens.map(
			token =>
				`${token.implementations[0].chain}:${token.implementations[0].contract}`
		)
	)

	return {
		tokens: tokens
			.map(token => {
				const balance = token.implementations.reduce(
					(acc, implementation) =>
						acc +
						implementation.balances.reduce((acc, balance) => {
							return acc + balance.balance
						}, 0),
					0
				)

				const implementations = token.implementations
					.map(implementation => {
						const implementationBalance =
							implementation.balances.reduce(
								(acc, balance) => acc + balance.balance,
								0
							)

						return {
							...implementation,
							balance: implementationBalance,
							percentage: (implementationBalance / balance) * 100
						}
					})
					.sort((a, b) => b.percentage - a.percentage)

				const { price, change } =
					prices.find(
						price =>
							price.id ===
							`${token.implementations[0].chain}:${token.implementations[0].contract}`
					) || {}

				return {
					...token,
					implementations,
					balance,
					price,
					change,
					value: price ? balance * price : undefined
				}
			})
			.sort((a, b) => (b.value || 0) - (a.value || 0)),
		defi: {}
	}

	// const positions = await db.position.findMany({
	// 	where: { cacheId: socketId, type: { in: types } },
	// 	omit: {
	// 		fungibleName: true,
	// 		fungibleSymbol: true,
	// 		protocolName: true,
	// 		createdAt: true,
	// 		updatedAt: true,
	// 		cacheId: true
	// 	},
	// 	include: {
	// 		fungible: {
	// 			omit: { createdAt: true, updatedAt: true },
	// 			include: {
	// 				implementations: {
	// 					where: { chain: { in: chains } },
	// 					select: {
	// 						chain: true,
	// 						contract: true,
	// 						decimals: true,
	// 						balances: {
	// 							where: { socketId }
	// 						}
	// 					}
	// 				}
	// 			}
	// 		},
	// 		protocol: {
	// 			omit: { createdAt: true, updatedAt: true }
	// 		}
	// 	}
	// })

	// return positions.reduce(
	// 	(acc, position) => {
	// 		const { type } = position
	// 		const _type = type === "wallet" ? "tokens" : "defi"

	// 		if (_type === "tokens") {
	// 			acc[_type].push(position)
	// 		} else if (_type === "defi") {
	// 			const protocolName = position.protocol?.name

	// 			if (protocolName === undefined) return acc

	// 			if (acc.defi[protocolName] === undefined) {
	// 				acc.defi[protocolName] = {
	// 					name: "",
	// 					icon: "",
	// 					url: "",
	// 					positions: [],
	// 					assets: []
	// 				}
	// 			}

	// 			if (
	// 				acc.defi[protocolName].name === "" &&
	// 				position.protocol?.name
	// 			) {
	// 				acc.defi[protocolName].name = position.protocol?.name
	// 				acc.defi[protocolName].icon = position.protocol?.icon
	// 				acc.defi[protocolName].url = position.protocol?.url
	// 			}

	// 			acc.defi[protocolName].positions.push(position)

	// 			if (
	// 				acc.defi[protocolName].assets.find(
	// 					asset => asset.name === position.fungible.name
	// 				) === undefined
	// 			)
	// 				acc.defi[protocolName].assets.push(position.fungible)
	// 		}

	// 		return acc
	// 	},
	// 	{
	// 		tokens: [],
	// 		defi: {}
	// 	} as PositionsResponse
	// )
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

	const data: ZerionPositionsResponse = response.data

	await db.$transaction(async tx => {
		const positions = data.data

		// Create the protocols for positions that are DeFi based.
		await tx.protocol.createMany({
			data: positions
				.map(position => {
					const { attributes } = position

					return {
						name: attributes?.application_metadata?.name ?? "",
						icon: attributes?.application_metadata?.icon?.url ?? "",
						url: attributes?.application_metadata?.url ?? ""
					}
				})
				.filter((protocol: { name: string }) => protocol.name !== ""),
			skipDuplicates: true
		})

		const fungibleData = positions.map(position => {
			const { attributes } = position

			// Only save records for chains that we support.
			const implementations = attributes.fungible_info.implementations
				.filter(implementation =>
					chains.includes(implementation.chain_id)
				)
				.map(implementation => ({
					chain: implementation.chain_id,
					contract: implementation.address || nativeTokenAddress,
					decimals: implementation.decimals
				}))

			// If Zerion does not have an icon for the fungible, try to find a static token.
			let icon = attributes.fungible_info.icon?.url
			if (icon === undefined) {
				const staticToken = tokens.find(
					t => t.symbol === attributes.fungible_info.symbol
				)

				if (staticToken !== undefined) icon = staticToken.logoURI
			}

			return {
				name: attributes.fungible_info.name,
				symbol: attributes.fungible_info.symbol,
				icon,
				verified: attributes.fungible_info.flags.verified,
				implementations: implementations
			}
		})

		// Create the fungible assets for each position.
		await tx.fungible.createMany({
			data: fungibleData.map(fungible => ({
				...fungible,
				// Do not create the implementations for the fungibles because
				// we will do it right below this in bulk.
				implementations: undefined
			})),
			skipDuplicates: true
		})

		// Create all of the implementation references for each fungible.
		await tx.implementation.createMany({
			data: fungibleData.flatMap(fungible =>
				fungible.implementations.map(implementation => ({
					chain: implementation.chain,
					contract: implementation.contract,
					decimals: implementation.decimals,
					fungibleName: fungible.name,
					fungibleSymbol: fungible.symbol
				}))
			),
			skipDuplicates: true
		})

		// Make sure the position cache exists in the database.
		await tx.positionCache.upsert({
			where: { socketId },
			create: {
				socketId
			},
			update: {}
		})

		await tx.implementationBalance.deleteMany({
			where: { socketId }
		})

		// Update the balances for every fungible position that is held.
		await Promise.all(
			positions.map(async position => {
				const { attributes, relationships } = position

				if (attributes.position_type !== "wallet") return

				const implementationChain = relationships.chain.data.id
				const implementationContract =
					attributes.fungible_info.implementations.find(
						implementation =>
							implementation.chain_id ===
							relationships.chain.data.id
					)?.address ||
					(attributes.fungible_info.name === "Ethereum" &&
						nativeTokenAddress) ||
					undefined
				const balance = attributes.quantity.float

				if (!implementationContract) return

				await tx.implementationBalance.upsert({
					where: {
						socketId_implementationChain_implementationContract: {
							socketId,
							implementationChain,
							implementationContract
						}
					},
					create: {
						socketId,
						implementationChain,
						implementationContract,
						balance
					},
					update: {
						balance
					}
				})
			})
		)

		// TODO: Need to delete the positions that are no longer in the list.

		// Update all of the positions into the cache.
		// await tx.positionCache.update({
		// 	where: { socketId },
		// 	data: {
		// 		updatedAt: new Date(),
		// 		positions: {
		// 			upsert: positions.map(position => {
		// 				const { attributes, relationships } = position

		// 				return {
		// 					where: { id: position.id },
		// 					create: {
		// 						id: `${socketId}-${position.id}`,
		// 						chain: relationships.chain.data.id,
		// 						type: attributes.position_type,
		// 						balance: attributes.quantity.float,
		// 						protocolName:
		// 							attributes?.application_metadata?.name ??
		// 							undefined,
		// 						fungibleName: attributes.fungible_info.name,
		// 						fungibleSymbol: attributes.fungible_info.symbol
		// 					},
		// 					update: {
		// 						balance: attributes.quantity.float
		// 					}
		// 				}
		// 			}),
		// 			deleteMany: {
		// 				cacheId: socketId,
		// 				id: {
		// 					notIn: positions.map(position => position.id)
		// 				}
		// 			}
		// 		}
		// 	}
		// })
	})
}

export const getPositions = async (
	address: string,
	chains = ["ethereum", "optimism", "base"]
) => {
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
		return await findPositions(socket.id)

	await getFungiblePositions(socket.id, socket.socketAddress, chains)

	return await findPositions(socket.id)
}
