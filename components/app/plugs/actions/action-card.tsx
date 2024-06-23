import { FC, useMemo } from "react"

import { usePlugs } from "@/contexts"
import { actionCategories, actions as staticActions } from "@/lib/constants"
import { formatTitle, getValues } from "@/lib/functions"

export const ActionCard: FC<{
	categoryName: keyof typeof actionCategories
	category: (typeof actionCategories)[keyof typeof actionCategories]
	handleVisibleToggle: () => void
}> = ({ categoryName, category, handleVisibleToggle }) => {
	const { id, actions, handle } = usePlugs()

	const primaryActions = useMemo(() => {
		return Object.keys(staticActions[categoryName]).reduce(
			(acc, actionName) => {
				const { primary, ...action } =
					staticActions[categoryName][actionName]

				if (primary) acc[actionName] = action

				return acc
			},
			{} as Record<
				string,
				(typeof staticActions)[typeof categoryName][string]
			>
		)
	}, [categoryName])

	if (!id) return null

	return (
		<div
			className="flex h-36 w-full flex-row items-center rounded-[20px] p-4"
			style={{
				background: `linear-gradient(45deg, ${category.gradientFrom}, ${category.gradientTo})`
			}}
		>
			{Object.keys(primaryActions).map(actionName => {
				const { icon: Icon } = primaryActions[actionName]

				return (
					<button
						key={actionName}
						className="group mx-auto flex cursor-pointer flex-col items-center gap-2 text-center text-white"
						onClick={() => {
							handle.action.edit({
								id,
								actions: JSON.stringify([
									...actions,
									{
										categoryName,
										actionName,
										values: getValues(
											categoryName,
											actionName
										)
									}
								])
							})

							handleVisibleToggle()
						}}
					>
						<div className="flex h-16 w-16 items-center justify-center rounded-full bg-white/20 transition-all duration-200 ease-in-out group-hover:bg-white/40">
							{Icon && <Icon size={24} />}
						</div>
						<p className="max-w-[120px] text-sm">
							{formatTitle(actionName)}
						</p>
					</button>
				)
			})}
		</div>
	)
}
