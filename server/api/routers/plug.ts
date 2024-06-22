import { z } from "zod"

import { Prisma } from "@prisma/client"
import { TRPCError } from "@trpc/server"
import { observable } from "@trpc/server/observable"

import { colors } from "@/lib/constants"
import {
	createTRPCRouter,
	protectedProcedure,
	publicProcedure
} from "@/server/api/trpc"

import { action } from "./action"

const workflow = Prisma.validator<Prisma.WorkflowDefaultArgs>()({
	include: { versions: { include: { actions: true } } }
})
export type Workflow = Prisma.WorkflowGetPayload<typeof workflow>

export const events = {
	add: "add-plug",
	rename: "rename-plug",
	edit: "edit-plug",
	delete: "delete-plug"
} as const

const orderBy = { updatedAt: "desc" } as const

const include = {
	versions: {
		include: {
			actions: { orderBy: { index: "asc" } }
		},
		orderBy: { version: "desc" }
	}
} as const

const includeAndOrderBy = {
	include,
	orderBy
} as const

const subsription = (event: string) =>
	protectedProcedure.subscription(async ({ ctx }) => {
		return observable<Workflow>(emit => {
			// Only send events to the user that created the Plug.
			const handleSubscription = (data: Workflow) => {
				if (data.userAddress === ctx.session.address) emit.next(data)
			}

			ctx.emitter.on(event, handleSubscription)

			return () => ctx.emitter.off(event, handleSubscription)
		})
	})

export const plug = createTRPCRouter({
	all: protectedProcedure
		.input(
			z
				.object({
					search: z.string().optional(),
					tag: z.string().optional()
				})
				.optional()
		)
		.query(async ({ input, ctx }) => {
			try {
				if (input) {
					// if (input.search && input.tag) return

					if (input.search)
						return await ctx.db.workflow.findMany({
							where: {
								userAddress: ctx.session.address,
								name: {
									contains: input.search,
									mode: "insensitive"
								}
							},
							...includeAndOrderBy
						})

					// if (input.tag) return
				}

				return await ctx.db.workflow.findMany({
					where: { userAddress: ctx.session.address },
					...includeAndOrderBy
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
					},
					include
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
						name: `${forking.name} (Fork)`,
						userAddress: ctx.session.address
					},
					include
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

				const plug = await ctx.db.workflow.update({
					where: {
						id: input.id
					},
					data: {
						name: input.name,
						color: input.color,
						isPrivate: input.isPrivate
					},
					include
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

	onAdd: subsription(events.add),
	onEdit: subsription(events.edit),
	onDelete: subsription(events.delete),

	action
})
