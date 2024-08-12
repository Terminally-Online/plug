import { FC, useEffect, useRef, useState } from "react"

import { Grip, X } from "lucide-react"

import { Draggable } from "@hello-pangea/dnd"

import {
	ConsoleColumnAddOptions,
	Header,
	PageDiscover,
	PageMine,
	Plugs,
	SocketActivity,
	SocketAssets,
	SocketCollectionList,
	SocketEarnings,
	SocketPositionList,
	SocketTokenList
} from "@/components"
import { useSockets } from "@/contexts"
import { cn, formatTitle } from "@/lib"
import { ConsoleColumnModel } from "@/prisma/types"

const DEFAULT_COLUMN_WIDTH = 420
const MIN_COLUMN_WIDTH = 380
const MAX_COLUMN_WIDTH = 920

const getBoundedWidth = (width: number) =>
	Math.min(Math.max(width, MIN_COLUMN_WIDTH), MAX_COLUMN_WIDTH)

export const ConsoleColumn: FC<{
	column: ConsoleColumnModel
}> = ({ column }) => {
	const resizeRef = useRef<HTMLDivElement>(null)

	const { handle } = useSockets()

	const [isResizing, setIsResizing] = useState(false)

	useEffect(() => {
		const handleMouseMove = (e: MouseEvent) => {
			if (!resizeRef.current || !isResizing) return

			handle.columns.resize({
				id: column.id,
				width: getBoundedWidth(
					e.clientX - resizeRef.current.getBoundingClientRect().left
				)
			})
		}

		const handleMouseUp = () => {
			setIsResizing(false)
		}

		if (isResizing) {
			window.addEventListener("mousemove", handleMouseMove)
			window.addEventListener("mouseup", handleMouseUp)
		}

		return () => {
			window.removeEventListener("mousemove", handleMouseMove)
			window.removeEventListener("mouseup", handleMouseUp)
		}
	}, [isResizing, column.id, handle.columns])

	return (
		<div className="relative my-2 select-none">
			<Draggable
				draggableId={column.index.toString()}
				index={column.index}
			>
				{(provided, snapshot) => (
					<div
						ref={provided.innerRef}
						className="relative ml-2 flex h-full w-full flex-col rounded-lg border-[1px] border-grayscale-100 bg-white"
						{...provided.draggableProps}
						style={{
							...provided.draggableProps.style,
							width: `${column.width ?? DEFAULT_COLUMN_WIDTH}px`
						}}
					>
						{/* <div
							ref={resizeRef}
							className={cn(
								"relative my-2 w-full select-none overflow-y-auto rounded-lg border-[1px] border-grayscale-100 bg-white",
								snapshot.isDragging && "opacity-60"
							)}
						> */}
						<div
							className={cn(
								"group z-[11] flex cursor-pointer flex-row items-center gap-4 rounded-t-lg border-b-[1px] border-grayscale-100 bg-white px-4 transition-all duration-200 ease-in-out",
								snapshot.isDragging
									? "bg-grayscale-0"
									: "hover:bg-grayscale-0"
							)}
							{...provided.dragHandleProps}
						>
							<Header
								size="md"
								icon={
									<Grip
										size={14}
										className="opacity-40 transition-all duration-200 ease-in-out group-hover:opacity-60"
									/>
								}
								label={formatTitle(
									column.key.replace("_", " ").toLowerCase()
								)}
								nextPadded={false}
								nextOnClick={() =>
									handle.columns.remove(column.id)
								}
								nextLabel={<X size={14} />}
							/>
						</div>
						{/* </div> */}

						<div className="overflow-y-scroll">
							{column.key === "ADD" ? (
								<ConsoleColumnAddOptions id={column.id} />
							) : column.key === "PLUGS" ? (
								<Plugs className="px-4" />
							) : column.key === "DISCOVER" ? (
								<PageDiscover className="pt-4" />
							) : column.key === "MY_PLUGS" ? (
								<PageMine className="pt-4" column={true} />
							) : column.key === "ACTIVITY" ? (
								<SocketActivity
									id={column.id}
									className="px-4 pt-4"
								/>
							) : column.key === "ASSETS" ? (
								<SocketAssets id={column.id} className="px-4" />
							) : column.key === "TOKENS" ? (
								<SocketTokenList
									id={column.id}
									className="px-4 pt-4"
									expanded={true}
								/>
							) : column.key === "COLLECTIBLES" ? (
								<SocketCollectionList
									id={column.id}
									className="px-4 pt-4"
								/>
							) : column.key === "POSITIONS" ? (
								<SocketPositionList
									id={column.id}
									className="px-4 pt-4"
								/>
							) : column.key === "EARNINGS" ? (
								<SocketEarnings className="px-4 pt-4" />
							) : column.key === "SETTINGS" ? (
								<></>
							) : (
								<></>
							)}
						</div>

						{/* <div
							className="h-full cursor-col-resize px-2"
							onMouseDown={e => {
								e.preventDefault()
								setIsResizing(true)
							}}
						>
							<div
								className={cn(
									"h-full w-[1px] bg-grayscale-100",
									snapshot.isDragging && "opacity-0"
								)}
							/>
						</div> */}
					</div>
				)}
			</Draggable>
		</div>
	)
}
