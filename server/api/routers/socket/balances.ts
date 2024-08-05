import { z } from "zod"

import { getCollectibles, getTokens } from "@/lib"

import { createTRPCRouter, protectedProcedure } from "../../trpc"

export const balances = createTRPCRouter({
	tokens: protectedProcedure
		.input(z.string().optional())
		.query(async ({ input }) => {
			if (input === undefined) return []

			return await getTokens(input)
		}),
	collectibles: protectedProcedure
		.input(z.string().optional())
		.query(async ({ input }) => {
			if (input === undefined) return []

			return await getCollectibles(input)
		})
})
