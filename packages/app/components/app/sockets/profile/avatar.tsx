import { FC, HTMLAttributes, useMemo } from "react"

import BoringAvatar from "boring-avatars"

import { cn } from "@/lib"
import { useSocket } from "@/state/authentication"

const breakpoints = 12
const angles = 360
const angle = angles / breakpoints

// TODO(#413): This is a temporary implementation of the Avatar component that implements color shifting
// with hue rotation. We will probably change it in the future because it does not look very good.
export const Avatar: FC<HTMLAttributes<HTMLDivElement> & { name: string; rotation?: number }> = ({
	name,
	rotation,
	className,
	style,
	...props
}) => {
	const { socket } = useSocket()

	const hueRotation = useMemo(() => {
		if (socket && socket.id === name) return 0
		if (rotation !== undefined) return rotation

		const seed = Math.abs(name.split("").reduce((hash, char) => (hash << 5) - hash + char.charCodeAt(0), 0))
		return Math.floor((seed % angles) / angle) * angle
	}, [name, socket, rotation])

	return (
		<div
			className={cn("relative overflow-hidden rounded-sm", className)}
			style={{ filter: `hue-rotate(${hueRotation}deg)`, ...style }}
			{...props}
		>
			<BoringAvatar variant="beam" name={name} size={"100%"} colors={["#385842", "#D2F38A"]} square />
		</div>
	)
}
