import { Session } from "next-auth"
import { useSession } from "next-auth/react"
import { ContextType, createContext, FC, PropsWithChildren, useContext, useMemo, useState } from "react"

import { Workflow } from "@prisma/client"

import { Actions } from "@/lib"
import { categories, tags } from "@/lib/constants"
import { api } from "@/server/client"
import { COLUMN_KEYS, useColumns } from "@/state"

const spread = (plugs: Array<Workflow> | undefined, plug: Workflow) => (!plugs ? [plug] : [plug, ...plugs])

export const ACTION_REGEX = /({\d+(?:=>\d+)?})/g

type WorkflowData = Pick<Workflow, "name" | "color" | "isPrivate">

export type Option = {
	icon: JSX.Element | undefined
	label: string
	value: string | number
	imagePath?: string
}

export type Value = string | Option | undefined | null

export const PlugContext = createContext<{
	plugs: Array<Workflow>

	search: string
	tag: (typeof tags)[number]
	handle: {
		search: (data: string) => void
		tag: (data: (typeof tags)[number]) => void
		plug: {
			add: (data?: { index?: number; from?: string }) => void
			edit: (data: { id: string } & WorkflowData) => void
			delete: (data: { plug: string; index: number; from?: string | null }) => void
			fork: (data: { plug: string; index: number; from: string }) => void
			queue: (data: { workflowId: string; startAt: Date; endAt?: Date; frequency?: number }) => void
		}
		action: {
			edit: (data: { id?: string; actions: string }) => void
		}
	}
}>({
	plugs: [],
	search: "",
	tag: tags[0],
	handle: {
		search: () => {},
		tag: () => {},
		plug: {
			add: () => {},
			edit: () => {},
			delete: () => {},
			fork: () => {},
			queue: () => {}
		},
		action: {
			edit: () => {}
		}
	}
})

export const PlugProvider: FC<PropsWithChildren<{ session: Session | null }>> = ({ session, children }) => {
	const { columns, add, navigate } = useColumns()

	const [search, handleSearch] = useState("")
	const [tag, handleTag] = useState<(typeof tags)[number]>(tags[0])

	const ids =
		(columns &&
			(columns
				.map(column => (column.item === null ? undefined : column.item))
				.filter(column => Boolean(column)) as Array<string>)) ||
		[]

	const { data: apiPlugs } = api.plugs.get.useQuery(ids, {
		enabled: session && ids.length > 0 ? true : false,
		onSuccess: data => setPlugs(prev => (prev ? data.map(plug => prev.find(p => p.id === plug.id) ?? plug) : data))
	})

	const [plugs, setPlugs] = useState<ContextType<typeof PlugContext>["plugs"]>(apiPlugs || [])

	const handleCreate = (
		data: Parameters<NonNullable<NonNullable<Parameters<typeof api.plugs.add.useMutation>[0]>["onSuccess"]>>[0],
		redirect = false
	) => {
		if (!plugs?.find(plug => plug.id === data.plug.id)) setPlugs(prev => spread(prev, data.plug))

		if (redirect)
			if (data.index === 0)
				add({
					index: data.index,
					key: COLUMN_KEYS.PLUG,
					from: data.from,
					item: data.plug.id
				})
			else if (data.index)
				navigate({
					key: COLUMN_KEYS.PLUG,
					index: data.index,
					from: data.from,
					item: data.plug.id
				})
	}

	api.plugs.onAdd.useSubscription(undefined, {
		onData: data => handleCreate({ plug: data, index: undefined, from: undefined }, false)
	})

	api.plugs.onEdit.useSubscription(undefined, {
		onData: data =>
			setPlugs(prev =>
				prev
					? prev.map(p => (p.id === data.id && p.updatedAt < data.updatedAt ? { ...p, ...data } : p))
					: [data]
			)
	})

	api.plugs.onDelete.useSubscription(undefined, {
		onData: (data: Workflow) => setPlugs(prev => (prev ? prev.filter(plug => plug.id !== data.id) : []))
	})

	const handle = {
		plug: {
			add: api.plugs.add.useMutation({
				onSuccess: data => handleCreate(data, true)
			}),
			edit: api.plugs.edit.useMutation({
				onMutate: data => {
					setPlugs(prev =>
						prev.map(p =>
							p.id === data.id
								? {
										...p,
										...data,
										updatedAt: new Date()
									}
								: p
						)
					)
				}
			}),
			delete: api.plugs.delete.useMutation({
				onMutate: data => {
					setPlugs(prev => prev.filter(plug => plug.id !== data.plug))
					navigate({
						index: data.index,
						key: data.from || COLUMN_KEYS.MY_PLUGS
					})
				}
			}),
			fork: api.plugs.fork.useMutation({
				onSuccess: data => handleCreate(data)
			}),
			queue: api.plugs.activity.queue.useMutation({
				onMutate: data => {
					setPlugs(prev =>
						prev.map(p =>
							p.id === data.workflowId
								? {
										...p,
										queuedAt: new Date(),
										frequency: data.frequency ?? p.frequency,
										updatedAt: new Date()
									}
								: p
						)
					)
				},
				onSuccess: data => {
					setPlugs(prev =>
						prev.map(p =>
							p.id === data.id
								? {
										...p,
										...data,
										updatedAt: new Date()
									}
								: p
						)
					)
				}
			})
		},
		action: {
			edit: api.plugs.action.edit.useMutation({
				onMutate: data => {
					setPlugs(previous =>
						previous.map(p =>
							p.id === data.id
								? {
										...p,
										actions: data.actions,
										updatedAt: new Date()
									}
								: p
						)
					)
				}
			})
		}
	}

	return (
		<PlugContext.Provider
			value={{
				plugs,
				search,
				tag,
				handle: {
					search: handleSearch,
					tag: handleTag,
					plug: {
						add: data => handle.plug.add.mutate(data),
						edit: data => handle.plug.edit.mutate(data),
						delete: data => handle.plug.delete.mutate(data),
						fork: data => handle.plug.fork.mutate(data),
						queue: data => handle.plug.queue.mutate(data)
					},
					action: {
						edit: data => handle.action.edit.mutate(data)
					}
				}
			}}
		>
			{children}
		</PlugContext.Provider>
	)
}

export const usePlugs = (id?: string) => {
	const context = useContext(PlugContext)

	const { data: session } = useSession()

	const { plugs } = context

	// Find the plug that is being utilized within the context based on the plug
	// id provided or the item id that is stored in the column. This way, we do not
	// need to drill down two props every time we want to use a plug since it can
	// be found in the context of the columns.
	const plug = useMemo(() => plugs?.find(plug => plug.id === id) || undefined, [plugs, id])

	const own = plug && session && session.address === plug.socketId

	const actions: Actions = useMemo(() => (plug ? JSON.parse(plug.actions) : []), [plug])

	return {
		...context,
		plug,
		own,
		actions
	}
}
