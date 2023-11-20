import { authedProcedure, router, publicProcedure } from "../trpc";
import { observable } from "@trpc/server/observable";
import { clearInterval } from "timers";

import { p } from "../prisma";

import { z } from "zod";
import { TRPCError } from "@trpc/server";

export const appRouter = router({
	healthcheck: publicProcedure.query(() => "yay!"),
	randomNumber: publicProcedure.subscription(() => {
		return observable<number>((emit) => {
			const int = setInterval(() => {
				emit.next(Math.random());
			}, 500);

			return () => {
				clearInterval(int);
			};
		});
	}),
	all: authedProcedure.query(async ({ ctx }) => {
		const userId = ctx.session?.user?.name;

		if (!userId) throw new TRPCError({ code: "UNAUTHORIZED" });

		// * Get the canvases from the database.
		const canvases = await p.canvas.findMany({
			where: {
				userId,
			},
		});

		// * Return the canvases.
		return canvases;
	}),
	// TODO: I want to implement a .infinite() method that returns a stream of
	//       canvases that are created by the user so that we do not send a massive
	//       amount of data to the client.
	get: authedProcedure.input(z.string()).query(async ({ ctx, input }) => {
		const userId = ctx.session?.user?.name;

		const canvas = await p.canvas.findUnique({
			where: {
				id: input,
			},
			include: { components: true },
		});

		if (!canvas) throw new TRPCError({ code: "NOT_FOUND" });

		if (canvas.public) return canvas;

		if (!userId) throw new TRPCError({ code: "UNAUTHORIZED" });

		if (canvas.userId !== userId)
			throw new TRPCError({ code: "FORBIDDEN" });

		return canvas;
	}),
	create: authedProcedure
		.input(
			z.object({
				name: z.string(),
				public: z.boolean(),
			})
		)
		.mutation(async ({ ctx, input }) => {
			const userId = ctx.session?.user?.name;

			if (!userId) throw new TRPCError({ code: "UNAUTHORIZED" });

			// * Create the canvas in the database.
			return await p.canvas.create({
				data: {
					name: input.name,
					public: input.public,
					user: {
						connectOrCreate: {
							where: { id: userId },
							create: { id: userId },
						},
					},
				},
				include: { components: true },
			});
		}),
	// TODO: Update
	// TODO: Delete (needs to revoke all the active signatures as well onchain)
});

export type AppRouter = typeof appRouter;
