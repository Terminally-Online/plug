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

	const add = useCallback(
		({ key, index, from, item }: Partial<Column>) => {
			updateColumns(prev => {
				const newColumn = { key, index: index ?? prev.length, from, item } as Column
				const updatedColumns = [...prev]
				if (index !== undefined) {
					updatedColumns.splice(index, 0, newColumn)
					for (let i = index + 1; i < updatedColumns.length; i++) updatedColumns[i].index++
				} else {
					updatedColumns.push(newColumn)
				}
				return updatedColumns
			})
		},
		[updateColumns]
	)

	const navigate = useCallback(
		({ index, key, item, from }: Partial<Column> & { index: number }) =>
			updateColumns(prev =>
				prev.map((col, idx) => (idx === index ? { ...col, key: key ?? col.key, item, from } : col))
			),
		[updateColumns]
	)

	const remove = useCallback(
		(index: number) =>
			updateColumns(prev => prev.filter((_, idx) => idx !== index).map((col, idx) => ({ ...col, index: idx }))),
		[updateColumns]
	)

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

	const resize = useCallback(
		({ index, width }: { index: number; width: number }) =>
			updateColumns(prev => {
				const updatedColumns = [...prev]
				const [movedColumn] = updatedColumns.splice(index, 1)
				updatedColumns.splice(index, 0, { ...movedColumn, width })
				return updatedColumns.map((col, idx) => ({ ...col, index: idx }))
			}),
		[updateColumns]
	)

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
