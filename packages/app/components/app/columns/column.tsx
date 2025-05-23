import React, { FC, memo, useEffect, useRef, useState } from "react"

import { Check, ChevronLeft, GitFork, Plus, Share, X } from "lucide-react"

import { useAtom } from "jotai"

import { Draggable } from "@hello-pangea/dnd"

import { ColumnSettings } from "@/components/app/columns/settings/column"
import { ColumnAdd, OPTIONS } from "@/components/app/columns/utils/column-add"
import { ColumnApplication } from "@/components/app/columns/utils/column-application"
import { Header } from "@/components/app/layout/header"
import { PlugsDiscover } from "@/components/app/plugs/discover"
import { PlugsMine } from "@/components/app/plugs/mine"
import { Plug } from "@/components/app/plugs/plug"
import { SocketActivity } from "@/components/app/sockets/activity/activity-list"
import { SocketCollectionList } from "@/components/app/sockets/collectibles/collection-list"
import { SocketPositionList } from "@/components/app/sockets/position/position-list"
import { SocketTokenList } from "@/components/app/sockets/tokens/token-list"
import { Button } from "@/components/shared/buttons/button"
import { cardColors, cn, formatTitle } from "@/lib"
import { useSocket } from "@/state/authentication"
import { columnByIndexAtom, COLUMNS, useColumnActions } from "@/state/columns"
import { plugByIdAtom, usePlugActions } from "@/state/plugs"

import { SparklingText } from "../utils/sparkling-text"

const MIN_COLUMN_WIDTH = 420
const MAX_COLUMN_WIDTH = 680

const getBoundedWidth = (width: number) => Math.min(Math.max(width, MIN_COLUMN_WIDTH), MAX_COLUMN_WIDTH)

