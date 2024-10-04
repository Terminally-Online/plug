import { AnimatePresence } from "framer-motion"

import { DragDropContext, Droppable, DropResult } from "@hello-pangea/dnd"

import { ColumnAdd, ConsoleColumn } from "@/components"
import { cn, MOBILE_INDEX } from "@/lib"
import { useColumns, useSidebar } from "@/state"

export const ConsoleColumnRow = () => {
	const { is } = useSidebar()
	const { columns, move } = useColumns()

	const onDragEnd = (result: DropResult) => {
		if (!columns || !result.destination) return

		move({
			from: result.source.index,
			to: result.destination.index
		})
	}

	return (
		<div
			className={cn(
				"flex h-full flex-row overflow-x-auto overflow-y-hidden",
				(is.searching || is.viewingAs) === false && "ml-2"
			)}
		>
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
