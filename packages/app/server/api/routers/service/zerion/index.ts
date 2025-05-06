import { createTRPCRouter } from "@/server/api/trpc"

import { fungibles } from "./fungibles"
import { gas } from "./gas"
import { nfts } from "./nfts"
import { wallet } from "./wallet"

export const zerion = createTRPCRouter({
	wallet,
	gas,
	fungibles,
	nfts
})
