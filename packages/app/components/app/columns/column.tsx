import React, { FC, useCallback, useEffect, useRef, useState } from "react"

import { Check, ChevronLeft, Settings, Share, X } from "lucide-react"

import { Draggable } from "@hello-pangea/dnd"

import { ConsoleAdmin } from "@/components/app/columns/admin/column-admin"
import { ConsoleSettings } from "@/components/app/columns/admin/console.settings"
import { ColumnAdd } from "@/components/app/columns/utils/column-add"
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
import { COLUMNS, useColumnStore } from "@/state/columns"
import { usePlugStore } from "@/state/plugs"

const MIN_COLUMN_WIDTH = 420
const MAX_COLUMN_WIDTH = 680

const getBoundedWidth = (width: number) => Math.min(Math.max(width, MIN_COLUMN_WIDTH), MAX_COLUMN_WIDTH)

// Memoized column content component to prevent unnecessary re-renders
const ColumnContent: FC<{
	columnKey: string
	index: number
	item?: string
	from?: string
}> = React.memo(({ columnKey, index, item, from }) => {
	if (columnKey === COLUMNS.KEYS.ADD) return <ColumnAdd index={index} />
	if (columnKey === COLUMNS.KEYS.DISCOVER) return <PlugsDiscover index={index} className="pt-4" />
	if (columnKey === COLUMNS.KEYS.MY_PLUGS) return <PlugsMine index={index} className="pt-4" />
	if (columnKey === COLUMNS.KEYS.PLUG) return <Plug index={index} item={item} from={from} className="px-4 pt-4" />
	if (columnKey === COLUMNS.KEYS.ACTIVITY) return <SocketActivity index={index} className="p-4" />
	if (columnKey === COLUMNS.KEYS.TOKENS) return <SocketTokenList index={index} expanded={true} className="p-4" />
	if (columnKey === COLUMNS.KEYS.COLLECTIBLES)
		return <SocketCollectionList index={index} expanded={true} className="p-4" />
	if (columnKey === COLUMNS.KEYS.POSITIONS) return <SocketPositionList index={index} className="p-4" />
	if (columnKey === COLUMNS.KEYS.ADMIN) return <ConsoleAdmin index={index} className="p-4" />
	if (columnKey === COLUMNS.KEYS.SETTINGS) return <ConsoleSettings index={index} className="p-4" />
	if (columnKey === COLUMNS.KEYS.APPLICATION) return <ColumnApplication index={index} />
	return null
})
ColumnContent.displayName = "ColumnContent"

// Memoized header actions to prevent unnecessary re-renders
const HeaderActions: FC<{
	plug?: any
	socket?: any
	onFrame: () => void
	onRemove: () => void
}> = React.memo(({ plug, socket, onFrame, onRemove }) => {
	const [copied, setCopied] = useState(false)

	const handleShare = useCallback(async () => {
		try {
			const shareUrl = `${window.location.origin}/app?plug=${plug.id}&rfid=${socket?.identity?.referralCode}`
			await navigator.clipboard.writeText(shareUrl)
			setCopied(true)
			setTimeout(() => setCopied(false), 2000)
		} catch (err) {
			console.error("Failed to copy link:", err)
		}
	}, [plug, socket])

	if (!plug) return null

	return (
		<div className="flex flex-row items-center justify-end gap-4">
			<Button variant="secondary" className="rounded-sm p-1" onClick={handleShare}>
				{copied ? (
					<Check size={14} className="opacity-60 transition-all" />
				) : (
					<Share size={14} className="opacity-60 transition-opacity group-hover:opacity-100" />
				)}
			</Button>

			<Button variant="secondary" className="rounded-sm p-1" onClick={onFrame}>
				<Settings size={14} className="opacity-60 transition-opacity group-hover:opacity-100" />
			</Button>

			<Button variant="secondary" className="rounded-sm p-1" onClick={onRemove}>
				<X size={14} className="opacity-60 transition-opacity group-hover:opacity-100" />
			</Button>
		</div>
	)
})
HeaderActions.displayName = "HeaderActions"

export const ConsoleColumn: FC<{
	id: number
}> = React.memo(({ id }) => {
	const resizeRef = useRef<HTMLDivElement>(null)
	const { column, handle } = useColumnStore(id)
	const { plug } = usePlugStore(column?.item ?? "")
	const { socket } = useSocket()

	const [width, setWidth] = useState(column?.width ?? 0)
	const [isResizing, setIsResizing] = useState(false)

	// Handle column resizing
	useEffect(() => {
		if (!column || !resizeRef.current) return

		const handleMouseMove = (e: MouseEvent) => {
			if (!isResizing || !resizeRef.current) return
			const rect = resizeRef.current.getBoundingClientRect()
			const newWidth = getBoundedWidth(e.clientX - rect.left)
			setWidth(newWidth)
		}

		const handleMouseUp = () => {
			if (!isResizing) return
			setIsResizing(false)

			// Only update the column width when we finish resizing
			handle.resize({
				index: column.index,
				width
			})
		}

		if (isResizing) {
			window.addEventListener("mousemove", handleMouseMove)
			window.addEventListener("mouseup", handleMouseUp)
			document.body.style.userSelect = "none" // Prevent text selection while resizing
		}

		return () => {
			window.removeEventListener("mousemove", handleMouseMove)
			window.removeEventListener("mouseup", handleMouseUp)
			document.body.style.userSelect = "" // Reset user select
		}
	}, [column, isResizing, width, handle])

	// Initialize width from column
	useEffect(() => {
		if (column?.width && !isResizing) {
			setWidth(column.width)
		}
	}, [column?.width, isResizing])

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
											{column.from && (
												<Button
													variant="secondary"
													onClick={() =>
														handle.navigate({ index: column.index, key: column.from })
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

											<HeaderActions
												plug={plug}
												socket={socket}
												onFrame={() => handle.frame("manage")}
												onRemove={() => handle.remove(column.index)}
											/>
										</div>
									}
									nextPadded={false}
									nextOnClick={plug === undefined ? () => handle.remove(column.index) : undefined}
									nextLabel={
										<X
											size={14}
											className="opacity-60 transition-opacity group-hover:opacity-100"
										/>
									}
								/>
							</div>

							<div className="flex-1 overflow-y-auto rounded-b-lg">
								<ColumnContent
									columnKey={column.key ?? ""}
									index={column.index}
									item={column.item}
									from={column.from}
								/>
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
})
ConsoleColumn.displayName = "ConsoleColumn"
