import React, { FC, useEffect, useRef, useState } from "react"

import { ChevronLeft, GitFork, Settings, Star, X } from "lucide-react"

import { Draggable } from "@hello-pangea/dnd"

import {
	ADMIN_OPTIONS,
	Avatar,
	Button,
	ColumnAdd,
	ColumnApplication,
	ConsoleAdmin,
	Header,
	Image,
	Plug,
	PlugsDiscover,
	PlugsMine,
	SocketActivity,
	SocketCollectionList,
	SocketPositionList,
	SocketTokenList
} from "@/components"
import { usePlugs } from "@/contexts"
import { cardColors, cn, formatTitle } from "@/lib"
import { COLUMN_KEYS, useColumns, useSocket } from "@/state"

const MIN_COLUMN_WIDTH = 420
const MAX_COLUMN_WIDTH = 680

const getBoundedWidth = (width: number) => Math.min(Math.max(width, MIN_COLUMN_WIDTH), MAX_COLUMN_WIDTH)

export const ConsoleColumn: FC<{
	index: number
}> = ({ index }) => {
	const resizeRef = useRef<HTMLDivElement>(null)

	const { socket } = useSocket()
	const { column, navigate, resize, remove, frame } = useColumns(index)

	const { plug, handle } = usePlugs(column?.item ?? "")

	const [width, setWidth] = useState(column?.width ?? 0)
	const [isResizing, setIsResizing] = useState(false)

	useEffect(() => {
		if (!column) return

		const handleMouseMove = (e: MouseEvent) => {
			if (!resizeRef.current || !isResizing) return

			setWidth(getBoundedWidth(e.clientX - resizeRef.current.getBoundingClientRect().left))
		}

		const handleMouseUp = () => {
			setIsResizing(false)

			resize({
				index: column.index,
				width
			})
		}

		if (isResizing) {
			window.addEventListener("mousemove", handleMouseMove)
			window.addEventListener("mouseup", handleMouseUp)
		}

		return () => {
			window.removeEventListener("mousemove", handleMouseMove)
			window.removeEventListener("mouseup", handleMouseUp)
		}
	}, [column, width, isResizing, resize])

	if (!column) return null

	return (
		<div className={cn("relative select-none", column.index === 0 && "ml-2")}>
			<Draggable draggableId={String(column.id)} index={column.index}>
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
											<p className="rounded-sm p-1">
												{ADMIN_OPTIONS.find(option => option.label === column.key)?.icon ?? (
													<Star size={14} className="opacity-40" />
												)}
											</p>

											{column.from && (
												<Button
													variant="secondary"
													onClick={() =>
														navigate({
															index: column.index,
															key: column.from
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

											{socket && column.viewAs && column.viewAs.id !== socket.id && (
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
														<Avatar
															name={column.viewAs?.id ?? socket.id}
															className="rounded-sm"
														/>
													)}
												</div>
											)}

											<div className="relative mr-auto overflow-hidden truncate overflow-ellipsis whitespace-nowrap">
												<p className="overflow-hidden truncate overflow-ellipsis text-lg font-bold">
													{formatTitle(
														plug
															? plug.name
															: (column.key?.replace("_", " ").toLowerCase() ?? "ERROR")
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
																index: column.index,
																from: column.key
															})
														}
													>
														<GitFork size={14} className="opacity-60 hover:opacity-100" />
													</Button>

													<Button
														variant="secondary"
														className="group rounded-sm p-1"
														onClick={() => frame(`${column.item}-manage`)}
													>
														<Settings size={14} className="opacity-60 hover:opacity-100" />
													</Button>

													<Button
														variant="secondary"
														className="group rounded-sm p-1"
														onClick={() => remove(column.index)}
													>
														<X size={14} className="opacity-60 hover:opacity-100" />
													</Button>
												</div>
											)}
										</div>
									}
									nextPadded={false}
									nextOnClick={plug === undefined ? () => remove(column.index) : undefined}
									nextLabel={<X size={14} />}
								/>
							</div>

							<div className="h-full overflow-y-scroll">
								{column.key === COLUMN_KEYS.ADD ? (
									<ColumnAdd />
								) : column.key === COLUMN_KEYS.DISCOVER ? (
									<PlugsDiscover index={column.index} className="pt-4" />
								) : column.key === COLUMN_KEYS.MY_PLUGS ? (
									<PlugsMine index={column.index} className="pt-4" />
								) : column.key === COLUMN_KEYS.PLUG ? (
									<Plug
										index={column.index}
										item={column.item}
										from={column.from}
										className="px-4 pt-4"
									/>
								) : column.key === COLUMN_KEYS.ACTIVITY ? (
									<SocketActivity index={column.index} className="px-4 pt-4" />
								) : column.key === COLUMN_KEYS.TOKENS ? (
									<SocketTokenList index={column.index} expanded={true} className="px-4 pt-4" />
								) : column.key === COLUMN_KEYS.COLLECTIBLES ? (
									<SocketCollectionList index={column.index} expanded={true} className="px-4 pt-4" />
								) : column.key === COLUMN_KEYS.POSITIONS ? (
									<SocketPositionList index={column.index} className="px-4 pt-4" />
								) : column.key === COLUMN_KEYS.ADMIN ? (
									<ConsoleAdmin index={column.index} className="px-4 pt-4" />
								) : column.key === COLUMN_KEYS.APPLICATION ? (
									<ColumnApplication index={column.index} className="pt-4" />
								) : (
									<React.Fragment></React.Fragment>
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
