import Link from "next/link"

import { ExternalLink } from "lucide-react"

import { useAtom, useAtomValue } from "jotai"

import { DragDropContext, Droppable, DropResult } from "@hello-pangea/dnd"

import { ConsoleColumn } from "@/components/app/columns/column"
import { ColumnAdd } from "@/components/app/columns/utils/column-add"
import { moveColumnAtom, primaryColumnsAtom } from "@/state/columns"

import { ConsoleSidebarPane } from "../sidebar"

export const ConsoleColumnRow = () => {
	const columns = useAtomValue(primaryColumnsAtom)

	const [, move] = useAtom(moveColumnAtom)

	const onDragEnd = (result: DropResult) => {
		if (!columns || !result.destination) return

		move({
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
							{columns
								.filter(column => column?.index >= 0)
								.sort((a, b) => a.index - b.index)
								.map(column => (
									<ConsoleColumn key={String(column.id)} index={column.index} />
								))}
							{provided.placeholder}
						</div>
					)}
				</Droppable>
			</DragDropContext>

			<ColumnAdd index={columns.length - 2} />

			<div className="mx-4 my-2 mr-48 flex w-max flex-col items-start justify-end font-bold">
				<Link
					href="/terms"
					className="group flex flex-row items-center gap-1 whitespace-nowrap transition-all duration-200 ease-in-out hover:opacity-100"
				>
					<span className="opacity-40 group-hover:opacity-100">Terms of Service</span>
					<ExternalLink size={14} className="opacity-40 group-hover:opacity-100" />
				</Link>
				<Link
					href="/privacy"
					className="group flex flex-row items-center gap-1 whitespace-nowrap transition-all duration-200 ease-in-out hover:opacity-100"
				>
					<span className="opacity-40 group-hover:opacity-100">Privacy Policy</span>
					<ExternalLink size={14} className="opacity-40 group-hover:opacity-100" />
				</Link>
				<p className="whitespace-nowrap opacity-40">A ⚫ Terminally Online project.</p>
			</div>
		</div>
	)
}
