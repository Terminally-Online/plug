import { AnimatePresence } from "framer-motion"

import { DragDropContext, Droppable, DropResult } from "@hello-pangea/dnd"

import { ColumnAdd, ConsoleColumn } from "@/components"
import { MOBILE_INDEX, useColumns } from "@/state"

export const ConsoleColumnRow = () => {
	const { columns, move } = useColumns()

	const onDragEnd = (result: DropResult) => {
		if (!columns || !result.destination) return

		move({
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
							<AnimatePresence>
								{columns
									.filter(column => column.index !== MOBILE_INDEX)
									.sort((a, b) => a.index - b.index)
									.map(column => (
										<ConsoleColumn key={column.index} column={column} />
									))}
							</AnimatePresence>
							{provided.placeholder}
						</div>
					)}
				</Droppable>
			</DragDropContext>

			<ColumnAdd />
		</div>
	)
}
