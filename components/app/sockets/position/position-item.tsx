import { FC } from "react"

import Image from "next/image"

import { motion } from "framer-motion"

import { Accordion, Counter } from "@/components/shared"
import { useFrame } from "@/contexts"
import { cn, formatTitle, getChainImage } from "@/lib"
import { RouterOutputs } from "@/server/client"

export const SocketPositionItem: FC<{
	id: string
	protocol: RouterOutputs["socket"]["balances"]["positions"]["protocols"][number]
}> = ({ id, protocol }) => {
	const { handleFrame } = useFrame({
		id,
		key: `position-${protocol.name}`
	})

	const change =
		protocol.positions.reduce((acc, position) => acc + (position.change ?? 0), 0) /
		protocol.positions.filter(position => position.change !== undefined).length

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
				onExpand={() => handleFrame()}
				accordion={
					<>
						{protocol.positions.map((position, index) => (
							<div key={index} className="relative">
								<div className="flex flex-row items-center gap-2">
									<Image src={position.fungible.icon ?? ""} alt="" width={24} height={24} />
									<div className="flex flex-col font-bold">
										<div className="flex w-full flex-row items-center justify-between gap-2">
											<p>{position.fungible.name}</p>
											<p className="ml-auto w-full text-right">{position.balance}</p>
										</div>
										<p className="mr-auto font-bold opacity-40">{position.type}</p>
									</div>
								</div>
							</div>
						))}
					</>
				}
			>
				<div className="flex w-full flex-row items-center gap-4">
					<div className="relative h-10 min-w-10">
						<Image
							className="absolute left-1/2 top-1/2 h-48 w-48 -translate-x-1/2 -translate-y-1/2 animate-fade-in rounded-full blur-2xl filter transition-all duration-200 ease-in-out"
							src={protocol.icon}
							alt=""
							width={140}
							height={140}
						/>
						<div
							className="absolute left-1/2 top-1/2 h-10 w-10 min-w-10 -translate-x-1/2 -translate-y-1/2 animate-fade-in rounded-full bg-grayscale-100"
							style={{
								backgroundImage: `url(${protocol.icon})`,
								backgroundSize: "cover",
								backgroundPosition: "center",
								backgroundRepeat: "no-repeat"
							}}
						/>
					</div>

					<div className="relative flex w-full flex-col">
						<div className="flex w-full flex-row justify-between">
							<p className="mr-auto font-bold">{protocol.name}</p>
							<div className="ml-auto flex flex-row font-bold">
								<p>$</p>
								<Counter count={protocol.positions.reduce((acc, position) => acc + (position.value ?? 0), 0)} />
							</div>
						</div>
						<div className="flex w-full flex-row items-center gap-2 text-sm font-bold ">
							<p className="opacity-40">{protocol.positions.length} Positions</p>
							<div
								className={cn(
									"ml-auto flex flex-row items-center text-sm",
									change === undefined ? "opacity-60" : change > 0 ? "text-plug-green" : "text-red-500"
								)}
							>
								<>
									{change !== undefined ? (
										<>
											<Counter count={change} decimals={2} />%
										</>
									) : (
										"-"
									)}
								</>
							</div>
						</div>
					</div>
				</div>
			</Accordion>
		</motion.div>
	)
}
