import { FC } from "react"

import Image from "next/image"

import { X } from "lucide-react"
import { AbiParameter } from "viem"

import { usePlugs } from "@/contexts"
import { actionCategories } from "@/lib/constants"
import { cn } from "@/lib/utils"

import { Fragments } from "./fragments"

export type Action = NonNullable<
	ReturnType<typeof usePlugs>["plugs"]
>[number]["versions"][number]["actions"][number]

export type Option = {
	icon: JSX.Element | undefined
	label: string
	value: string | number
}

export type Value = string | Option | undefined

export type StaticAction = {
	address: string
	abi: string
	inputs: Array<AbiParameter>
	options: Array<Array<Option> | { [key: string]: Array<Option> }> | undefined
	sentence: string
	info: string
	icon: string
	primary: true
}

type Props = {
	action: Action
	preview?: boolean
}

export const Sentence: FC<Props> = ({ action, preview = false }) => {
	const { handleRemoveAction } = usePlugs()

	return (
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
						src={
							actionCategories[
								action.categoryName as keyof typeof actionCategories
							].image
						}
						alt={`Icon for ${action.categoryName}`}
						width={24}
						height={24}
					/>
				)}

				<Fragments action={action} />
			</p>

			{preview === false && (
				<button
					className="group mb-auto ml-4 mt-[4px] cursor-pointer rounded-full border-[1px] border-grayscale-100 p-1 hover:bg-grayscale-100"
					onClick={() => handleRemoveAction(action)}
				>
					<X
						size={14}
						className="opacity-60 group-hover:opacity-80"
					/>
				</button>
			)}
		</div>
	)
}
