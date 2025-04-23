import { createTRPCRouter } from "@/server/api/trpc"

import { chart } from "./chart"
import { nftCollections } from "./nft-collections"
import { nftPortfolio } from "./nft-portfolio"
import { nftPositions } from "./nft-positions"
import { pnl } from "./pnl"
import { portfolio } from "./portfolio"
import { positions } from "./positions"
import { transactions } from "./transactions"

export const wallet = createTRPCRouter({
	chart,
	portfolio,
	positions,
	transactions,
	nftPositions,
	nftCollections,
	nftPortfolio,
	pnl
})
