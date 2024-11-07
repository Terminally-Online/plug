import { useCallback, useMemo } from "react"

import { MinimalUserSocketModel } from "@/prisma/types"

import { useAtom } from "jotai"

import { Column, Schedule, Transfer } from "@/lib"
import { useSocket } from "@/state"

import { atomWithStorage } from "jotai/utils"

export const MOBILE_INDEX = -1
export const DEMO_VIEW_AS = "0x62180042606624f02d8a130da8a3171e9b33894d"

export const COLUMN_KEYS = {
	ACTIVITY: "ACTIVITY",
	ADD: "ADD",
	ADMIN: "ADMIN",
	ALERTS: "ALERTS",
	APPLICATION: "APPLICATION",
	ASSETS: "ASSETS",
	AUTHENTICATE: "AUTHENTICATE",
	COLLECTIBLES: "COLLECTIBLES",
	DISCOVER: "DISCOVER",
	EARNINGS: "EARNINGS",
	HOME: "HOME",
	MY_PLUGS: "MY_PLUGS",
	PANE: "PANE",
	PLUG: "PLUG",
	PLUGS: "PLUGS",
	POSITIONS: "POSITIONS",
	PROFILE: "PROFILE",
	SETTINGS: "SETTINGS",
	TOKENS: "TOKENS"
}

export const DEFAULT_COLUMN_WIDTH = 420
export const DEFAULT_COLUMNS = [
	{ key: COLUMN_KEYS.PANE, index: -2 },
	{ key: COLUMN_KEYS.HOME, index: -1 },
	{ key: COLUMN_KEYS.DISCOVER, index: 0 },
	{ key: COLUMN_KEYS.MY_PLUGS, index: 1 },
	{ key: COLUMN_KEYS.ACTIVITY, index: 2 },
	{ key: COLUMN_KEYS.TOKENS, index: 3 }
].map(column => ({ ...column, width: DEFAULT_COLUMN_WIDTH }))
export const DEFAULT_DEMO_COLUMNS = DEFAULT_COLUMNS.map(column => ({ ...column, viewAsId: DEMO_VIEW_AS }))

const COLUMN_OFFSET = DEFAULT_COLUMNS.filter(column => column.index < 0).length

const columnsAtom = atomWithStorage<Array<Column>>("socketColumns", DEFAULT_COLUMNS)
export const useColumns = (index?: number, key?: string) => {
	const { socket } = useSocket()

	const [columns, setColumns] = useAtom(columnsAtom)

	const updateColumns = useCallback((updater: (prev: Column[]) => Column[]) => setColumns(updater), [setColumns])

	/**
	 * Add a column either to the end or to a specific index.
	 * @param key The key of the column to add.
	 * @param index The index of the column to add.
	 * @param from The from of the column to add.
	 * @param item The item of the column to add.
	 */
	const add = useCallback(
		({ index, key, from, item }: Partial<Column>) => {
		  updateColumns(prev => {
			const newColumn = {
			  key,
			  index: prev.length,
			  from,
			  item,
			  width: DEFAULT_COLUMN_WIDTH
			} as Column
			const updatedColumns = [...prev, newColumn]
	  
			requestAnimationFrame(() => {
			  const container = document.querySelector(".flex.h-full.flex-row.overflow-x-auto")
			  if (container) {
				container.scrollTo({
				  left: container.scrollWidth,
				  behavior: "smooth"
				})
			  }
			})
	  
			return updatedColumns.map((col, idx) => ({ ...col, index: idx - COLUMN_OFFSET }))
		  })
		},
		[updateColumns]
	  )

	/**
	 * Update the state (navigate within the context) of a column.
	 * @param index The index of the column to update.
	 * @param key The new key of the column.
	 * @param item The new item of the column.
	 * @param from The new from of the column.
	 */
	const navigate = useCallback(
		({ index, key, item, from }: Partial<Column> & { index: number }) =>
			updateColumns(prev =>
				prev.map(col => (col.index === index ? { ...col, key: key ?? col.key, item, from } : col))
			),
		[updateColumns]
	)

	/**
	 * Remove a column from the state.
	 * @param index The index of the column to remove.
	 */
	const remove = useCallback(
		(index: number) =>
			updateColumns(prev =>
				prev
					.filter(col => col.index !== index)
					.sort((a, b) => a.index - b.index)
					.map((col, idx) => ({ ...col, index: idx - COLUMN_OFFSET }))
			),
		[updateColumns]
	)

	/**
	 * Move one column (in from index) to another index (to index).
	 * @param from The index of the column to move the column from.
	 * @param to The index of the column to move the column to.
	 */
	const move = useCallback(
		({ from, to }: { from: number; to: number }) =>
			updateColumns(prev => {
				const updatedColumns = [...prev]
				const [movedColumn] = updatedColumns.splice(from + COLUMN_OFFSET, 1)
				updatedColumns.splice(to + COLUMN_OFFSET, 0, movedColumn)
				return updatedColumns.map((col, idx) => ({ ...col, index: idx - COLUMN_OFFSET }))
			}),
		[updateColumns]
	)

	/**
	 * Resize a single column.
	 * @param index The index of the column to resize.
	 * @param width The new width of the column.
	 */
	const resize = useCallback(
		({ index, width }: { index: number; width: number }) =>
			updateColumns(prev => prev.map(col => (col.index === index ? { ...col, width } : col))),
		[updateColumns]
	)

	/**
	 * View the state of a column as another socket (user).
	 * @param index The index of the column to view as.
	 * @param as The socket to view as.
	 */
	const as = useCallback(
		({ index, as }: { index: number; as: MinimalUserSocketModel }) =>
			updateColumns(prev =>
				prev.map(col =>
					col.index === index ? { ...col, viewAs: as === col.viewAs ? undefined : as, frame: undefined } : col
				)
			),
		[updateColumns]
	)

	/**
	 * Toggle the frame visibility of a column.
	 * @param index The index of the column to toggle the frame state in.
	 * @param key The id of the frame to toggole.
	 */
	const frame = (columnKey?: string) => {
		const frameKey = columnKey || key
		updateColumns(prev =>
			prev.map(col =>
				col.index === index ? { ...col, frame: col.frame === frameKey ? undefined : frameKey } : col
			)
		)
	}

	const schedule = (schedule?: Schedule) => {
		updateColumns(prev => prev.map(col => (col.index === index ? { ...col, schedule } : col)))
	}

	const transfer = (updater: Transfer | undefined | ((prev: Transfer | undefined) => Transfer)) => {
		updateColumns(prev =>
			prev.map(col =>
				col.index === index
					? { ...col, transfer: typeof updater === "function" ? updater(col.transfer) : updater }
					: col
			)
		)
	}

	const column = useMemo(
		() => (index !== undefined ? columns.find(column => column.index === index) : undefined),
		[columns, index]
	)

	const isExternal = useMemo(() => {
		if (socket === undefined) return false

		return (column !== undefined && column.viewAs && column.viewAs.id !== socket.id) || false
	}, [column, socket])

	const isFrame = useMemo(() => {
		if (column === undefined) return false

		return column.frame === key
	}, [column, key])

	return {
		columns,
		column,
		isExternal,
		isFrame,
		add,
		navigate,
		remove,
		move,
		resize,
		as,
		schedule,
		transfer,
		frame
	}
}
