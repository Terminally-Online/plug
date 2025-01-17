import React from "react"

import { DragDropContext, Droppable, DropResult } from "@hello-pangea/dnd"

import { ConsoleColumn } from "@/components/app/columns/column"
import { ColumnAdd } from "@/components/app/columns/utils/column-add"
import { useColumnList } from "@/state/columns"

import { ConsoleSidebarPane } from "../sidebar"

// Memoized column component that only re-renders when necessary
const MemoizedColumn = React.memo(({ id, index }: { id: number; index: number }) => <ConsoleColumn id={index} />)
MemoizedColumn.displayName = "MemoizedColumn"

export const ConsoleColumnRow = () => {
	const { columnIds, handle } = useColumnList()

	const onDragEnd = (result: DropResult) => {
		if (!result.destination) return

		handle.move({
			from: result.source.index,
			to: result.destination.index
		})
	}

	return (
		<div className="flex h-full flex-row overflow-x-auto overflow-y-hidden">
			<ConsoleSidebarPane />

			<DragDropContext onDragEnd={onDragEnd}>
				<Droppable droppableId="droppable" direction="horizontal">
					{provided => (
						<div ref={provided.innerRef} className="flex flex-row" {...provided.droppableProps}>
							{columnIds.map((id, index) => (
								<MemoizedColumn key={id} id={id} index={index} />
							))}
							{provided.placeholder}
						</div>
					)}
				</Droppable>
			</DragDropContext>

			<ColumnAdd index={columnIds.length - 2} />
		</div>
	)
}
