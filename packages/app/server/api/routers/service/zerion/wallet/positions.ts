import { z } from "zod"

import { protectedProcedure } from "@/server/api/trpc"

import { buildQueryParams, zerion, zerionApi } from "../lib"

const PositionsInputSchema = z.object({
	path: z.object({
		address: z.string().optional()
	}),
	query: z
		.object({
			currency: z.string().default("usd"),
			filter: z
				.object({
					positions: z.enum(["only_simple", "only_complex", "no_filter"]).default("only_simple"),
					chainIds: z.array(z.string()).optional(),
					positionTypes: z.array(z.string()).optional(),
					fungibleIds: z.array(z.string()).optional(),
					dappIds: z.array(z.string()).optional(),
					trash: z.enum(["only_trash", "only_non_trash"]).default("only_non_trash")
				})
				.optional(),
			sort: z.enum(["value", "recent"]).default("value"),
			aggregate: z.boolean().optional().default(false)
		})
		.optional()
		.default({})
})

export const PositionSchema = z.object({
	type: z.string(),
	id: z.string(),
	attributes: z.object({
		parent: z.any().nullable(),
		protocol: z.any().nullable(),
		poolAddress: z.string().nullable().optional(),
		groupId: z.string().nullable().optional(),
		name: z.string(),
		position_type: z.string(),
		quantity: z.object({
			int: z.string(),
			decimals: z.number(),
			float: z.number(),
			numeric: z.string()
		}),
		value: z.number().nullable(),
		price: z.number().nullable(),
		changes: z
			.object({
				absolute_1d: z.number(),
				percent_1d: z.number()
			})
			.nullable(),
		fungible_info: z.object({
			name: z.string(),
			symbol: z.string(),
			icon: z
				.object({
					url: z.string()
				})
				.nullable(),
			flags: z.object({
				verified: z.boolean()
			}),
			implementations: z.array(
				z
					.object({
						chain_id: z.string(),
						address: z.string().nullable(),
						decimals: z.number()
					})
					.extend({
						balance: z.number().optional(),
						value: z.number().optional(),
						percentage: z.number().optional()
					})
			)
		}),
		flags: z.object({
			displayable: z.boolean(),
			is_trash: z.boolean()
		}),
		updated_at: z.string(),
		updated_at_block: z.number().nullable(),
		application_metadata: z
			.object({
				name: z.string(),
				icon: z
					.object({
						url: z.string()
					})
					.nullable(),
				url: z.string().nullable().optional()
			})
			.nullable()
			.optional()
	}),
	relationships: z.object({
		chain: z.object({
			links: z.object({
				related: z.string()
			}),
			data: z.object({
				type: z.string(),
				id: z.string()
			})
		}),
		fungible: z.object({
			links: z.object({
				related: z.string()
			}),
			data: z.object({
				type: z.string(),
				id: z.string()
			})
		}),
		dapp: z
			.object({
				data: z.object({
					type: z.string(),
					id: z.string()
				})
			})
			.optional()
	})
})

const PositionsOutputSchema = z.object({
	links: z.object({
		self: z.string()
	}),
	data: z.array(PositionSchema)
})

const groupByPositionType = (
	positions: Array<z.infer<typeof PositionSchema>> = []
): Record<string, Array<z.infer<typeof PositionSchema>>> => {
	return positions.reduce<Record<string, Array<z.infer<typeof PositionSchema>>>>((groups, position) => {
		const type = position.attributes.position_type
		if (!groups[type]) {
			groups[type] = []
		}
		groups[type].push(position)
		return groups
	}, {})
}

