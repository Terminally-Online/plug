import { createTRPCRouter } from "../../trpc"
import { maintenance } from "./maintenance"
import { simulation } from "./simulation"

export const jobs = createTRPCRouter({
	maintenance,
	simulation
})
