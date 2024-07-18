import {
	ContextType,
	createContext,
	FC,
	PropsWithChildren,
	useContext,
	useEffect,
	useMemo,
	useState
} from "react"

import { useRouter } from "next/navigation"

import { useFrame } from "@/contexts"
import {
	categories,
	routes,
	actions as staticActions,
	tags
} from "@/lib/constants"
import { Workflow } from "@/server/api/routers/plug"
import { api } from "@/server/client"

const spread = (plugs: Array<Workflow> | undefined, plug: Workflow) =>
	!plugs ? [plug] : [plug, ...plugs]

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
	plugs?: Array<Workflow>
	id?: string
	plug?: Workflow

	search: string
	tag: (typeof tags)[number]

	handle: {
		select: (data?: string) => void
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
	id: undefined,
	plug: undefined,

	search: "",
	tag: tags[0],

	handle: {
		select: () => {},
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

export const PlugProvider: FC<PropsWithChildren> = ({ children }) => {
	const router = useRouter()

	const { handleFrameVisible } = useFrame()

	const [id, handleId] = useState<ContextType<typeof PlugContext>["id"]>()
	const [search, handleSearch] = useState("")
	const [tag, handleTag] = useState<(typeof tags)[number]>(tags[0])

	const { data: apiPlugs } = api.plug.all.useQuery({ target: "mine" })
	const [plugs, setPlugs] =
		useState<ContextType<typeof PlugContext>["plugs"]>(apiPlugs)
	const { data: apiPlug } = api.plug.get.useQuery(id)

	const plug = useMemo(
		() => plugs && plugs.find(plug => plug.id === id),
		[plugs, id]
	)

	const handleCreate = (
		data: Parameters<
			NonNullable<
				NonNullable<
					Parameters<typeof api.plug.add.useMutation>[0]
				>["onSuccess"]
			>
		>[0],
		redirect = true
	) => {
		if (!plugs?.find(plug => plug.id === data.plug.id))
			setPlugs(prev => spread(prev, data.plug))

		if (redirect)
			router.push(
				`/app/plugs/${data.plug.id}${data.from ? `?from=${data.from}` : ""}`
			)
	}

	api.plug.onAdd.useSubscription(undefined, {
		onData: data => handleCreate({ plug: data, from: "" }, false)
	})

	api.plug.onEdit.useSubscription(undefined, {
		onData: data =>
			setPlugs(prev =>
				prev
					? prev.map(p =>
							p.id === data.id && p.updatedAt < data.updatedAt
								? { ...p, ...data }
								: p
						)
					: [data]
			)
	})

	api.plug.onDelete.useSubscription(undefined, {
		onData: (data: Workflow) => {
			// If the user was viewing the deleted plug, show a confirmation frame to signal that.
			if (id === data.id) {
				handleFrameVisible("deleted")
				router.push(routes.app.index)
			}

			setPlugs(prev =>
				prev ? prev.filter(plug => plug.id !== data.id) : []
			)
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
				},
				onError: (_, __, context) => setPlugs(context)
			}),
			delete: api.plug.delete.useMutation({
				onMutate: data => {
					const previous = plugs ?? []

					setPlugs(previous.filter(plug => plug.id !== data.id))

					router.push(data.from ?? `/app/plugs/`)

					return previous
				},
				onError: (_, __, context) => setPlugs(context)
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
				},
				onError: (_, __, context) => setPlugs(context)
			})
		}
	}

	useEffect(() => {
		if (!apiPlug || !plugs || plugs.some(plug => plug.id === apiPlug.id))
			return

		setPlugs(prev => (prev ? prev.concat(apiPlug) : [apiPlug]))
	}, [apiPlug, plugs])

	return (
		<PlugContext.Provider
			value={{
				plugs,
				id,
				plug,

				search,
				tag,

				handle: {
					select: handleId,
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
	const { frameVisible } = useFrame()

	const context = useContext(PlugContext)

	const { plug } = context

	const [chains, setChains] = useState<string[]>([])

	const actions: Array<{
		categoryName: keyof typeof categories
		actionName: keyof (typeof staticActions)[keyof typeof categories]
		values: Array<Value>
	}> = useMemo(() => (plug ? JSON.parse(plug.actions) : []), [plug])

	const chainsAvailable = useMemo(() => {
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
			const staticAction =
				staticActions[action.categoryName][action.actionName]

			return staticAction
				? (staticAction["sentence"].split(ACTION_REGEX) as string[])
				: []
		})
	}, [actions])

	const dynamic = useMemo(() => {
		return fragments.map(sentence =>
			sentence.filter(fragment => fragment.match(ACTION_REGEX))
		)
	}, [fragments])

	const handleChainSelect = (chain: string) => {
		setChains(prev =>
			prev.includes(chain)
				? prev.filter(c => c !== chain)
				: [...prev, chain]
		)
	}

	useEffect(() => {
		if (id) context.handle.select(id)
	}, [id, context.handle])

	// When there is only one chain available, select it by default. The user
	// will first go to the Socket frame so we are preloading the chain.
	useEffect(() => {
		if (chainsAvailable.length === 1) setChains([chainsAvailable[0]])
	}, [frameVisible, chainsAvailable])

	return {
		...context,
		actions,
		chains,
		chainsAvailable,
		fragments,
		dynamic,
		handle: {
			...context.handle,
			chain: {
				select: handleChainSelect
			}
		}
	}
}
