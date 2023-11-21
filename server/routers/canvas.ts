import { EventEmitter } from "stream";

import { ComponentType, Prisma } from "@prisma/client";

import { authedProcedure, router, publicProcedure } from "../trpc";
import { observable } from "@trpc/server/observable";

import { p } from "../prisma";

import { z } from "zod";
import { TRPCError } from "@trpc/server";

import componentRouter, { ComponentSchema } from "./component";

const emitter = new EventEmitter();

export const canvasWithComponents =
  Prisma.validator<Prisma.CanvasDefaultArgs>()({
    include: { components: true },
  });

export type CanvasWithComponents = Prisma.CanvasGetPayload<
  typeof canvasWithComponents
>;

export const CanvasSchema = z.object({
  id: z.string(),
  name: z.string(),
  public: z.boolean(),
  color: z.string(),
  components: z.array(ComponentSchema),
  createdAt: z.string().optional(),
  updatedAt: z.string().optional(),
});

export default router({
  all: authedProcedure.query(async ({ ctx }) => {
    const userId = ctx.user.name;

    // * Get the canvases from the database.
    const canvases = await p.canvas.findMany({
      where: {
        userId,
      },
    });

    // * Return the canvases.
    return canvases;
  }),
  get: publicProcedure.input(z.string()).query(async ({ ctx, input }) => {
    const canvas = await p.canvas.findUnique({
      where: {
        id: input,
      },
      include: { components: true },
    });

    if (!canvas) throw new TRPCError({ code: "NOT_FOUND" });

    if (canvas.public) return canvas;

    const userId = ctx.session?.user?.name;

    // if (!userId) throw new TRPCError({ code: "UNAUTHORIZED" });
    //
    // if (canvas.userId !== userId) throw new TRPCError({ code: "FORBIDDEN" });

    return canvas as CanvasWithComponents;
  }),
  create: authedProcedure
    .input(
      z.object({
        name: z.string(),
        public: z.boolean(),
      })
    )
    .mutation(async ({ ctx, input }) => {
      const userId = ctx.user.name;

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
  update: authedProcedure
    .input(
      z.object({
        id: z.string(),
        name: z.string().optional(),
        color: z.string().optional(),
        public: z.boolean().optional(),
        components: z.array(ComponentSchema),
      })
    )
    .mutation(async ({ ctx, input }) => {
      const userId = ctx.user.name;

      const canvas = await p.canvas.findUnique({
        where: {
          id: input.id,
        },
        include: { components: true },
      });

      if (!canvas) throw new TRPCError({ code: "NOT_FOUND" });

      if (canvas.userId !== userId) throw new TRPCError({ code: "FORBIDDEN" });

      // * Update the fields that were passed in.
      const updatedCanvas: CanvasWithComponents = await p.canvas.update({
        where: {
          id: input.id,
        },
        data: {
          ...canvas,
          ...input,
          // TODO: For now we are not updating components
          components: undefined,
        },
        include: { components: true },
      });

      // * Emit an update event.
      emitter.emit("update", updatedCanvas);

      return updatedCanvas;
    }),
  onUpdate: authedProcedure.subscription(() => {
    return observable<CanvasWithComponents>((emit) => {
      emitter.on("update", emit.next);

      return () => {
        emitter.off("update", emit.next);
      };
    });
  }),
  // TODO: onUpdate
  component: componentRouter,
});
