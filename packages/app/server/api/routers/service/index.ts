import { createTRPCRouter } from "@/server/api/trpc"

import { zerion } from "./zerion"

export const service = createTRPCRouter({
	zerion
})