export const ConsoleColumn: FC<{
	index: number
}> = memo(({ index }) => {
	const resizeRef = useRef<HTMLDivElement>(null)

	const { socket } = useSocket()

	const [column] = useAtom(columnByIndexAtom(index))
	const { remove, resize, navigate } = useColumnActions(index)

	const [plug] = useAtom(plugByIdAtom(column?.item ?? "__non-existant__"))
	const { fork } = usePlugActions()

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
		<div className={cn("relative select-none")}>
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
							className="relative flex w-full select-none flex-col overflow-hidden bg-white"
						>
							<div
								className={cn(
									"group relative z-[999999] flex w-full cursor-pointer flex-row items-center gap-4 overflow-hidden overflow-y-auto border-b-[1px] border-plug-green/10 bg-white px-4 transition-all duration-200 ease-in-out",
									snapshot.isDragging ? "bg-plug-green/5" : "hover:bg-plug-green/5"
								)}
								{...provided.dragHandleProps}
							>
								<Header
									size="md"
									label={
										<div className="flex w-full flex-row items-center gap-4">
											{column.key !== COLUMNS.KEYS.ADD && (
												<Button
													variant="secondary"
													onClick={() =>
														navigate({
															index: column.index,
															key: column.from ?? COLUMNS.KEYS.ADD
														})
													}
													className="rounded-sm p-1"
												>
													<ChevronLeft size={14} />
												</Button>
											)}

											{plug ? (
												<div
													className="h-6 w-6 min-w-6 rounded-sm bg-plug-green/10"
													style={{
														backgroundImage: cardColors[plug.color]
													}}
												/>
											) : (
												(OPTIONS.find(option => option.label === column.key)?.icon ?? (
													<Plus size={18} className="opacity-40" />
												))
											)}

											<div className="top-0 z-[31] flex w-max flex-row items-center gap-2 overflow-hidden">
												<SparklingText
													className="text-lg font-bold"
													sparkles={Boolean(
														plug?.renamedAt &&
														plug.renamedAt > (plug.createdAt ?? 0) &&
														plug.renamedAt !== plug.createdAt
													)}
													sparkleKey={new Date(plug?.renamedAt ?? "")?.getTime()}
													color={cardColors[plug?.color ?? "yellow"]}
													item={column.item ?? ""}
												>
													{formatTitle(
														plug &&
															column.key === COLUMNS.KEYS.PLUG &&
															column.item !== undefined
															? plug.name
															: (column.key?.replace("_", " ").toLowerCase() ?? "ERROR")
													)}
												</SparklingText>
											</div>

											{plug && (
												<div className="ml-auto flex w-max flex-row items-center justify-end gap-4">
													{plug.actions.length !== 0 && <>
														<Button
															variant="secondary"
															className="rounded-sm p-1"
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
																<Check
																	size={14}
																	className="opacity-60 transition-all"
																/>
															) : (
																<Share
																	size={14}
																	className="opacity-60 transition-opacity group-hover:opacity-100"
																/>
															)}
														</Button>

														<Button
															variant="secondary"
															className="rounded-sm p-1"
															onClick={() =>
																fork({
																	index: column.index,
																	from: column.from ?? COLUMNS.KEYS.ADD,
																	plug: plug?.id ?? ""
																})
															}
														>
															<GitFork
																size={14}
																className="opacity-60 transition-opacity group-hover:opacity-100"
															/>
														</Button>
													</>}

													<Button
														variant="secondary"
														className="rounded-sm p-1"
														onClick={() => remove(column.index)}
													>
														<X
															size={14}
															className="opacity-60 transition-opacity group-hover:opacity-100"
														/>
													</Button>
												</div>
											)}
										</div>
									}
									nextPadded={false}
									nextOnClick={
										column.key !== COLUMNS.KEYS.PLUG ? () => remove(column.index) : undefined
									}
									nextLabel={
										<X
											size={14}
											className="opacity-60 transition-opacity group-hover:opacity-100"
										/>
									}
								/>
							</div>

							<div className="flex-1 overflow-y-auto overflow-x-hidden rounded-b-lg">
								{column.key === COLUMNS.KEYS.ADD ? (
									<ColumnAdd index={column.index} />
								) : column.key === COLUMNS.KEYS.DISCOVER ? (
									<PlugsDiscover index={column.index} className="pt-4" />
								) : column.key === COLUMNS.KEYS.MY_PLUGS ? (
									<PlugsMine index={column.index} className="pt-4" />
								) : column.key === COLUMNS.KEYS.PLUG && column.item ? (
									<Plug
										index={column.index}
										item={column.item}
										from={column.from}
										className="px-4 pt-4"
									/>
								) : column.key === COLUMNS.KEYS.ACTIVITY ? (
									<SocketActivity index={column.index} className="p-4" />
								) : column.key === COLUMNS.KEYS.TOKENS ? (
									<SocketTokenList index={column.index} isExpanded className="p-4" />
								) : column.key === COLUMNS.KEYS.COLLECTIBLES ? (
									<SocketCollectionList index={column.index} isExpanded className="p-4" />
								) : column.key === COLUMNS.KEYS.POSITIONS ? (
									<SocketPositionList index={column.index} className="p-4" />
								) : column.key === COLUMNS.KEYS.SETTINGS ? (
									<ColumnSettings index={column.index} className="p-4" />
								) : column.key === COLUMNS.KEYS.APPLICATION ? (
									<ColumnApplication index={column.index} />
								) : (
									<React.Fragment></React.Fragment>
								)}
							</div>
						</div>

						<div className="relative h-full cursor-col-resize">
							<div
								className={cn("h-full w-[1px] bg-plug-green/10", snapshot.isDragging && "opacity-0")}
							/>
							<div
								className="absolute -left-4 -right-4 bottom-0 top-0 z-[999]"
								onMouseDown={e => {
									e.preventDefault()
									setIsResizing(true)
								}}
							/>
						</div>
					</div>
				)
				}
			</Draggable >
		</div >
	)
})

ConsoleColumn.displayName = "ConsoleColumn"
