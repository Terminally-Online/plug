import { FC, useEffect, useRef, useState } from "react"

import { Grip, X } from "lucide-react"

import { Draggable } from "@hello-pangea/dnd"

import {
	Header,
	PlugsDiscover,
	SocketActivity,
	SocketAssets,
	SocketCollectionList,
	SocketPositionList,
	SocketTokenList
} from "@/components/app"
import { PlugsMine } from "@/components/app/plugs/mine"
import { ConsoleColumnAddOptions } from "@/components/console"
import { Plugs } from "@/components/shared/framework/plugs"
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
		<div className={cn(column.index === 0 && "ml-2")}>
			<Draggable
				draggableId={column.index.toString()}
				index={column.index}
			>
				{(provided, snapshot) => (
					<div
						className="flex h-full flex-row"
						ref={provided.innerRef}
						{...provided.draggableProps}
						style={{
							...provided.draggableProps.style,
							width: `${column.width ?? DEFAULT_COLUMN_WIDTH}px`
						}}
					>
						<div
							className={cn(
								"relative my-2 w-full select-none overflow-y-auto rounded-lg border-[1px] border-grayscale-100 bg-white",
								snapshot.isDragging && "opacity-60"
							)}
							ref={resizeRef}
						>
							<div
								className={cn(
									"group flex cursor-pointer flex-row items-center gap-4 rounded-t-lg border-b-[1px] border-grayscale-100 px-4 transition-all duration-200 ease-in-out",
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
										column.key
											.replace("_", " ")
											.toLowerCase()
									)}
									nextPadded={false}
									nextOnClick={() =>
										handle.columns.remove(column.id)
									}
									nextLabel={<X size={14} />}
								/>
							</div>

							<div className="overflow-y-scrol">
								{column.key === "ADD" ? (
									<ConsoleColumnAddOptions id={column.id} />
								) : column.key === "PLUGS" ? (
									<Plugs className="px-4" />
								) : column.key === "DISCOVER" ? (
									<PlugsDiscover
										className="pt-4"
										column={true}
									/>
								) : column.key === "MY_PLUGS" ? (
									<PlugsMine className="pt-4" column={true} />
								) : column.key === "ASSETS" ? (
									<SocketAssets className="px-4" />
								) : column.key === "ACTIVITY" ? (
									<SocketActivity className="px-4" />
								) : column.key === "TOKENS" ? (
									<SocketTokenList
										className="px-4 pt-4"
										expanded={true}
									/>
								) : column.key === "COLLECTIBLES" ? (
									<SocketCollectionList className="px-4 pt-4" />
								) : column.key === "POSITIONS" ? (
									<SocketPositionList className="px-4 pt-4" />
								) : column.key === "EARNINGS" ? (
									<></>
								) : column.key === "SETTINGS" ? (
									<></>
								) : (
									<></>
								)}
							</div>
						</div>

						<div
							className="h-full cursor-col-resize px-2"
							onMouseDown={e => {
								e.preventDefault()
								setIsResizing(true)
							}}
						>
							<div
								className={cn(
									"aniamte-fade-in h-full w-[1px] bg-grayscale-100 transition-all duration-200 ease-in-out",
									snapshot.isDragging && "opacity-0"
								)}
							/>
						</div>
					</div>
				)}
			</Draggable>
		</div>
	)
}
