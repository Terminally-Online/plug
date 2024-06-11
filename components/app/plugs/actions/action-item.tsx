import { type FC, useState } from "react"

import Image from "next/image"

import { Info } from "lucide-react"
import { set } from "zod"

import { Button } from "@/components/buttons"
import { usePlugs } from "@/contexts"
import { actionCategories, actions } from "@/lib/constants"
import { formatAddress, formatTitle } from "@/lib/functions"

import { Frame } from "../../frames/base"

type Props = {
	categoryName: keyof typeof actionCategories
	actionName: keyof (typeof actions)[keyof typeof actionCategories]
	handleVisibleToggle: () => void
}

export const ActionItem: FC<Props> = ({
	categoryName,
	actionName,
	handleVisibleToggle
}) => {
	const { handleAddAction } = usePlugs()

	const [actionVisible, setActionVisible] = useState(false)
	// @ts-ignore
	const { icon, ...action } = actions[categoryName][actionName]

	return (
		<>
			<div className="flex flex-row items-center gap-2">
				<Button
					variant="secondary"
					sizing="md"
					className="w-full px-6 text-left"
					onClick={() => {
						handleVisibleToggle()
						// @ts-ignore
						handleAddAction({
							categoryName,
							actionName,
							data: JSON.stringify(action)
						})
					}}
				>
					{formatTitle(actionName)}
				</Button>
				<button
					className="ml-2"
					onClick={() => setActionVisible(!actionVisible)}
				>
					<Info size={14} className="opacity-60" />
				</button>
			</div>

			<Frame
				className="scrollbar-hide z-[3] h-[calc(100vh-80px)] overflow-y-auto"
				icon={
					<Image
						src={actionCategories[categoryName].image}
						alt={categoryName}
						width={24}
						height={24}
						className="rounded-md"
					/>
				}
				label={formatTitle(actionName)}
				visible={actionVisible}
				handleBack={() => setActionVisible(false)}
				handleVisibleToggle={() => handleVisibleToggle()}
			>
				<div className="flex flex-col gap-8">
					<p className="opacity-60">{action.info}</p>

					<div className="flex flex-col gap-2">
						<p className="font-bold">Input Data</p>
						{action.inputs.length > 0 &&
							// @ts-ignore
							action.inputs.map((input, index) => {
								return (
									<p
										key={index}
										className="flex w-full flex-row gap-2"
									>
										<span className="font-bold opacity-40">
											{formatTitle(
												input.name?.replace("$", "") ??
													""
											)}
										</span>
										<span className="ml-auto opacity-60">
											{input.type}
										</span>
									</p>
								)
							})}
					</div>

					<div className="flex flex-col gap-2">
						<p className="font-bold">Fuse</p>
						<p className="flex flex-row gap-2">
							<span className="font-bold opacity-40">
								Address
							</span>
							<span className="ml-auto opacity-60">
								{formatAddress(action.address)}
							</span>
						</p>
						<p className="flex flex-row gap-2">
							<span className="font-bold opacity-40">
								Function
							</span>
							<div className="ml-auto opacity-60">plug</div>
						</p>
						<p className="flex flex-row gap-2">
							<span className="font-bold opacity-40">
								Supported Chains
							</span>
							<span className="ml-auto opacity-60">
								Supported Chains
							</span>
						</p>
					</div>
				</div>
			</Frame>
		</>
	)
}
