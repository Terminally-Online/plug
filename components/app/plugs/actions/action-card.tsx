import { FC, useMemo } from "react"

import { usePlugs } from "@/contexts"
import { categories, formatTitle, getValues, actions as staticActions } from "@/lib"
import { useColumns } from "@/state"

export const ActionCard: FC<{
	index: number
	item: string
	categoryName: keyof typeof categories
	category: (typeof categories)[keyof typeof categories]
}> = ({ index, item, categoryName, category }) => {
	const { frame } = useColumns(index)
	const { plug, actions, handle } = usePlugs(item)

	const primaryActions = useMemo(() => {
		return Object.keys(staticActions[categoryName]).reduce(
			(acc, actionName) => {
				const { primary, ...action } = staticActions[categoryName][actionName]

				if (primary) acc[actionName] = action

				return acc
			},
			{} as Record<string, (typeof staticActions)[typeof categoryName][string]>
		)
	}, [categoryName])

	if (!plug) return null

	return (
		<div
			className="flex h-36 w-full flex-row items-center rounded-[20px] p-4"
			style={{
				background: `linear-gradient(30deg, ${category.gradientFrom}, ${category.gradientTo})`
			}}
		>
			{Object.keys(primaryActions).map(actionName => {
				const { icon: Icon } = primaryActions[actionName]

				return (
					<button
						key={actionName}
						className="group relative z-[3] mx-auto flex cursor-pointer flex-col items-center gap-2 text-center text-white"
						onClick={() => {
							handle.action.edit({
								id: plug.id,
								actions: JSON.stringify([
									...actions,
									{
										categoryName,
										actionName,
										values: getValues(categoryName, actionName)
									}
								])
							})

							frame()
						}}
					>
						<div className="flex h-16 w-16 items-center justify-center rounded-full bg-white/20 transition-all duration-200 ease-in-out group-hover:bg-white/40">
							{Icon && <Icon size={24} />}
						</div>
						<p className="max-w-[120px] text-sm font-bold">{formatTitle(actionName)}</p>
					</button>
				)
			})}
		</div>
	)
}
