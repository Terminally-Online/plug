import type { FC } from "react"

import { ActionListItem } from "@/components/app/plugs/actions/action-list-item"
import { categories } from "@/lib/constants"

export const ActionList = () => (
	<>
		{Object.keys(categories).map(categoryName => (
			<ActionListItem key={categoryName} categoryName={categoryName} />
		))}
	</>
)
