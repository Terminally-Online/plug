"use client"

import { DragDropContext, Droppable, DropResult } from "@hello-pangea/dnd"

import { ConsoleColumn, ConsoleColumnAdd } from "@/components/console"
import { useSockets } from "@/contexts"

export const ConsoleColumnRow = () => {
	const { socket, handle } = useSockets()

	const onDragEnd = (result: DropResult) => {
		if (!result.destination) return

		handle.columns.move({
			from: result.source.index,
			to: result.destination.index
		})
	}

	if (socket === undefined) return null

	return (
		<div className="flex h-screen flex-row overflow-x-auto overflow-y-hidden">
			<DragDropContext onDragEnd={onDragEnd}>
				<Droppable droppableId="droppable" direction="horizontal">
					{provided => (
						<div
							ref={provided.innerRef}
							className="flex flex-row"
							{...provided.droppableProps}
						>
							{socket.columns
								.sort((a, b) => a.index - b.index)
								.map(column => (
									<ConsoleColumn
										key={column.index}
										column={column}
									/>
								))}
							{provided.placeholder}
						</div>
					)}
				</Droppable>
			</DragDropContext>

			<ConsoleColumnAdd />
		</div>
	)
}
