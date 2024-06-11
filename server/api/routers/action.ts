import { z } from "zod"

import { TRPCError } from "@trpc/server"

import { createTRPCRouter, protectedProcedure } from "@/server/api/trpc"

import { events } from "./plug"

export const action = createTRPCRouter({
	add: protectedProcedure
		.input(
			z.object({
				id: z.string(),
				version: z.number(),
				action: z.object({
					categoryName: z.string(),
					actionName: z.string(),
					data: z.string()
				})
			})
		)
		.mutation(async ({ input, ctx }) => {
			try {
				const versionNumber =
					(await ctx.db.version.count({
						where: { workflowId: input.id }
					})) + 1

				// If there are not any previous versions, we do not have any actions to
				// carry over, so we create a new version with the current action
				if (versionNumber === 1) {
					await ctx.db.version.create({
						data: {
							workflowId: input.id,
							version: versionNumber,
							actions: {
								create: {
									index: 1,
									...input.action
								}
							}
						}
					})
				}
				// Create a new version based on the most recent version and add
				// an action to it at the end of actions.
				else {
					const previousVersion = await ctx.db.version.findFirst({
						where: {
							workflowId: input.id,
							version: input.version
						},
						include: { actions: { orderBy: { index: "asc" } } },
						orderBy: { version: "desc" }
					})

					if (previousVersion == null)
						throw new TRPCError({ code: "BAD_REQUEST" })

					await ctx.db.version.create({
						data: {
							workflowId: input.id,
							version: versionNumber,
							actions: {
								create: [
									...previousVersion.actions.map(action => ({
										...action,
										id: undefined,
										versionId: undefined,
										index: action.index
									})),
									{
										index:
											previousVersion.actions.length + 1,
										...input.action
									}
								]
							}
						}
					})
				}

				const plug = await ctx.db.workflow.findUnique({
					where: { id: input.id },
					include: {
						versions: {
							include: {
								actions: { orderBy: { index: "asc" } }
							},
							orderBy: { version: "desc" }
						}
					}
				})

				ctx.emitter.emit(events.edit, plug)

				return plug
			} catch (error) {
				throw new TRPCError({ code: "BAD_REQUEST" })
			}
		}),
	remove: protectedProcedure
		.input(z.string())
		.mutation(async ({ input, ctx }) => {
			try {
				const action = await ctx.db.action.findUnique({
					where: { id: input },
					include: { version: { include: { workflow: true } } }
				})

				if (action === null)
					throw new TRPCError({ code: "BAD_REQUEST" })

				const versionNumber =
					(await ctx.db.version.count({
						where: { workflowId: action.version.workflowId }
					})) + 1

				const version = await ctx.db.version.findFirst({
					where: {
						workflowId: action.version.workflowId,
						version: action.version.version
					},
					include: {
						actions: { orderBy: { index: "asc" } }
					}
				})

				if (version == null)
					throw new TRPCError({ code: "BAD_REQUEST" })

				await ctx.db.version.create({
					data: {
						workflowId: action.version.workflowId,
						version: versionNumber,
						actions: {
							create: version.actions
								.filter(a => a.id !== action.id)
								.map((a, i) => ({
									...a,
									id: undefined,
									versionId: undefined,
									index: i + 1
								}))
						}
					}
				})

				const plug = await ctx.db.workflow.findUnique({
					where: { id: action.version.workflowId },
					include: {
						versions: {
							include: {
								actions: { orderBy: { index: "asc" } }
							},
							orderBy: { version: "desc" }
						}
					}
				})

				ctx.emitter.emit(events.edit, plug)

				return plug
			} catch (error) {
				throw new TRPCError({ code: "BAD_REQUEST" })
			}
		})
})
