import { useSession } from "next-auth/react"
import React, { FC, HTMLAttributes, useEffect, useMemo, useState } from "react"

import { AnimatePresence, motion } from "framer-motion"
import { Carrot, Clock, Egg, Heart, PawPrintIcon } from "lucide-react"

import { Button } from "@/components/shared/buttons/button"
import { Counter } from "@/components/shared/utils/counter"
import { DateSince } from "@/components/shared/utils/date-since"
import { greenGradientStyle, sunGradientStyle } from "@/lib"
import { api } from "@/server/client"
import { useSocket } from "@/state/authentication"

export const ColumnCompanion: FC<HTMLAttributes<HTMLDivElement> & { index: number }> = ({ index }) => {
	const { data: session } = useSession()
	const { socket } = useSocket()

	const [iconIndex, setIconIndex] = useState(0)
	const [treatsAnimation, setTreatsAnimation] = useState<{ count: number; key: number } | null>(null)

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

	const [timeSinceFeed, setTimeSinceFeed] = useState(
		Number(new Date().getTime()) -
			Number(feed?.lastFeedAt?.getTime() || socket?.identity?.companion?.lastFeedAt?.getTime() || "0")
	)

	const [hours, minutes, seconds] = useMemo(() => {
		const timeUntilFeed = Math.max(24 * 60 * 60 * 1000 - timeSinceFeed, 0) / 1000

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

	useEffect(() => {
		const interval = setInterval(() => {
			setIconIndex(prev => (prev === 0 ? 1 : 0))
		}, 3000)

		return () => clearInterval(interval)
	}, [])

	if (!socket || !session?.user.id) return null

	return (
		<>
			<div className="flex h-full flex-col items-center gap-4 px-4 py-4">
				<div className="relative mx-4 flex h-full min-h-96 w-full flex-col items-center justify-center gap-1 rounded-lg bg-gradient-to-tr from-plug-green/5 to-white p-8 py-16 text-center">
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

				<Button
					variant={canFeed ? "primary" : "primaryDisabled"}
					className="flex w-full items-center justify-center py-4"
					onClick={canFeed ? () => feedMutation.mutate() : () => {}}
					disabled={feedMutation.isPending || !canFeed}
				>
					{feedMutation.isPending ? (
						"..."
					) : canFeed && !feed ? (
						"Feed"
					) : (
						<span className="tabular-nums">
							<Counter count={`${hours}:${minutes}:${seconds}`} />
						</span>
					)}
				</Button>
			</div>
		</>
	)
}
