import { useEffect, useState } from "react"

import { AnimatePresence } from "framer-motion"

import { DragDropContext, Droppable, DropResult } from "@hello-pangea/dnd"

import { ColumnAdd, ConsoleColumn } from "@/components"
import { useSockets } from "@/contexts"

// This component has a bunch of hacky workarounds to make it work due to the
// complexities of TRPC combined with commonly requested animation frames. The
// resulting synchronization loss is not perfect resulting in a flash of the
// previous state before onMutate fires.
//
// Right now we lose the ability to have all of our state logic within the context.
// Thankfully, column indexes are not used anywhere but in this component so it
// is not a huge deal. Even in this state though, if the request to the server
// errors out, the state will be reverted back to the socket state. Even then,
// realistically, there should never be a server error since it's a pretty a
// pretty minimal endpoint that only handles the reordering of columns.
//
// Nothing happens for this state in the columns themselves because they are
// simply rendered in the order they are in the array and all drag events live
// on the Draggable component.
//
// If there is a better way to do this and manage it in an onMutate callback,
// please let me know
//
// - CHANCE (08/12/2024)
export const ConsoleColumnRow = () => {
	const { socket, handle } = useSockets()

	const [columns, setColumns] = useState<any[]>(socket?.columns || [])

	const onDragEnd = (result: DropResult) => {
		if (!columns || !result.destination) return

		// This is done here instead of onMutate because there is a perceivable
		// delay between the drag finishing and the onMutate callback being
		// fired resulting in a unintended "flash" even though the state change
		// is already done and propagating. There is no way to prevent this
		// without a hacky workaround so we just manage the state here instead.
		handle.columns.move({
			from: result.source.index,
			to: result.destination.index
		})

		const reorderedColumns = columns?.sort((a, b) => a.index - b.index).slice(1, columns.length) ?? []

		const [removed] = reorderedColumns.splice(result.source.index, 1)
		reorderedColumns.splice(result.destination.index, 0, removed)

		setColumns([
			columns[0],
			...reorderedColumns.map((column, index) => ({
				...column,
				index: index
			}))
		])
	}

	// This supports the local state management of the column dragging since
	// the other endpoints update the state of the columns in the context.
	// This means that if one is added, deleted, or navigated within, the
	// state will be updated while still maintaining the ability to have local
	// state here that instantly updates the UI without a perceivable delay.
	useEffect(() => {
		if (socket === undefined) return
		setColumns(socket?.columns)
	}, [socket])

	return (
		<div className="flex h-full flex-row overflow-x-auto overflow-y-hidden">
			<DragDropContext onDragEnd={onDragEnd}>
				<Droppable droppableId="droppable" direction="horizontal">
					{provided => (
						<div ref={provided.innerRef} className="flex flex-row" {...provided.droppableProps}>
							<AnimatePresence>
								{columns
									.filter(column => column.index !== -1)
									.sort((a, b) => a.index - b.index)
									.map(column => (
										<ConsoleColumn key={column.id} column={column} />
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
