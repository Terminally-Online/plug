import { NextApiRequest, NextApiResponse } from "next"
import { type Session } from "next-auth"

import { initTRPC, TRPCError } from "@trpc/server"
import { type CreateNextContextOptions } from "@trpc/server/adapters/next"

import superjson from "superjson"
import { ZodError } from "zod"

import { env } from "@/env"
import { getServerAuthSession } from "@/server/auth"
import { db } from "@/server/db"
import { emitter } from "@/server/emitter"

interface CreateContextOptions {
	res: NextApiResponse
	session: Session | null
	headers?: NextApiRequest["headers"]
}

export const createInnerTRPCContext = ({ res, session, headers }: CreateContextOptions) => {
	return {
		res,
		session,
		db,
		emitter,
		headers
	}
}

export const createTRPCContext = async (opts: CreateNextContextOptions) => {
	const { req, res } = opts
	const session = await getServerAuthSession({ req, res })

	return createInnerTRPCContext({
		res,
		session,
		headers: req.headers
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

export const createTRPCRouter = t.router

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
export const apiKeyProcedure = t.procedure.use(
	t.middleware(({ ctx, next }) => {
		if (!ctx.headers) {
			throw new TRPCError({
				code: "UNAUTHORIZED",
				message: "No headers found"
			})
		}

		const apiKey = ctx.headers["x-api-key"]

		if (!apiKey || apiKey !== env.SOLVER_API_KEY) {
			throw new TRPCError({
				code: "UNAUTHORIZED",
				message: "Invalid API key"
			})
		}

		return next({
			ctx: {
				isAdmin: true
			}
		})
	})
)
