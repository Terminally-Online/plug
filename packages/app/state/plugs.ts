import { useSession } from "next-auth/react"
import { useMemo } from "react"

import { atom, useAtom, useAtomValue } from "jotai"

import { Workflow } from "@prisma/client"

import { Actions } from "@/lib"
import { Column } from "@/lib"
import { tags } from "@/lib/constants"
import { api } from "@/server/client"

import { columnAtomsMapAtom, COLUMNS, useColumnList, useColumnStore } from "./columns"
import { atomWithStorage } from "jotai/utils"

// Core storage atoms
export const plugsMapAtom = atomWithStorage<Record<string, Workflow>>("plug.map", {})
export const plugOrderAtom = atomWithStorage<string[]>("plug.order", [])
export const viewedPlugsAtom = atomWithStorage<Set<string>>("plug.viewed", new Set<string>())

// UI state atoms
export const searchAtom = atom("")
export const tagAtom = atom<(typeof tags)[number]>(tags[0])

// Derived atoms
export const plugsAtom = atom(get => {
	const order = get(plugOrderAtom)
	const plugsMap = get(plugsMapAtom)
	return order.map(id => plugsMap[id]).filter(Boolean)
})

// Filtered plugs atom
export const filteredPlugsAtom = atom(get => {
	const plugs = get(plugsAtom)
	const search = get(searchAtom).toLowerCase()
	const tag = get(tagAtom)

	return plugs.filter(plug => {
		if (search && !plug.name.toLowerCase().includes(search)) return false
		if (tag !== tags[0] && !plug.tags?.includes(tag)) return false
		return true
	})
})

// Types
export type WorkflowData = Pick<Workflow, "name" | "color" | "isPrivate">
export type Option = {
	label: string
	value: string | number
	icon: string
}
export type Value = string | Option | undefined | null

export const ACTION_REGEX = /({\d+(?:=>\d+)?})/g

// Hooks
const usePlugActions = () => {
	const { column, handle } = useColumnStore()
	const [plugsMap, setPlugsMap] = useAtom(plugsMapAtom)
	const [order, setOrder] = useAtom(plugOrderAtom)

	const addMutation = api.plugs.add.useMutation({
		onSuccess: result => {
			const { plug } = result
			if (!plugsMap[plug.id]) {
				setPlugsMap(prev => ({ ...prev, [plug.id]: plug }))
				setOrder(prev => [plug.id, ...prev])
			}

			if (result.index) {
				handle.navigate({
					key: COLUMNS.KEYS.PLUG,
					index: result.index,
					from: result.from,
					item: plug.id
				})
			} else {
				handle.add({
					key: COLUMNS.KEYS.PLUG,
					from: result.from,
					item: plug.id
				})
			}
		}
	})

	const editMutation = api.plugs.edit.useMutation({
		onSuccess: result => {
			setPlugsMap(prev => ({
				...prev,
				[result.id]: { ...prev[result.id], ...result, updatedAt: new Date() }
			}))
		}
	})

	const deleteMutation = api.plugs.delete.useMutation({
		onSuccess: (_, variables) => {
			setPlugsMap(prev => {
				const { [variables.plug]: _, ...rest } = prev
				return rest
			})
			setOrder(prev => prev.filter(id => id !== variables.plug))
		}
	})

	const forkMutation = api.plugs.fork.useMutation({
		onSuccess: result => {
			const { plug } = result
			if (!plugsMap[plug.id]) {
				setPlugsMap(prev => ({ ...prev, [plug.id]: plug }))
				setOrder(prev => [plug.id, ...prev])
			}

			handle.navigate({
				key: COLUMNS.KEYS.PLUG,
				index: result.index,
				from: result.from,
				item: plug.id
			})
		}
	})

	const queueMutation = api.plugs.activity.queue.useMutation({
		onSuccess: result => {
			setPlugsMap(prev => ({
				...prev,
				[result.id]: {
					...prev[result.id],
					queuedAt: new Date(),
					frequency: result.frequency
				}
			}))
		}
	})

	return {
		add: addMutation.mutate,
		edit: editMutation.mutate,
		delete: deleteMutation.mutate,
		fork: forkMutation.mutate,
		queue: queueMutation.mutate
	}
}

