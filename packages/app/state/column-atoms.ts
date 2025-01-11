import { atom } from "jotai"

import { Column } from "@/lib"

import { atomFamily, atomWithStorage } from "jotai/utils"

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

const DEFAULT_COLUMNS = [
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

// Base storage atom
const columnsStorageAtom = atomWithStorage<Column[]>("plug.columns", DEFAULT_COLUMNS)

// Individual column atom family
const columnAtomFamily = atomFamily((id: number) =>
	atom(
		get => get(columnsStorageAtom).find(col => col.id === id),
		(get, set, update: Partial<Column>) => {
			set(columnsStorageAtom, prev => prev.map(col => (col.id === id ? { ...col, ...update } : col)))
		}
	)
)

// Visible columns order atom
const columnOrderAtom = atom(get => {
	const columns = get(columnsStorageAtom)
	return columns
		.filter(col => col.index >= 0)
		.sort((a, b) => a.index - b.index)
		.map(col => col.id)
})

// Column operations atom
const columnOpsAtom = atom(
	null,
	(
		get,
		set,
		action: {
			type: "add" | "remove" | "move" | "update"
			payload: any
		}
	) => {
		const columns = get(columnsStorageAtom)

		switch (action.type) {
			case "add": {
				const { key, index, from, item } = action.payload
				const newColumn = {
					id: Math.random() * 1e18,
					key,
					index: index ?? columns.length - COLUMNS.OFFSET,
					from,
					item,
					width: COLUMNS.DEFAULT_WIDTH
				}
				set(columnsStorageAtom, [...columns, newColumn])
				break
			}

			case "remove": {
				const { id } = action.payload
				set(
					columnsStorageAtom,
					columns.filter(col => col.id !== id)
				)
				break
			}

			case "move": {
				const { fromIndex, toIndex } = action.payload
				const reordered = [...columns]
				const [moved] = reordered.splice(fromIndex + COLUMNS.OFFSET, 1)
				reordered.splice(toIndex + COLUMNS.OFFSET, 0, moved)
				set(
					columnsStorageAtom,
					reordered.map((col, idx) => ({
						...col,
						index: idx - COLUMNS.OFFSET
					}))
				)
				break
			}

			case "update": {
				const { id, ...changes } = action.payload
				set(
					columnsStorageAtom,
					columns.map(col => (col.id === id ? { ...col, ...changes } : col))
				)
				break
			}
		}
	}
)

export { columnAtomFamily, columnOrderAtom, columnOpsAtom, columnsStorageAtom }
