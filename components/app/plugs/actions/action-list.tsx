import { ActionListItem } from "@/components/app"
import { categories } from "@/lib/constants"

export const ActionList = () => (
	<>
		{Object.keys(categories).map(categoryName => (
			<ActionListItem key={categoryName} categoryName={categoryName} />
		))}
	</>
)
