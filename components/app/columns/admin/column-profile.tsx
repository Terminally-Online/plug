import { FC, HTMLAttributes } from "react"

import { useColumns, useSocket } from "@/state"

export const ColumnProfile: FC<HTMLAttributes<HTMLDivElement>> = ({ ...props }) => {
	const { socket } = useSocket()
	const { columns } = useColumns()

	return (
		<div {...props}>
			<pre className="text-xs">{JSON.stringify(socket, null, 2)}</pre>
			<pre className="text-xs">{JSON.stringify(columns, null, 2)}</pre>
		</div>
	)
}
