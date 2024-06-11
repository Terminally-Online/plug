import { type FC, useMemo } from "react"

import { usePlugs } from "@/contexts"
import { actionCategories, actions } from "@/lib/constants"
import { formatTitle } from "@/lib/functions"

type Props = {
	categoryName: keyof typeof actionCategories
	category: (typeof actionCategories)[keyof typeof actionCategories]
	handleVisibleToggle: () => void
}

export const ActionCard: FC<Props> = ({
	categoryName,
	category,
	handleVisibleToggle
}) => {
	const { handleAddAction } = usePlugs()

	const primaryActions = useMemo(() => {
		const categoryActions = actions[categoryName]
		const primaryActions: Record<
			keyof typeof categoryActions,
			(typeof categoryActions)[keyof typeof categoryActions]
		> = {}

		for (const actionName in categoryActions) {
			if (categoryActions[actionName].primary) {
				primaryActions[actionName] = categoryActions[actionName]
			}
		}

		return primaryActions
	}, [categoryName])

	return (
		<div
			className="flex h-36 w-full flex-row items-center rounded-[20px] p-4"
			style={{
				background: `linear-gradient(45deg, ${category.gradientFrom}, ${category.gradientTo})`
			}}
		>
			{Object.keys(primaryActions).map(actionName => {
				const { icon: Icon, ...action } =
					primaryActions[actionName as keyof typeof primaryActions]

				return (
					<button
						key={actionName}
						className="group mx-auto flex cursor-pointer flex-col items-center gap-2 text-center text-white"
						onClick={() => {
							handleVisibleToggle()

							handleAddAction({
								categoryName,
								actionName,
								data: JSON.stringify(action)
							})
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
