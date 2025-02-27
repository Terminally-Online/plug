import { TRPCError } from "@trpc/server"

import { z } from "zod"

import { SOCKET_BASE_QUERY } from "@/lib"

import { anonymousProtectedProcedure, apiKeyProcedure, createTRPCRouter } from "../../trpc"

export const onboard = createTRPCRouter({
	onboard: anonymousProtectedProcedure
		.input(
			z.object({
				onboardingColor: z.string().optional()
			})
		)
		.mutation(async ({ ctx, input }) => {
			await ctx.db.socketIdentity.update({
				where: { socketId: ctx.session.address },
				data: {
					onboardingAt: new Date(),
					onboardingColor: input.onboardingColor
				}
			})

			const socket = await ctx.db.socket.findUnique({
				where: { id: ctx.session.address },
				...SOCKET_BASE_QUERY
			})

			if (!socket) {
				throw new TRPCError({
					code: "NOT_FOUND",
					message: "Socket not found"
				})
			}

			return socket
		}),

	onboarding: apiKeyProcedure.query(async ({ ctx }) => {
		return await ctx.db.socketIdentity.findMany({
			where: {
				onboardedAt: null,
				onboardingAt: { not: null }
			}
		})
	}),

	onboarded: apiKeyProcedure.input(z.array(z.string())).mutation(async ({ ctx, input }) => {
		await ctx.db.socketIdentity.updateMany({
			where: {
				socketId: {
					in: input
				}
			},
			data: {
				onboardedAt: new Date()
			}
		})
	})
})
