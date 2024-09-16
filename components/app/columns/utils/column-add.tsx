import { useMemo } from "react"

import { Plus } from "lucide-react"

import { Button } from "@/components"
import { useColumns } from "@/state"

export const ColumnAdd = () => {
	const { columns, add } = useColumns()

	const isAdding = useMemo(() => {
		return columns.find(column => column.key === "ADD") !== undefined
	}, [columns])

	if (isAdding) return null

	return (
		<div
			className="relative my-2 mr-2 flex select-none items-center justify-center rounded-lg border-[1px] border-grayscale-100 bg-grayscale-0 p-4"
			style={{ minWidth: `${380}px` }}
		>
			<Button
				variant="secondary"
				className="flex flex-row items-center gap-2 font-bold"
				onClick={() => add({ key: "ADD" })}
			>
				<Plus size={14} />
				Add Column
			</Button>
		</div>
	)
}
