import { FC, useState } from "react"

import Image from "next/image"

import { motion } from "framer-motion"

import { Accordion } from "@/components/shared"
import { useFrame } from "@/contexts"
import { cn } from "@/lib"
import { RouterOutputs } from "@/server/client"

export const SocketPositionItem: FC<{
	id: string
	position: RouterOutputs["socket"]["balances"]["positions"]["defi"][string]
}> = ({ id, position }) => {
	const { handleFrame } = useFrame({
		id,
		key: `position-${position.name}`
	})

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
						{position.positions.map(position => (
							<div key={position.id} className="relative">
								<div className="flex flex-row items-center gap-2">
									<Image
										src={position.fungible.icon}
										alt=""
										width={24}
										height={24}
									/>
									<div className="flex flex-col font-bold">
										<div className="flex w-full flex-row items-center justify-between gap-2">
											<p>{position.fungible.name}</p>
											<p className="ml-auto w-full text-right">
												{position.balance}
											</p>
										</div>
										<p className="mr-auto font-bold opacity-40">
											{position.type}
										</p>
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
							src={position.icon}
							alt=""
							width={140}
							height={140}
						/>
						<div
							className="absolute left-1/2 top-1/2 h-10 w-10 min-w-10 -translate-x-1/2 -translate-y-1/2 animate-fade-in rounded-full bg-grayscale-100"
							style={{
								backgroundImage: `url(${position.icon})`,
								backgroundSize: "cover",
								backgroundPosition: "center",
								backgroundRepeat: "no-repeat"
							}}
						/>
					</div>

					<div className="relative flex flex-col">
						<p className="mr-auto font-bold">{position.name}</p>
						<div className="flex w-full flex-row items-center gap-2">
							{position.assets.map((asset, index) => (
								<Image
									key={asset.name}
									className={cn(
										"h-4 w-4 rounded-full",
										index !== 0 && "-ml-3"
									)}
									src={asset.icon}
									alt=""
									width={24}
									height={24}
								/>
							))}

							<p className="text-sm font-bold opacity-40">
								{position.assets.length} Token
								{position.assets.length > 1 ? "s" : ""}
							</p>
						</div>
					</div>
				</div>
			</Accordion>
		</motion.div>
	)
}
