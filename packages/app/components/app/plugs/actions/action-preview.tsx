import { FC, useMemo } from "react"

import { useAtom } from "jotai"

import { Sentence } from "@/components/app/plugs/sentences/sentence"
import { SchemasRequestActions, SchemasResponseCoils } from "@/lib"
import { useActions } from "@/state/actions"
import { plugByIdAtom } from "@/state/plugs"

export const ActionPreview: FC<{
	index: number
	item: string
	actions?: SchemasRequestActions
	errors?: Array<string | null>
}> = ({ index, item, actions, errors = [] }) => {
	const [plug] = useAtom(plugByIdAtom(item))
	const [solverActions] = useActions()

	actions = actions ?? plug?.actions ?? []

	// Calculate available coils for each action
	const availableCoils = useMemo(() => {
		if (!plug || !solverActions) return {}

		const coils: Record<string, { type: string; actionIndex: number }> = {}

		plug.actions.forEach((action, actionIndex) => {
			const actionSchema = solverActions[action.protocol]?.schema[action.action]

			if (!actionSchema || !actionSchema.coils) return

			Object.keys(actionSchema.coils).forEach(name => {
				if (!actionSchema?.coils?.[name]) return

				coils[name] = {
					type: actionSchema.coils[name],
					actionIndex
				}
			})
		})

		return coils
	}, [plug, solverActions])

	// Build a mapping of previous coils for each action
	const getCoilsForAction = (actionIndex: number): SchemasResponseCoils => {
		if (!plug || !solverActions) return {}

		// Only include coils from previous actions
		const prevCoils: SchemasResponseCoils = {}

		for (let i = 0; i < actionIndex; i++) {
			const prevAction = plug.actions[i]
			const prevSchema = solverActions[prevAction.protocol]?.schema[prevAction.action]

			if (prevSchema?.coils) {
				Object.entries(prevSchema.coils).forEach(([name, type]) => {
					prevCoils[name] = type
				})
			}
		}

		return prevCoils
	}

	// Function to validate coil types
	const validateType = (coilName: string, expectedType: string): boolean => {
		if (!availableCoils[coilName]) return false
		return availableCoils[coilName].type === expectedType
	}

	return (
		<div className="flex flex-col">
			{actions.map((action, actionIndex) => {
				// Get coils from previous actions that can be linked in this action
				const prevCoils = getCoilsForAction(actionIndex)

				return (
					<div key={`${index}-${actionIndex}`}>
						<Sentence
							index={index}
							item={item}
							actionIndex={actionIndex}
							action={action}
							preview
							error={errors && errors[actionIndex] ? true : false}
							prevCoils={prevCoils}
							validateType={validateType}
							availableCoils={availableCoils}
						/>

						{errors && errors[actionIndex] && (
							<p className="text-sm font-bold text-plug-red">Error: {errors[actionIndex]}</p>
						)}
					</div>
				)
			})}
		</div>
	)
}
