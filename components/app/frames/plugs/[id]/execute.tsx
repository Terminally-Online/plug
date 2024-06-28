import { useEffect } from "react"

import Image from "next/image"

import BlockiesSvg from "blockies-react-svg"
import { CheckCircle, Eye, LoaderCircle } from "lucide-react"

import { Frame } from "@/components/app/frames/base"
import { ActionPreview } from "@/components/app/plugs/actions"
import { Button } from "@/components/buttons"
import { useFrame, usePlugs, useSockets } from "@/contexts"

import { ChainFrame } from "./execute/chain"
import { RanFrame } from "./execute/ran"
import { RunFrame } from "./execute/run"
import { RunningFrame } from "./execute/running"
import { ScheduleFrame } from "./execute/schedule"
import { SocketFrame } from "./execute/socket"

export const ExecuteFrame = () => {
	const { socket } = useSockets()
	const { chains, plug } = usePlugs()
	const { frameVisible, handleFrameVisible } = useFrame()

	if (!plug) return null

	return (
		<>
			<SocketFrame />
			<ChainFrame />
			<ScheduleFrame />

			<RunFrame />
			<RunningFrame />
			<RanFrame />
		</>
	)
}
