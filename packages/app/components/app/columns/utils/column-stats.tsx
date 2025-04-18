import { useSession } from "next-auth/react"
import React, { FC, useState } from "react"

import { Clipboard } from "lucide-react"

import { Search } from "@/components/app/inputs/search"
import { Button } from "@/components/shared/buttons/button"
import { Counter } from "@/components/shared/utils/counter"
import { cn } from "@/lib"
import { api } from "@/server/client"
import { useSocket } from "@/state/authentication"

const COLORS = ["#F3B08A", "#F3EF8A", "#9F8AF3", "#8AF3E6"]

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
				"hover:plug-green/5 group relative flex w-full flex-col items-center filter transition-all duration-200 ease-in-out",
				isActive ? "plug-green/5" : "",
			)}
			onMouseEnter={() => onHover(index)}
			onMouseLeave={() => onHover(undefined)}
		>
			<div className="absolute bottom-0 left-1/2 top-8 w-[2px] -translate-x-1/2 bg-plug-green/10" />
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
					<>
						<div
							className={cn(
								"relative w-full transition-all duration-200 ease-in-out",
								hovering === undefined && isActive || hovering !== i && "grayscale filter",
								i === 0 && "rounded-b-lg",
								i === stats.length - 1 && "rounded-t-lg"
							)}
							style={{
								height: stat === null ? 0 : `${(stat / total) * 100}%`,
								minHeight: 8,
								backgroundColor: COLORS[i % COLORS.length]
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
				))}
			</div>
		</div>
	)
}

const ProfileStats = () => {
	const [hoveredPeriod, setHoveredPeriod] = useState<number | undefined>(undefined)
	const [toggledStats, setToggledStats] = useState<boolean[]>([false, false, false, false])

	const { data: statsData } = api.socket.stats.get.useQuery(undefined)

	const stats: number[][] =
		statsData?.periods.map((_, index) => [
			statsData.counts.plugs[index] ?? 0,
			statsData.counts.forks[index] ?? 0,
			statsData.counts.views[index] ?? 0,
			statsData.counts.referrals[index] ?? 0
		]) ?? Array(4).fill([0, 0, 0, 0])

	const max = Math.max(...stats.map(period => period.reduce((sum: number, value: number) => sum + (value ?? 0), 0)))
	const currentStats = hoveredPeriod !== undefined ? stats[hoveredPeriod] : stats[stats.length - 1]

	const startDate = statsData?.periods[0]?.weekStart
		? new Date(statsData.periods[0].weekStart).toLocaleDateString("en-US", {
			month: "2-digit",
			year: "2-digit"
		})
		: "09/24"
	const endDate = statsData?.periods[statsData.periods.length - 1]?.weekEnd
		? new Date(statsData.periods[statsData.periods.length - 1].weekEnd).toLocaleDateString("en-US", {
			month: "2-digit",
			year: "2-digit"
		})
		: "10/13"

	const handleToggle = (statIndex: number) => {
		setToggledStats(prev => [...prev.slice(0, statIndex), !prev[statIndex], ...prev.slice(statIndex + 1)])
	}

	return (
		<div className="flex h-full flex-col gap-8">
			<div className="flex flex-col gap-2">
				<div className="flex flex-row gap-2">
					<div
						className={`relative flex w-full cursor-pointer flex-col items-start justify-center rounded-md border-[1px] px-6 py-4 text-left ${toggledStats[0] === false ? "border-white bg-plug-green/5" : "border-plug-green/10 bg-white"
							}`}
						onClick={() => handleToggle(0)}
					>
						<p className="text-[32px] font-bold">
							<Counter count={currentStats[0]} />
						</p>
						<p className="flex flex-row items-center gap-2 font-bold text-black/40">
							<span className="h-2 w-2 rounded-full" style={{ backgroundColor: COLORS[0] }} />
							Plugs
						</p>
					</div>
					<div
						className={`relative flex w-full cursor-pointer flex-col items-start justify-center rounded-md border-[1px] px-6 py-4 text-left ${toggledStats[1] === false ? "border-white bg-plug-green/5" : "border-plug-green/10 bg-white"
							}`}
						onClick={() => handleToggle(1)}
					>
						<p className="text-[32px] font-bold">
							<Counter count={currentStats[1]} />
						</p>
						<p className="flex flex-row items-center gap-2 font-bold text-black/40">
							<span className="h-2 w-2 rounded-full" style={{ backgroundColor: COLORS[1] }} />
							Forks
						</p>
					</div>
				</div>
				<div className="flex flex-row gap-2">
					<div
						className={`relative flex w-full cursor-pointer flex-col items-start justify-center rounded-md border-[1px] px-6 py-4 text-left ${toggledStats[2] === false ? "border-white bg-plug-green/5" : "border-plug-green/10 bg-white"
							}`}
						onClick={() => handleToggle(2)}
					>
						<p className="text-[32px] font-bold">
							<Counter count={currentStats[2]} />
						</p>
						<p className="flex flex-row items-center gap-2 font-bold text-black/40">
							<span className="h-2 w-2 rounded-full" style={{ backgroundColor: COLORS[2] }} />
							Views
						</p>
					</div>
					<div
						className={`relative flex w-full cursor-pointer flex-col items-start justify-center rounded-md border-[1px] px-6 py-4 text-left ${toggledStats[3] === false ? "border-white bg-plug-green/5" : "border-plug-green/10 bg-white"
							}`}
						onClick={() => handleToggle(3)}
					>
						<p className="text-[32px] font-bold">
							<Counter count={currentStats[3]} />
						</p>
						<p className="flex flex-row items-center gap-2 font-bold text-black/40">
							<span className="h-2 w-2 rounded-full" style={{ backgroundColor: COLORS[3] }} />
							Referrals
						</p>
					</div>
				</div>
			</div>
			<div className="flex h-full flex-col">
				<div className="relative flex h-full flex-row gap-2">
					{Array.from({ length: stats.length }).map((_, i) => (
						<ProfileStat
							key={i}
							index={i}
							isActive={i === stats.length - 1}
							stats={stats[i].map((stat: number, j: number) => (toggledStats[j] ? 0 : stat))}
							max={max}
							onHover={setHoveredPeriod}
						/>
					))}
				</div>
				<div className="flex flex-row items-center justify-between py-2">
					<p className="font-bold opacity-40">{startDate}</p>
					<p className="font-bold opacity-40">{endDate}</p>
				</div>
			</div>
		</div>
	)
}

export const ColumnStats: FC<{ index: number }> = () => {
	const { data: session } = useSession()
	const { socket } = useSocket()
	const [copied, setCopied] = useState(false)

	if (!socket || !session?.user.id) return null

	const handleCopy = async () => {
		try {
			await navigator.clipboard.writeText(window.location.href)
			setCopied(true)
			setTimeout(() => setCopied(false), 2000)
		} catch (err) {
			console.error("Failed to copy:", err)
		}
	}

	return (
		<div className="flex h-full flex-col justify-between gap-4 overflow-y-scroll px-4 py-4 text-center">
			<ProfileStats />

			<div className="flex flex-col gap-2">
				<Search
					className="select-none pointer-events-none"
					icon={<Clipboard size={14} className="opacity-60" />}
					placeholder="Copy Referral Link"
					search={socket?.identity?.referralCode ?? ""}
					handleSearch={() => { }}
				/>
				<Button
					variant={copied ? "primaryDisabled" : "primary"}
					className="flex w-full flex-row items-center justify-center gap-2 truncate py-4"
					onClick={handleCopy}
				>
					{copied ? "Copied!" : "Copy Referral Link"}
				</Button>
			</div>
		</div>
	)
}
