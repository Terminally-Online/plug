import { FC, HTMLAttributes } from "react"

import { FileCog } from "lucide-react"

import { ActivityList, Header } from "@/components"

export const SocketActivity: FC<HTMLAttributes<HTMLDivElement> & { id: string }> = ({ id, ...props }) => {
	return (
		<div {...props}>
			<ActivityList id={id} />
		</div>
	)
}
