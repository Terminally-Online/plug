import { db } from "@/server/db"
import { TRPCError } from "@trpc/server"

import axios from "axios"

import { NATIVE_TOKEN_ADDRESS, TOKENS } from "@/lib/constants"
import { ZerionPositions } from "@/lib/types"

import { getChainId } from "../blockchain"
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

const prohibitedSymbolInclusions = [...prohibitedNameInclusions, "claim", "airdrop", "visit"]

const MINUTE = 60 * 1000
const HOUR = 60 * MINUTE
const POSITIONS_CACHE_TIME = 12 * HOUR

const getAPIKey = () => {
	return `Basic ${process.env.ZERION_API_KEY}`
}

const findPositions = async (socketId: string, search: string = "") => {
	const tokens = await db.fungible.findMany({
		where: {
			AND: [
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
				})),
				{
					OR: [
						{
							name: {
								contains: search,
								mode: "insensitive"
							}
						},
						{
							symbol: {
								contains: search,
								mode: "insensitive"
							}
						}
					]
				},
				{
					implementations: {
						some: { balances: { some: { socketId, balance: { gt: 0 } } } }
					}
				}
			]
		},
		select: {
			name: true,
			symbol: true,
			icon: true,
			verified: true,
			implementations: {
				where: { balances: { some: { socketId, balance: { gt: 0 } } } },
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

	const protocols = await db.protocol.findMany({
		where: {
			positions: { some: { cacheId: socketId } }
		},
		omit: { createdAt: true, updatedAt: true },
		include: {
			positions: {
				where: { cacheId: socketId },
				omit: {
					id: true,
					createdAt: true,
					updatedAt: true,
					cacheId: true,
					fungibleName: true,
					fungibleSymbol: true
				},
				include: {
					fungible: {
						omit: { createdAt: true, updatedAt: true },
						include: {
							implementations: {
								omit: {
									createdAt: true,
									updatedAt: true,
									fungibleName: true,
									fungibleSymbol: true
								}
							}
						}
					}
				}
			}
		}
	})

	const prices = await getPrices(
		[...tokens.map(token => `${token.implementations[0].chain}:${token.implementations[0].contract}`)].concat(
			protocols.flatMap(protocol =>
				protocol.positions.map(
					position =>
						`${position.fungible.implementations[0].chain}:${position.fungible.implementations[0].contract}`
				)
			)
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
						const implementationBalance = implementation.balances.reduce(
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
						price => price.id === `${token.implementations[0].chain}:${token.implementations[0].contract}`
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
		protocols: protocols.map(protocol => {
			return {
				...protocol,
				positions: protocol.positions.map(position => {
					const { fungible, balance } = position

					const { price, change } =
						prices.find(
							price =>
								price.id ===
								`${fungible.implementations[0].chain}:${fungible.implementations[0].contract}`
						) || {}

					return {
						...position,
						price,
						change,
						value: price && balance ? balance * price : undefined
					}
				})
			}
		})
	}
}

const getFungiblePositions = async (socketId: string, socketAddress: string, chains: string[]) => {
	const response = await axios.get(
		`https://api.zerion.io/v1/wallets/${socketAddress}/positions/?filter[positions]=no_filter&currency=usd&filter[chain_ids]=${chains.join(",")}&filter[trash]=only_non_trash&sort=value`,
		{
			headers: {
				accept: "application/json",
				authorization: getAPIKey()
			}
		}
	)

	if (response.status !== 200) throw new TRPCError({ code: "INTERNAL_SERVER_ERROR" })

	const data: ZerionPositions = response.data

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
				.filter(implementation => chains.includes(implementation.chain_id))
				.map(implementation => ({
					chain: implementation.chain_id,
					contract: implementation.address || NATIVE_TOKEN_ADDRESS,
					decimals: implementation.decimals
				}))

			// If Zerion does not have an icon for the fungible, try to find a static token.
			let icon = attributes.fungible_info.icon?.url
			if (icon === undefined) {
				const staticToken = TOKENS.find(t => t.symbol === attributes.fungible_info.symbol)
				if (staticToken !== undefined) icon = staticToken.logoURI
			}
			// If we could not find it as a static token, try to find it from the llamas.
			if (icon === undefined) {
				const implementation = implementations.sort((a, b) => getChainId(b.chain) - getChainId(a.chain))[0]
				icon = `https://token-icons.llamao.fi/icons/tokens/${getChainId(implementation.chain)}/${implementation.contract}?h=240&w=240`
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
			update: {
				updatedAt: new Date()
			}
		})

		// Delete all the implementation balances for the Socket.
		// TODO: This is not the ideal way to do this because we end up writing
		//       a lot of data to the database that was likely already there.
		//       We should really only be deleting the ones that the user no
		//       longer has.
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
						implementation => implementation.chain_id === relationships.chain.data.id
					)?.address ||
					(attributes.fungible_info.name === "Ethereum" && NATIVE_TOKEN_ADDRESS) ||
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

		const defi = positions.filter(position => position.attributes.position_type !== "wallet")

		// Update all of the positions into the cache.
		await tx.positionCache.update({
			where: { socketId },
			data: {
				updatedAt: new Date(),
				positions: {
					upsert: defi.map(position => {
						const { attributes, relationships } = position

						return {
							where: { id: `${socketId}-${position.id}` },
							create: {
								id: `${socketId}-${position.id}`,
								chain: relationships.chain.data.id,
								type: attributes.position_type,
								balance: attributes.quantity.float,
								protocolName: attributes?.application_metadata?.name ?? undefined,
								fungibleName: attributes.fungible_info.name,
								fungibleSymbol: attributes.fungible_info.symbol
							},
							update: {
								balance: attributes.quantity.float
							}
						}
					}),
					deleteMany: {
						cacheId: socketId,
						id: {
							notIn: defi.map(position => `${socketId}-${position.id}`)
						}
					}
				}
			}
		})
	})
}

export const getPositions = async (address: string, search?: string, chains = ["ethereum"]) => {
	const socket = await db.userSocket.findFirst({
		where: { id: address }
	})

	if (!socket)
		return {
			tokens: [],
			protocols: []
		}

	const cachedPositions = await db.positionCache.findUnique({
		where: { socketId: socket.id },
		select: { updatedAt: true }
	})

	if (cachedPositions && cachedPositions.updatedAt > new Date(Date.now() - POSITIONS_CACHE_TIME))
		return await findPositions(socket.id, search)

	await getFungiblePositions(socket.id, socket.socketAddress, chains)

	return await findPositions(socket.id, search)
}
