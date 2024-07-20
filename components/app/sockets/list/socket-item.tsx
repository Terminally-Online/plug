import { FC } from "react"

import Link from "next/link"
import { useRouter } from "next/navigation"

import BlockiesSvg from "blockies-react-svg"
import { motion } from "framer-motion"

import { Accordion } from "@/components"
import { cn, formatAddress } from "@/lib"
import { UserSocket } from "@/server/api/routers/socket"

type Props = { socket: UserSocket }

export const SocketItem: FC<Props> = ({ socket }) => {
	const router = useRouter()

	// TODO: Implement display of the deployed chains.
	// TODO: Implement the backend functionality for this.
	const activity = Array.from({ length: 7 }, () => Math.random())

	return (
		<motion.div
			variants={{
				hidden: { opacity: 0, y: 10 },
				visible: {
					opacity: 1,
					y: 0,
					transition: {
						type: "spring",
						stiffness: 100,
						damping: 10
					}
				}
			}}
		>
			<Accordion
				onExpand={() => {
					router.push(`/app/sockets/${socket.socketAddress}`)
				}}
			>
				<div className="flex w-full flex-row items-center gap-4 text-left">
					<div className="relative h-10 w-10">
						<BlockiesSvg
							address={socket.socketAddress}
							className="absolute left-1/2 top-1/2 h-24 w-24 -translate-x-1/2 -translate-y-1/2 animate-fade-in rounded-full blur-2xl filter transition-all duration-200 ease-in-out"
						/>
						<BlockiesSvg
							address={socket.socketAddress}
							className="absolute left-1/2 top-1/2 h-10 w-10 min-w-10 -translate-x-1/2 -translate-y-1/2 animate-fade-in rounded-md bg-grayscale-100"
						/>
					</div>

					<div className="flex flex-col">
						<p className="font-bold">{socket.name}</p>
						<p className="text-sm text-black/65">
							{formatAddress(socket.socketAddress)}
						</p>
					</div>

					<div className="ml-auto mt-auto flex flex-row gap-1">
						{activity.map((activity, index) => (
							<div
								key={index}
								className={cn(
									"rounded-xs mt-auto w-4",
									index === 6
										? "bg-gradient-to-tr from-[#00E100] to-[#A3F700]"
										: "bg-grayscale-100"
								)}
								style={{
									height: `${Math.max(10, activity * 32)}px`
								}}
							/>
						))}
					</div>
				</div>
			</Accordion>
		</motion.div>
	)
}
