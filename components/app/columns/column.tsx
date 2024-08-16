import { FC, useEffect, useRef, useState } from "react"

import { ChevronLeft, Cog, GitFork, Grip, Share, X } from "lucide-react"

import { Draggable } from "@hello-pangea/dnd"

import {
	Button,
	ConsoleColumnAddOptions,
	Header,
	Plug,
	Plugs,
	PlugsDiscover,
	PlugsMine,
	SocketActivity,
	SocketAssets,
	SocketCollectionList,
	SocketEarnings,
	SocketPositionList,
	SocketTokenList
} from "@/components"
import { useFrame, usePlugs, useSockets } from "@/contexts"
import { cardColors, cn, formatTitle, VIEW_KEYS } from "@/lib"
import { ConsoleColumnModel } from "@/prisma/types"

import { ConsoleAlerts } from "./column-alerts"
import { ConsoleSearch } from "./console-search"

const DEFAULT_COLUMN_WIDTH = 420
const MIN_COLUMN_WIDTH = 380
const MAX_COLUMN_WIDTH = 920

const getBoundedWidth = (width: number) =>
	Math.min(Math.max(width, MIN_COLUMN_WIDTH), MAX_COLUMN_WIDTH)

export const ConsoleColumn: FC<{
	column: ConsoleColumnModel
}> = ({ column }) => {
	const resizeRef = useRef<HTMLDivElement>(null)

	const { handleFrame } = useFrame({ id: column.id })
	const { handle } = useSockets()
	const { plug } = usePlugs(column.id)

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
			<Draggable draggableId={column.id.toString()} index={column.index}>
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
								"group z-[11] flex cursor-pointer flex-row items-center gap-4 overflow-hidden rounded-t-lg border-b-[1px] border-grayscale-100 bg-white px-4 transition-all duration-200 ease-in-out",
								snapshot.isDragging
									? "bg-grayscale-0"
									: "hover:bg-grayscale-0"
							)}
							{...provided.dragHandleProps}
						>
							<Header
								size="md"
								label={
									<div className="flex w-full flex-row items-center gap-4">
										<Button
											variant="none"
											onClick={() => {}}
											className="rounded-sm p-1"
										>
											<Grip
												size={14}
												className="opacity-40 transition-all duration-200 ease-in-out group-hover:opacity-60"
											/>
										</Button>

										{column.from && (
											<Button
												variant="secondary"
												onClick={() =>
													handle.columns.navigate({
														id: column.id,
														key: column.from!
													})
												}
												className="rounded-sm p-1"
											>
												<ChevronLeft size={14} />
											</Button>
										)}

										{plug && (
											<div
												className="h-6 w-6 min-w-6 rounded-sm bg-grayscale-100"
												style={{
													backgroundImage:
														cardColors[plug.color]
												}}
											/>
										)}

										<div className="relative mr-auto overflow-hidden truncate overflow-ellipsis whitespace-nowrap">
											<p className="overflow-hidden truncate overflow-ellipsis font-bold">
												{formatTitle(
													plug
														? plug.name
														: column.key
																.replace(
																	"_",
																	" "
																)
																.toLowerCase()
												)}
											</p>
										</div>

										{plug && (
											<div className="flex flex-row items-center justify-end gap-4">
												<Button
													variant="secondary"
													className="group rounded-sm p-1"
													onClick={
														() => {}
														// handleFrame("share")
													}
												>
													<Share
														size={14}
														className="opacity-60 hover:opacity-100"
													/>
												</Button>

												<Button
													variant="secondary"
													className="group rounded-sm p-1"
													onClick={
														() => {}
														// handleFrame("share")
													}
												>
													<GitFork
														size={14}
														className="opacity-60 hover:opacity-100"
													/>
												</Button>

												<Button
													variant="secondary"
													className="group rounded-sm p-1"
													onClick={() =>
														handleFrame("manage")
													}
												>
													<Cog
														size={14}
														className="opacity-60 hover:opacity-100"
													/>
												</Button>

												<Button
													variant="secondary"
													className="group rounded-sm p-1"
													onClick={() =>
														handle.columns.remove(
															column.id
														)
													}
												>
													<X
														size={14}
														className="opacity-60 hover:opacity-100"
													/>
												</Button>
											</div>
										)}
									</div>
								}
								nextPadded={false}
								nextOnClick={
									plug === undefined
										? () => handle.columns.remove(column.id)
										: undefined
								}
								nextLabel={<X size={14} />}
							/>
						</div>
						{/* </div> */}

						<div className="h-full overflow-y-scroll">
							{column.key === VIEW_KEYS.ADD ? (
								<ConsoleColumnAddOptions id={column.id} />
							) : column.key === VIEW_KEYS.SEARCH ? (
								<ConsoleSearch
									className="px-4 pt-4"
									id={column.id}
								/>
							) : column.key === VIEW_KEYS.ALERTS ? (
								<ConsoleAlerts
									id={column.id}
									className="px-4 pt-4"
								/>
							) : column.key === VIEW_KEYS.PLUGS ? (
								<Plugs className="px-4" id={column.id} />
							) : column.key === VIEW_KEYS.DISCOVER ? (
								<PlugsDiscover
									className="pt-4"
									id={column.id}
								/>
							) : column.key === VIEW_KEYS.MY_PLUGS ? (
								<PlugsMine
									className="pt-4"
									column={true}
									id={column.id}
								/>
							) : column.key === VIEW_KEYS.PLUG ? (
								<Plug
									className="px-4 pt-4"
									id={column.id}
									item={column.item}
								/>
							) : column.key === VIEW_KEYS.ACTIVITY ? (
								<SocketActivity
									id={column.id}
									className="px-4 pt-4"
								/>
							) : column.key === VIEW_KEYS.ASSETS ? (
								<SocketAssets id={column.id} className="px-4" />
							) : column.key === "TOKENS" ? (
								<SocketTokenList
									id={column.id}
									className="px-4 pt-4"
									expanded={true}
								/>
							) : column.key === VIEW_KEYS.COLLECTIBLES ? (
								<SocketCollectionList
									id={column.id}
									className="px-4 pt-4"
								/>
							) : column.key === VIEW_KEYS.POSITIONS ? (
								<SocketPositionList
									id={column.id}
									className="px-4 pt-4"
								/>
							) : column.key === VIEW_KEYS.EARNINGS ? (
								<SocketEarnings
									id={column.id}
									className="px-4 pt-4"
								/>
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
