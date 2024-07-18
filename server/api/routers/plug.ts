import { z } from "zod"

import { Prisma } from "@prisma/client"
import { TRPCError } from "@trpc/server"
import { observable } from "@trpc/server/observable"

import { colors } from "@/lib"
import {
	createTRPCRouter,
	protectedProcedure,
	publicProcedure
} from "@/server/api/trpc"

import { action } from "./action"

const workflow = Prisma.validator<Prisma.WorkflowDefaultArgs>()({})
export type Workflow = Prisma.WorkflowGetPayload<typeof workflow>

export const events = {
	add: "add-plug",
	rename: "rename-plug",
	edit: "edit-plug",
	delete: "delete-plug",
	view: "view-plug"
} as const

const views: Record<string, number> = {}

const subscription = (event: string) =>
	protectedProcedure.subscription(async ({ ctx }) => {
		return observable<Workflow>(emit => {
			// Only send events to the user that created the Plug.
			const handleSubscription = (data: Workflow) => {
				emit.next(data)
			}

			ctx.emitter.on(event, handleSubscription)

			return () => ctx.emitter.off(event, handleSubscription)
		})
	})

export const plug = createTRPCRouter({
	get: publicProcedure
		.input(z.string().optional())
		.query(async ({ input, ctx }) => {
			try {
				if (input === undefined)
					throw new TRPCError({ code: "BAD_REQUEST" })

				const plug = await ctx.db.workflow.findUnique({
					where: {
						id: input,
						isPrivate: false
					}
				})

				return plug
			} catch (error) {
				throw new TRPCError({ code: "BAD_REQUEST" })
			}
		}),

	infinite: protectedProcedure
		.input(
			z.object({
				mine: z.boolean().optional(),
				cursor: z.string().nullish(),
				limit: z.number().optional().default(10),
				search: z.string().optional(),
				tag: z.string().optional()
			})
		)
		.query(async ({ input, ctx }) => {
			const { cursor, search, tag } = input

			const sessionWhere =
				input.mine === true
					? {
							userAddress: ctx.session.address
						}
					: {
							isPrivate: false,
							isCurated: false,
							userAddress: {
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
								not: input.mine ? undefined : "Untitled Plug"
							}
						: {
								not: input.mine ? undefined : "Untitled Plug"
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
								not: input.mine ? undefined : "Untitled Plug"
							}
						: {
								not: input.mine ? undefined : "Untitled Plug"
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
				orderBy: { createdAt: "asc" }
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

	all: protectedProcedure
		.input(
			z.object({
				target: z.union([
					z.literal("mine"),
					z.literal("others"),
					z.literal("curated")
				]),
				search: z.string().optional(),
				tag: z.string().optional(),
				limit: z.number().optional()
			})
		)
		.query(async ({ input, ctx }) => {
			// .query = GET
			try {
				if (input.target === "mine")
					return await ctx.db.workflow.findMany({
						where: {
							userAddress: ctx.session.address
						},
						take: input.limit ? input.limit : undefined,
						orderBy: {
							updatedAt: "desc"
						}
					})

				if (input.target === "curated")
					return await ctx.db.workflow.findMany({
						where: {
							isPrivate: false,
							isCurated: true,
							actions: { not: "[]" }
						},
						take: input.limit ? input.limit : undefined,
						orderBy: {
							updatedAt: "desc"
						}
					})

				return await ctx.db.workflow.findMany({
					where: {
						isPrivate: false,
						userAddress: {
							not: ctx.session.address
						},
						actions: { not: "[]" }
					},
					take: input.limit ? input.limit : undefined,
					orderBy: {
						updatedAt: "desc"
					}
				})
			} catch (error) {
				throw new TRPCError({ code: "BAD_REQUEST" })
			}
		}),
	add: protectedProcedure
		.input(z.string().optional())
		.mutation(async ({ input, ctx }) => {
			try {
				const plug = await ctx.db.workflow.create({
					data: {
						name: "Untitled Plug",
						userAddress: ctx.session.address,
						color: Object.keys(colors)[
							Math.floor(
								Math.random() * Object.keys(colors).length
							)
						] as keyof typeof colors
					}
				})

				ctx.emitter.emit(events.add, plug)

				return { plug, from: input }
			} catch (error) {
				throw new TRPCError({ code: "BAD_REQUEST" })
			}
		}),
	fork: protectedProcedure
		.input(z.object({ id: z.string(), from: z.string().optional() }))
		.mutation(async ({ input, ctx }) => {
			try {
				const forking = await ctx.db.workflow.findUnique({
					where: { id: input.id }
				})

				if (forking == null)
					throw new TRPCError({ code: "BAD_REQUEST" })

				const { id, ...forkingData } = forking

				const plug = await ctx.db.workflow.create({
					data: {
						...forkingData,
						name: forking.name,
						userAddress: ctx.session.address
					}
				})

				ctx.emitter.emit(events.add, plug)

				return { plug, from: input.from }
			} catch (error) {
				throw new TRPCError({ code: "BAD_REQUEST" })
			}
		}),
	edit: protectedProcedure
		.input(
			z.object({
				id: z.string().optional(),
				name: z.string(),
				color: z.string(),
				isPrivate: z.boolean()
			})
		)
		.mutation(async ({ input, ctx }) => {
			try {
				if (input.id === undefined)
					throw new TRPCError({ code: "BAD_REQUEST" })

				const { id, ...data } = input

				const editingPlug = await ctx.db.workflow.findUniqueOrThrow({
					where: {
						id
					}
				})

				if (editingPlug.userAddress !== ctx.session.address)
					throw new TRPCError({ code: "UNAUTHORIZED" })

				const plug = await ctx.db.workflow.update({
					where: {
						id
					},
					data
				})

				ctx.emitter.emit(events.edit, plug)

				return plug
			} catch (error) {
				throw new TRPCError({ code: "BAD_REQUEST" })
			}
		}),
	delete: protectedProcedure
		.input(z.object({ id: z.string(), from: z.string().optional() }))
		.mutation(async ({ input, ctx }) => {
			try {
				const deletingPlug = await ctx.db.workflow.findUniqueOrThrow({
					where: {
						id: input.id
					}
				})

				if (deletingPlug.userAddress !== ctx.session.address)
					throw new TRPCError({ code: "UNAUTHORIZED" })

				const plug = await ctx.db.workflow.delete({
					where: {
						id: input.id
					}
				})

				ctx.emitter.emit(events.delete, plug)

				return { plug, from: input.from }
			} catch (error) {
				throw new TRPCError({ code: "BAD_REQUEST" })
			}
		}),

	// below are subscriptions
	onAdd: subscription(events.add), // sends message to subscribers
	onEdit: subscription(events.edit),
	onDelete: subscription(events.delete),

	onView: protectedProcedure // when subscription is opened, user is logged in
		.input(z.string().optional())
		.subscription(async ({ input, ctx }) => {
			// subscription logic
			try {
				if (input === undefined)
					throw new TRPCError({ code: "BAD_REQUEST" })

				return observable<number>(emit => {
					// in memory state manager, this gives us a constant stream
					const handleSubscription = () => {
						// define
						emit.next(views[input])
					}

					views[input] = (views[input] ?? 0) + 1 // increment
					ctx.emitter.on(events.view, handleSubscription) // call handleSubscription
					ctx.emitter.emit(events.view) // show results from handleSubscription

					return () => {
						views[input] = views[input] - 1 // decrement
						ctx.emitter.emit(events.view)
						ctx.emitter.off(events.view, handleSubscription) //
					}
				})
			} catch (error) {
				throw new TRPCError({ code: "BAD_REQUEST" })
			}
		}),

	action
})
