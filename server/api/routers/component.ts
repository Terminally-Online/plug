import { z } from "zod";
import { protectedProcedure, publicProcedure, createTRPCRouter } from "../trpc";
import { TRPCError } from "@trpc/server";

export const ComponentSchema = z.object({
  id: z.string(),
  top: z.number(),
  left: z.number(),
  type: z.union([z.literal("PLUG"), z.literal("BOX"), z.literal("MARKDOWN")]),
  width: z.number(),
  height: z.number(),
  content: z.string(),
  createdAt: z.string().optional(),
  updatedAt: z.string().optional(),
});

export default createTRPCRouter({
  add: protectedProcedure
    .input(
      z.object({
        id: z.string(),
        component: ComponentSchema.omit({ id: true }),
      })
    )
    .mutation(async ({ ctx, input }) => {
      const canvas = await ctx.db.canvas.findUnique({
        where: {
          id: input.id,
        },
        include: { components: true },
      });

      if (!canvas) throw new TRPCError({ code: "NOT_FOUND" });

      const component = await ctx.db.component.create({
        data: {
          ...input.component,
          canvasId: input.id,
        },
      });

      return component;
    }),
  move: publicProcedure
    .input(
      z.object({
        id: z.string(),
        component: z.object({
          id: z.string(),
          top: z.number(),
          left: z.number(),
        }),
      })
    )
    .mutation(async ({ ctx, input }) => {
      const component = await ctx.db.component.update({
        where: {
          id: input.component.id,
          canvasId: input.id,
        },
        data: {
          top: input.component.top,
          left: input.component.left,
        },
      });

      return component;
    }),
  selecting: publicProcedure
    .input(
      z.object({
        id: z.string(),
        component: z.object({
          id: z.string(),
          selecting: z.string(),
        }),
      })
    )
    .mutation(async ({ ctx, input }) => {
      const component = await ctx.db.component.update({
        where: {
          id: input.component.id,
          canvasId: input.id,
        },
        data: {
          selecting: {
            connect: {
              id: input.component.selecting,
            },
          },
        },
      });

      return component;
    }),
});
