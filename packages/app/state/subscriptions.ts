import { useSession } from "next-auth/react"

import { useAtom } from "jotai"

import { api } from "@/server/client"
import { plugsAtom, spreadPlugs } from "@/state"

export const useSubscriptions = () => {
	const session = useSession()
	const [, setPlugs] = useAtom(plugsAtom)

	api.plugs.onAdd.useSubscription(undefined, {
		enabled: Boolean(session.data),
		onData: data => setPlugs(prev => spreadPlugs(prev, data))
	})

	api.plugs.onEdit.useSubscription(undefined, {
		enabled: Boolean(session.data),
		onData: data =>
			setPlugs(prev => prev.map(p => (p.id === data.id && p.updatedAt < data.updatedAt ? { ...p, ...data } : p)))
	})

	api.plugs.onDelete.useSubscription(undefined, {
		enabled: Boolean(session.data),
		onData: data => setPlugs(prev => prev.filter(plug => plug.id !== data.id))
	})
}
