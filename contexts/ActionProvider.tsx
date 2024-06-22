import type { FC, PropsWithChildren } from "react"
import { createContext, useContext, useEffect, useMemo, useState } from "react"

import { Value } from "@/components/app/sentences"
import { getIndexes } from "@/components/app/sentences/fragments"
import { actionCategories, actions } from "@/lib/constants"
import { Workflow } from "@/server/api/routers/plug"
import { api } from "@/server/client"

import { usePlugs } from "./PlugProvider"

type ActionData = Omit<
	Workflow["versions"][number]["actions"][number],
	| "id"
	| "index"
	| "workflowId"
	| "categoryName"
	| "actionName"
	| "versionId"
	| "createdAt"
	| "updatedAt"
> & {
	categoryName: keyof typeof actionCategories
	actionName: keyof (typeof actions)[keyof typeof actionCategories]
}

type ActionContextProps = {
	actions: Array<Workflow["versions"][number]["actions"][number]>
	handleAdd: (action: ActionData) => void
	handleEdit: (action: Array<ActionData>) => void
	handleRemove: (
		action: Workflow["versions"][number]["actions"][number]
	) => void
}

export const ACTION_REGEX = /({\d+(?:=>\d+)?})/g

export const ActionContext = createContext<ActionContextProps>({
	actions: [],
	handleAdd: () => {},
	handleEdit: () => {},
	handleRemove: () => {}
})

const getActions = (plug?: Workflow, version?: number) =>
	plug && version && plug.versions[plug.versions.length - version]
		? plug.versions[plug.versions.length - version].actions
		: []

export const ActionProvider: FC<PropsWithChildren> = ({ children }) => {
	const { id, plug, version } = usePlugs()

	const [actions, setActions] = useState(getActions(plug, version))

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

	// This loops through the fragments and updates the respective value based
	// on the stringified index value of the fragment. This admittedly is a bit
	// confusing to read on first glance, but it's a way to update child values
	// when a parent value changes.
	//
	// Example:
	// {0}    will update the value at index 0.
	// {0=>1} will update the value at index 1 and set the value to undefined
	//        when the value at index 0 changes which is signalled by the
	//        upper index value.
	const handleValue = (
		sentenceIndex: number,
		fragment: string,
		value: Value
	) => {
		const [, fragmentIndex] = getIndexes(fragment)

		// setValues(prev =>
		// 	prev.map((v, i) =>
		// 		fragmentIndex === getIndexes(dynamicFragments[i])[0]
		// 			? undefined
		// 			: i === fragmentIndex
		// 				? value
		// 				: v
		// 	)
		// )
	}

	// console.log(
	// 	"render",
	// 	plug?.versions[plug.versions.length - version - 1]?.actions
	// )

	// TODO: These should update local state.
	const handleAddAction = api.plug.action.add.useMutation({
		// onSuccess: action => {
		// 	if (!plug) return
		// 	const previous =
		// 		plug.versions[plug.versions.length - version].actions ?? []
		// 	console.log("mutation", action)
		// 	console.log("actions", actions, plug?.versions.length, version)
		// 	console.log(
		// 		"mutate",
		// 		plug?.versions[plug.versions.length - version].actions
		// 	)
		// 	return previous
		// },
		// onError: (error, newData, context) => setActions(context ?? [])
	})
	const handleEditAction = api.plug.action.edit.useMutation({
		onMutate: action => {
			// if (!plug) return
			// // const previous =
			// // 	plug.versions[plug.versions.length - version].actions ?? []
			// // console.log("mutation", action)
			// // console.log("actions", actions, plug?.versions.length, version)
			// // console.log(
			// // 	"mutate",
			// // 	plug?.versions[plug.versions.length - version].actions
			// // )
			// return previous
		},
		onError: (error, newData, context) => setActions(context ?? [])
	})
	const handleRemoveAction = api.plug.action.remove.useMutation({
		onMutate: action => {}
	})

	useEffect(() => {
		if (!plug) return

		setActions(getActions(plug, version))
	}, [plug, version])

	return (
		<ActionContext.Provider
			value={{
				actions,
				handleAdd: action =>
					handleAddAction.mutate({
						id,
						version,
						action
					}),
				handleEdit: actions =>
					handleEditAction.mutate({
						id,
						actions
					}),
				handleRemove: action => handleRemoveAction.mutate(action.id)
			}}
		>
			{children}
		</ActionContext.Provider>
	)
}

export const useActions = () => useContext(ActionContext)
