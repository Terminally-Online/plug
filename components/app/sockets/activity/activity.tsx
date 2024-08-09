import { FC, HTMLAttributes } from "react"

import { FileCog } from "lucide-react"

import { ActivityList, Header } from "@/components"

export const SocketActivity: FC<
	HTMLAttributes<HTMLDivElement> & { id: string }
> = ({ id, ...props }) => {
	return (
		<div {...props}>
			<Header
				size="md"
				icon={<FileCog size={14} className="opacity-40" />}
				label="Runs"
			/>

			<ActivityList id={id} />
		</div>
	)
}
