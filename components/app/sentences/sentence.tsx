import { FC } from "react"

import Image from "next/image"

import { X } from "lucide-react"

import { usePlugs } from "@/contexts"
import { actionCategories } from "@/lib/constants"
import { cn } from "@/lib/utils"

import { Fragments } from "./fragments"

export const Sentence: FC<{
	index: number
	preview?: boolean
}> = ({ index, preview = false }) => {
	const { id, actions, handle } = usePlugs()

	const { categoryName } = actions[index]

	return (
		<>
			<div
				className={cn(
					"flex flex-row items-center font-bold",
					preview === false && "rounded-lg bg-grayscale-0 p-4"
				)}
			>
				<p className="flex w-full flex-wrap items-center gap-[8px]">
					{preview === false && (
						<Image
							className="mr-2 h-6 w-6 rounded-md"
							src={actionCategories[categoryName].image}
							alt={`Icon for ${categoryName}`}
							width={24}
							height={24}
						/>
					)}

					<Fragments index={index} />
				</p>

				{preview === false && (
					<button
						className="group mb-auto ml-4 mt-[4px] cursor-pointer rounded-full border-[1px] border-grayscale-100 p-1 hover:bg-grayscale-100"
						onClick={() =>
							handle.action.edit({
								id,
								actions: JSON.stringify(
									actions.filter((_, i) => i !== index)
								)
							})
						}
					>
						<X
							size={14}
							className="opacity-60 group-hover:opacity-80"
						/>
					</button>
				)}
			</div>

			{index < actions.length - 1 && (
				<div className="mx-auto h-4 w-[2px] bg-grayscale-100" />
			)}
		</>
	)
}
