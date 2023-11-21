import { TRPCError, initTRPC } from "@trpc/server";

import { Context } from "./context";

const t = initTRPC.context<Context>().create({
	errorFormatter({ shape }) {
		return shape;
	},
});

export const router = t.router;

export const publicProcedure = t.procedure;

export const middleware = t.middleware;

export const mergeRouters = t.mergeRouters;

const isAuthed = middleware(({ next, ctx }) => {
	const user = ctx.session?.user;

	console.log("context", ctx);

	if (!user?.name) {
		throw new TRPCError({ code: "UNAUTHORIZED" });
	}

	return next({
		ctx: {
			user: {
				...user,
				name: user.name,
			},
		},
	});
});

export const authedProcedure = t.procedure.use(isAuthed);