const reduceByRelationships = (
	positions: Array<z.infer<typeof PositionSchema>> = []
): Array<z.infer<typeof PositionSchema>> => {
	const groupedByFungible: Record<string, Array<z.infer<typeof PositionSchema>>> = {}

	positions.forEach(position => {
		const fungibleId = position.relationships.fungible.data.id
		if (!groupedByFungible[fungibleId]) {
			groupedByFungible[fungibleId] = []
		}
		groupedByFungible[fungibleId].push(position)
	})

	return Object.values(groupedByFungible).map(positionGroup => {
		if (positionGroup.length === 1) {
			const position = positionGroup[0]
			const chainId = position.relationships.chain.data.id
			const value = position.attributes.value || 0

			return {
				...position,
				attributes: {
					...position.attributes,
					fungible_info: {
						...position.attributes.fungible_info,
						implementations: position.attributes.fungible_info.implementations.map(impl => {
							if (impl.chain_id === chainId) {
								return {
									...impl,
									balance: position.attributes.quantity.float || 0,
									value,
									percentage: 100
								}
							}
							return impl
						})
					}
				}
			}
		}

		const basePosition = { ...positionGroup[0] }
		const totalValue = positionGroup.reduce((sum, pos) => sum + (pos.attributes.value || 0), 0)

		const chainPositionMap: Record<string, z.infer<typeof PositionSchema>> = {}
		positionGroup.forEach(pos => {
			const chainId = pos.relationships.chain.data.id
			chainPositionMap[chainId] = pos
		})

		basePosition.id = `${basePosition.relationships.fungible.data.id}-multichain`
		basePosition.attributes = {
			...basePosition.attributes,
			value: totalValue,
			quantity: basePosition.attributes.quantity,
			name: `${basePosition.attributes.name} (${positionGroup.length} chains)`,
			fungible_info: {
				...basePosition.attributes.fungible_info,
				implementations: basePosition.attributes.fungible_info.implementations.map(impl => {
					const chainId = impl.chain_id
					const chainPosition = chainPositionMap[chainId]

					return {
						...impl,
						balance: chainPosition?.attributes.quantity.float || 0,
						value: chainPosition?.attributes.value || 0,
						percentage: chainPosition ? ((chainPosition.attributes.value || 0) / totalValue) * 100 : 0
					}
				})
			}
		}

		return basePosition
	})
}

const groupByDapp = (
	positions: Array<z.infer<typeof PositionSchema>> = []
): Record<string, Array<z.infer<typeof PositionSchema>>> => {
	return positions.reduce<Record<string, Array<z.infer<typeof PositionSchema>>>>((groups, position) => {
		const dappId = position.relationships.dapp?.data.id || "unknown"
		if (!groups[dappId]) {
			groups[dappId] = []
		}
		groups[dappId].push(position)
		return groups
	}, {})
}

export const positions = protectedProcedure
	.input(PositionsInputSchema)
	.output(PositionsOutputSchema)
	.query(async ({ ctx, input }) => {
		const address = input.path.address || ctx.session.address

		const queryParams = buildQueryParams({
			currency: input.query.currency,
			"filter[positions]": input.query?.filter?.positions,
			"filter[chain_ids]": input.query?.filter?.chainIds,
			"filter[position_types]": input.query?.filter?.positionTypes,
			"filter[fungible_ids]": input.query?.filter?.fungibleIds,
			"filter[dapp_ids]": input.query?.filter?.dappIds,
			"filter[trash]": input.query?.filter?.trash,
			sort: input.query.sort
		})

		type PositionsResponse = z.infer<typeof PositionsOutputSchema>

		const response = (await zerionApi(() =>
			zerion.get(`/wallets/${address}/positions/${queryParams}`)
		)) as PositionsResponse

		if (!input.query.aggregate) return response

		if (input.query?.filter?.positions === "only_complex") {
			const groupedByDapp = groupByDapp(response.data)
			const groupedData = Object.values(groupedByDapp).flat()
			return {
				links: response.links,
				data: groupedData
			}
		}

		const responseData = response.data
		const groupedPositions = groupByPositionType(responseData)
		const walletPositions = groupedPositions["wallet"] || []
		const aggregatedWalletPositions = reduceByRelationships(walletPositions)

		const otherPositions = responseData.filter(pos => pos.attributes.position_type !== "wallet")

		return {
			links: response.links,
			data: [...aggregatedWalletPositions, ...otherPositions]
		}
	})
