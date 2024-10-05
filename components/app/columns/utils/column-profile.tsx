import { useSession } from "next-auth/react"
import React, { FC, useEffect, useMemo, useState } from "react"

import { AnimatePresence, motion } from "framer-motion"
import { Carrot, Clock, Egg, Heart, PawPrintIcon, Sun } from "lucide-react"

import { api } from "@/server/client"

import { Avatar, Button, Counter, DateSince, Image } from "@/components"
import { cn, formatAddress, greenGradientStyle, sunGradientStyle } from "@/lib"
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

const ProfileStats = () => {
	const [hoveredPeriod, setHoveredPeriod] = useState<number | undefined>(undefined)
	const [toggledStats, setToggledStats] = useState<boolean[]>([false, false, false, false])

	const max = Math.max(...stats.map(period => period.reduce((sum, value) => sum + (value ?? 0), 0)))
	const currentStats = hoveredPeriod !== undefined ? stats[hoveredPeriod] : stats[stats.length - 1]

	const handleToggle = (statIndex: number) => {
		setToggledStats(prev => [...prev.slice(0, statIndex), !prev[statIndex], ...prev.slice(statIndex + 1)])
	}

	return (
		<>
			<div className="flex flex-col">
				<div className="relative flex h-48 flex-row gap-2 px-4">
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
				<div className="flex flex-row items-center justify-between px-4 py-2">
					<p className="font-bold opacity-40">09/24</p>
					<p className="font-bold opacity-40">10/13</p>
				</div>
			</div>

			<div className="flex flex-col gap-2 px-4">
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
		</>
	)
}

