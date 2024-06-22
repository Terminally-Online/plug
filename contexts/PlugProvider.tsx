import type { FC, PropsWithChildren } from "react"
import { createContext, useContext, useEffect, useMemo, useState } from "react"

import { useRouter } from "next/navigation"

import { Value } from "@/components/app/sentences"
import { tags } from "@/lib/constants"
import { useDebounce } from "@/lib/hooks"
import { Workflow } from "@/server/api/routers/plug"
import { api } from "@/server/client"

type PlugContextProps = {
	// Single plug fields and methods
	plugs?: Array<Workflow>
	id?: string
	plug?: Workflow
	version?: number

	search: string
	tag: (typeof tags)[number]

	// Action methods for a single plug
	actions: {
		plug: {
			handleSelect: (id?: string) => void
			handleSearch: (search: string) => void
			handleTag: (tag: (typeof tags)[number]) => void
			handleAdd: (from: string) => void
			handleEdit: ({
				name,
				color,
				isPrivate
			}: Pick<Workflow, "name" | "color" | "isPrivate">) => void
			handleDelete: ({ id, from }: { id: string; from?: string }) => void
			handleFork: ({ id, from }: { id: string; from?: string }) => void
			handleVersion: (direction: 1 | -1) => void
		}
	}
}

const spread = (plugs: Array<Workflow> | undefined, plug: Workflow) =>
	!plugs ? [plug] : [plug, ...plugs]

const getActions = (plug: Workflow | undefined, version: number) =>
	plug && plug.versions[plug.versions.length - version]
		? plug.versions[plug.versions.length - version].actions
		: []

export const ACTION_REGEX = /({\d+(?:=>\d+)?})/g

export const PlugContext = createContext<PlugContextProps>({
	plugs: [],
	id: undefined,
	plug: undefined,
	version: undefined,

	search: "",
	tag: tags[0],

	actions: {
		plug: {
			handleSelect: () => {},
			handleSearch: () => {},
			handleTag: () => {},
			handleAdd: () => {},
			handleEdit: () => {},
			handleDelete: () => {},
			handleFork: () => {},
			handleVersion: () => {}
		}
	}
})

export const PlugProvider: FC<PropsWithChildren> = ({ children }) => {
	const router = useRouter()

	const [id, setId] = useState<PlugContextProps["id"]>()
	const [version, setVersion] = useState(0)
	const [search, debouncedSearch, handleSearch] = useDebounce("")
	const [tag, setTag] = useState<(typeof tags)[number]>(tags[0])

	const { data: apiPlugs } = api.plug.all.useQuery()

	const [plugs, setPlugs] = useState<PlugContextProps["plugs"]>(apiPlugs)

	const plug = useMemo(
		() => plugs && plugs.find(plug => plug.id === id),
		[plugs, id]
	)

	const actions = useMemo(() => getActions(plug, version), [plug, version])

	const fragments = useMemo(
		() =>
			!actions
				? []
				: actions.map(action => {
						const parsed = action.data
							? JSON.parse(action.data)
							: undefined

						return !parsed || !parsed["sentence"]
							? []
							: (parsed["sentence"].split(
									ACTION_REGEX
								) as string[])
					}),
		[actions]
	)

	// Filter down to only the dynamic fragments that match the regex pattern
	// so that we can use the carried index value to update the correct indexes
	// when one is updated by the user.
	const dynamic = useMemo(
		() =>
			fragments.map(fragment =>
				fragment.filter(fragment => fragment.match(ACTION_REGEX))
			),
		[fragments]
	)

	const [values, setValues] = useState<Array<Array<Value>>>(
		dynamic.map(fragment => fragment.map(() => undefined))
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

	const handleAdd = api.plug.add.useMutation({
		onSuccess: data => handleCreate(data)
	})
	const handleFork = api.plug.fork.useMutation({
		onSuccess: data => handleCreate(data)
	})
	api.plug.onAdd.useSubscription(undefined, {
		onData: data => handleCreate({ plug: data, from: "" }, false)
	})

	const handleEdit = api.plug.edit.useMutation({
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

	const handleDelete = api.plug.delete.useMutation({
		onMutate: data => {
			const previous = plugs ?? []

			setPlugs(previous.filter(plug => plug.id !== data.id))

			router.push(data.from ?? `/app/plugs/`)

			return previous
		},
		onError: (_, __, context) => setPlugs(context)
	})
	api.plug.onDelete.useSubscription(undefined, {
		onData: (data: Workflow) =>
			setPlugs(prev =>
				prev ? prev.filter(plug => plug.id !== data.id) : []
			)
	})

	useEffect(() => {
		// When a new plug is selected, we want to set it to the latest version
		// and when it is updated (in this client or another client), we automatically
		// update the local state to use the latest version.
		if (plug) setVersion(plug.versions[0]?.version ?? 1)
	}, [plug])

	const handleAddAction = api.plug.action.add.useMutation()
	const handleEditAction = api.plug.action.edit.useMutation()
	const handleRemoveAction = api.plug.action.remove.useMutation()

	return (
		<PlugContext.Provider
			value={{
				plugs,
				id,
				plug,
				version,

				search,
				tag,

				actions: {
					plug: {
						handleSelect: setId,
						handleSearch,
						handleTag: setTag,
						handleAdd: (from: string) => handleAdd.mutate(from),
						handleEdit: ({ name, color, isPrivate }) =>
							handleEdit.mutate({
								id,
								name,
								color,
								isPrivate
							}),
						handleDelete: ({ id, from }) =>
							handleDelete.mutate({ id, from }),
						handleFork: ({ id, from }) =>
							handleFork.mutate({ id, from }),
						handleVersion: (direction: 1 | -1) =>
							setVersion(prev => prev + direction)
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

	useEffect(() => {
		if (id) context.actions.plug.handleSelect(id)
	}, [context.actions.plug, id])

	return context
}
