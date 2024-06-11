import type { FC } from "react"

import { ActionListItem } from "@/components/app/plugs/actions/action-list-item"
import { actionCategories } from "@/lib/constants"

type Props = {
	handleNestedToggle: () => void
}

export const ActionList: FC<Props> = ({ handleNestedToggle }) => (
	<>
		{Object.keys(actionCategories).map(categoryName => (
			<ActionListItem
				key={categoryName}
				categoryName={categoryName}
				category={actionCategories[categoryName]}
				handleVisibleToggle={handleNestedToggle}
			/>
		))}
	</>
)
