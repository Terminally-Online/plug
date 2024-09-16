import { FC } from "react"

import { ActionListItem } from "@/components"
import { categories } from "@/lib"

export const ActionList: FC<{ index: number; item: string }> = ({ index, item }) => (
	<>
		{Object.keys(categories).map(categoryName => (
			<ActionListItem key={categoryName} index={index} item={item} categoryName={categoryName} />
		))}
	</>
)
