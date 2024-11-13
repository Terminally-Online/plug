import { motion } from "framer-motion"

import { InfoCard } from "./cards"
import { LandingContainer } from "./layout"

const ITEMS = 12

export const Platform = () => {
	return (
		<LandingContainer className="relative mb-[80px] grid grid-cols-3 grid-rows-2 gap-8">
			<InfoCard
				icon={<span className="text-white">🌐</span>}
				text="Protocol Abstraction."
				description="All the functionality of major protocols are available to you without having to bounce between apps."
				className="relative z-[99999] h-[320px] sm:h-[320px] 2xl:h-[300px]"
			>
				<div className="absolute bottom-[50%] left-0 right-0 top-0 bg-gradient-to-b from-plug-white/0 to-plug-white" />
				<div className="absolute bottom-0 left-0 right-0 top-[50%] bg-plug-white" />
			</InfoCard>
			<InfoCard
				icon={<span className="text-white">🌐</span>}
				text="Customizable Layouts."
				description="However you prefer it, your layout can be resized, re-ordered, and refocused to what matters most to you."
				className="relative z-[99999] col-span-2 row-span-2 h-[320px] sm:h-[320px] 2xl:h-[300px]"
			>
				<div className="absolute inset-0 flex flex-row items-center justify-center">
					<motion.div
						className="h-full border-r-[2px] border-plug-green/10"
						initial={{ width: "20%" }}
						animate={{ width: "30%" }}
						transition={{
							duration: 0.5,
							repeat: Infinity,
							repeatType: "reverse",
							ease: "easeInOut",
							delay: 3,
							repeatDelay: 6
						}}
					>
						<p className="p-2 pl-4 font-bold">Discover</p>
						<div className="h-[2px] w-full bg-plug-green/10" />
						<div className="flex grid h-full grid-cols-2 gap-1 px-2 pt-2">
							{Array.from({ length: ITEMS }).map((_, index) => (
								<div
									className="gradient-animated h-[52px] w-full rounded-lg bg-grayscale-100"
									key={index}
								/>
							))}
						</div>
					</motion.div>
					<motion.div
						className="h-full border-r-[2px] border-plug-green/10"
						initial={{ width: "20%" }}
						animate={{ width: "30%" }}
						transition={{
							duration: 0.5,
							repeat: Infinity,
							repeatType: "reverse",
							ease: "easeInOut",
							delay: 2,
							repeatDelay: 10
						}}
					>
						<p className="p-2 pl-4 font-bold">Tokens</p>
						<div className="h-[2px] w-full bg-plug-green/10" />
						<div className="flex h-full flex-col gap-1 px-2 pt-2">
							{Array.from({ length: ITEMS / 2 }).map((_, index) => (
								<div
									className="gradient-animated h-[48px] w-full rounded-lg bg-grayscale-100"
									key={index}
								/>
							))}
						</div>
					</motion.div>
					<motion.div
						className="h-full border-r-[2px] border-plug-green/10"
						initial={{ width: "30%" }}
						animate={{ width: "20%" }}
						transition={{
							duration: 0.5,
							repeat: Infinity,
							repeatType: "reverse",
							ease: "easeInOut",
							delay: 4,
							repeatDelay: 10
						}}
					>
						<p className="p-2 pl-4 font-bold">Collectibles</p>
						<div className="h-[2px] w-full bg-plug-green/10" />
						<div className="flex h-full flex-col gap-1 px-2 pt-2">
							{Array.from({ length: ITEMS / 2 }).map((_, index) => (
								<div
									className="gradient-animated h-[48px] w-full rounded-lg bg-grayscale-100"
									key={index}
								/>
							))}
						</div>
					</motion.div>
					<motion.div
						className="h-full border-r-[2px] border-plug-green/10"
						initial={{ width: "30%" }}
						animate={{ width: "20%" }}
						transition={{
							duration: 0.5,
							repeat: Infinity,
							repeatType: "reverse",
							ease: "easeInOut",
							delay: 6,
							repeatDelay: 6
						}}
					>
						<p className="p-2 pl-4 font-bold">Strategies</p>
						<div className="h-[2px] w-full bg-plug-green/10" />
						<div className="flex grid h-full grid-cols-2 gap-1 px-2 pt-2">
							{Array.from({ length: ITEMS }).map((_, index) => (
								<div
									className="gradient-animated h-[52px] w-full rounded-lg bg-grayscale-100"
									key={index}
								/>
							))}
						</div>
					</motion.div>

					<motion.div
						className="h-full border-r-[2px] border-plug-green/10"
						initial={{ width: "20%" }}
						animate={{ width: "40%" }}
						transition={{
							duration: 0.5,
							repeat: Infinity,
							repeatType: "reverse",
							ease: "easeInOut",
							delay: 6,
							repeatDelay: 6
						}}
					>
						<p className="p-2 pl-4 font-bold">Activity</p>
						<div className="h-[2px] w-full bg-plug-green/10" />
						<div className="flex h-full flex-col gap-1 px-2 pt-2">
							{Array.from({ length: ITEMS / 2 }).map((_, index) => (
								<div
									className="gradient-animated h-[48px] w-full rounded-lg bg-grayscale-100"
									key={index}
								/>
							))}
						</div>
					</motion.div>
				</div>

				<div className="absolute bottom-[50%] left-0 right-0 top-0 bg-gradient-to-b from-plug-white/0 to-plug-white" />
				<div className="absolute bottom-0 left-0 right-0 top-[50%] bg-plug-white" />
			</InfoCard>
			<InfoCard
				icon={<span className="text-white">🌐</span>}
				text="Portfolio Management."
				description="Swap between different chains with ease."
				className="relative z-[99999] h-[320px] sm:h-[320px] 2xl:h-[300px]"
			>
				<div className="absolute bottom-[50%] left-0 right-0 top-0 bg-gradient-to-b from-plug-white/0 to-plug-white" />
				<div className="absolute bottom-0 left-0 right-0 top-[50%] bg-plug-white" />
			</InfoCard>
		</LandingContainer>
	)
}
