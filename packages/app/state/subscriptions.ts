import { useSession } from "next-auth/react"

import { useAtom } from "jotai"

import { Workflow } from "@prisma/client"

import { api } from "@/server/client"
import { plugOrderAtom, plugsMapAtom } from "@/state/plugs"

type PlugsMap = Record<string, Workflow>

export const useSubscriptions = () => {
	const session = useSession()
	const [plugsMap, setPlugsMap] = useAtom(plugsMapAtom)
	const [order, setOrder] = useAtom(plugOrderAtom)

	api.plugs.onAdd.useSubscription(undefined, {
		enabled: Boolean(session.data),
		onData: (data: Workflow) => {
			if (!plugsMap[data.id]) {
				setPlugsMap((prev: PlugsMap) => ({ ...prev, [data.id]: data }))
				setOrder((prev: string[]) => [data.id, ...prev])
			}
		}
	})

	api.plugs.onEdit.useSubscription(undefined, {
		enabled: Boolean(session.data),
		onData: (data: Workflow) => {
			const existing = plugsMap[data.id]
			if (existing && existing.updatedAt < data.updatedAt) {
				setPlugsMap((prev: PlugsMap) => ({
					...prev,
					[data.id]: { ...prev[data.id], ...data }
				}))
			}
		}
	})

	api.plugs.onDelete.useSubscription(undefined, {
		enabled: Boolean(session.data),
		onData: (data: Workflow) => {
			setPlugsMap((prev: PlugsMap) => {
				const { [data.id]: _, ...rest } = prev
				return rest
			})
			setOrder((prev: string[]) => prev.filter(id => id !== data.id))
		}
	})
}
