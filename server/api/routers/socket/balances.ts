import { z } from "zod"

import { getBalances, getTokens } from "@/lib"
import { getCollectibles } from "@/lib/functions/opensea"

import { createTRPCRouter, protectedProcedure } from "../../trpc"

export const balances = createTRPCRouter({
	balances: protectedProcedure.input(z.string()).query(async ({ input }) => {
		if (input === "") return []

		return await getBalances(input)
	}),
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
