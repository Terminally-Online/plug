import { useCallback } from "react"

import { atom, useAtom } from "jotai"

import { Column, Schedule, Transfer } from "@/lib"

import { atomFamily, atomWithStorage } from "jotai/utils"

export const COLUMNS = {
	SIDEBAR_INDEX: -2,
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
}

export const DEFAULT_COLUMNS = [
	{ key: COLUMNS.KEYS.PANE, index: COLUMNS.SIDEBAR_INDEX },
	{ key: COLUMNS.KEYS.HOME, index: COLUMNS.MOBILE_INDEX },
	{ key: COLUMNS.KEYS.DISCOVER, index: 0 },
	{ key: COLUMNS.KEYS.MY_PLUGS, index: 1 },
	{ key: COLUMNS.KEYS.ACTIVITY, index: 2 },
	{ key: COLUMNS.KEYS.TOKENS, index: 3 }
].map(column => ({
	...column,
	id: Math.random() * 1e18,
	width: COLUMNS.DEFAULT_WIDTH
}))

export const columnsStorageAtom = atomWithStorage<Column[]>("plug.columns", DEFAULT_COLUMNS)
export const primaryColumnsAtom = atom(
	get => get(columnsStorageAtom),
	(get, set, update: Column[] | ((prev: Column[]) => Column[])) => {
		set(columnsStorageAtom, typeof update === "function" ? update(get(columnsStorageAtom)) : update)
	}
)
const columnIndexMapAtom = atom(get => {
	const columns = get(columnsStorageAtom)
	return new Map(columns.map(col => [col.index, col]))
})
export const columnByIndexAtom = atomFamily((index: number) => atom(get => get(columnIndexMapAtom).get(index)))
export const isFrameAtom = atom(_ => (column?: Column, key?: string) => column?.frame === key)

export const addColumnAtom = atom(null, (get, set, { index, key, from, item }: Partial<Column>) => {
	const columns = get(primaryColumnsAtom)

	const newColumn = {
		id: Math.random() * 1e18,
		key,
		index: index ?? columns[columns.length - 1].index + 1,
		from,
		item,
		width: COLUMNS.DEFAULT_WIDTH
	} as Column

	set(primaryColumnsAtom, prev => {
		const updatedColumns = [...prev]
		const insertIndex = index !== undefined ? index + COLUMNS.OFFSET : updatedColumns.length
		updatedColumns.splice(insertIndex, 0, newColumn)

		return updatedColumns.map((col, idx) => ({
			...col,
			index: idx - COLUMNS.OFFSET
		}))
	})
})

export const navigateColumnAtom = atom(
	null,
	(_, set, { index, key, item, from }: Partial<Column> & { index: number }) => {
		set(primaryColumnsAtom, prev => {
			const targetColumn = prev.find(col => col.index === index)
			if (
				!targetColumn ||
				(targetColumn.key === key && targetColumn.item === item && targetColumn.from === from)
			) {
				return prev
			}
			return prev.map(col => (col.index === index ? { ...col, key: key ?? col.key, item, from } : col))
		})
	}
)

export const removeColumnAtom = atom(null, (_, set, index: number) => {
	set(primaryColumnsAtom, prev => {
		const updatedColumns = [...prev]
		updatedColumns.splice(index + COLUMNS.OFFSET, 1)
		return updatedColumns.map((col, idx) => ({
			...col,
			index: idx - COLUMNS.OFFSET
		}))
	})
})

export const moveColumnAtom = atom(null, (_, set, { from, to }: { from: number; to: number }) => {
	if (from === to) return

	set(primaryColumnsAtom, prev => {
		const updatedColumns = [...prev]
		const [movedColumn] = updatedColumns.splice(from + COLUMNS.OFFSET, 1)
		updatedColumns.splice(to + COLUMNS.OFFSET, 0, movedColumn)
		return updatedColumns.map((col, idx) => ({
			...col,
			index: idx - COLUMNS.OFFSET
		}))
	})
})

