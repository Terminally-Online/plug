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

import {
	actionCategories,
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
			add: (data: string) => void
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

	const { data: apiPlugs } = api.plug.all.useQuery()

	const [plugs, setPlugs] =
		useState<ContextType<typeof PlugContext>["plugs"]>(apiPlugs)
	const [id, handleId] = useState<ContextType<typeof PlugContext>["id"]>()

	const [search, handleSearch] = useState("")
	const [tag, handleTag] = useState<(typeof tags)[number]>(tags[0])

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
		onData: (data: Workflow) =>
			setPlugs(prev =>
				prev ? prev.filter(plug => plug.id !== data.id) : []
			)
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
	const context = useContext(PlugContext)

	const { plug } = context

	const actions: Array<{
		categoryName: keyof typeof actionCategories
		actionName: keyof (typeof staticActions)[keyof typeof actionCategories]
		values: Array<Value>
	}> = useMemo(() => (plug ? JSON.parse(plug.actions) : []), [plug])

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

	useEffect(() => {
		if (id) context.handle.select(id)
	}, [context.handle, id])

	return {
		...context,
		actions,
		fragments,
		dynamic
	}
}
