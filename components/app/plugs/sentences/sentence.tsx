import { FC } from "react"

import Image from "next/image"

import { useSession } from "next-auth/react"

import { X } from "lucide-react"

import { Button, Fragments } from "@/components"
import { usePlugs } from "@/contexts"
import { categories, cn } from "@/lib"

export const Sentence: FC<{
	index: number
	preview?: boolean
}> = ({ index, preview = false }) => {
	const { data: session } = useSession()
	const { id, plug, actions, handle } = usePlugs()

	const { categoryName } = actions[index]

	const own = plug && session && session.address === plug.userAddress

	return (
		<>
			<div
				className={cn(
					"flex flex-row items-center font-bold",
					preview === false && "rounded-lg bg-grayscale-0 p-4"
				)}
			>
				<p className="flex w-full flex-wrap items-center gap-2">
					{preview === false && (
						<Image
							className="mr-2 h-6 w-6 rounded-md"
							src={categories[categoryName].image}
							alt={`Icon for ${categoryName}`}
							width={24}
							height={24}
						/>
					)}

					<Fragments index={index} />
				</p>

				{preview === false && own && (
					<Button
						variant="secondary"
						className="mb-auto ml-4 mt-[4px] p-1"
						onClick={() =>
							handle.action.edit({
								id,
								actions: JSON.stringify(
									actions.filter((_, i) => i !== index)
								)
							})
						}
					>
						<X size={14} />
					</Button>
				)}
			</div>

			{index < actions.length - 1 && (
				<div className="mx-auto h-4 w-[2px] bg-grayscale-100" />
			)}
		</>
	)
}
