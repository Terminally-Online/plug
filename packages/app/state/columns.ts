import { useCallback, useMemo } from "react"

import { atom, useAtom, useAtomValue } from "jotai"

import { Column, Schedule, Transfer } from "@/lib"

import { atomWithStorage } from "jotai/utils"

export const COLUMNS = {
	MOBILE_INDEX: -1,
	DEFAULT_WIDTH: 520,
	OFFSET: 2,
	KEYS: {
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
} as const

export const DEFAULT_COLUMNS = [
	{ key: COLUMNS.KEYS.PANE, index: -2 },
	{ key: COLUMNS.KEYS.HOME, index: -1 },
	{ key: COLUMNS.KEYS.DISCOVER, index: 0 },
	{ key: COLUMNS.KEYS.MY_PLUGS, index: 1 },
	{ key: COLUMNS.KEYS.ACTIVITY, index: 2 },
	{ key: COLUMNS.KEYS.TOKENS, index: 3 }
].map(column => ({
	...column,
	id: Math.random() * 1e18,
	width: COLUMNS.DEFAULT_WIDTH
}))

// Individual column atoms
export const columnAtomsMapAtom = atomWithStorage<Record<number, Column>>(
	"plug.columns.map",
	DEFAULT_COLUMNS.reduce((acc, col) => ({ ...acc, [col.id]: col }), {})
)

// Column order atom (just IDs)
export const columnOrderAtom = atomWithStorage<number[]>(
	"plug.columns.order",
	DEFAULT_COLUMNS.map(col => col.id)
)

// Index to ID mapping for faster lookups
export const columnIndexMapAtom = atom(get => {
	const columnsMap = get(columnAtomsMapAtom)
	return Object.values(columnsMap).reduce(
		(acc, col) => {
			acc[col.index] = col.id
			return acc
		},
		{} as Record<number, number>
	)
})

// Derived atom for all columns in order
const columnsAtom = atom(get => {
	const order = get(columnOrderAtom)
	const columnsMap = get(columnAtomsMapAtom)
	return order.map(id => columnsMap[id]).filter(Boolean)
})

// Visible columns atom (filtered and sorted)
const visibleColumnsAtom = atom(get => {
	const columns = get(columnsAtom)
	return columns.filter(column => column?.index >= 0).sort((a, b) => a.index - b.index)
})

// Visible column IDs atom (just the order)
const visibleColumnIdsAtom = atom(get => {
	const order = get(columnOrderAtom)
	const columnsMap = get(columnAtomsMapAtom)
	return order
		.filter(id => columnsMap[id]?.index >= 0)
		.sort((a, b) => (columnsMap[a]?.index ?? 0) - (columnsMap[b]?.index ?? 0))
})

// Create a stable handlers atom that doesn't change reference
const handlersAtom = atom(
	null,
	(
		get,
		set,
		action: {
			type: "add" | "remove" | "move" | "resize" | "frame" | "navigate" | "schedule" | "transfer"
			payload: any
		}
	) => {
		const columnsMap = get(columnAtomsMapAtom)
		const indexMap = get(columnIndexMapAtom)

		const getColumnByIndex = (index: number) => {
			const id = indexMap[index]
			return id !== undefined ? columnsMap[id] : undefined
		}

		switch (action.type) {
			case "add": {
				const { index, key, from, item } = action.payload
				const newColumn = {
					id: Math.floor(Math.random() * 1e18),
					key,
					index: index ?? Object.values(columnsMap).length - COLUMNS.OFFSET,
					from,
					item,
					width: COLUMNS.DEFAULT_WIDTH
				} as Column

				set(columnAtomsMapAtom, { ...columnsMap, [newColumn.id]: newColumn })
				set(columnOrderAtom, prev => {
					const insertIndex = index !== undefined ? index + COLUMNS.OFFSET : prev.length
					const newOrder = [...prev]
					newOrder.splice(insertIndex, 0, newColumn.id)
					return newOrder
				})
				break
			}

			case "remove": {
				const { index } = action.payload
				const targetColumn = getColumnByIndex(index)
				if (!targetColumn) return

				const { [targetColumn.id]: _, ...rest } = columnsMap
				set(columnAtomsMapAtom, rest)
				set(columnOrderAtom, prev => prev.filter(id => id !== targetColumn.id))

				// Update indices
				set(columnAtomsMapAtom, prev => {
					const updated = { ...prev }
					Object.values(updated).forEach(col => {
						if (col.index > index) {
							updated[col.id] = { ...col, index: col.index - 1 }
						}
					})
					return updated
				})
				break
			}

			case "move": {
				const { from, to } = action.payload
				if (from === to) return

				set(columnOrderAtom, prev => {
					const newOrder = [...prev]
					const [moved] = newOrder.splice(from + COLUMNS.OFFSET, 1)
					newOrder.splice(to + COLUMNS.OFFSET, 0, moved)
					return newOrder
				})

				// Update indices
				set(columnAtomsMapAtom, prev => {
					const updated = { ...prev }
					Object.values(updated).forEach(col => {
						const newIndex =
							col.index === from
								? to
								: col.index > from && col.index <= to
									? col.index - 1
									: col.index < from && col.index >= to
										? col.index + 1
										: col.index
						updated[col.id] = { ...col, index: newIndex }
					})
					return updated
				})
				break
			}

			case "resize": {
				const { index, width } = action.payload
				const targetColumn = getColumnByIndex(index)
				if (!targetColumn || targetColumn.width === width) return

				set(columnAtomsMapAtom, {
					...columnsMap,
					[targetColumn.id]: { ...targetColumn, width }
				})
				break
			}

			case "frame": {
				const { index, key } = action.payload
				const targetColumn = getColumnByIndex(index)
				if (!targetColumn) return

				const newFrame = targetColumn.frame === key ? undefined : key

				set(columnAtomsMapAtom, {
					...columnsMap,
					[targetColumn.id]: { ...targetColumn, frame: newFrame }
				})
				break
			}

			case "navigate": {
				const { index, key, item, from } = action.payload
				const targetColumn = getColumnByIndex(index)
				if (!targetColumn) return

				if (targetColumn.key === key && targetColumn.item === item && targetColumn.from === from) return

				set(columnAtomsMapAtom, {
					...columnsMap,
					[targetColumn.id]: { ...targetColumn, key: key ?? targetColumn.key, item, from }
				})
				break
			}

			case "schedule": {
				const { index, schedule } = action.payload
				const targetColumn = getColumnByIndex(index)
				if (!targetColumn || targetColumn.schedule === schedule) return

				set(columnAtomsMapAtom, {
					...columnsMap,
					[targetColumn.id]: { ...targetColumn, schedule }
				})
				break
			}

			case "transfer": {
				const { index, updater } = action.payload
				const targetColumn = getColumnByIndex(index)
				if (!targetColumn) return

				const newTransfer = typeof updater === "function" ? updater(targetColumn.transfer) : updater
				if (targetColumn.transfer === newTransfer) return

				set(columnAtomsMapAtom, {
					...columnsMap,
					[targetColumn.id]: { ...targetColumn, transfer: newTransfer }
				})
				break
			}
		}
	}
)

// Hook to get a column and its stable handlers
export const useColumnStore = (index?: number, key?: string) => {
	const [columnsMap] = useAtom(columnAtomsMapAtom)
	const indexMap = useAtomValue(columnIndexMapAtom)
	const [, dispatch] = useAtom(handlersAtom)

	const column = useMemo(
		() => (index !== undefined ? columnsMap[indexMap[index]] : undefined),
		[columnsMap, indexMap, index]
	)

	const isFrame = useMemo(() => {
		if (!column || !key) return false
		return column.frame === key
	}, [column, key])

	const handle = useMemo(
		() => ({
			add: (payload: Partial<Column>) => dispatch({ type: "add", payload }),
			remove: (index: number) => dispatch({ type: "remove", payload: { index } }),
			move: (payload: { from: number; to: number }) => dispatch({ type: "move", payload }),
			resize: (payload: { index: number; width: number }) => dispatch({ type: "resize", payload }),
			frame: (key?: string) => {
				if (index === undefined) return
				dispatch({ type: "frame", payload: { index, key } })
			},
			navigate: (payload: Partial<Column> & { index: number }) => dispatch({ type: "navigate", payload }),
			schedule: (schedule?: Schedule) => dispatch({ type: "schedule", payload: { index, schedule } }),
			transfer: (updater: Transfer | undefined | ((prev: Transfer | undefined) => Transfer)) =>
				dispatch({ type: "transfer", payload: { index, updater } })
		}),
		[dispatch, index]
	)

	return { column, isFrame, handle }
}

export const useColumnList = () => {
	const columnIds = useAtomValue(visibleColumnIdsAtom)
	const [, dispatch] = useAtom(handlersAtom)

	const handle = useMemo(
		() => ({
			move: (payload: { from: number; to: number }) => dispatch({ type: "move", payload })
		}),
		[dispatch]
	)

	return {
		columnIds,
		handle
	}
}

export const useColumn = (id: number) => {
	const [columnsMap] = useAtom(columnAtomsMapAtom)
	return useMemo(() => columnsMap[id], [columnsMap, id])
}
