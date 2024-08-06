import { FC, useCallback, useState } from "react"

import Image from "next/image"

import { motion } from "framer-motion"

import { Accordion, Counter, SocketTokenPercentages } from "@/components"
import { useFrame } from "@/contexts"
import { cn, formatTitle, getChainImage } from "@/lib"
import { RouterOutputs } from "@/server/client"

import { TokenFrame } from "../../frames/assets/token"
import { TokenImage } from "./token-image"

export const SocketTokenItem: FC<{
	token?: NonNullable<RouterOutputs["socket"]["balances"]["tokens"]>[number]
	handleSelect?: (
		token: NonNullable<
			RouterOutputs["socket"]["balances"]["tokens"]
		>[number]
	) => void
}> = ({ token, handleSelect }) => {
	const { handleFrameVisible } = useFrame()

	const priceChange = token?.chains[0].change

	return (
		<>
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
					loading={token === undefined}
					expanded={false}
					onExpand={
						token === undefined
							? () => {}
							: () => handleFrameVisible(`token-${token.symbol}`)
					}
				>
					{token === undefined ? (
						<div className="invisible">
							<p>.</p>
							<p>.</p>
						</div>
					) : (
						<div className="flex flex-row items-center gap-4">
							<TokenImage
								logo={token.logo}
								symbol={token.symbol}
							/>

							<div className="flex w-full flex-col items-center tabular-nums">
								<div className="flex w-full flex-row font-bold">
									<p>{token.name}</p>
									<p className="ml-auto flex flex-row items-center">
										{token.value ? (
											<>
												$
												<Counter
													count={token.value}
													decimals={2}
												/>
											</>
										) : (
											"-"
										)}
									</p>
								</div>

								<div className="flex w-full flex-row">
									<p className="flex flex-row items-center gap-2">
										<SocketTokenPercentages
											chains={token.chains}
										/>
										<span className="flex w-max flex-row items-center gap-1 text-sm opacity-60">
											<Counter count={token.balance} />
											<span className="w-max whitespace-nowrap">
												{token.symbol.toUpperCase()}
											</span>
										</span>
									</p>

									<p
										className={cn(
											"ml-auto text-sm",
											priceChange === undefined
												? "opacity-60"
												: priceChange > 0
													? "text-plug-green"
													: "text-red-500"
										)}
									>
										<span className="ml-auto flex flex-row items-center">
											{priceChange !== undefined ? (
												<>
													<Counter
														count={priceChange}
														decimals={2}
													/>
													%
												</>
											) : (
												"-"
											)}
										</span>
									</p>
								</div>
							</div>
						</div>
					)}
				</Accordion>
			</motion.div>

			<TokenFrame symbol={token?.symbol ?? ""} />
		</>
	)
}
