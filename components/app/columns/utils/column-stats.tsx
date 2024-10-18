import { useSession } from "next-auth/react"
import React, { FC, useState } from "react"

import { Avatar, Button, Counter, Image } from "@/components"
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
				"group relative flex w-full flex-col items-center grayscale filter transition-all duration-200 ease-in-out hover:grayscale-0",
				isActive && "grayscale-0"
			)}
			onMouseEnter={() => onHover(index)}
			onMouseLeave={() => onHover(undefined)}
		>
			<div className="absolute bottom-0 left-1/2 top-8 w-[2px] -translate-x-1/2 bg-grayscale-100" />
			<p
				className={cn(
					"mb-8 font-bold opacity-40 transition-all duration-200 ease-in-out group-hover:opacity-100",
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
										height: stat === null ? 0 : (stat / total) * 400,
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

const ProfileStats = () => {
	const [hoveredPeriod, setHoveredPeriod] = useState<number | undefined>(undefined)
	const [toggledStats, setToggledStats] = useState<boolean[]>([false, false, false, false])

	const max = Math.max(...stats.map(period => period.reduce((sum, value) => sum + (value ?? 0), 0)))
	const currentStats = hoveredPeriod !== undefined ? stats[hoveredPeriod] : stats[stats.length - 1]

	const handleToggle = (statIndex: number) => {
		setToggledStats(prev => [...prev.slice(0, statIndex), !prev[statIndex], ...prev.slice(statIndex + 1)])
	}

	return (
		<div className="flex flex-col gap-8">
			<div className="flex flex-col gap-2">
				<div className="flex flex-row gap-2">
					<div
						className={`relative flex w-full cursor-pointer flex-col items-start justify-center rounded-md border-[1px] px-6 py-4 text-left ${
							toggledStats[0] === false ? "border-white bg-grayscale-0" : "border-grayscale-100 bg-white"
						}`}
						onClick={() => handleToggle(0)}
					>
						<p className="text-[32px] font-bold">
							<Counter count={currentStats[0]} />
						</p>
						<p className="flex flex-row items-center gap-2 font-bold text-black/40">
							<span className="h-2 w-2 rounded-full bg-gradient-to-tr from-plug-green to-plug-yellow" />
							Users
						</p>
					</div>
					<div
						className={`relative flex w-full cursor-pointer flex-col items-start justify-center rounded-md border-[1px] px-6 py-4 text-left ${
							toggledStats[1] === false ? "border-white bg-grayscale-0" : "border-grayscale-100 bg-white"
						}`}
						onClick={() => handleToggle(1)}
					>
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
					<div
						className={`relative flex w-full cursor-pointer flex-col items-start justify-center rounded-md border-[1px] px-6 py-4 text-left ${
							toggledStats[2] === false ? "border-white bg-grayscale-0" : "border-grayscale-100 bg-white"
						}`}
						onClick={() => handleToggle(2)}
					>
						<p className="text-[32px] font-bold">
							<Counter count={currentStats[2]} />
						</p>
						<p className="flex flex-row items-center gap-2 font-bold text-black/40">
							<span className="h-2 w-2 rounded-full bg-gradient-to-tr from-ocean-blue to-ocean-purple" />
							Views
						</p>
					</div>
					<div
						className={`relative flex w-full cursor-pointer flex-col items-start justify-center rounded-md border-[1px] px-6 py-4 text-left ${
							toggledStats[3] === false ? "border-white bg-grayscale-0" : "border-grayscale-100 bg-white"
						}`}
						onClick={() => handleToggle(3)}
					>
						<p className="text-[32px] font-bold">
							<Counter count={currentStats[3]} />
						</p>
						<p className="flex flex-row items-center gap-2 font-bold text-black/40">
							<span className="h-2 w-2 rounded-full bg-gradient-to-tr from-pink-pink to-pink-purple" />
							Referrals
						</p>
					</div>
				</div>
			</div>

			<div className="flex flex-col">
				<div className="relative flex flex-row gap-2">
					{Array.from({ length: stats.length }).map((_, i) => (
						<ProfileStat
							key={i}
							index={i}
							isActive={i === stats.length - 1}
							stats={stats[i].map((stat, j) => (toggledStats[j] ? null : stat))}
							max={max}
							onHover={setHoveredPeriod}
						/>
					))}
				</div>
				<div className="flex flex-row items-center justify-between py-2">
					<p className="font-bold opacity-40">09/24</p>
					<p className="font-bold opacity-40">10/13</p>
				</div>
			</div>
		</div>
	)
}

export const ColumnStats: FC<{ index: number }> = () => {
	const { data: session } = useSession()
	const { socket } = useSocket()

	if (!socket || !session?.user.id) return null

	return (
		<div className="flex h-full flex-col justify-between gap-4 overflow-y-scroll px-6 py-4 text-center">
			<ProfileStats />
			<Button className="w-full" onClick={() => {}}>
				Share Link
			</Button>
		</div>
	)
}
