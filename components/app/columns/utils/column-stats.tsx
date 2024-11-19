import { useSession } from "next-auth/react"
import React, { FC, useState } from "react"

import { CheckCircle, Clipboard } from "lucide-react"

import { Button, Counter, Search } from "@/components"
import { cn } from "@/lib"
import { api } from "@/server/client"
import { useSocket } from "@/state"

const GRADIENTS = ["#00E100, #A3F700", "#FFA800, #FAFF00", "#4E7FFD, #9E62FF", "#F94EFD, #FD4ECC"]

type StatsResponse = {
	counts: {
		referrals: number[]
		views: number[]
		runs: number[]
		users: number[]
	}
	periods: {
		weekStart: string
		weekEnd: string
	}[]
}

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
					<>
						<div
							className={cn(
								"relative w-full bg-gradient-to-tr transition-all duration-200 ease-in-out",
								hovering !== undefined && hovering !== i && "grayscale filter",
								i === 0 && "rounded-b-lg",
								i === stats.length - 1 && "rounded-t-lg"
							)}
							style={{
								height: stat === null ? 0 : `${(stat / total) * 100}%`,
								minHeight: 8,
								background: `linear-gradient(30deg, ${GRADIENTS[i % GRADIENTS.length]})`
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

	const { data: statsData } = api.socket.stats.get.useQuery(undefined, {
		refetchInterval: 60 * 1000
	})

	// Construct stats array with real data
	const stats: number[][] =
		statsData?.periods.map((_, index) => [
			statsData.counts.users[index] ?? 0,
			statsData.counts.runs[index] ?? 0,
			statsData.counts.views[index] ?? 0,
			statsData.counts.referrals[index] ?? 0
		]) ?? Array(4).fill([0, 0, 0, 0])

	const max = Math.max(...stats.map(period => period.reduce((sum: number, value: number) => sum + (value ?? 0), 0)))
	const currentStats = hoveredPeriod !== undefined ? stats[hoveredPeriod] : stats[stats.length - 1]

	// Format dates from the API response
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
		<div className="flex h-full flex-col justify-between gap-4 overflow-y-scroll px-6 py-4 text-center">
			<ProfileStats />

			<div className="flex flex-col gap-2">
				<Search
					icon={<Clipboard size={14} className="opacity-60" />}
					placeholder="Copy Referral Link"
					search={socket?.identity?.referralCode ?? ""}
					handleSearch={() => {}}
				/>
				<Button
					variant={copied ? "primaryDisabled" : "primary"}
					className="flex w-full flex-row items-center justify-center gap-2 truncate py-4"
					onClick={handleCopy}
				>
					{copied ? (
						<>
							<CheckCircle size={14} className="opacity-60" />
							Copied!
						</>
					) : (
						<>
							<Clipboard size={14} className="opacity-60" />
							Copy Referral Link
						</>
					)}
				</Button>
			</div>
		</div>
	)
}
