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

// const action = Prisma.validator<Prisma.ActionDefaultArgs>()
// export type Action = Prisma.ActionGetPayload<typeof action>

const workflow = Prisma.validator<Prisma.WorkflowDefaultArgs>()({
	include: { versions: { include: { actions: true } } }
})
export type Workflow = Prisma.WorkflowGetPayload<typeof workflow>

export const events = {
	add: "add-plug",
	rename: "rename-plug",
	edit: "edit-plug"
}

export const plug = createTRPCRouter({
	preview: publicProcedure.input(z.string()).query(async ({ ctx }) => {
		// TODO: When there is no account logged in, show the top curated ones.
		// if (ctx.session == null) return []

		// TODO: When there is an account logged in, show the plugs that they
		// 	     have used most / created most recently.
		// TODO: When there is an account logged in, and the owned plugs is less than
		//       4, show a mix of curated and owned.

		return []
	}),
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
							orderBy: { updatedAt: "desc" },
							include: {
								versions: {
									include: {
										actions: { orderBy: { index: "asc" } }
									},
									orderBy: { version: "desc" }
								}
							}
						})

					// if (input.tag) return
				}

				return await ctx.db.workflow.findMany({
					where: { userAddress: ctx.session.address },
					orderBy: { updatedAt: "desc" },
					include: {
						versions: {
							include: {
								actions: { orderBy: { index: "asc" } }
							},
							orderBy: { version: "desc" }
						}
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
				const color = Object.keys(colors)[
					Math.floor(Math.random() * Object.keys(colors).length)
				] as keyof typeof colors

				const plug = await ctx.db.workflow.create({
					data: {
						name: "Untitled Plug",
						userAddress: ctx.session.address,
						color
					},
					include: {
						versions: {
							include: {
								actions: { orderBy: { index: "asc" } }
							},
							orderBy: { version: "desc" }
						}
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
						name: `${forking.name} (Fork)`,
						userAddress: ctx.session.address
					},
					include: {
						versions: {
							include: {
								actions: { orderBy: { index: "asc" } }
							},
							orderBy: { version: "desc" }
						}
					}
				})

				ctx.emitter.emit(events.add, plug)

				return { plug, from: input.from }
			} catch (error) {
				throw new TRPCError({ code: "BAD_REQUEST" })
			}
		}),
	onAdd: protectedProcedure.subscription(async ({ ctx }) => {
		return observable<Workflow>(emit => {
			const handleAdd = (data: Workflow) => {
				if (data.userAddress === ctx.session.address) emit.next(data)
			}
			ctx.emitter.on(events.add, handleAdd)
			return () => ctx.emitter.off(events.add, handleAdd)
		})
	}),
	edit: protectedProcedure
		.input(
			z.object({
				id: z.string(),
				name: z.string(),
				color: z.string(),
				isPrivate: z.boolean()
			})
		)
		.mutation(async ({ input, ctx }) => {
			try {
				const plug = await ctx.db.workflow.update({
					where: {
						id: input.id
					},
					data: {
						name: input.name,
						color: input.color,
						isPrivate: input.isPrivate
					},
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
	onEdit: protectedProcedure.subscription(async ({ ctx }) => {
		return observable<Workflow>(emit => {
			const handleEdit = (data: Workflow) => {
				if (data.userAddress === ctx.session.address) emit.next(data)
			}
			ctx.emitter.on(events.edit, handleEdit)
			return () => ctx.emitter.off(events.edit, handleEdit)
		})
	}),
	action
})
