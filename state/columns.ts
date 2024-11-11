import { useCallback, useMemo } from "react"

import { atom, useAtom, useAtomValue } from "jotai"

import { Column, Schedule, Transfer } from "@/lib"

import { atomWithStorage, splitAtom } from "jotai/utils"

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
	} as const
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

const columnsStorageAtom = atomWithStorage<Column[]>("socketColumns", DEFAULT_COLUMNS)
const primaryColumnsAtom = atom(
	get => get(columnsStorageAtom),
	(get, set, update: Column[] | ((prev: Column[]) => Column[])) => {
		set(columnsStorageAtom, typeof update === "function" ? update(get(columnsStorageAtom)) : update)
	}
)

const columnAtomsAtom = splitAtom(columnsStorageAtom)

const columnByIndexAtom = atom(get => {
	const columnAtoms = get(columnAtomsAtom)
	return (searchIndex: number): Column | undefined => {
		for (const columnAtom of columnAtoms) {
			const column = get(columnAtom)
			if (column && (column.index === searchIndex || column.id === searchIndex)) {
				return column
			}
		}
		return undefined
	}
})

const isFrameAtom = atom(_ => (column?: Column, key?: string) => {
	if (!column) return false
	return column.frame === key
})

export const useColumnData = (index?: number) => {
	const getColumnByIndex = useAtomValue(columnByIndexAtom)
	const column = useMemo(() => (index !== undefined ? getColumnByIndex(index) : undefined), [index, getColumnByIndex])
	return { column }
}

export const useColumn = (index: number) => {
	const getColumnByIndex = useAtomValue(columnByIndexAtom)
	return useMemo(() => getColumnByIndex(index), [index, getColumnByIndex])
}

export const useColumnStore = (index?: number, key?: string) => {
	const [columns, setColumns] = useAtom(primaryColumnsAtom)
	const [columnAtoms, columnDispatch] = useAtom(columnAtomsAtom)
	const getColumnByIndex = useAtomValue(columnByIndexAtom)
	const isFrame = useAtomValue(isFrameAtom)

	// setColumns(DEFAULT_COLUMNS)

	const column = useMemo(() => (index !== undefined ? getColumnByIndex(index) : undefined), [index, getColumnByIndex])

	const handleScroll = useCallback(() => {
		queueMicrotask(() => {
			requestAnimationFrame(() => {
				const container = document.querySelector(".flex.h-full.flex-row.overflow-x-auto")
				if (container) {
					const lastColumnOffset = container.scrollWidth - COLUMNS.DEFAULT_WIDTH
					container.scrollTo({ left: lastColumnOffset, behavior: "smooth" })
				}
			})
		})
	}, [])

	const handle = {
		add: useCallback(
			({ index, key, from, item }: Partial<Column>) => {
				const newColumn = {
					id: Math.random() * 1e18,
					key,
					index: index ?? columns.length - COLUMNS.OFFSET,
					from,
					item,
					width: COLUMNS.DEFAULT_WIDTH
				} as Column

				// Find position based on array position instead
				const insertIndex = index !== undefined ? index + COLUMNS.OFFSET : undefined
				const targetAtom = insertIndex !== undefined ? columnAtoms[insertIndex] : undefined

				columnDispatch({
					type: "insert",
					value: newColumn,
					before: targetAtom
				})

				handleScroll()
			},
			[columns.length, columnAtoms, columnDispatch, handleScroll]
		),

		navigate: useCallback(
			({ index, key, item, from }: Partial<Column> & { index: number }) => {
				setColumns(prev => {
					const targetColumn = prev.find(col => col.index === index)
					if (!targetColumn) return prev

					if (targetColumn.key === key && targetColumn.item === item && targetColumn.from === from) {
						return prev
					}

					return prev.map(col => (col.index === index ? { ...col, key: key ?? col.key, item, from } : col))
				})
			},
			[setColumns]
		),

		remove: useCallback(
			(index: number) => {
				// Find by array position with OFFSET
				const targetAtom = columnAtoms[index + COLUMNS.OFFSET]
				if (!targetAtom) return

				columnDispatch({
					type: "remove",
					atom: targetAtom
				})

				// Update indices after removal
				setColumns(prev =>
					prev.map((col, idx) => ({
						...col,
						index: idx - COLUMNS.OFFSET
					}))
				)
			},
			[columnAtoms, columnDispatch, setColumns]
		),

		move: useCallback(
			({ from, to }: { from: number; to: number }) => {
				if (from === to) return

				// const fromAtom = columnAtoms[from + COLUMNS.OFFSET]
				// const toAtom = columnAtoms[to + COLUMNS.OFFSET]
				// if (!fromAtom || !toAtom) return

				// columnDispatch({
				// 	type: "move",
				// 	atom: fromAtom,
				// 	before: toAtom
				// })

				setColumns(prev => {
					const updatedColumns = [...prev]
					const [movedColumn] = updatedColumns.splice(from + COLUMNS.OFFSET, 1)
					updatedColumns.splice(to + COLUMNS.OFFSET, 0, movedColumn)
					return updatedColumns.map((col, idx) => ({ ...col, index: idx - COLUMNS.OFFSET }))
				})
			},
			[, setColumns]
		),

		resize: useCallback(
			({ index, width }: { index: number; width: number }) => {
				setColumns(prev => {
					const targetColumn = prev.find(col => col.index === index)
					if (!targetColumn || targetColumn.width === width) return prev

					return prev.map(col => (col.index === index ? { ...col, width } : col))
				})
			},
			[setColumns]
		),

		frame: useCallback(
			(columnKey?: string) => {
				if (!index) return

				setColumns(prev => {
					const targetColumn = prev.find(col => col.index === index)
					if (!targetColumn) return prev

					const frameKey = columnKey || key
					const newFrame = targetColumn.frame === frameKey ? undefined : frameKey

					if (targetColumn.frame === newFrame) return prev

					return prev.map(col => (col.index === index ? { ...col, frame: newFrame } : col))
				})
			},
			[index, key, setColumns]
		),

		schedule: useCallback(
			(schedule?: Schedule) => {
				if (!index) return

				setColumns(prev => {
					const targetColumn = prev.find(col => col.index === index)
					if (!targetColumn || targetColumn.schedule === schedule) return prev

					return prev.map(col => (col.index === index ? { ...col, schedule } : col))
				})
			},
			[index, setColumns]
		),

		transfer: useCallback(
			(updater: Transfer | undefined | ((prev: Transfer | undefined) => Transfer)) => {
				if (!index) return

				setColumns(prev => {
					const targetColumn = prev.find(col => col.index === index)
					if (!targetColumn) return prev

					const newTransfer = typeof updater === "function" ? updater(targetColumn.transfer) : updater

					if (targetColumn.transfer === newTransfer) return prev

					return prev.map(col => (col.index === index ? { ...col, transfer: newTransfer } : col))
				})
			},
			[index, setColumns]
		)
	}

	return {
		columns,
		column,
		isFrame: useMemo(() => isFrame(column, key), [isFrame, column, key]),
		handle
	}
}