export const usePlugSubscriptions = () => {
	const session = useSession()
	const [plugsMap, setPlugsMap] = useAtom(plugsMapAtom)
	const [order, setOrder] = useAtom(plugOrderAtom)

	api.plugs.onAdd.useSubscription(undefined, {
		enabled: Boolean(session.data),
		onData: data => {
			if (!plugsMap[data.id]) {
				setPlugsMap(prev => ({ ...prev, [data.id]: data }))
				setOrder(prev => [data.id, ...prev])
			}
		}
	})

	api.plugs.onEdit.useSubscription(undefined, {
		enabled: Boolean(session.data),
		onData: data => {
			const existing = plugsMap[data.id]
			if (existing && existing.updatedAt < data.updatedAt) {
				setPlugsMap(prev => ({
					...prev,
					[data.id]: { ...prev[data.id], ...data }
				}))
			}
		}
	})

	api.plugs.onDelete.useSubscription(undefined, {
		enabled: Boolean(session.data),
		onData: data => {
			setPlugsMap(prev => {
				const { [data.id]: _, ...rest } = prev
				return rest
			})
			setOrder(prev => prev.filter(id => id !== data.id))
		}
	})
}

export const usePlugStore = (id?: string, action?: { protocol: string; action: string }) => {
	const session = useSession()
	const { columnIds } = useColumnList()
	const [columnsMap] = useAtom(columnAtomsMapAtom)
	const [plugsMap, setPlugsMap] = useAtom(plugsMapAtom)
	const [order, setOrder] = useAtom(plugOrderAtom)
	const [search, setSearch] = useAtom(searchAtom)
	const [tag, setTag] = useAtom(tagAtom)
	const [viewedPlugs, setViewedPlugs] = useAtom(viewedPlugsAtom)
	const filteredPlugs = useAtomValue(filteredPlugsAtom)

	// Get all plug IDs from columns
	const columnPlugIds = Object.values(columnsMap)
		.map(column => (column as Column | undefined)?.item)
		.filter((item): item is string => Boolean(item))

	// Include the requested plug ID if it's not in the columns
	const ids = id && !columnPlugIds.includes(id) ? [...columnPlugIds, id] : columnPlugIds

	const { data: solverActions } = api.solver.actions.getSchemas.useQuery(
		{ protocol: action?.protocol, action: action?.action, chainId: 1 },
		{ enabled: Boolean(action) }
	)

	// Query for all plugs including the requested one
	api.plugs.get.useQuery(
		{ ids, viewed: Array.from(viewedPlugs) },
		{
			enabled: Boolean(session.data) && ids.length > 0,
			onSuccess: data => {
				// Update plugs state
				const updates: Record<string, Workflow> = {}
				const newIds: string[] = []

				data.forEach(plug => {
					updates[plug.id] = plug
					if (!order.includes(plug.id)) {
						newIds.push(plug.id)
					}
				})

				if (Object.keys(updates).length > 0) {
					setPlugsMap(prev => ({ ...prev, ...updates }))
				}
				if (newIds.length > 0) {
					setOrder(prev => [...newIds, ...prev])
				}

				// Update viewed plugs
				setViewedPlugs(prev => {
					const newSet = new Set([...Array.from(prev)].slice(-49))
					data.forEach(plug => newSet.add(plug.id))
					if (newSet.size > 50) {
						const entries = Array.from(newSet)
						newSet.clear()
						entries.slice(-50).forEach(id => newSet.add(id))
					}
					return newSet
				})
			}
		}
	)

	const plug = id ? plugsMap[id] : undefined
	const own = (plug && session.data && session.data.address === plug.socketId) || false
	const actions: Actions = useMemo(() => {
		if (!plug) return []
		try {
			return JSON.parse(plug.actions)
		} catch {
			return []
		}
	}, [plug])

	const actionMutation = api.plugs.action.edit.useMutation({
		onSuccess: result => {
			setPlugsMap(prev => ({
				...prev,
				[result.id]: { ...prev[result.id], ...result }
			}))
		}
	})

	return {
		plugs: filteredPlugs,
		plug,
		own,
		actions,
		search,
		tag,
		handle: {
			search: setSearch,
			tag: setTag,
			plug: usePlugActions(),
			action: {
				edit: actionMutation.mutate
			}
		},
		solver: {
			actions: solverActions
		}
	}
}

export const usePlugData = (id?: string) => {
	const [plugsMap] = useAtom(plugsMapAtom)
	const session = useSession()

	const plug = id ? plugsMap[id] : undefined
	const own = plug && session.data && session.data.address === plug.socketId
	const actions: Actions = useMemo(() => {
		if (!plug) return []
		try {
			return JSON.parse(plug.actions)
		} catch {
			return []
		}
	}, [plug])

	return { plug, own, actions }
}
