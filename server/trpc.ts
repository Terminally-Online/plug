import { inferAsyncReturnType, initTRPC, TRPCError } from '@trpc/server'

import { createContext } from './context'

export const t = initTRPC
	.context<inferAsyncReturnType<typeof createContext>>()
	.create()

const isAdminMiddleware = t.middleware(({ ctx, next }) => {
	if (!ctx.isAdmin) throw new TRPCError({ code: 'UNAUTHORIZED' })

	return next({ ctx: { ...ctx, isAdmin: true } })
})

export const router = t.router
export const procedure = t.procedure
export const adminProcedure = procedure.use(isAdminMiddleware)
