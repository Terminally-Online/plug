import { observable } from "@trpc/server/observable"

import { anonymousProtectedProcedure, protectedProcedure } from "@/server/api/trpc"

const scopes = ["anonymous", "protected"] as const

const procedure = (scope: (typeof scopes)[number] = "anonymous") =>
	scope === "anonymous" ? anonymousProtectedProcedure : protectedProcedure

export const subscription = <T>(scope: (typeof scopes)[number], event: string) =>
	procedure(scope).subscription(async ({ ctx }) => {
		return observable<T>(emit => {
			const handleSubscription = (data: T) => {
				emit.next(data)
			}

			ctx.emitter.on(event, handleSubscription)

			return () => ctx.emitter.off(event, handleSubscription)
		})
	})

export const subscriptions = {
	socket: {},
	plugs: {
		add: "add-plug",
		rename: "rename-plug",
		edit: "edit-plug",
		delete: "delete-plug",
		queue: "queue-plug"
	}
} as const
