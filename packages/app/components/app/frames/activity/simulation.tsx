import { FC } from "react"

import { Frame } from "@/components/app/frames/base"
import { RouterOutputs } from "@/server/client"
import { useColumnStore } from "@/state/columns"

export const SimulationFrame: FC<{
	index: number
	icon: JSX.Element
	simulation: RouterOutputs["plugs"]["activity"]["get"][number]["simulations"][number]
}> = ({ index, icon, simulation }) => {
	const { isFrame } = useColumnStore(index, `${simulation.id}-simulation`)

	// const actions = useMemo(() => JSON.parse(activity.actions), [activity])

	return <Frame index={index} icon={icon} label={simulation.id} visible={isFrame} hasOverlay={true}></Frame>
}
