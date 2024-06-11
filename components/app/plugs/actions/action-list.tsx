import type { FC } from "react"

import { actionCategories } from "@/lib/constants"

import { ActionListItem } from "./action-list-item"

type Props = {
	handleNestedToggle: () => void
}

export const ActionList: FC<Props> = ({ handleNestedToggle }) => (
	<>
		{Object.keys(actionCategories).map(categoryName => {
			const category =
				actionCategories[categoryName as keyof typeof actionCategories]

			return (
				<ActionListItem
					key={categoryName}
					categoryName={categoryName as keyof typeof actionCategories}
					category={category}
					handleVisibleToggle={handleNestedToggle}
				/>
			)
		})}
	</>
)
