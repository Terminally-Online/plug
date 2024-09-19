import { useSession } from "next-auth/react"
import Image from "next/image"

import Avatar from "boring-avatars"
import { motion } from "framer-motion"
import { Bell, HousePlug, Plus, Search } from "lucide-react"

import { usePlugs } from "@/contexts"
import { MOBILE_INDEX, VIEW_KEYS } from "@/lib"
import { useColumns, useSocket } from "@/state"

const ProgressiveBlur = () => {
	return (
		<>
			<style jsx>{`
				.gradient-blur {
					position: absolute;
					z-index: 5;
					inset: auto 0 0 0;
					height: 100%;
					pointer-events: none;
				}
				.gradient-blur > div,
				.gradient-blur::before,
				.gradient-blur::after {
					position: absolute;
					inset: 0;
				}
				.gradient-blur::before {
					content: "";
					z-index: 1;
					backdrop-filter: blur(0.5px);
					mask: linear-gradient(
						to bottom,
						rgba(0, 0, 0, 0) 0%,
						rgba(0, 0, 0, 1) 12.5%,
						rgba(0, 0, 0, 1) 25%,
						rgba(0, 0, 0, 0) 37.5%
					);
				}
				.gradient-blur > div:nth-of-type(1) {
					z-index: 2;
					backdrop-filter: blur(1px);
					mask: linear-gradient(
						to bottom,
						rgba(0, 0, 0, 0) 12.5%,
						rgba(0, 0, 0, 1) 25%,
						rgba(0, 0, 0, 1) 37.5%,
						rgba(0, 0, 0, 0) 50%
					);
				}
				.gradient-blur > div:nth-of-type(2) {
					z-index: 3;
					backdrop-filter: blur(2px);
					mask: linear-gradient(
						to bottom,
						rgba(0, 0, 0, 0) 25%,
						rgba(0, 0, 0, 1) 37.5%,
						rgba(0, 0, 0, 1) 50%,
						rgba(0, 0, 0, 0) 62.5%
					);
				}
				.gradient-blur > div:nth-of-type(3) {
					z-index: 4;
					backdrop-filter: blur(4px);
					mask: linear-gradient(
						to bottom,
						rgba(0, 0, 0, 0) 37.5%,
						rgba(0, 0, 0, 1) 50%,
						rgba(0, 0, 0, 1) 62.5%,
						rgba(0, 0, 0, 0) 75%
					);
				}
				.gradient-blur > div:nth-of-type(4) {
					z-index: 5;
					backdrop-filter: blur(8px);
					mask: linear-gradient(
						to bottom,
						rgba(0, 0, 0, 0) 50%,
						rgba(0, 0, 0, 1) 62.5%,
						rgba(0, 0, 0, 1) 75%,
						rgba(0, 0, 0, 0) 87.5%
					);
				}
				.gradient-blur > div:nth-of-type(5) {
					z-index: 6;
					backdrop-filter: blur(16px);
					mask: linear-gradient(
						to bottom,
						rgba(0, 0, 0, 0) 62.5%,
						rgba(0, 0, 0, 1) 75%,
						rgba(0, 0, 0, 1) 87.5%,
						rgba(0, 0, 0, 0) 100%
					);
				}
				.gradient-blur > div:nth-of-type(6) {
					z-index: 7;
					backdrop-filter: blur(32px);
					mask: linear-gradient(
						to bottom,
						rgba(0, 0, 0, 0) 75%,
						rgba(0, 0, 0, 1) 87.5%,
						rgba(0, 0, 0, 1) 100%
					);
				}
				.gradient-blur::after {
					content: "";
					z-index: 8;
					backdrop-filter: blur(64px);
					mask: linear-gradient(to bottom, rgba(0, 0, 0, 0) 87.5%, rgba(0, 0, 0, 1) 100%);
				}
			`}</style>

			<div className="gradient-blur">
				<div></div>
				<div></div>
				<div></div>
				<div></div>
				<div></div>
				<div></div>
			</div>
		</>
	)
}

export const PageNavbar = () => {
	const { data: session } = useSession()
	const { avatar } = useSocket()
	const { column, navigate } = useColumns(MOBILE_INDEX)
	const { handle } = usePlugs()

	if (!column) return null

	return (
		<div className="fixed bottom-0 left-0 right-0 z-[10] ">
			<ProgressiveBlur />
			<div className="relative z-[11] flex flex-row items-center justify-between gap-2  px-8 py-6">
				<button
					className="group flex h-8 w-8 items-center justify-center"
					onClick={() => navigate({ index: MOBILE_INDEX, key: VIEW_KEYS.HOME })}
				>
					<HousePlug
						size={24}
						className="text-black text-opacity-40 transition-all duration-200 ease-in-out group-hover:text-opacity-100"
					/>
				</button>
				<button
					className="group flex h-8 w-8 items-center justify-center"
					onClick={() => navigate({ index: MOBILE_INDEX, key: VIEW_KEYS.SEARCH })}
				>
					<Search
						size={24}
						className="text-black text-opacity-40 transition-all duration-200 ease-in-out group-hover:text-opacity-100"
					/>
				</button>
				<button
					className="group flex h-8 w-8 items-center justify-center rounded-sm bg-gradient-to-tr from-plug-green to-plug-yellow"
					onClick={() => handle.plug.add({ index: MOBILE_INDEX })}
				>
					<Plus
						size={24}
						className="text-white text-opacity-80 transition-all duration-200 ease-in-out group-hover:text-opacity-100"
					/>
				</button>
				<button
					className="group flex h-8 w-8 items-center justify-center"
					onClick={() => navigate({ index: MOBILE_INDEX, key: VIEW_KEYS.ACTIVITY })}
				>
					<Bell
						size={24}
						className="text-black text-opacity-40 transition-all duration-200 ease-in-out group-hover:text-opacity-100"
					/>
				</button>
				<button
					className="group h-8 w-8"
					onClick={() => navigate({ index: MOBILE_INDEX, key: VIEW_KEYS.ACTIVITY })}
				>
					{session && (
						<button
							className="relative h-8 w-8 rounded-sm bg-grayscale-0 transition-all duration-200 ease-in-out"
							onClick={() => {}}
						>
							<motion.div
								initial={{ opacity: 0 }}
								animate={{ opacity: 1 }}
								exit={{ opacity: 0 }}
								transition={{ duration: 0.2 }}
							>
								{avatar ? (
									<Image
										src={avatar}
										alt="ENS Avatar"
										width={64}
										height={64}
										className="h-full w-full rounded-sm"
									/>
								) : (
									<div className="overflow-hidden rounded-sm">
										<Avatar
											name={session?.address}
											variant="beam"
											size={"100%"}
											colors={["#00E100", "#A3F700"]}
											square
										/>
									</div>
								)}
							</motion.div>
						</button>
					)}
				</button>
			</div>{" "}
		</div>
	)
}
