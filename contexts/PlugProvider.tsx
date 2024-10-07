import { Session } from "next-auth"
import { useSession } from "next-auth/react"
import { ContextType, createContext, FC, PropsWithChildren, useContext, useEffect, useMemo, useState } from "react"

import { Workflow } from "@/server/api/routers/plug"
import { api } from "@/server/client"

import { categories, actions as staticActions, tags, VIEW_KEYS } from "@/lib/constants"
import { useColumns } from "@/state"

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
			fork: () => {}
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

	const { data: apiPlugs } = api.plug.get.useQuery(ids, {
		enabled: session && ids.length > 0 ? true : false
	})

	const [plugs, setPlugs] = useState<ContextType<typeof PlugContext>["plugs"]>(apiPlugs || [])

	const handleCreate = (
		data: Parameters<NonNullable<NonNullable<Parameters<typeof api.plug.add.useMutation>[0]>["onSuccess"]>>[0],
		redirect = false
	) => {
		if (!plugs?.find(plug => plug.id === data.plug.id)) setPlugs(prev => spread(prev, data.plug))

		if (redirect)
			if (data.index === 0)
				add({
					index: data.index,
					key: VIEW_KEYS.PLUG,
					from: data.from,
					item: data.plug.id
				})
			else if (data.index)
				navigate({
					key: VIEW_KEYS.PLUG,
					index: data.index,
					from: data.from,
					item: data.plug.id
				})
	}

	api.plug.onAdd.useSubscription(undefined, {
		onData: data => handleCreate({ plug: data, index: undefined, from: undefined }, false)
	})

	api.plug.onEdit.useSubscription(undefined, {
		onData: data =>
			setPlugs(prev =>
				prev
					? prev.map(p => (p.id === data.id && p.updatedAt < data.updatedAt ? { ...p, ...data } : p))
					: [data]
			)
	})

	api.plug.onDelete.useSubscription(undefined, {
		onData: (data: Workflow) => setPlugs(prev => (prev ? prev.filter(plug => plug.id !== data.id) : []))
	})

	const handle = {
		plug: {
			add: api.plug.add.useMutation({
				onSuccess: data => handleCreate(data, true)
			}),
			edit: api.plug.edit.useMutation({
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
			delete: api.plug.delete.useMutation({
				onMutate: data => {
					setPlugs(prev => prev.filter(plug => plug.id !== data.plug))
					navigate({
						index: data.index,
						key: data.from || VIEW_KEYS.MY_PLUGS
					})
				}
			}),
			fork: api.plug.fork.useMutation({
				onSuccess: data => handleCreate(data)
			})
		},
		action: {
			edit: api.plug.action.edit.useMutation({
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

	useEffect(() => {
		if (!apiPlugs) return

		setPlugs(prev => (prev ? apiPlugs.map(plug => prev.find(p => p.id === plug.id) ?? plug) : apiPlugs))
	}, [apiPlugs])

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
						fork: data => handle.plug.fork.mutate(data)
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

	const actions: Array<{
		categoryName: keyof typeof categories
		actionName: keyof (typeof staticActions)[keyof typeof categories]
		values: Array<Value>
	}> = useMemo(() => (plug ? JSON.parse(plug.actions) : []), [plug])

	const chains = useMemo(() => {
		if (!actions) return []

		const set = actions
			.map(action => new Set(categories[action.categoryName].chains))
			// @ts-ignore -- Don't feel like properly typing this right now.
			.reduce((acc, curr) => {
				if (acc === null) return curr

				return new Set([...acc].filter(chain => curr.has(chain)))
			}, null)

		return set ? Array.from(set) : []
	}, [actions])

	// Split all of the sentence fragments into an appropriate array based on the
	// regex shape that enables the f-string like syntax.
	const fragments = useMemo(() => {
		return actions.map(action => {
			const staticAction = staticActions[action.categoryName][action.actionName]

			return staticAction ? (staticAction["sentence"].split(ACTION_REGEX) as string[]) : []
		})
	}, [actions])

	const dynamic = useMemo(() => {
		return fragments.map(sentence => sentence.filter(fragment => fragment.match(ACTION_REGEX)))
	}, [fragments])

	return {
		...context,
		plug,
		own,
		actions,
		chains,
		fragments,
		dynamic
	}
}
