import { FC } from "react"

import { ActionListItem } from "@/components"
import { categories } from "@/lib"

export const ActionList: FC<{ id: string }> = ({ id }) => (
	<>
		{Object.keys(categories).map(categoryName => (
			<ActionListItem
				id={id}
				key={categoryName}
				categoryName={categoryName}
			/>
		))}
	</>
)
