import { FC, useMemo } from "react"

import BoringAvatar from "boring-avatars"

import { useSocket } from "@/state"

const breakpoints = 4
const angles = 360
const angle = angles / breakpoints

export const Avatar: FC<{ name: string; rotation?: number }> = ({ name, rotation }) => {
	const { socket } = useSocket()

	const hueRotation = useMemo(() => {
		if (socket && socket.id === name) return 0
		if (rotation !== undefined) return rotation
		const seed = Math.abs(name.split("").reduce((hash, char) => (hash << 5) - hash + char.charCodeAt(0), 0))
		return Math.floor((seed % angles) / angle) * angle
	}, [name, socket])

	return (
		<div className="relative overflow-hidden rounded-sm" style={{ filter: `hue-rotate(${hueRotation}deg)` }}>
			<BoringAvatar variant="beam" name={name} size={"100%"} colors={["#00E100", "#A3F700"]} square />
		</div>
	)
}
