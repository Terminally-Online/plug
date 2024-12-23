import React, { FC, useEffect, useRef, useState } from "react"

import { Check, ChevronLeft, PlugIcon, Settings, Share, Star, X } from "lucide-react"

import { Draggable } from "@hello-pangea/dnd"

import {
	ADMIN_OPTIONS,
	Button,
	ColumnAdd,
	ColumnApplication,
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
import { cardColors, cn, formatTitle } from "@/lib"
import { COLUMNS, useColumnStore, usePlugStore, useSocket } from "@/state"

const MIN_COLUMN_WIDTH = 420
const MAX_COLUMN_WIDTH = 680

const getBoundedWidth = (width: number) => Math.min(Math.max(width, MIN_COLUMN_WIDTH), MAX_COLUMN_WIDTH)

export const ConsoleColumn: FC<{
	id: number
}> = ({ id }) => {
	const resizeRef = useRef<HTMLDivElement>(null)

	const {
		column,
		handle: { frame, remove, resize, navigate }
	} = useColumnStore(id)
	const { plug } = usePlugStore(column?.item ?? "")
	const { socket } = useSocket()

	const [width, setWidth] = useState(column?.width ?? 0)
	const [isResizing, setIsResizing] = useState(false)
	const [copied, setCopied] = useState(false)

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
	}, [resize, column, width, isResizing])

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
							className="relative my-2 flex w-full select-none flex-col overflow-hidden rounded-lg border-[1px] border-plug-green/10 bg-white"
						>
							<div
								className={cn(
									"group relative z-[30] flex cursor-pointer flex-row items-center gap-4 overflow-hidden overflow-y-auto rounded-t-lg border-b-[1px] border-plug-green/10 bg-white px-4 transition-all duration-200 ease-in-out",
									snapshot.isDragging ? "bg-plug-green/5" : "hover:bg-plug-green/5"
								)}
								{...provided.dragHandleProps}
							>
								<Header
									size="md"
									label={
										<div className="flex w-full flex-row items-center gap-4">
											<p className="rounded-sm p-1">
												{ADMIN_OPTIONS.find(option => option.label === column.key)?.icon ?? (
													<>
														{column.key === COLUMNS.KEYS.PLUG ? (
															<PlugIcon size={14} className="opacity-40" />
														) : (
															<Star size={14} className="opacity-40" />
														)}
													</>
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
													className="h-6 w-6 min-w-6 rounded-sm bg-plug-green/10"
													style={{
														backgroundImage: cardColors[plug.color]
													}}
												/>
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
														onClick={async () => {
															try {
																const shareUrl = `${window.location.origin}/app?plug=${plug.id}&rfid=${socket?.identity?.referralCode}`
																await navigator.clipboard.writeText(shareUrl)
																setCopied(true)
																setTimeout(() => setCopied(false), 2000)
															} catch (err) {
																console.error("Failed to copy link:", err)
															}
														}}
													>
														{copied ? (
															<Check size={14} className="opacity-60 transition-all" />
														) : (
															<Share
																size={14}
																className="opacity-60 transition-opacity group-hover:opacity-100"
															/>
														)}
													</Button>

													<Button
														variant="secondary"
														className="group rounded-sm p-1"
														onClick={() => frame("manage")}
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

							<div className="flex-1 overflow-y-auto rounded-b-lg">
								{column.key === COLUMNS.KEYS.ADD ? (
									<ColumnAdd />
								) : column.key === COLUMNS.KEYS.DISCOVER ? (
									<PlugsDiscover index={column.index} className="pt-4" />
								) : column.key === COLUMNS.KEYS.MY_PLUGS ? (
									<PlugsMine index={column.index} className="pt-4" />
								) : column.key === COLUMNS.KEYS.PLUG ? (
									<Plug
										index={column.index}
										item={column.item}
										from={column.from}
										className="px-4 pt-4"
									/>
								) : column.key === COLUMNS.KEYS.ACTIVITY ? (
									<SocketActivity index={column.index} className="px-4 pt-4" />
								) : column.key === COLUMNS.KEYS.TOKENS ? (
									<SocketTokenList index={column.index} expanded={true} className="px-4 pt-4" />
								) : column.key === COLUMNS.KEYS.COLLECTIBLES ? (
									<SocketCollectionList index={column.index} expanded={true} className="px-4 pt-4" />
								) : column.key === COLUMNS.KEYS.POSITIONS ? (
									<SocketPositionList index={column.index} className="px-4 pt-4" />
								) : column.key === COLUMNS.KEYS.ADMIN ? (
									<ConsoleAdmin index={column.index} className="px-4 pt-4" />
								) : column.key === COLUMNS.KEYS.APPLICATION ? (
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
								className={cn("h-full w-[1px] bg-plug-green/10", snapshot.isDragging && "opacity-0")}
							/>
						</div>
					</div>
				)}
			</Draggable>
		</div>
	)
}
