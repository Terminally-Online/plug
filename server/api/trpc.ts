import { type Session } from "next-auth"

import superjson from "superjson"
import { ZodError } from "zod"

import { getServerAuthSession } from "@/server/auth"
import { db } from "@/server/db"
import { emitter } from "@/server/emitter"
import { initTRPC, TRPCError } from "@trpc/server"
import { type CreateNextContextOptions } from "@trpc/server/adapters/next"

interface CreateContextOptions {
	session: Session | null
}

export const createInnerTRPCContext = (opts: CreateContextOptions) => {
	return {
		session: opts.session,
		db,
		emitter
	}
}

export const createTRPCContext = async (opts: CreateNextContextOptions) => {
	const { req, res } = opts
	const session = await getServerAuthSession({ req, res })

	return createInnerTRPCContext({
		session
	})
}

const t = initTRPC.context<typeof createTRPCContext>().create({
	transformer: superjson,
	errorFormatter({ shape, error }) {
		return {
			...shape,
			data: {
				...shape.data,
				zodError: error.cause instanceof ZodError ? error.cause.flatten() : null
			}
		}
	}
})

/**
 * 3. ROUTER & PROCEDURE (THE IMPORTANT BIT)
 *
 * These are the pieces you use to build your tRPC API. You should import these a lot in the
 * "/src/server/api/routers" directory.
 */

/**
 * This is how you create new routers and sub-routers in your tRPC API.
 *
 * @see https://trpc.io/docs/router
 */
export const createTRPCRouter = t.router

/**
 * Public (unauthenticated) procedure
 *
 * This is the base piece you use to build new queries and mutations on your tRPC API. It does not
 * guarantee that a user querying is authorized, but you can still access user session data if they
 * are logged in.
 */
export const publicProcedure = t.procedure

export const anonymousProtectedProcedure = t.procedure.use(
	t.middleware(({ ctx, next }) => {
		if (!ctx.session?.user)
			throw new TRPCError({
				code: "UNAUTHORIZED",
				message: "isAuthenticated() failed"
			})

		return next({
			ctx: {
				session: {
					...ctx.session,
					user: ctx.session.user
				}
			}
		})
	})
)

export const protectedProcedure = t.procedure.use(
	t.middleware(({ ctx, next }) => {
		if (!ctx.session?.user)
			throw new TRPCError({
				code: "UNAUTHORIZED",
				message: "isAuthenticated() failed"
			})

		if (!ctx.session?.address || ctx.session.address.startsWith("0x") === false || ctx.session.user.anonymous)
			throw new TRPCError({
				code: "UNAUTHORIZED",
				message: "isNonAnonymousAuthenticated() failed"
			})

		return next({
			ctx: {
				session: {
					...ctx.session,
					address: ctx.session.address,
					user: ctx.session.user
				}
			}
		})
	})
)
