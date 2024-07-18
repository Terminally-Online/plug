import { ActionListItem } from "@/components"
import { categories } from "@/lib"

export const ActionList = () => (
	<>
		{Object.keys(categories).map(categoryName => (
			<ActionListItem key={categoryName} categoryName={categoryName} />
		))}
	</>
)
