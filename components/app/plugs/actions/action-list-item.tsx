import Image from "next/image"
import { FC } from "react"

import { ChevronRight } from "lucide-react"

import { ActionCard, Button } from "@/components"
import { categories, formatTitle } from "@/lib"
import { useColumns } from "@/state"

export const ActionListItem: FC<{
	index: number
	item: string
	categoryName: string
}> = ({ index, item, categoryName }) => {
	const { frame } = useColumns(index, categoryName)

	const category = categories[categoryName]

	return (
		<div className="flex flex-col gap-4">
			<div className="flex flex-col items-center gap-2">
				<button className="group flex w-full flex-row items-center gap-4" onClick={() => frame()}>
					<Image
						src={category.image}
						alt={categoryName}
						width={32}
						height={32}
						className="h-6 w-6 rounded-md"
					/>

					<p className="text-lg font-bold">{formatTitle(categoryName)}</p>

					<Button
						variant="secondary"
						className="ml-auto p-1 group-hover:bg-grayscale-100"
						onClick={() => frame()}
					>
						<ChevronRight size={14} />
					</Button>
				</button>

				<ActionCard index={index} item={item} categoryName={categoryName} category={category} />
			</div>
		</div>
	)
}
