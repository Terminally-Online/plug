import { Prisma } from "@prisma/client"

import { SOCKET_BASE_QUERY } from "@/lib"

const socketModel = Prisma.validator<Prisma.SocketDefaultArgs>()({
	...SOCKET_BASE_QUERY
})
export type SocketModel = Prisma.SocketGetPayload<typeof socketModel>

const minimalSocketModel = Prisma.validator<Prisma.SocketDefaultArgs>()({
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
export type MinimalSocketModel = Prisma.SocketGetPayload<typeof minimalSocketModel>

// ---------------------------------------------------------------------------
// Plugs
// ---------------------------------------------------------------------------
const plugModel = Prisma.validator<Prisma.PlugDefaultArgs>()({})
export type PlugModel = Prisma.PlugGetPayload<typeof plugModel>
