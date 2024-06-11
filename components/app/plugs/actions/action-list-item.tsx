import type { FC } from "react"
import { useMemo, useState } from "react"

import Image from "next/image"

import { ChevronRight } from "lucide-react"

import { ActionCard } from "@/components/app/plugs/actions/action-card"
import { Button } from "@/components/buttons"
import { actionCategories, actions } from "@/lib/constants"
import { formatTitle } from "@/lib/functions"

import { Frame } from "../../frames/base"
import { ActionItem } from "./action-item"

type Props = {
	categoryName: keyof typeof actionCategories
	category: (typeof actionCategories)[keyof typeof actionCategories]
	handleVisibleToggle: () => void
}

export const ActionListItem: FC<Props> = ({
	categoryName,
	category,
	handleVisibleToggle
}) => {
	const [actionItemsVisible, setActionItemsVisible] = useState(false)

	return (
		<div className="flex flex-col gap-4">
			<button
				className="flex flex-col items-center gap-2"
				onClick={() => setActionItemsVisible(!actionItemsVisible)}
			>
				<div className="flex w-full flex-row items-center gap-2">
					<Image
						src={category.image}
						alt={categoryName}
						width={24}
						height={24}
						className="rounded-md"
					/>
					<p className="text-lg font-bold">
						{formatTitle(categoryName)}
					</p>

					<Button
						variant="secondary"
						className="ml-auto p-1"
						onClick={() =>
							setActionItemsVisible(!actionItemsVisible)
						}
					>
						<ChevronRight size={14} className="opacity-60" />
					</Button>
				</div>

				<ActionCard
					categoryName={categoryName}
					category={category}
					handleVisibleToggle={handleVisibleToggle}
				/>
			</button>

			<Frame
				className="scrollbar-hide z-[2] h-[calc(100vh-80px)] overflow-y-auto"
				icon={
					<Image
						src={category.image}
						alt={categoryName}
						width={24}
						height={24}
						className="rounded-md"
					/>
				}
				label={formatTitle(categoryName)}
				visible={actionItemsVisible}
				handleBack={() => setActionItemsVisible(!actionItemsVisible)}
				handleVisibleToggle={() => handleVisibleToggle()}
			>
				<div className="flex flex-col gap-8">
					<ActionCard
						categoryName={categoryName}
						category={category}
						handleVisibleToggle={handleVisibleToggle}
					/>

					<div className="flex flex-col gap-2">
						{Object.keys(actions[categoryName]).map(actionName => (
							<ActionItem
								key={actionName}
								categoryName={categoryName}
								actionName={actionName}
								handleVisibleToggle={() =>
									handleVisibleToggle()
								}
							/>
						))}
					</div>
				</div>
			</Frame>
		</div>
	)
}
