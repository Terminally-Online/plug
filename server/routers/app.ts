import { router, publicProcedure } from "../trpc";
import { observable } from "@trpc/server/observable";
import { clearInterval } from "timers";

import canvasRouter from "./canvas";

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
	canvas: canvasRouter,
});

export type AppRouter = typeof appRouter;
