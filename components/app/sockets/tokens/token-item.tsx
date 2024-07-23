import { FC, useCallback, useState } from "react"

import Image from "next/image"

import { motion } from "framer-motion"

import { Accordion, Counter, SocketTokenPercentages } from "@/components"
import { cn, getChainImage } from "@/lib"
import { RouterOutputs } from "@/server/client"

type Props = {
	token?: NonNullable<RouterOutputs["socket"]["tokens"]>[number]
	handleSelect?: (
		token: NonNullable<RouterOutputs["socket"]["tokens"]>[number]
	) => void
}

export const SocketTokenItem: FC<Props> = ({ token, handleSelect }) => {
	const [expanded, setExpanded] = useState(false)

	const priceChange = token?.chains[0].change

	const handleClick = useCallback(() => {
		if (token === undefined) return

		if (handleSelect !== undefined) handleSelect(token)
		else setExpanded(!expanded)
	}, [token, expanded, handleSelect])

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
				loading={token === undefined}
				expanded={expanded}
				onExpand={token === undefined ? () => {} : handleClick}
				accordion={
					token && (
						<span className="relative flex w-full flex-col gap-2 border-t-[1px] border-grayscale-100 pt-4">
							{token.chains.map((chain, index) => (
								<div
									key={index}
									className="flex flex-row items-center gap-4"
								>
									<Image
										src={getChainImage(chain.chainId)}
										alt="Ethereum"
										className="h-4 w-4 rounded-full"
										width={16}
										height={16}
									/>

									<p className="mr-auto font-bold">
										{chain.chainName}
									</p>

									<p className="flex flex-col tabular-nums opacity-60">
										<Counter
											count={
												expanded
													? chain.balanceFormatted
													: 0
											}
										/>
									</p>

									<p className="flex min-w-[60px] flex-row items-center text-right font-bold tabular-nums">
										<Counter
											count={
												expanded ? chain.percentage : 0
											}
										/>
										%
									</p>
								</div>
							))}
						</span>
					)
				}
			>
				{token === undefined ? (
					<div className="invisible">
						<p>.</p>
						<p>.</p>
					</div>
				) : (
					<div className="flex w-full flex-row items-center gap-4 text-left tabular-nums">
						<div className="relative h-10 w-10">
							<Image
								src={token.logoURI ?? ""}
								alt={token.symbol}
								className="absolute left-1/2 top-1/2 h-48 w-48 -translate-x-1/2 -translate-y-1/2 animate-fade-in rounded-full blur-2xl filter transition-all duration-200 ease-in-out"
								width={140}
								height={140}
							/>
							<Image
								src={token.logoURI ?? ""}
								alt={token.symbol}
								className="absolute left-1/2 top-1/2 h-10 w-10 min-w-10 -translate-x-1/2 -translate-y-1/2 animate-fade-in rounded-full bg-grayscale-100"
								width={140}
								height={140}
							/>
						</div>

						<div className="flex flex-col text-left">
							<p className="font-bold">{token.name}</p>
							<p className="flex flex-row items-center gap-2">
								<SocketTokenPercentages chains={token.chains} />
								<span className="flex flex-row items-center gap-2 text-sm opacity-60">
									<Counter count={token.balanceFormatted} />
									{token.symbol}
								</span>
							</p>
						</div>

						<div className="ml-auto flex flex-col text-right">
							<span className="flex flex-row items-center gap-2 font-bold">
								<span className="ml-auto flex flex-row items-center">
									{token.totalValue ? (
										<>
											$
											<Counter count={token.totalValue} />
										</>
									) : (
										"-"
									)}
								</span>
							</span>

							<span
								className={cn(
									"text-sm",
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
											<Counter count={priceChange} />%
										</>
									) : (
										"-"
									)}
								</span>
							</span>
						</div>
					</div>
				)}
			</Accordion>
		</motion.div>
	)
}
