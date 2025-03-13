import { on } from "node:events"

import { anonymousProtectedProcedure, protectedProcedure } from "@/server/api/trpc"

const scopes = ["anonymous", "protected"] as const

const procedure = (scope: (typeof scopes)[number] = "anonymous") =>
	scope === "anonymous" ? anonymousProtectedProcedure : protectedProcedure

export const subscription = <T>(scope: (typeof scopes)[number], event: string) =>
	procedure(scope).subscription(async function* (opts) {
		for await (const [data] of on(opts.ctx.emitter, event, {
			signal: opts.signal
		})) {
			const post = data as T
			yield post
		}
	})

export const subscriptions = {
	socket: {},
	plugs: {
		add: "add-plug",
		rename: "rename-plug",
		edit: "edit-plug",
		delete: "delete-plug"
	},
	execution: {
		update: "update-execution",
		delete: "delete-execution"
	}
} as const
