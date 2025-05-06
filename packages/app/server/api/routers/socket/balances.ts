import { TRPCError } from "@trpc/server"

import { z } from "zod"

import { getMetadataForToken } from "@/lib/opensea/metadata"

import { anonymousProtectedProcedure, createTRPCRouter } from "../../trpc"

export const balances = createTRPCRouter({
	metadata: anonymousProtectedProcedure
		.input(
			z.object({
				type: z.union([z.literal("ERC20"), z.literal("ERC721"), z.literal("ERC1155")]),
				address: z.string(),
				tokenId: z.string(),
				chain: z.string()
			})
		)
		.query(async ({ input }) => {
			if (input.type === "ERC20") throw new TRPCError({ code: "NOT_IMPLEMENTED" })

			return await getMetadataForToken({
				address: input.address,
				chain: input.chain,
				tokenId: input.tokenId
			})
		})
})
