import { useEffect } from "react"

import BlockiesSvg from "blockies-react-svg"
import { ChevronRight, Users } from "lucide-react"

import { Button, Frame } from "@/components"
import { useFrame, useSockets } from "@/contexts"

export const SocketFrame = () => {
	const { frameVisible, handleFrameVisible } = useFrame()
	const { socket, sockets, handleSelect } = useSockets()

	const isFrame = frameVisible
		? frameVisible.split("-")[0] === "socket"
		: false

	const nextFrame = frameVisible
		? `chain-${frameVisible.split("-")[1]}`
		: undefined

	const handleSocketSelect = (socketAddress: string) => {
		handleSelect(socketAddress)
		handleFrameVisible(nextFrame)
	}

	useEffect(() => {
		if (socket === undefined && sockets && sockets.length === 1) {
			handleSelect(sockets[0].socketAddress)
		}

		if (isFrame) handleFrameVisible(nextFrame)
	}, [socket, sockets, isFrame, nextFrame, handleSelect, handleFrameVisible])

	return (
		<Frame
			className="z-[2]"
			icon={<Users size={18} />}
			label="Choose Socket"
			visible={isFrame}
		>
			<div className="flex flex-col gap-4">
				{sockets && sockets.length > 0 ? (
					sockets.map((socketInMap, index) => (
						<div
							key={index}
							className="group flex cursor-pointer flex-row items-center gap-4"
							onClick={() =>
								handleSocketSelect(socketInMap.socketAddress)
							}
						>
							<BlockiesSvg
								address={socketInMap.socketAddress}
								className="h-6 w-6 rounded-md"
							/>
							<p className="mr-auto font-bold">
								{socketInMap.name}
							</p>

							<Button
								variant="secondary"
								className="ml-auto p-1 group-hover:bg-grayscale-100"
								onClick={() =>
									handleSocketSelect(
										socketInMap.socketAddress
									)
								}
							>
								<ChevronRight size={14} />
							</Button>
						</div>
					))
				) : (
					<p>TODO</p>
				)}
			</div>
		</Frame>
	)
}
