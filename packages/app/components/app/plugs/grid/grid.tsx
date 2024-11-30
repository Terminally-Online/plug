import { FC } from "react"

import { Workflow } from "@prisma/client"

import { PlugGridItem } from "@/components"

export const PlugGrid: FC<
	React.HTMLAttributes<HTMLDivElement> & {
		index: number
		from: string
		plugs: Array<Workflow | undefined> | undefined
		count?: number
	}
> = ({ index, from, plugs, count, ...props }) => {
	if (plugs === undefined) return null

	return (
		<div {...props}>
			<div
				className="grid gap-2"
				style={{
					gridTemplateColumns: `repeat(auto-fit, minmax(220px, 1fr))`
				}}
			>
				{plugs.slice(0, count || plugs.length).map((plug, plugIndex) => (
					<PlugGridItem key={plugIndex} index={index} from={from} plug={plug} />
				))}
			</div>
		</div>
	)
}
