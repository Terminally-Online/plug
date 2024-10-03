import Image from "next/image"
import { FC, useEffect, useRef, useState } from "react"

import { ChevronLeft, GitFork, Grip, Settings, X } from "lucide-react"

import { Draggable } from "@hello-pangea/dnd"

import {
	Avatar,
	Button,
	ColumnAddOptions,
	ColumnAuthenticate,
	ColumnProfile,
	ColumnSearch,
	ColumnViewAs,
	ConsoleAdmin,
	Header,
	Plug,
	PlugsDiscover,
	PlugsMine,
	SocketActivity,
	SocketCollectionList,
	SocketPositionList,
	SocketTokenList
} from "@/components"
import { usePlugs } from "@/contexts"
import { cardColors, cn, Column, formatTitle, VIEW_KEYS } from "@/lib"
import { useColumns, useSocket } from "@/state"

const MIN_COLUMN_WIDTH = 380
const MAX_COLUMN_WIDTH = 620

const getBoundedWidth = (width: number) => Math.min(Math.max(width, MIN_COLUMN_WIDTH), MAX_COLUMN_WIDTH)

export const ConsoleColumn: FC<{
	column: Column
}> = ({ column }) => {
	const { key, index, item, from, width } = column

	const resizeRef = useRef<HTMLDivElement>(null)

	const { socket } = useSocket()
	const { navigate, resize, remove, frame } = useColumns(index)

	const { plug, handle } = usePlugs(item)

	const [isResizing, setIsResizing] = useState(false)

	useEffect(() => {
		const handleMouseMove = (e: MouseEvent) => {
			if (!resizeRef.current || !isResizing) return

			resize({
				index,
				width: getBoundedWidth(e.clientX - resizeRef.current.getBoundingClientRect().left)
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
	}, [index, isResizing, resize])

	return (
		<div className="relative select-none">
			<Draggable draggableId={`${index}-${key}`} index={index}>
				{(provided, snapshot) => (
					<div
						ref={provided.innerRef}
						className="relative flex h-full w-full flex-row rounded-lg"
						{...provided.draggableProps}
						style={{
							...provided.draggableProps.style,
							width: `${width}px`
						}}
					>
						<div
							ref={resizeRef}
							className="relative my-2 w-full select-none overflow-y-hidden rounded-lg border-[1px] border-grayscale-100 bg-white"
						>
							<div
								className={cn(
									"group relative z-[30] flex cursor-pointer flex-row items-center gap-4 overflow-hidden overflow-y-auto rounded-t-lg border-b-[1px] border-grayscale-100 bg-white px-4 transition-all duration-200 ease-in-out",
									snapshot.isDragging ? "bg-grayscale-0" : "hover:bg-grayscale-0"
								)}
								{...provided.dragHandleProps}
							>
								<Header
									size="md"
									label={
										<div className="flex w-full flex-row items-center gap-4">
											<Button variant="none" onClick={() => {}} className="rounded-sm p-1">
												<Grip
													size={14}
													className="opacity-40 transition-all duration-200 ease-in-out group-hover:opacity-60"
												/>
											</Button>

											{from && (
												<Button
													variant="secondary"
													onClick={() =>
														navigate({
															index,
															key: from
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
														backgroundImage: cardColors[plug.color]
													}}
												/>
											)}

											{socket &&
												column.viewAs &&
												column.viewAs.socketAddress !== socket.socketAddress && (
													<div className="relative h-6 w-6 min-w-6 overflow-hidden rounded-sm">
														{column.viewAs.identity?.ens?.avatar ? (
															<Image
																src={column.viewAs.identity.ens.avatar}
																alt="ENS Avatar"
																width={240}
																height={240}
																className="rounded-sm"
															/>
														) : (
															<Avatar name={column.viewAs?.id ?? socket.id} className="rounded-sm" />
														)}
													</div>
												)}

											<div className="relative mr-auto overflow-hidden truncate overflow-ellipsis whitespace-nowrap">
												<p className="overflow-hidden truncate overflow-ellipsis text-lg font-bold">
													{formatTitle(
														plug
															? plug.name
															: (key?.replace("_", " ").toLowerCase() ?? "ERROR")
													)}
												</p>
											</div>

											{plug && (
												<div className="flex flex-row items-center justify-end gap-4">
													<Button
														variant="secondary"
														className="group rounded-sm p-1"
														onClick={() =>
															handle.plug.fork({
																plug: plug.id,
																index,
																from: key
															})
														}
													>
														<GitFork size={14} className="opacity-60 hover:opacity-100" />
													</Button>

													<Button
														variant="secondary"
														className="group rounded-sm p-1"
														onClick={() => frame(`${index}-${item}-manage`)}
													>
														<Settings size={14} className="opacity-60 hover:opacity-100" />
													</Button>

													<Button
														variant="secondary"
														className="group rounded-sm p-1"
														onClick={() => remove(index)}
													>
														<X size={14} className="opacity-60 hover:opacity-100" />
													</Button>
												</div>
											)}
										</div>
									}
									nextPadded={false}
									nextOnClick={plug === undefined ? () => remove(index) : undefined}
									nextLabel={<X size={14} />}
								/>
							</div>

							{/* TODO(#416): Have the column.frame check to disable the scroll when a frame is open. Need a better solution. */}
							<div className={cn("h-full", !column.frame && "overflow-y-scroll")}>
								{key === VIEW_KEYS.AUTHENTICATE ? (
									<ColumnAuthenticate index={index} />
								) : key === VIEW_KEYS.ADD ? (
									<ColumnAddOptions index={index} className="px-4 pt-4" />
								) : key === VIEW_KEYS.SEARCH ? (
									<ColumnSearch index={index} className="px-4 pt-4" />
								) : key === VIEW_KEYS.VIEW_AS ? (
									<ColumnViewAs />
								) : key === VIEW_KEYS.DISCOVER ? (
									<PlugsDiscover index={index} className="pt-4" />
								) : key === VIEW_KEYS.MY_PLUGS ? (
									<PlugsMine index={index} className="pt-4" />
								) : key === VIEW_KEYS.PLUG ? (
									<Plug index={index} item={item} from={from} className="px-4 pt-4" />
								) : key === VIEW_KEYS.ACTIVITY ? (
									<SocketActivity index={index} className="px-4 pt-4" />
								) : key === VIEW_KEYS.TOKENS ? (
									<SocketTokenList index={index} expanded={true} className="px-4 pt-4" />
								) : key === VIEW_KEYS.COLLECTIBLES ? (
									<SocketCollectionList index={index} expanded={true} className="px-4 pt-4" />
								) : key === VIEW_KEYS.POSITIONS ? (
									<SocketPositionList index={index} className="px-4 pt-4" />
								) : key === VIEW_KEYS.ADMIN ? (
									<ConsoleAdmin index={index} className="px-4 pt-4" />
								) : key === VIEW_KEYS.PROFILE ? (
									<ColumnProfile className="px-4 py-4" />
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
								className={cn("h-full w-[1px] bg-grayscale-100", snapshot.isDragging && "opacity-0")}
							/>
						</div>
					</div>
				)}
			</Draggable>
		</div>
	)
}
