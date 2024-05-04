// import { z } from "zod"
//
// import { Prisma } from "@prisma/client"
// import { TRPCError } from "@trpc/server"
// import { observable } from "@trpc/server/observable"
//
// import { createTRPCRouter, protectedProcedure } from "@/server/api/trpc"
// import { emitter } from "@/server/emitter"
//
// export const component = Prisma.validator<Prisma.ComponentDefaultArgs>()({})
//
// export type Component = Prisma.ComponentGetPayload<typeof component>
//
// export const ComponentSchema = z.object({
// 	id: z.string(),
// 	top: z.number(),
// 	left: z.number(),
// 	width: z.number(),
// 	height: z.number(),
// 	content: z.string()
// })
//
// const events = {
// 	add: "add-component",
// 	move: "move-component",
// 	update: "update-component"
// }
//
// export default createTRPCRouter({
// 	add: protectedProcedure
// 		.input(
// 			z.object({
// 				id: z.string(),
// 				component: ComponentSchema.omit({ id: true })
// 			})
// 		)
// 		.mutation(async ({ ctx, input }) => {
// 			try {
// 				const canvas = await ctx.db.canvas.findUnique({
// 					where: {
// 						id: input.id
// 					}
// 				})
//
// 				if (!canvas) throw new TRPCError({ code: "NOT_FOUND" })
//
// 				if (canvas.userId !== ctx.session.user.name)
// 					throw new TRPCError({ code: "UNAUTHORIZED" })
//
// 				const component = await ctx.db.component.create({
// 					data: {
// 						...input.component,
// 						canvasId: input.id
// 					}
// 				})
//
// 				emitter.emit(events.add, component)
//
// 				return component
// 			} catch (e) {
// 				throw new TRPCError({ code: "BAD_REQUEST" })
// 			}
// 		}),
//
// 	move: protectedProcedure
// 		.input(
// 			z
// 				.object({
// 					id: z.string(),
// 					component: z.object({
// 						id: z.string(),
// 						top: z.number().optional(),
// 						left: z.number().optional()
// 					})
// 				})
// 				.transform(data => {
// 					if (!data.component.left && !data.component.top)
// 						throw new TRPCError({ code: "BAD_REQUEST" })
//
// 					return data
// 				})
// 		)
// 		.mutation(async ({ ctx, input }) => {
// 			try {
// 				const canvas = await ctx.db.canvas.findUnique({
// 					where: {
// 						id: input.id
// 					}
// 				})
//
// 				if (!canvas) throw new TRPCError({ code: "NOT_FOUND" })
//
// 				if (canvas.userId !== ctx.session.user.name)
// 					throw new TRPCError({ code: "UNAUTHORIZED" })
//
// 				const component = await ctx.db.component.update({
// 					where: {
// 						id: input.component.id,
// 						canvasId: input.id
// 					},
// 					data: {
// 						top: input.component.top,
// 						left: input.component.left
// 					}
// 				})
//
// 				emitter.emit(events.move, component)
//
// 				return component
// 			} catch (e) {
// 				throw new TRPCError({ code: "BAD_REQUEST" })
// 			}
// 		}),
//
// 	onAdd: protectedProcedure
// 		.input(z.string())
// 		.subscription(async ({ ctx, input }) => {
// 			try {
// 				const canvas = await ctx.db.canvas.findUnique({
// 					where: {
// 						id: input
// 					}
// 				})
//
// 				if (!canvas) throw new TRPCError({ code: "NOT_FOUND" })
//
// 				if (
// 					canvas.userId !== ctx.session.user.name &&
// 					canvas.public === false
// 				)
// 					throw new TRPCError({ code: "UNAUTHORIZED" })
//
// 				return observable<Component>(emit => {
// 					const handleAdd = (data: Component) => {
// 						if (input !== data.canvasId) return
//
// 						emit.next(data)
// 					}
//
// 					emitter.on(events.add, handleAdd)
//
// 					return () => {
// 						emitter.off(events.add, handleAdd)
// 					}
// 				})
// 			} catch (e) {
// 				throw new TRPCError({ code: "BAD_REQUEST" })
// 			}
// 		}),
//
// 	onMove: protectedProcedure
// 		.input(z.string())
// 		.subscription(async ({ ctx, input }) => {
// 			try {
// 				const canvas = await ctx.db.canvas.findUnique({
// 					where: {
// 						id: input
// 					}
// 				})
//
// 				if (!canvas) throw new TRPCError({ code: "NOT_FOUND" })
//
// 				if (
// 					canvas.userId !== ctx.session.user.name &&
// 					canvas.public === false
// 				)
// 					throw new TRPCError({ code: "UNAUTHORIZED" })
//
// 				return observable<Component>(emit => {
// 					const handleMove = (data: Component) => {
// 						if (input !== data.canvasId) return
//
// 						emit.next(data)
// 					}
//
// 					emitter.on(events.move, handleMove)
//
// 					return () => {
// 						emitter.off(events.move, handleMove)
// 					}
// 				})
// 			} catch (e) {
// 				throw new TRPCError({ code: "BAD_REQUEST" })
// 			}
// 		})
// })
