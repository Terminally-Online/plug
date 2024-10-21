import { TRPCError } from "@trpc/server"
import { z } from "zod"

import { getCollectibles, getPositions } from "@/lib/functions/zerion"
import { getMetadataForToken } from "@/lib/opensea/metadata"

import { anonymousProtectedProcedure, createTRPCRouter } from "../../trpc"

export const balances = createTRPCRouter({
	collectibles: anonymousProtectedProcedure
		.input(z.string().optional())
		.query(async ({ input, ctx }) => await getCollectibles(ctx.session.address, input)),
	positions: anonymousProtectedProcedure
		.input(z.string().optional())
		.query(async ({ input, ctx }) => await getPositions(ctx.session.address, input)),
	metadata: anonymousProtectedProcedure
		.input(
			z.object({
				type: z.union([z.literal("ERC20"), z.literal("ERC721"), z.literal("ERC1155")]),
				address: z.string(),
				tokenId: z.string(),
				chain: z.string()
			})
		)
		.query(async ({ input, ctx }) => {
			try {
				return await getMetadataForToken(ctx.db, {
					type: input.type,
					address: input.address,
					chain: input.chain,
					tokenId: input.tokenId
				})
			} catch (error) {
				if (error instanceof TRPCError) {
					throw error
				}
				throw new TRPCError({ code: "INTERNAL_SERVER_ERROR" })
			}
		})
})
