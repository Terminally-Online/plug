import type { FC, PropsWithChildren } from "react"
import { createContext, useContext, useEffect, useMemo, useState } from "react"

import { useRouter } from "next/navigation"

import { actionCategories, actions, tags } from "@/lib/constants"
import { useDebounce } from "@/lib/hooks"
import { Workflow } from "@/server/api/routers/plug"
import { api } from "@/server/client"

type PlugContextProps = {
	plugs: Array<Workflow> | undefined
	filteredPlugs: Array<Workflow> | undefined
	search: string
	tag: (typeof tags)[number]
	handleSearch: (search: string) => void
	handleTag: (tag: (typeof tags)[number]) => void

	// Single plug fields and methods
	id?: string
	plug?: Workflow
	version: number
	handleAdd: (from: string) => void

	// Action methods for a single plug
	handleSelect: (id: string) => void
	handleEdit: ({
		name,
		color,
		isPrivate
	}: Pick<Workflow, "name" | "color" | "isPrivate">) => void
	handleFork: ({ id, from }: { id: string; from?: string }) => void
	handleVersionChange: (direction: 1 | -1) => void
	handleAddAction: (
		action: Omit<
			Workflow["versions"][number]["actions"][number],
			"id" | "index" | "workflowId" | "categoryName" | "actionName"
		> & {
			categoryName: keyof typeof actionCategories
			actionName: keyof (typeof actions)[keyof typeof actionCategories]
		}
	) => void
	handleRemoveAction: (
		action: Workflow["versions"][number]["actions"][number]
	) => void
}

export const PlugContext = createContext<PlugContextProps>({
	plugs: [],
	filteredPlugs: [],
	search: "",
	tag: tags[0],
	handleSearch: () => {},
	handleTag: () => {},

	// Single plug fields and methods
	id: undefined,
	plug: undefined,
	version: 0,
	handleAdd: () => {},

	// Action methods for a single plug
	handleSelect: () => {},
	handleEdit: () => {},
	handleFork: () => {},
	handleVersionChange: () => {},
	handleAddAction: () => {},
	handleRemoveAction: () => {}
})

const spread = (plugs: Array<Workflow> | undefined, plug: Workflow) =>
	!plugs ? [plug] : [plug, ...plugs]

const map = (plugs: Array<Workflow> | undefined, plug: Workflow) =>
	!plugs ? [plug] : plugs.map(p => (p.id === plug.id ? plug : p))

export const PlugProvider: FC<PropsWithChildren> = ({ children }) => {
	const router = useRouter()

	const [id, setId] = useState<PlugContextProps["id"]>()
	const [version, setVersion] = useState(0)

	const {
		debounce: handleSearch,
		value: search,
		debounced: debouncedSearch
	} = useDebounce({ initial: "" })

	const [tag, setTag] = useState<(typeof tags)[number]>(tags[0])

	const { data: apiPlugs } = api.plug.all.useQuery()
	const { data: apiFilteredPlugs } = api.plug.all.useQuery({
		search: debouncedSearch,
		tag
	})

	const [plugs, setPlugs] = useState<PlugContextProps["plugs"]>(apiPlugs)

	const plug = useMemo(
		() => plugs && plugs.find(plug => plug.id === id),
		[plugs, id]
	)

	const filteredPlugs = useMemo(
		() =>
			apiFilteredPlugs && apiFilteredPlugs.filter(plug => plug.id !== id),
		[apiFilteredPlugs, id]
	)

	const handleAdd = api.plug.add.useMutation({
		onSuccess: data =>
			router.push(`/app/plugs/${data.plug.id}?from=${data.from}`)
	})
	const handleEdit = api.plug.edit.useMutation()
	const handleFork = api.plug.fork.useMutation({
		onSuccess: data =>
			router.push(
				`/app/plugs/${data.plug.id}${data.from ? `?from=${data.from}` : ""}`
			)
	})

	api.plug.onAdd.useSubscription(undefined, {
		onData: (data: Workflow) => setPlugs(prev => spread(prev, data))
	})
	api.plug.onEdit.useSubscription(undefined, {
		onData: (data: Workflow) => setPlugs(prev => map(prev, data))
	})

	const handleAddAction = api.plug.action.add.useMutation()
	const handleRemoveAction = api.plug.action.remove.useMutation()

	useEffect(() => {
		if (plug) setVersion(plug.versions[0]?.version ?? 1)
	}, [plug])

	return (
		<PlugContext.Provider
			value={{
				plugs,
				filteredPlugs,
				search,
				tag,
				handleSearch,
				handleTag: setTag,

				// Single plug fields and methods
				id,
				plug,
				version,
				handleAdd: (from: string) => handleAdd.mutate(from),

				// Action methods for a single plug
				handleSelect: setId,
				handleEdit: ({ name, color, isPrivate }) =>
					handleEdit.mutate({
						id: id || "",
						name,
						color,
						isPrivate
					}),
				handleFork: ({ id, from }) => handleFork.mutate({ id, from }),
				handleVersionChange: (direction: 1 | -1) =>
					setVersion(prev => prev + direction),
				handleAddAction: action =>
					handleAddAction.mutate({
						id: id ?? "",
						version,
						action
					}),
				handleRemoveAction: action =>
					handleRemoveAction.mutate(action.id)
			}}
		>
			{children}
		</PlugContext.Provider>
	)
}

export const usePlugs = () => useContext(PlugContext)
