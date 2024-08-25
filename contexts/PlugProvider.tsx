import { ContextType, createContext, FC, PropsWithChildren, useContext, useEffect, useMemo, useState } from "react"

import { useSession } from "next-auth/react"

import { useFrame, useSockets } from "@/contexts"
import { categories, actions as staticActions, tags, VIEW_KEYS } from "@/lib/constants"
import { Workflow } from "@/server/api/routers/plug"
import { api } from "@/server/client"

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
			add: (data?: string) => void
			edit: (data: { id: string } & WorkflowData) => void
			delete: (data: { id: string; from?: string }) => void
			fork: (data: { id: string; from?: string }) => void
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

// TODO: Pick back up here in the morning. We are removing the single slot id of the
//       plug that is being selected so that we can have multiple visible and accessible
//       throughout a set of columns. Right now, frames retrieve the single id that is
//       stored and when a parameter is provided to the hook it selects it. Instead,
//       we should automatically retrieve the plug based on all of the `item` values
//       that are stored in the columns.
// TODO: Remove search and tag from here because we may have multiple contexts open now.

export const PlugProvider: FC<PropsWithChildren> = ({ children }) => {
	const { socket, handle: handleSocket } = useSockets()

	const [search, handleSearch] = useState("")
	const [tag, handleTag] = useState<(typeof tags)[number]>(tags[0])

	const ids =
		(socket?.columns
			.map(column => (column.item === null ? undefined : column.item))
			.filter(column => Boolean(column)) as Array<string>) || []

	const { data: apiPlugs } = api.plug.get.useQuery(ids, {
		enabled: ids.length > 0 ? true : false
	})

	const [plugs, setPlugs] = useState<ContextType<typeof PlugContext>["plugs"]>(apiPlugs || [])

	const handleCreate = (
		data: Parameters<NonNullable<NonNullable<Parameters<typeof api.plug.add.useMutation>[0]>["onSuccess"]>>[0],
		redirect = true
	) => {
		if (!plugs?.find(plug => plug.id === data.plug.id)) setPlugs(prev => spread(prev, data.plug))

		if (redirect)
			handleSocket.columns.add({
				key: VIEW_KEYS.PLUG,
				index: 0,
				item: data.plug.id
			})
	}

	api.plug.onAdd.useSubscription(undefined, {
		onData: data => handleCreate({ plug: data, from: "" }, false)
	})

	api.plug.onEdit.useSubscription(undefined, {
		onData: data =>
			setPlugs(prev => (prev ? prev.map(p => (p.id === data.id && p.updatedAt < data.updatedAt ? { ...p, ...data } : p)) : [data]))
	})

	api.plug.onDelete.useSubscription(undefined, {
		onData: (data: Workflow) => {
			// If the user was viewing the deleted plug, show a confirmation frame to signal that.
			// if (id === data.id) {
			// handlePage({ key: "home" })
			// handleFrame()
			// }

			setPlugs(prev => (prev ? prev.filter(plug => plug.id !== data.id) : []))
		}
	})

	const handle = {
		plug: {
			add: api.plug.add.useMutation({
				onSuccess: data => handleCreate(data)
			}),
			edit: api.plug.edit.useMutation({
				onMutate: data => {
					const previous = plugs ?? []

					setPlugs(
						previous.map(p =>
							p.id === data.id
								? {
										...p,
										...data,
										updatedAt: new Date()
									}
								: p
						)
					)

					return previous
				}
				// onError: (_, __, context) => setPlugs(context)
			}),
			delete: api.plug.delete.useMutation({
				onMutate: data => {
					const previous = plugs ?? []

					setPlugs(previous.filter(plug => plug.id !== data.id))

					// handlePage({ key: data.from ?? "/app/plugs/" })

					return previous
				}
				// onError: (_, __, context) => setPlugs(context)
			}),
			fork: api.plug.fork.useMutation({
				onSuccess: data => handleCreate(data)
			})
		},
		action: {
			edit: api.plug.action.edit.useMutation({
				onMutate: data => {
					const previous = plugs ?? []

					setPlugs(
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

					return previous
				}
				// onError: (_, __, context) => setPlugs(context)
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

export const usePlugs = (id: string) => {
	const context = useContext(PlugContext)

	const { data: session } = useSession()
	const { socket } = useSockets()

	const { plugs } = context

	// Find the plug that is being utilized within the context based on the plug
	// id provided or the item id that is stored in the column. This way, we do not
	// need to drill down two props every time we want to use a plug since it can
	// be found in the context of the columns.
	const plug = useMemo(
		() =>
			plugs?.find(plug => plug.id === id) ||
			plugs?.find(plug => plug.id === socket?.columns.find(column => column.id === id)?.item) ||
			undefined,
		[plugs, id, socket]
	)

	const own = plug && session && session.address === plug.userAddress

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
