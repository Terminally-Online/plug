import { FC } from "react"

import { X } from "lucide-react"

import { Accordion, Button, Fragments, Image } from "@/components"
import { usePlugs } from "@/contexts"
import { categories, cn } from "@/lib"

export const Sentence: FC<{
	index: number
	item: string
	actionIndex: number
	preview?: boolean
}> = ({ index, item, actionIndex, preview = false }) => {
	const { plug, own, actions, handle } = usePlugs(item)

	const { categoryName } = actions[actionIndex]

	if (plug === undefined || actions === undefined) return null

	return (
		<>
			<Accordion className="hover:cursor-auto hover:border-grayscale-100 hover:bg-white">
				<div className={cn("flex flex-row items-center font-bold")}>
					<p className="flex w-full flex-wrap items-center gap-[4px]">
						{preview === false && (
							<div className="relative h-6 w-10">
								<Image
									className="absolute mr-2 h-6 w-6 rounded-sm blur-xl filter"
									src={categories[categoryName].image}
									alt={`Icon for ${categoryName}`}
									width={64}
									height={64}
								/>
								<Image
									className="absolute mr-2 h-6 w-6 rounded-sm"
									src={categories[categoryName].image}
									alt={`Icon for ${categoryName}`}
									width={64}
									height={64}
								/>
							</div>
						)}

						<Fragments index={index} item={item} actionIndex={actionIndex} />
					</p>

					{preview === false && own && (
						<Button
							variant="secondary"
							className="mb-auto ml-4 mt-[4px] p-1"
							onClick={() =>
								handle.action.edit({
									id: plug.id,
									actions: JSON.stringify(actions.filter((_, i) => i !== actionIndex))
								})
							}
						>
							<X size={14} />
						</Button>
					)}
				</div>
			</Accordion>

			{actionIndex < actions.length - 1 && <div className="mx-auto h-2 w-[2px] bg-grayscale-0" />}
		</>
	)
}
