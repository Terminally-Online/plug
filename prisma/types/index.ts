import { Prisma } from "@prisma/client"

const consoleColumnModel = Prisma.validator<Prisma.ConsoleColumnDefaultArgs>()(
	{}
)
export type ConsoleColumnModel = Prisma.ConsoleColumnGetPayload<
	typeof consoleColumnModel
>

const userSocketModel = Prisma.validator<Prisma.UserSocketDefaultArgs>()({
	include: { columns: true }
})
export type UserSocketModel = Prisma.UserSocketGetPayload<
	typeof userSocketModel
>

// ---------------------------------------------------------------------------
// Tokens
// ---------------------------------------------------------------------------
const tokenPriceModel = Prisma.validator<Prisma.TokenPriceCreateArgs>()({})
export type TokenPriceModel = Prisma.TokenPriceGetPayload<
	typeof tokenPriceModel
>

const tokenBalanceModel = Prisma.validator<Prisma.TokenBalanceCreateArgs>()({})
export type TokenBalanceModel = Prisma.TokenBalanceGetPayload<
	typeof tokenBalanceModel
>

const tokenBalanceCacheModel =
	Prisma.validator<Prisma.TokenBalanceCacheDefaultArgs>()({
		include: { tokens: true }
	})
export type TokenBalanceCacheModel = Prisma.TokenBalanceCacheGetPayload<
	typeof tokenBalanceCacheModel
>

// ---------------------------------------------------------------------------
// Collectibles
// ---------------------------------------------------------------------------
const openseaCollectionModel =
	Prisma.validator<Prisma.OpenseaCollectionDefaultArgs>()({})
export type OpenseaCollectionModel = Prisma.OpenseaCollectionGetPayload<
	typeof openseaCollectionModel
>
const openseaCollectibleModel =
	Prisma.validator<Prisma.OpenseaCollectibleDefaultArgs>()({
		include: { collection: true }
	})
export type OpenseaCollectibleModel = Prisma.OpenseaCollectibleGetPayload<
	typeof openseaCollectibleModel
>
const openseaCollectibleCacheModel =
	Prisma.validator<Prisma.OpenseaCollectibleCacheDefaultArgs>()({
		include: { collectibles: { include: { collection: true } } }
	})
export type OpenseaCollectibleCacheModel =
	Prisma.OpenseaCollectibleCacheGetPayload<
		typeof openseaCollectibleCacheModel
	>
