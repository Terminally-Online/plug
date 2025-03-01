import { useSession } from "next-auth/react"
import { useMemo } from "react"

import { atom, useAtom, useAtomValue, useSetAtom } from "jotai"

import { Plug } from "@prisma/client"

import { Actions } from "@/lib"
import { api } from "@/server/client"

import { COLUMNS, primaryColumnsAtom, useColumnActions } from "./columns"
import { atomWithStorage } from "jotai/utils"
import { useResponse } from "@/lib/hooks/useResponse"

export const plugsAtom = atom<Plug[]>([])
export const viewedPlugsAtom = atomWithStorage<Set<string>>("plug.viewed", new Set<string>())

export const workflowByIdAtom = atom(get => (id: string) => get(plugsAtom).find(plug => plug.id === id))

export const ACTION_REGEX = /({\d+(?:=>\d+)?})/g

export type PlugData = Pick<Plug, "name" | "color" | "isPrivate">
export type Option = {
	label: string
	value: string | number
	icon: string
}
export type Value = string | Option | undefined | null

export const spreadPlugs = (plugs: Array<Plug> | undefined, plug: Plug) => (!plugs ? [plug] : [plug, ...plugs])

const usePlugActions = () => {
	const columns = useAtomValue(primaryColumnsAtom)
	const { add, frame, navigate } = useColumnActions()

	const [plugs, setPlugs] = useAtom(plugsAtom)

	const addMutation = api.plugs.add.useMutation({
		onSuccess: data => {
			if (data.index !== undefined)
				navigate({
					key: COLUMNS.KEYS.PLUG,
					index: data.index,
					from: data.from,
					item: data.plug.id
				})
			else
				add({
					index: columns[columns.length - 1].index + 1,
					key: COLUMNS.KEYS.PLUG,
					from: data.from,
					item: data.plug.id
				})
		}
	})

	const editMutation = api.plugs.edit.useMutation({
		onSuccess: result =>
			setPlugs(prev => prev.map(p => (p.id === result.id ? { ...p, ...result, updatedAt: new Date() } : p)))
	})

	const deleteMutation = api.plugs.delete.useMutation({
		onSuccess: (result, variables) => {
			setPlugs(prev => prev.filter(plug => plug.id !== variables.plug))

			navigate({
				key: COLUMNS.KEYS.MY_PLUGS,
				index: result.index
			})

			frame()
		}
	})

	const forkMutation = api.plugs.fork.useMutation({
		onSuccess: result => {
			if (!plugs?.find(plug => plug.id === result.plug.id)) setPlugs(prev => spreadPlugs(prev, result.plug))

			navigate({
				key: COLUMNS.KEYS.PLUG,
				index: result.index,
				from: result.from,
				item: result.plug.id
			})
		}
	})

	const queueMutation = api.plugs.activity.queue.useMutation({
		onSuccess: result =>
			setPlugs(prev =>
				prev.map(p => (p.id === result.id ? { ...p, queuedAt: new Date(), frequency: result.frequency } : p))
			)
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

	const setPlugs = useSetAtom(plugsAtom)

	api.plugs.onAdd.useSubscription(undefined, {
		enabled: Boolean(session.data),
		onData: data =>
			setPlugs(prev => {
				const exists = prev.some(p => p.id === data.id)
				if (!exists) return [...prev, data]
				return prev.map(p => (p.id === data.id && p.updatedAt < data.updatedAt ? { ...p, ...data } : p))
			})
	})

	api.plugs.onEdit.useSubscription(undefined, {
		enabled: Boolean(session.data),
		onData: data =>
			setPlugs(prev => {
				const exists = prev.some(p => p.id === data.id)
				if (!exists) return [...prev, data]
				return prev.map(p => (p.id === data.id && p.updatedAt < data.updatedAt ? { ...p, ...data } : p))
			})
	})

	api.plugs.onDelete.useSubscription(undefined, {
		enabled: Boolean(session.data),
		onData: data => setPlugs(prev => prev.filter(plug => plug.id !== data.id))
	})
}

export const usePlugStore = (
	id?: string,
	action?: { protocol: string; action: string; search: Record<number, string | undefined> }
) => {
	const session = useSession()

	// const columns = useAtomValue(primaryColumnsAtom)

	const [plugs, setPlugs] = useAtom(plugsAtom)
	const [viewedPlugs, setViewedPlugs] = useAtom(viewedPlugsAtom)

	// const ids = (columns?.map(column => column?.item).filter(Boolean) as string[]) || []

	const { data: solverActions } = api.solver.actions.schemas.useQuery(
		{
			chainId: 8453,
			protocol: action?.protocol,
			action: action?.action,
			search: Object.entries(action?.search ?? {}).map(([key, value]) => `search[${key}]=${value}`)
		},
		{ enabled: Boolean(action?.protocol ?? action?.action) ?? false, placeholderData: (prev) => prev }
	)

	useResponse(() => api.plugs.all.useQuery(
		{ target: "mine" },
		{ enabled: Boolean(session.data) }
	), {
		onSuccess: data =>
			setPlugs(prev => {
				const uniqueData = data.filter(d => !prev.some(p => p.id === d.id))

				return [...prev, ...uniqueData]
			})
	})

	// useResponse(() => api.plugs.get.useQuery(
	// 	{ ids, viewed: Array.from(viewedPlugs) },
	// 	{ enabled: Boolean(session.data) && ids.length > 0 }
	// ), {
	// 	onSuccess: data => {
	// 		setPlugs(prev => {
	// 			const uniqueData = data.filter(d => !prev.some(p => p.id === d.id))
	// 			return [...prev, ...uniqueData]
	// 		})
	// 		setViewedPlugs(prev => {
	// 			const newSet = new Set([...Array.from(prev)].slice(-49))
	// 			data.forEach(plug => newSet.add(plug.id))
	// 			if (newSet.size > 50) {
	// 				const entries = Array.from(newSet)
	// 				newSet.clear()
	// 				entries.slice(-50).forEach(id => newSet.add(id))
	// 			}
	// 			return newSet
	// 		})
	// 	}
	// })

	const plug = plugs.find(p => p.id === id)
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
		onSuccess: result => setPlugs(prev => prev.map(p => (p.id === result.id ? { ...p, ...result } : p)))
	})

	return {
		plugs,
		plug,
		own,
		actions,
		handle: {
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
