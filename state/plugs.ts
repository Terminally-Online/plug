import { useSession } from "next-auth/react"
import { useMemo } from "react"

import { atom, useAtom, useAtomValue } from "jotai"

import { Workflow } from "@prisma/client"

import { Actions } from "@/lib"
import { tags } from "@/lib/constants"
import { api } from "@/server/client"

import { COLUMNS, useColumnStore } from "./columns"
import { atomWithStorage } from "jotai/utils"

export const plugsAtom = atom<Workflow[]>([])
export const searchAtom = atom("")
export const tagAtom = atom<(typeof tags)[number]>(tags[0])
export const viewedPlugsAtom = atomWithStorage<Set<string>>("plug.viewed", new Set<string>())

export const workflowByIdAtom = atom(get => (id: string) => get(plugsAtom).find(plug => plug.id === id))

export const ACTION_REGEX = /({\d+(?:=>\d+)?})/g

export type WorkflowData = Pick<Workflow, "name" | "color" | "isPrivate">
export type Option = {
	label: string
	value: string | number
	icon: string
}
export type Value = string | Option | undefined | null

const spread = (plugs: Array<Workflow> | undefined, plug: Workflow) => (!plugs ? [plug] : [plug, ...plugs])

const usePlugActions = () => {
	const { columns, handle } = useColumnStore()

	const [plugs, setPlugs] = useAtom(plugsAtom)

	const addMutation = api.plugs.add.useMutation({
		onSuccess: result => {
			if (!plugs?.find(plug => plug.id === result.plug.id)) setPlugs(prev => spread(prev, result.plug))

			if (result.index)
				handle.navigate({
					key: COLUMNS.KEYS.PLUG,
					index: result.index,
					from: result.from,
					item: result.plug.id
				})
			else
				handle.add({
					index: columns[columns.length - 1].index + 1,
					key: COLUMNS.KEYS.PLUG,
					from: result.from,
					item: result.plug.id
				})
		}
	})

	const editMutation = api.plugs.edit.useMutation({
		onSuccess: result =>
			setPlugs(prev => prev.map(p => (p.id === result.id ? { ...p, ...result, updatedAt: new Date() } : p)))
	})

	const deleteMutation = api.plugs.delete.useMutation({
		onSuccess: (_, variables) => setPlugs(prev => prev.filter(plug => plug.id !== variables.plug))
	})

	const forkMutation = api.plugs.fork.useMutation({
		onSuccess: result => {
			if (!plugs?.find(plug => plug.id === result.plug.id)) setPlugs(prev => spread(prev, result.plug))

			handle.navigate({
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
		add: (data?: { index?: number; from?: string }) => addMutation.mutate(data),
		edit: (data: { id: string } & WorkflowData) => editMutation.mutate(data),
		delete: (data: { plug: string; index: number; from?: string | null }) => deleteMutation.mutate(data),
		fork: (data: { plug: string; index: number; from: string }) => forkMutation.mutate(data),
		queue: (data: { workflowId: string; startAt: Date; endAt?: Date; frequency?: number }) =>
			queueMutation.mutate(data)
	}
}

export const usePlugStore = (id?: string) => {
	const session = useSession()
	const { columns } = useColumnStore()

	const [plugs, setPlugs] = useAtom(plugsAtom)
	const [search, setSearch] = useAtom(searchAtom)
	const [tag, setTag] = useAtom(tagAtom)
	const [viewedPlugs, setViewedPlugs] = useAtom(viewedPlugsAtom)

	const ids = (columns?.map(column => column?.item).filter(Boolean) as string[]) || []

	api.plugs.get.useQuery(
		{ ids, viewed: Array.from(viewedPlugs) },
		{
			enabled: Boolean(session.data) && ids.length > 0,
			onSuccess: data => {
				setPlugs(prev => data.map(plug => prev.find(p => p.id === plug.id) ?? plug))
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

	api.plugs.onAdd.useSubscription(undefined, {
		onData: data => setPlugs(prev => spread(prev, data))
	})

	api.plugs.onEdit.useSubscription(undefined, {
		onData: data =>
			setPlugs(prev => prev.map(p => (p.id === data.id && p.updatedAt < data.updatedAt ? { ...p, ...data } : p)))
	})

	api.plugs.onDelete.useSubscription(undefined, {
		onData: data => setPlugs(prev => prev.filter(plug => plug.id !== data.id))
	})

	const plug = plugs.find(p => p.id === id)
	const own = plug && session.data && session.data.address === plug.socketId
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
		search,
		tag,
		handle: {
			search: setSearch,
			tag: setTag,
			plug: usePlugActions(),
			action: {
				edit: (data: { id?: string; actions: string }) => actionMutation.mutate(data)
			}
		}
	}
}

export const usePlugData = (id?: string) => {
	const getWorkflowById = useAtomValue(workflowByIdAtom)
	const plug = id ? getWorkflowById(id) : undefined
	const session = useSession()

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
