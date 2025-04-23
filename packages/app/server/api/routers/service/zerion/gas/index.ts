import { createTRPCRouter } from "@/server/api/trpc"

import { prices } from "./prices"

export const gas = createTRPCRouter({
	prices
})
