import { DragDropContext, Droppable, DropResult } from "@hello-pangea/dnd"

import { ColumnAdd, ConsoleColumn } from "@/components"
import { useColumnStore } from "@/state/columns"

export const ConsoleColumnRow = () => {
	const { columns, handle } = useColumnStore()

	const onDragEnd = (result: DropResult) => {
		if (!columns || !result.destination) return

		handle.move({
			from: result.source.index,
			to: result.destination.index
		})
	}

	return (
		<div className="flex h-full flex-row overflow-x-auto overflow-y-hidden">
			<DragDropContext onDragEnd={onDragEnd}>
				<Droppable droppableId="droppable" direction="horizontal">
					{provided => (
						<div ref={provided.innerRef} className="flex flex-row" {...provided.droppableProps}>
							{columns
								.filter(column => column?.index >= 0)
								.sort((a, b) => a.index - b.index)
								.map(column => (
									<ConsoleColumn key={String(column.id)} id={column.index} />
								))}
							{provided.placeholder}
						</div>
					)}
				</Droppable>
			</DragDropContext>

			<ColumnAdd />
		</div>
	)
}
