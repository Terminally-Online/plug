import { EventEmitter } from "stream";

import { Prisma } from "@prisma/client";

import { protectedProcedure, createTRPCRouter, publicProcedure } from "../trpc";
import { observable } from "@trpc/server/observable";

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

export default createTRPCRouter({
  test: publicProcedure.query(() => "yay!"),
  all: protectedProcedure.query(async ({ ctx }) => {
    const userId = ctx.session.user.id;

    // * Get the canvases from the database.
    const canvases = await ctx.db.canvas.findMany({
      where: {
        userId,
      },
    });

    // * Return the canvases.
    return canvases;
  }),
  get: publicProcedure.input(z.string()).query(async ({ ctx, input }) => {
    const canvas = await ctx.db.canvas.findUnique({
      where: {
        id: input,
      },
      include: { components: true },
    });

    if (!canvas) throw new TRPCError({ code: "NOT_FOUND" });

    if (canvas.public) return canvas;

    const userId = ctx.session?.user?.name;

    console.log("userid when getting", ctx);

    // if (!userId) throw new TRPCError({ code: "UNAUTHORIZED" });
    //
    // if (canvas.userId !== userId) throw new TRPCError({ code: "FORBIDDEN" });

    return canvas as CanvasWithComponents;
  }),
  create: protectedProcedure
    .input(
      z.object({
        name: z.string(),
        public: z.boolean(),
      })
    )
    .mutation(async ({ ctx, input }) => {
      const userId = ctx.session.user.id;

      // * Create the canvas in the database.
      return await ctx.db.canvas.create({
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
  update: protectedProcedure
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
      const userId = ctx.session.user.id;

      const canvas = await ctx.db.canvas.findUnique({
        where: {
          id: input.id,
        },
        include: { components: true },
      });

      if (!canvas) throw new TRPCError({ code: "NOT_FOUND" });

      if (canvas.userId !== userId) throw new TRPCError({ code: "FORBIDDEN" });

      // * Update the fields that were passed in.
      const updatedCanvas: CanvasWithComponents = await ctx.db.canvas.update({
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
  onUpdate: protectedProcedure.subscription(() => {
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
