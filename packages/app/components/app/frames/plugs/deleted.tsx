import { FC } from "react"

import { ChevronLeft, Trash2 } from "lucide-react"

import { Button, Container, Frame } from "@/components"
import { useMediaQuery } from "@/lib"
import { COLUMNS, useColumnStore } from "@/state"
import { usePlugStore } from "@/state"

export const DeletedFrame: FC<{ index: number }> = ({ index }) => {
	const { isFrame, handle } = useColumnStore(index, "deleted")
	const { handle: plugHandle } = usePlugStore()
	const { md } = useMediaQuery()

	// Set frame state immediately when component mounts
	if (!isFrame) {
		handle.frame("deleted")
	}

	return (
		<>
			<Container className="border-grayscale-100 fixed left-0 right-0 top-0 z-[10] border-b-[1px] bg-white md:hidden">
				<div className="flex flex-row items-center gap-4 py-4">
					<Button
						variant="secondary"
						className="rounded-sm p-1"
						onClick={() =>
							handle.navigate({
								index: -1,
								key: COLUMNS.KEYS.HOME
							})
						}
					>
						<ChevronLeft size={14} />
					</Button>
					<span className="font-bold">Plug Unavailable</span>
				</div>
			</Container>
			<Frame
				index={index}
				className="z-[2] mt-16 md:mt-0"
				icon={<Trash2 size={18} className="opacity-40" />}
				label="Plug Unavailable"
				visible={isFrame}
				hasOverlay={true}
				next={<></>}
				nextEmpty={true}
			>
				<div className="flex flex-col gap-4">
					<p className="font-bold leading-6">
						<span className="opacity-40">
							This Plug is no longer available. It may have been deleted or made private by its creator.
						</span>
					</p>

					{md ? (
						<Button
							className="w-full py-4"
							onClick={() => {
								plugHandle.plug.add({ index, from: COLUMNS.KEYS.HOME })
							}}
						>
							Create New Plug
						</Button>
					) : (
						<div className="flex flex-row gap-2">
							<Button
								variant="secondary"
								className="w-full py-4"
								onClick={() => {
									handle.navigate({
										index: -1,
										key: COLUMNS.KEYS.HOME
									})
								}}
							>
								Return Home
							</Button>
							<Button
								className="w-full py-4"
								onClick={() => {
									plugHandle.plug.add({ index, from: COLUMNS.KEYS.HOME })
								}}
							>
								Create New Plug
							</Button>
						</div>
					)}
				</div>
			</Frame>
		</>
	)
}