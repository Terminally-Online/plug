import { useSession } from "next-auth/react"
import React, { FC, useState } from "react"

import { Avatar, Counter, Image } from "@/components"
import { cn, formatAddress } from "@/lib"
import { useSocket } from "@/state"

const stats = [
	[1900, 5123, 5200, 1234],
	[1300, 3123, 500, 2123],
	[927, 4123, 390, 1234],
	[201, 523, 1233, 1230]
]

const gradients = ["#00E100, #A3F700", "#FFA800, #FAFF00", "#4E7FFD, #9E62FF", "#F94EFD, #FD4ECC"]

const ProfileStat: FC<{
	index: number
	isActive: boolean
	stats: Array<number | null>
	max: number
	onHover: (index: number | undefined) => void
}> = ({ index, isActive, stats, max, onHover }) => {
	const [hovering, setHovering] = useState<number | undefined>(undefined)

	const total = stats.reduce((a, b) => (a ?? 0) + (b ?? 0), 0) ?? 0

	return (
		<div
			className={cn(
				"group relative flex h-48 w-full flex-col items-center grayscale filter transition-all duration-200 ease-in-out hover:grayscale-0",
				isActive && "grayscale-0"
			)}
			onMouseEnter={() => onHover(index)}
			onMouseLeave={() => onHover(undefined)}
		>
			<div className="absolute bottom-0 left-1/2 top-8 w-[2px] -translate-x-1/2 bg-grayscale-100" />
			<p
				className={cn(
					"font-bold opacity-40 transition-all duration-200 ease-in-out group-hover:opacity-100",
					isActive && "opacity-100"
				)}
			>
				<Counter count={hovering !== undefined ? (stats[hovering] ?? 0) : total} />
			</p>
			<div
				className="relative mt-auto flex h-full w-full flex-col-reverse"
				style={{ height: `${(total / max) * 100}%` }}
			>
				{stats.map((stat, i) => (
					<React.Fragment key={i}>
						{stat !== null && (
							<>
								<div
									className={cn(
										"relative w-full bg-gradient-to-tr transition-all duration-200 ease-in-out",
										hovering !== undefined && hovering !== i && "grayscale filter",
										i === 0 && "rounded-b-lg",
										i === stats.length - 1 && "rounded-t-lg"
									)}
									style={{
										height: stat === null ? 0 : (stat / total) * 100,
										background: `linear-gradient(30deg, ${gradients[i % gradients.length]})`
									}}
									onMouseEnter={() => {
										setHovering(i)
									}}
									onMouseLeave={() => {
										setHovering(undefined)
									}}
								/>
								<div className="relative h-[1px] w-full bg-white" />
							</>
						)}
					</React.Fragment>
				))}
			</div>
		</div>
	)
}

const ProfileStats: FC<{ onHover: (index: number | undefined) => void }> = ({ onHover }) => {
	const max = Math.max(...stats.map(period => period.reduce((sum, value) => sum + (value ?? 0), 0)))

	return (
		<>
			<div className="flex flex-col">
				<div className="relative flex h-48 flex-row gap-2 px-4">
					{Array.from({ length: stats.length }).map((_, i) => (
						<ProfileStat
							key={i}
							index={i}
							isActive={i === stats.length - 1}
							stats={stats[i]}
							max={max}
							onHover={onHover}
						/>
					))}
				</div>
				<div className="flex flex-row items-center justify-between px-4 py-2">
					<p className="font-bold opacity-40">09/24</p>
					<p className="font-bold opacity-40">10/13</p>
				</div>
			</div>
		</>
	)
}

export const ColumnProfile: FC<{ index: number }> = () => {
	const { data: session } = useSession()
	const { name, avatar, socket } = useSocket()
	const [hoveredPeriod, setHoveredPeriod] = useState<number | undefined>(undefined)

	if (!socket || !session?.user.id) return null

	const currentStats = hoveredPeriod !== undefined ? stats[hoveredPeriod] : stats[stats.length - 1]

	return (
		<div className="flex h-full flex-col gap-4 overflow-hidden py-4 text-center">
			<div className="flex flex-row items-center gap-8 px-4 pb-4">
				<div className="relative h-24 w-24">
					{avatar ? (
						<Image
							src={avatar}
							alt="ENS Avatar"
							width={64}
							height={64}
							className="h-full w-full rounded-sm"
						/>
					) : (
						<>
							<div className="absolute blur-[80px] filter">
								<Avatar name={socket?.id ?? ""} />
							</div>
							<Avatar name={socket?.id ?? ""} />
						</>
					)}
				</div>

				<div className="flex flex-col">
					<p className="mr-auto text-lg font-bold">{name !== "" ? name : formatAddress(socket.id, 6)}</p>
					<p className="mr-auto text-lg font-bold opacity-40">Joined: {new Date().toLocaleDateString()}</p>
				</div>
			</div>

			<ProfileStats onHover={setHoveredPeriod} />

			<div className="flex flex-col gap-2 px-4">
				<div className="flex flex-row gap-2">
					<div className="flex w-full flex-col items-start justify-center rounded-md bg-grayscale-0 px-6 py-4 text-left">
						<p className="text-[32px] font-bold">
							<Counter count={currentStats[0]} />
						</p>
						<p className="flex flex-row items-center gap-2 font-bold text-black/40">
							<span className="h-2 w-2 rounded-full bg-gradient-to-tr from-plug-green to-plug-yellow" />
							Users
						</p>
					</div>
					<div className="flex w-full flex-col items-start justify-center rounded-md bg-grayscale-0 px-6 py-4 text-left">
						<p className="text-[32px] font-bold">
							<Counter count={currentStats[1]} />
						</p>
						<p className="flex flex-row items-center gap-2 font-bold text-black/40">
							<span className="h-2 w-2 rounded-full bg-gradient-to-tr from-sun-orange to-sun-yellow" />
							Runs
						</p>
					</div>
				</div>
				<div className="flex flex-row gap-2">
					<div className="flex w-full flex-col items-start justify-center rounded-md bg-grayscale-0 px-6 py-4 text-left">
						<p className="text-[32px] font-bold">
							<Counter count={currentStats[2]} />
						</p>
						<p className="flex flex-row items-center gap-2 font-bold text-black/40">
							<span className="from-ocean-blue to-ocean-purple h-2 w-2 rounded-full bg-gradient-to-tr" />
							Views
						</p>
					</div>
					<div className="flex w-full flex-col items-start justify-center rounded-md bg-grayscale-0 px-6 py-4 text-left">
						<p className="text-[32px] font-bold">
							<Counter count={currentStats[3]} />
						</p>
						<p className="flex flex-row items-center gap-2 font-bold text-black/40">
							<span className="from-pink-pink to-pink-purple h-2 w-2 rounded-full bg-gradient-to-tr" />
							Referrals
						</p>
					</div>
				</div>
			</div>
		</div>
	)
}
