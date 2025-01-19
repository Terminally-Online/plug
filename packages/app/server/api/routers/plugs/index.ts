import { TRPCError } from "@trpc/server"

import { z } from "zod"

import { Prisma } from "@prisma/client"

import { colors } from "@/lib"
import { action } from "@/server/api/routers/plugs/action"
import { activity } from "@/server/api/routers/plugs/activity"
import { anonymousProtectedProcedure, createTRPCRouter, publicProcedure } from "@/server/api/trpc"
import { subscription, subscriptions } from "@/server/subscription"

const workflow = Prisma.validator<Prisma.WorkflowDefaultArgs>()({})
export type Workflow = Prisma.WorkflowGetPayload<typeof workflow>

export const plugs = createTRPCRouter({
	get: publicProcedure
		.input(
			z.object({
				ids: z.array(z.string()),
				viewed: z.array(z.string())
			})
		)
		.query(async ({ input, ctx }) => {
			if (input.ids.length === 0) return []

			const weekStart = new Date()
			weekStart.setDate(weekStart.getDate() - ((weekStart.getDay() + 6) % 7))
			weekStart.setUTCHours(0, 0, 0, 0)

			const workflows = await ctx.db.workflow.findMany({
				where: {
					id: { in: input.ids },
					isPrivate: false
				}
			})

			const newViews = workflows
				.filter(
					w => !input.viewed.includes(w.id) && w.socketId !== ctx.session?.address // Exclude creator views
				)
				.map(w => ({
					workflowId: w.id,
					date: weekStart,
					views: 1
				}))

			if (newViews.length > 0) {
				await ctx.db.view.createMany({
					data: newViews,
					skipDuplicates: true
				})

				await ctx.db.view.updateMany({
					where: {
						workflowId: { in: newViews.map(v => v.workflowId) },
						date: weekStart
					},
					data: {
						views: {
							increment: 1
						}
					}
				})
			}

			return workflows
		}),

	infinite: anonymousProtectedProcedure
		.input(
			z.object({
				address: z.string().optional(),
				cursor: z.string().nullish(),
				limit: z.number().optional().default(10),
				search: z.string().optional(),
				tag: z.string().optional()
			})
		)
		.query(async ({ input, ctx }) => {
			const { cursor, search, tag } = input

			const mine = input.address === ctx.session.address

			const sessionWhere =
				mine === true
					? {
							socketId: ctx.session.address
						}
					: {
							isPrivate: false,
							isCurated: false,
							socketId: {
								not: ctx.session.address
							},
							actions: { not: "[]" }
						}

			const count = await ctx.db.workflow.count({
				where: {
					name: search
						? {
								contains: search,
								mode: "insensitive",
								not: mine ? undefined : "Untitled Plug"
							}
						: {
								not: mine ? undefined : "Untitled Plug"
							},
					tags: tag
						? {
								has: tag
							}
						: undefined,
					...sessionWhere
				}
			})

			const plugs = await ctx.db.workflow.findMany({
				where: {
					name: search
						? {
								contains: search,
								mode: "insensitive",
								not: mine ? undefined : "Untitled Plug"
							}
						: {
								not: mine ? undefined : "Untitled Plug"
							},
					tags: tag
						? {
								has: tag
							}
						: undefined,
					...sessionWhere
				},
				cursor: cursor ? { id: cursor } : undefined,
				take: input.limit + 1,
				orderBy: { createdAt: "asc" },
				include: {
					socket: true
				}
			})

			let nextCursor: typeof cursor | undefined = undefined
			if (plugs.length > input.limit) {
				const nextItem = plugs.pop()
				nextCursor = nextItem!.id
			}

			return {
				plugs: plugs,
				nextCursor,
				count
			}
		}),

	all: anonymousProtectedProcedure
		.input(
			z.object({
				target: z.union([z.literal("mine"), z.literal("others"), z.literal("curated")]),
				search: z.string().optional(),
				tag: z.string().optional(),
				limit: z.number().optional()
			})
		)
		.query(async ({ input, ctx }) => {
			try {
				const addForkCounts = async (ctx: any, workflows: any[]) => {
					const forkCounts = await Promise.all(
						workflows.map(async workflow => ({
							id: workflow.id,
							forks: await ctx.db.workflow.count({
								where: { workflowForkedId: workflow.id }
							})
						}))
					)
					return workflows.map(workflow => ({
						...workflow,
						forkCount: forkCounts.find(f => f.id === workflow.id)?.forks ?? 0
					}))
				}

				if (input.target === "mine") {
					const workflows = await ctx.db.workflow.findMany({
						where: {
							socketId: ctx.session.address
						},
						take: input.limit ? input.limit : undefined,
						orderBy: {
							updatedAt: "desc"
						},
						include: {
							socket: { include: { identity: { include: { ens: true } } } },
							_count: {
								select: {
									executions: true
								}
							},
							views: {
								where: {
									date: {
										gte: new Date(new Date().setDate(new Date().getDate() - 7))
									}
								},
								select: {
									views: true
								}
							}
						}
					})
					return addForkCounts(ctx, workflows)
				}

				if (input.target === "curated")
					return addForkCounts(
						ctx,
						await ctx.db.workflow.findMany({
							where: {
								isPrivate: false,
								isCurated: true,
								actions: { not: "[]" }
							},
							take: input.limit ? input.limit : undefined,
							orderBy: {
								updatedAt: "desc"
							},
							include: {
								socket: { include: { identity: { include: { ens: true } } } },
								_count: {
									select: {
										executions: true
									}
								},
								views: {
									where: {
										date: {
											gte: new Date(new Date().setDate(new Date().getDate() - 7)) // Last 7 days
										}
									},
									select: {
										views: true
									}
								}
							}
						})
					)

				return addForkCounts(
					ctx,
					await ctx.db.workflow.findMany({
						where: {
							isPrivate: false,
							socketId: {
								not: ctx.session.address
							},
							actions: { not: "[]" }
						},
						take: input.limit ? input.limit : undefined,
						orderBy: {
							updatedAt: "desc"
						},
						include: {
							socket: { include: { identity: { include: { ens: true } } } },
							_count: {
								select: {
									executions: true
								}
							},
							views: {
								where: {
									date: {
										gte: new Date(new Date().setDate(new Date().getDate() - 7)) // Last 7 days
									}
								},
								select: { views: true }
							}
						}
					})
				)
			} catch (error) {
				throw new TRPCError({ code: "BAD_REQUEST" })
			}
		}),

	add: anonymousProtectedProcedure
		.input(z.object({ index: z.number().optional(), from: z.string().optional() }).optional())
		.mutation(async ({ input, ctx }) => {
			try {
				const plug = await ctx.db.workflow.create({
					data: {
						name: "Untitled Plug",
						socketId: ctx.session.address,
						color: Object.keys(colors)[
							Math.floor(Math.random() * Object.keys(colors).length)
						] as keyof typeof colors
					}
				})

				ctx.emitter.emit(subscriptions.plugs.add, plug)

				return { plug, index: input?.index, from: input?.from }
			} catch (error) {
				throw new TRPCError({ code: "BAD_REQUEST" })
			}
		}),

	fork: anonymousProtectedProcedure
		.input(z.object({ plug: z.string(), index: z.number(), from: z.string() }))
		.mutation(async ({ input, ctx }) => {
			try {
				const forking = await ctx.db.workflow.findUnique({
					where: { id: input.plug }
				})

				if (forking == null) throw new TRPCError({ code: "BAD_REQUEST" })

				const { id, ...forkingData } = forking

				const name = forking.name.replace(/ \(#\d+\)$/, "")
				const count = await ctx.db.workflow.count({
					where: { workflowForkedId: forking.workflowForkedId || input.plug }
				})
				const forkNumber = count + 1
				const plug = await ctx.db.workflow.create({
					data: {
						...forkingData,
						name: `${name} (#${forkNumber})`,
						socketId: ctx.session.address,
						workflowForkedId: forking.workflowForkedId || input.plug
					}
				})

				ctx.emitter.emit(subscriptions.plugs.add, plug)

				return { plug, index: input.index, from: input.from }
			} catch (error) {
				throw new TRPCError({ code: "BAD_REQUEST" })
			}
		}),

	edit: anonymousProtectedProcedure
		.input(
			z.object({
				id: z.string().optional(),
				name: z.string(),
				namedAt: z.date().optional(),
				color: z.string(),
				isPrivate: z.boolean()
			})
		)
		.mutation(async ({ input, ctx }) => {
			try {
				if (input.id === undefined) throw new TRPCError({ code: "BAD_REQUEST" })

				const { id, ...data } = input

				const plug = await ctx.db.workflow.update({
					where: {
						id,
						socketId: ctx.session.address
					},
					data
				})

				ctx.emitter.emit(subscriptions.plugs.edit, plug)

				return plug
			} catch (error) {
				throw new TRPCError({ code: "BAD_REQUEST" })
			}
		}),

	delete: anonymousProtectedProcedure
		.input(z.object({ plug: z.string(), index: z.number(), from: z.string().nullish() }))
		.mutation(async ({ input, ctx }) => {
			const plug = await ctx.db.workflow.delete({
				where: {
					id: input.plug,
					socketId: ctx.session.address
				}
			})

			ctx.emitter.emit(subscriptions.plugs.delete, plug)

			return { plug, index: input.index, from: input.from }
		}),

	onAdd: subscription<Workflow>("anonymous", subscriptions.plugs.add),
	onEdit: subscription<Workflow>("anonymous", subscriptions.plugs.edit),
	onDelete: subscription<Workflow>("anonymous", subscriptions.plugs.delete),

	action,
	activity
})
