import { FC } from "react"

import { Workflow } from "@prisma/client"

import { Animate, PlugGridItem } from "@/components"

type Props = React.HTMLAttributes<HTMLDivElement> & {
	index: number
	from: string
	plugs: Array<Workflow | undefined> | undefined
	count?: number
}

export const PlugGrid: FC<Props> = ({ index, from, plugs, count, ...props }) => {
	if (plugs === undefined) return null

	return (
		<div {...props}>
			<Animate.Grid>
				{plugs.slice(0, count || plugs.length).map((plug, plugIndex) => (
					<Animate.ListItem key={plugIndex}>
						<PlugGridItem index={index} from={from} plug={plug} />
					</Animate.ListItem>
				))}
			</Animate.Grid>
		</div>
	)
}
