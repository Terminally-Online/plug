import { Prisma } from "@prisma/client"

import { SOCKET_BASE_QUERY } from "@/lib"

const userSocketModel = Prisma.validator<Prisma.UserSocketDefaultArgs>()({
	...SOCKET_BASE_QUERY
})
export type UserSocketModel = Prisma.UserSocketGetPayload<typeof userSocketModel>

const minimalUserSocketModel = Prisma.validator<Prisma.UserSocketDefaultArgs>()({
	include: {
		identity: {
			include: {
				ens: {
					omit: { createdAt: true, updatedAt: true }
				}
			},
			omit: { createdAt: true, updatedAt: true, socketId: true, farcasterId: true }
		}
	},
	omit: { createdAt: true, updatedAt: true, admin: true }
})
export type MinimalUserSocketModel = Prisma.UserSocketGetPayload<typeof minimalUserSocketModel>

// ---------------------------------------------------------------------------
// Tokens
// ---------------------------------------------------------------------------
const tokenPriceModel = Prisma.validator<Prisma.TokenPriceCreateArgs>()({})
export type TokenPriceModel = Prisma.TokenPriceGetPayload<typeof tokenPriceModel>

// ---------------------------------------------------------------------------
// Positions
// ---------------------------------------------------------------------------
const fungibleModel = Prisma.validator<Prisma.FungibleDefaultArgs>()({
	select: {
		name: true,
		symbol: true,
		icon: true,
		verified: true,
		implementations: {
			omit: {
				createdAt: true,
				updatedAt: true,
				fungibleName: true,
				fungibleSymbol: true
			},
			include: {
				balances: {
					select: {
						balance: true
					}
				}
			}
		}
	}
})
export type FungibleModel = Prisma.FungibleGetPayload<typeof fungibleModel>

const positionModel = Prisma.validator<Prisma.PositionDefaultArgs>()({
	omit: {
		fungibleName: true,
		fungibleSymbol: true,
		protocolName: true,
		createdAt: true,
		updatedAt: true,
		cacheId: true
	},
	include: {
		fungible: {
			include: {
				implementations: {
					select: {
						chain: true,
						contract: true,
						decimals: true
					}
				}
			}
		},
		protocol: {
			omit: { createdAt: true, updatedAt: true }
		}
	}
})
export type PositionModel = Prisma.PositionGetPayload<typeof positionModel>

const positionCacheModel = Prisma.validator<Prisma.PositionCacheDefaultArgs>()({})
export type PositionCacheModel = Prisma.PositionCacheGetPayload<typeof positionCacheModel>

// ---------------------------------------------------------------------------
// Plugs
// ---------------------------------------------------------------------------
const workflowModel = Prisma.validator<Prisma.WorkflowDefaultArgs>()({})
export type WorkflowModel = Prisma.WorkflowGetPayload<typeof workflowModel>
