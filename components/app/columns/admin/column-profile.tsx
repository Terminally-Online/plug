import { FC, HTMLAttributes } from "react"

import { useSockets } from "@/contexts"

export const ColumnProfile: FC<HTMLAttributes<HTMLDivElement>> = ({ ...props }) => {
	const { socket } = useSockets()

	return (
		<div {...props}>
			<pre className="text-xs">{JSON.stringify(socket, null, 2)}</pre>
		</div>
	)
}
