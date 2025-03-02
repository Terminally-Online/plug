import { useSession } from "next-auth/react"
import { useCallback } from "react"
import { useMemo } from "react"

import { atom, useAtom, useAtomValue } from "jotai"

import { Plug } from "@prisma/client"

import { Actions } from "@/lib"
import { api } from "@/server/client"

import { COLUMNS, useColumnActions } from "./columns"
import { selectAtom } from "jotai/utils"
import { atomFamily, atomWithStorage } from "jotai/utils"

export type PlugData = Pick<Plug, "name" | "color" | "isPrivate">
export type Option = {
	label: string
	value: string | number
	icon: string
}
export type Value = string | Option | undefined | null

export const ACTION_REGEX = /({\d+(?:=>\d+)?})/g

export const plugsStorageAtom = atomWithStorage<Plug[]>("plug.plugs", [])
export const plugsAtom = atom(
	get => get(plugsStorageAtom),
	(get, set, update: Plug[] | ((prev: Plug[]) => Plug[])) => {
		set(plugsStorageAtom, typeof update === "function" ? update(get(plugsStorageAtom)) : update)
	}
)
const plugIdMapAtom = atom(get => {
	const plugs = get(plugsStorageAtom)
	return new Map(plugs.map(p => [p.id, p]))
})
export const plugByIdAtom = atomFamily((id: string) => atom(get => get(plugIdMapAtom).get(id)))
export const plugIdsAtom = selectAtom(
	plugsAtom,
	plugs => plugs.map(p => p.id),
	(a, b) => {
		if (a.length !== b.length) return false
		return a.every((id, i) => id === b[i])
	}
)
export const viewedPlugsAtom = atomWithStorage<Set<string>>("plug.viewed", new Set<string>())

export const addPlugAtom = atom(null, (_, set, plug: Plug) => {
	set(plugsAtom, prev => {
		const exists = prev.some(p => p.id === plug.id)
		if (exists) {
			return prev.map(p => (p.id === plug.id && p.updatedAt < plug.updatedAt ? { ...p, ...plug } : p))
		}
		return [plug, ...prev]
	})
})
export const editPlugAtom = atom(null, (_, set, update: Partial<Plug> & { id: string }) => {
	set(plugsAtom, prev => prev.map(p => (p.id === update.id ? { ...p, ...update, updatedAt: new Date() } : p)))
})
export const deletePlugAtom = atom(null, (_, set, id: string) => {
	set(plugsAtom, prev => prev.filter(plug => plug.id !== id))
})
export const queuePlugAtom = atom(null, (_, set, { id, frequency }: { id: string; frequency: number }) => {
	set(plugsAtom, prev => prev.map(p => (p.id === id ? { ...p, queuedAt: new Date(), frequency } : p)))
})

export const usePlugActions = () => {
	const { add, frame, navigate } = useColumnActions()
	const [, addPlug] = useAtom(addPlugAtom)
	const [, editPlug] = useAtom(editPlugAtom)
	const [, deletePlug] = useAtom(deletePlugAtom)
	const [, queuePlug] = useAtom(queuePlugAtom)

	const plugs = useAtomValue(plugsAtom)

	const addMutation = api.plugs.add.useMutation({
		onSuccess: data => {
			addPlug(data.plug)

			if (data.index !== undefined)
				navigate({
					key: COLUMNS.KEYS.PLUG,
					index: data.index,
					from: data.from,
					item: data.plug.id
				})
			else
				add({
					key: COLUMNS.KEYS.PLUG,
					from: data.from,
					item: data.plug.id
				})
		}
	})

	const editMutation = api.plugs.edit.useMutation({
		onSuccess: result => editPlug(result)
	})

	const deleteMutation = api.plugs.delete.useMutation({
		onSuccess: (result, variables) => {
			deletePlug(variables.plug)

			navigate({
				key: COLUMNS.KEYS.MY_PLUGS,
				index: result.index
			})

			frame()
		}
	})

	const forkMutation = api.plugs.fork.useMutation({
		onSuccess: result => {
			if (!plugs.find(plug => plug.id === result.plug.id)) {
				addPlug(result.plug)
			}

			navigate({
				key: COLUMNS.KEYS.PLUG,
				index: result.index,
				from: result.from,
				item: result.plug.id
			})
		}
	})

	const queueMutation = api.plugs.activity.queue.useMutation({
		onSuccess: result => queuePlug({ id: result.id, frequency: result.frequency })
	})

	return useMemo(
		() => ({
			add: (...params: Parameters<typeof addMutation.mutate>) => addMutation.mutate(...params),
			edit: (...params: Parameters<typeof editMutation.mutate>) => editMutation.mutate(...params),
			delete: (...params: Parameters<typeof deleteMutation.mutate>) => deleteMutation.mutate(...params),
			fork: (...params: Parameters<typeof forkMutation.mutate>) => forkMutation.mutate(...params),
			queue: (...params: Parameters<typeof queueMutation.mutate>) => queueMutation.mutate(...params)
		}),
		[addMutation, editMutation, deleteMutation, forkMutation, queueMutation]
	)
}

