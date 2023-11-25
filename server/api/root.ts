import { createTRPCRouter } from "@/server/api/trpc";
import canvasRouter from "@/server/api/routers/canvas";

export const appRouter = createTRPCRouter({
  canvas: canvasRouter,
});

export type AppRouter = typeof appRouter;
