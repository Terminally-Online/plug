import { FC } from "react"

import Image from "next/image"

import { X } from "lucide-react"

import { Button, Fragments } from "@/components"
import { usePlugs } from "@/contexts"
import { categories, cn } from "@/lib"

export const Sentence: FC<{
	id: string
	index: number
	preview?: boolean
}> = ({ id, index, preview = false }) => {
	const { plug, own, actions, handle } = usePlugs(id)

	const { categoryName } = actions[index]

	if (plug === undefined || actions === undefined) return null

	return (
		<>
			<div
				className={cn(
					"flex flex-row items-center font-bold",
					preview === false && "rounded-lg bg-grayscale-0 p-4"
				)}
			>
				<p className="flex w-full flex-wrap items-center gap-[4px]">
					{preview === false && (
						<Image
							className="mr-2 h-6 w-6 rounded-sm"
							src={categories[categoryName].image}
							alt={`Icon for ${categoryName}`}
							width={24}
							height={24}
						/>
					)}

					<Fragments id={id} index={index} />
				</p>

				{preview === false && own && (
					<Button
						variant="secondary"
						className="mb-auto ml-4 mt-[4px] p-1"
						onClick={() =>
							handle.action.edit({
								id: plug.id,
								actions: JSON.stringify(actions.filter((_, i) => i !== index))
							})
						}
					>
						<X size={14} />
					</Button>
				)}
			</div>

			{index < actions.length - 1 && <div className="mx-auto h-2 w-[2px] bg-grayscale-100" />}
		</>
	)
}