export const resizeColumnAtom = atom(null, (_, set, { index, width }: { index: number; width: number }) => {
	set(primaryColumnsAtom, prev => {
		const targetColumn = prev.find(col => col.index === index)
		if (!targetColumn || targetColumn.width === width) return prev
		return prev.map(col => (col.index === index ? { ...col, width } : col))
	})
})

export const frameColumnAtom = atom(
	null,
	(_, set, { index, key, columnKey }: { index: number; key?: string; columnKey?: string }) => {
		set(primaryColumnsAtom, prev => {
			const targetColumn = prev.find(col => col.index === index || col.id === index)
			if (!targetColumn) return prev

			const frameKey = columnKey || key
			const newFrame = targetColumn.frame === frameKey ? undefined : frameKey

			if (targetColumn.frame === newFrame) return prev

			return prev.map(col => (col.index === index ? { ...col, frame: newFrame } : col))
		})
	}
)

export const chainColumnAtom = atom(null, (_, set, { index, chain }: { index: number; chain?: number }) => {
	set(primaryColumnsAtom, prev => prev.map(col => (col.index === index ? { ...col, chain } : col)))
})

export const scheduleColumnAtom = atom(null, (_, set, { index, schedule }: { index: number; schedule?: Schedule }) => {
	set(primaryColumnsAtom, prev => prev.map(col => (col.index === index ? { ...col, schedule } : col)))
})

export const transferColumnAtom = atom(
	null,
	(
		_,
		set,
		{
			index,
			updater
		}: {
			index: number
			updater: Transfer | undefined | ((prev: Transfer | undefined) => Transfer)
		}
	) => {
		set(primaryColumnsAtom, prev => {
			const targetColumn = prev.find(col => col.index === index)
			if (!targetColumn) return prev

			const newTransfer = typeof updater === "function" ? updater(targetColumn.transfer) : updater

			if (targetColumn.transfer === newTransfer) return prev

			return prev.map(col => (col.index === index ? { ...col, transfer: newTransfer } : col))
		})
	}
)

export const useScrollToLastColumn = () => {
	return useCallback(() => {
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
}

export const useColumnActions = (index?: number, key?: string) => {
	const [, addColumn] = useAtom(addColumnAtom)
	const [, navigateColumn] = useAtom(navigateColumnAtom)
	const [, removeColumn] = useAtom(removeColumnAtom)
	const [, moveColumn] = useAtom(moveColumnAtom)
	const [, resizeColumn] = useAtom(resizeColumnAtom)
	const [, frameColumn] = useAtom(frameColumnAtom)
	const [, chainColumn] = useAtom(chainColumnAtom)
	const [, scheduleColumn] = useAtom(scheduleColumnAtom)
	const [, transferColumn] = useAtom(transferColumnAtom)

	const handleScroll = useScrollToLastColumn()

	return {
		add: useCallback(
			(params: Partial<Column>) => {
				addColumn(params)
				handleScroll()
			},
			[addColumn, handleScroll]
		),
		navigate: useCallback(
			(params: Partial<Column> & { index: number }) => navigateColumn(params),
			[navigateColumn]
		),
		remove: useCallback((idx: number) => removeColumn(idx), [removeColumn]),
		move: useCallback((params: { from: number; to: number }) => moveColumn(params), [moveColumn]),
		resize: useCallback((params: { index: number; width: number }) => resizeColumn(params), [resizeColumn]),
		frame: useCallback(
			(columnKey?: string) => {
				if (index === undefined) return
				frameColumn({ index, key, columnKey })
			},
			[frameColumn, index, key]
		),
		chain: useCallback(
			(chain?: number) => {
				if (index === undefined) return
				chainColumn({ index, chain })
			},
			[chainColumn, index]
		),
		schedule: useCallback(
			(schedule?: Schedule) => {
				if (index === undefined) return
				scheduleColumn({ index, schedule })
			},
			[scheduleColumn, index]
		),
		transfer: useCallback(
			(updater: Transfer | undefined | ((prev: Transfer | undefined) => Transfer)) => {
				if (index === undefined) return
				transferColumn({ index, updater })
			},
			[transferColumn, index]
		)
	}
}