export const ColumnProfile: FC<{ index: number }> = () => {
	const { data: session } = useSession()
	const { name, avatar, socket } = useSocket()
	const [iconIndex, setIconIndex] = useState(0)

	const feedMutation = api.socket.companion.feed.useMutation({
		onMutate: () => {
			setTimeSinceFeed(0)
		},
		onSuccess: data => {
			setTreatsAnimation({
				key: Date.now(),
				count: data.treatsFed - (socket?.identity?.companion?.treatsFed ?? 0)
			})
			setFeed(data)
		}
	}) 
	const [feed, setFeed] = useState<{
		createdAt: Date
		updatedAt: Date
		name: string
		feedCount: number
		treatsFed: number
		lastFeedAt: Date | null
		streak: number
		socketId: string
	}>()

	const [treatsAnimation, setTreatsAnimation] = useState<{ count: number; key: number } | null>(null)

	useEffect(() => {
		const interval = setInterval(() => {
			setIconIndex(prev => (prev === 0 ? 1 : 0))
		}, 3000)

		return () => clearInterval(interval)
	}, [])

	const [timeSinceFeed, setTimeSinceFeed] = useState(
		Number(new Date().getTime()) -
			Number(feed?.lastFeedAt?.getTime() || socket?.identity?.companion?.lastFeedAt?.getTime() || "0")
	)

	const [hours, minutes, seconds] = useMemo(() => {
		const timeUntilFeed = Math.max(24 * 60 * 60 * 1000 - timeSinceFeed,0) / 1000

		return [
			String(Math.floor(timeUntilFeed / 3600)).padStart(2, "0"),
			String(Math.floor((timeUntilFeed % 3600) / 60)).padStart(2, "0"),
			String(Math.floor(timeUntilFeed % 60)).padStart(2, "0")
		]
	}, [timeSinceFeed])

	const canFeed = useMemo(() => {
		if (!socket) return false

		return timeSinceFeed >= 24 * 60 * 60 * 1000
	}, [socket, timeSinceFeed])

	useEffect(() => {
		if (!socket && !feed) return

		const interval = setInterval(() => {
			setTimeSinceFeed(
				Number(new Date().getTime()) -
					Number(feed?.lastFeedAt?.getTime() || socket?.identity?.companion?.lastFeedAt?.getTime() || "0")
			)
		}, 1000)

		return () => clearInterval(interval)
	}, [socket, feed])

	if (!socket || !session?.user.id) return null

	return (
		<div className="flex h-full flex-col gap-4 overflow-y-scroll text-center">
			<div className="flex flex-row items-center gap-8 px-4 py-4">
				<div className="relative max-w-[64px]">
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

				<div className="relative flex w-full flex-col">
					<p className="mr-auto text-lg font-bold">
						{name !== "" ? name : formatAddress(socket.socketAddress, 6)}
					</p>
				</div>
			</div>

			<div className="relative mx-4 flex min-h-96 flex-col items-center justify-center gap-1 rounded-lg bg-gradient-to-tr from-grayscale-0 to-white p-8 py-16">
				<div className="absolute left-4 right-4 top-4">
					<div className="flex w-full flex-row items-center justify-between gap-2">
						<p
							className="flex w-full flex-row text-right font-bold"
							style={{
								...greenGradientStyle
							}}
						>
							#123
						</p>
						<p
							className="w-max whitespace-nowrap font-bold"
							style={{
								...sunGradientStyle
							}}
						>
							{feed?.streak ?? socket?.identity?.companion?.streak} day streak
						</p>
					</div>
				</div>
				<div className="relative mb-4 h-[48px] w-[48px]">
					<AnimatePresence mode="wait">
						<motion.div
							key={iconIndex}
							initial={{ opacity: 0, y: 10 }}
							animate={{ opacity: 1, y: 0 }}
							exit={{ opacity: 0, y: -10 }}
							transition={{
								duration: 0.5,
								bounce: 0.5,
								type: "spring"
							}}
							className="absolute inset-0"
						>
							{iconIndex === 0 ? (
								<Heart size={48} className="opacity-40" />
							) : (
								<>
									{canFeed ? (
										<PawPrintIcon size={48} className="opacity-40" />
									) : (
										<motion.div
											style={{ transformOrigin: "bottom center" }}
											animate={{
												rotate: [-6, 6, -6],
												transition: {
													duration: 0.5,
													repeat: Infinity,
													ease: "easeInOut"
												}
											}}
										>
											<Egg size={48} className="opacity-40" />
										</motion.div>
									)}
								</>
							)}
						</motion.div>
					</AnimatePresence>
				</div>
				<p className="font-bold">{canFeed ? "Feed your companion." : "Your companion is incubating."}</p>
				<p className="mx-auto max-w-[240px] text-sm font-bold opacity-40">
					{canFeed
						? "Feed your companion to boost the growth of their DNA."
						: "As you interact with protocols, increase your stats below, and feed your companion, the DNA will grow."}
				</p>

				<div className="absolute bottom-0 left-0 right-0 m-4 flex flex-row justify-between gap-2">
					<p className="relative flex flex-row items-center gap-2 font-bold opacity-40">
						<Carrot size={24} className="h-4 w-4 min-w-4 opacity-40" />
						<Counter count={feed?.treatsFed ?? socket?.identity?.companion?.treatsFed ?? 0} />
						{treatsAnimation && (
							<motion.span
								key={treatsAnimation.key}
								initial={{ opacity: 1, y: 0 }}
								animate={{ opacity: 0, y: -20 }}
								exit={{ opacity: 0 }}
								transition={{ duration: 1 }}
								className="absolute -right-8 text-opacity-100"
							>
								+{treatsAnimation.count}
							</motion.span>
						)}
					</p>
					<p className="flex flex-row items-center gap-2 font-bold opacity-40">
						<Clock size={24} className="h-4 w-4 opacity-40" />
						<DateSince date={socket?.createdAt ?? new Date()} ago={false} />
					</p>
				</div>
			</div>

			<div className="flex flex-row gap-2 px-4">
				<Button
					variant="secondary"
					className={cn(
						"w-max",
						canFeed === false && "cursor-auto bg-grayscale-0 hover:bg-grayscale-0 hover:text-opacity-60"
					)}
					onClick={
						canFeed
							? () => feedMutation.mutate()
							: () => {}
					}
					disabled={feedMutation.isLoading || !canFeed}
				>
					{feedMutation.isLoading ? (
						"..."
					) : canFeed && !feed ? (
						"Feed"
					) : (
						<span className="tabular-nums">
							<Counter count={`${hours}:${minutes}:${seconds}`} />
						</span>
					)}
				</Button>
				<Button className="w-full" onClick={() => {}}>
					Share Link
				</Button>
			</div>

			<ProfileStats />
		</div>
	)
}
