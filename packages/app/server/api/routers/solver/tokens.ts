import { z } from "zod"

import { getTokens } from "@/lib"
import { anonymousProtectedProcedure, createTRPCRouter } from "@/server/api/trpc"

export const tokens = createTRPCRouter({
	get: anonymousProtectedProcedure.input(z.string().optional()).query(async ({ input }) => await getTokens(input))
})
