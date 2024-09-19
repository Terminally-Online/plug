import { useCallback, useMemo } from "react"

import { MinimalUserSocketModel } from "@/prisma/types"

import { useAtom } from "jotai"

import { Column, DEFAULT_VIEWS } from "@/lib"
import { useSocket } from "@/state"

import { atomWithStorage } from "jotai/utils"

const columnsAtom = atomWithStorage<Column[]>("socketColumns", DEFAULT_VIEWS)
export const useColumns = (index?: number) => {
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
				const newColumn = { key, index: index ?? prev.length, from, item } as Column
				const updatedColumns = [...prev]
				if (index !== undefined) {
					updatedColumns.splice(index, 0, newColumn)
				} else {
					updatedColumns.push(newColumn)
				}
				return updatedColumns.map((col, idx) => ({ ...col, index: idx - 1 }))
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
					// NOTE: Account for the mobile index always being -1.
					.map((col, idx) => ({ ...col, index: idx - 1 }))
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
				const [movedColumn] = updatedColumns.splice(from, 1)
				updatedColumns.splice(to, 0, movedColumn)
				return updatedColumns.map((col, idx) => ({ ...col, index: idx }))
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
			updateColumns(prev => prev.map((col, idx) => (idx === index ? { ...col, width } : col))),
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
				prev.map((col, idx) => (idx === index ? { ...col, viewAs: as === col.viewAs ? undefined : as } : col))
			),
		[updateColumns]
	)

	const column = useMemo(
		() => (index !== undefined ? columns.find(column => column.index === index) : undefined),
		[columns, index]
	)

	const isExternal = useMemo(() => {
		if (socket === undefined) return false

		return (column !== undefined && column.viewAs && column.viewAs.id !== socket.id) || false
	}, [column, socket])

	return {
		columns,
		column,
		isExternal,
		add,
		navigate,
		remove,
		move,
		resize,
		as
	}
}