export const usePlugSubscriptions = ({ enabled }: { enabled: boolean }) => {
	const [, addPlug] = useAtom(addPlugAtom)
	const [, editPlug] = useAtom(editPlugAtom)
	const [, deletePlug] = useAtom(deletePlugAtom)
	const plugs = useAtomValue(plugsAtom)

	const plugIdSet = useMemo(() => new Set(plugs.map(p => p.id)), [plugs])

	api.plugs.onAdd.useSubscription(undefined, {
		enabled,
		onData: data => {
			// Only add if the plug is new or it's been updated
			if (
				!plugIdSet.has(data.id) ||
				(plugs.find(p => p.id === data.id)?.updatedAt || 0) < (data.updatedAt || 0)
			) {
				addPlug(data)
			}
		}
	})

	api.plugs.onEdit.useSubscription(undefined, {
		enabled,
		onData: data => {
			// Only update if plug exists and has been updated
			const existingPlug = plugs.find(p => p.id === data.id)
			if (
				existingPlug &&
				(!existingPlug.updatedAt || !data.updatedAt || existingPlug.updatedAt < data.updatedAt)
			) {
				editPlug(data)
			}
		}
	})

	api.plugs.onDelete.useSubscription(undefined, {
		enabled,
		onData: data => {
			// Only delete if the plug exists
			if (plugIdSet.has(data.id)) {
				deletePlug(data.id)
			}
		}
	})
}

// Helper function to safely use plugByIdAtom without conditional hook calls
const useSafePlugById = (id: string | undefined) => {
	// Create a dummy ID that will never match any real plug if id is undefined
	const safeId = id || "___nonexistent___"
	const plug = useAtomValue(plugByIdAtom(safeId))
	// Only return the plug if id is defined
	return id ? plug : undefined
}

export const usePlugStore = (
	id?: string,
	action?: { protocol: string; action: string; search: Record<number, string | undefined> }
) => {
	const session = useSession()
	const plugIds = useAtomValue(plugIdsAtom)
	const [, editPlug] = useAtom(editPlugAtom)

	const plug = useSafePlugById(id)

	const { data: solverActions } = api.solver.actions.schemas.useQuery(
		{
			chainId: 8453,
			protocol: action?.protocol,
			action: action?.action,
			search: Object.entries(action?.search ?? {}).map(([key, value]) => `search[${key}]=${value}`)
		},
		{ enabled: Boolean(action?.protocol ?? action?.action) ?? false, placeholderData: prev => prev }
	)

	const own = (plug && session.data && session.data.address === plug.socketId) || false

	// Memoize parsed actions to prevent unnecessary recalculation
	const actions: Actions = useMemo(() => {
		if (!plug) return []
		try {
			return JSON.parse(plug.actions)
		} catch {
			return []
		}
	}, [plug])

	const actionMutation = api.plugs.action.edit.useMutation({
		onSuccess: result => editPlug(result)
	})

	// Use a stable reference for the edit action
	const editAction = useCallback(
		(params: Parameters<typeof actionMutation.mutate>[0]) => actionMutation.mutate(params),
		[actionMutation]
	)

	const plugActions = usePlugActions()

	// Create a stable object reference
	return useMemo(
		() => ({
			plugIds,
			plug,
			own,
			actions,
			handle: {
				plug: plugActions,
				action: {
					edit: editAction
				}
			},
			solver: {
				actions: solverActions
			}
		}),
		[plugIds, plug, own, actions, solverActions, editAction, plugActions]
	)
}
