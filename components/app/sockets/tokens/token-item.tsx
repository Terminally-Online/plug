import { FC, useMemo } from "react"

import { motion } from "framer-motion"

import { Accordion, Counter, SocketTokenPercentages } from "@/components"
import { useFrame } from "@/contexts"
import { cn } from "@/lib"
import { RouterOutputs } from "@/server/client"

import { TokenFrame } from "../../frames/assets/token"
import { TokenImage } from "./token-image"

export const SocketTokenItem: FC<{
	id: string
	token?: NonNullable<RouterOutputs["socket"]["balances"]["positions"]>["tokens"][number]
}> = ({ id, token }) => {
	const { handleFrame } = useFrame({
		id,
		key: `token/${token?.symbol ?? ""}`
	})

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
				<Accordion loading={token === undefined} expanded={false} onExpand={token === undefined ? () => {} : () => handleFrame()}>
					{token === undefined ? (
						<div className="invisible">
							<p>.</p>
							<p>.</p>
						</div>
					) : (
						<div className="flex flex-row items-center gap-4">
							<TokenImage logo={token?.icon ?? ""} symbol={token?.symbol} />

							<div className="flex w-full flex-col items-center tabular-nums">
								<div className="flex w-full flex-row font-bold">
									<p>{token?.name}</p>
									<div className="ml-auto flex flex-row items-center">
										{token.value && (
											<>
												$
												<Counter count={token.value} decimals={2} />
											</>
										)}
									</div>
								</div>

								<div className="flex w-full flex-row font-bold">
									<div className="flex flex-row items-center gap-2">
										<SocketTokenPercentages implementations={token.implementations} />
										<div className="flex w-max flex-row items-center gap-1 text-sm opacity-40">
											<Counter count={token.balance} />
											<p className="w-max whitespace-nowrap">{token.symbol.toUpperCase()}</p>
										</div>
									</div>

									<div
										className={cn(
											"ml-auto flex flex-row items-center text-sm",
											token.change === undefined
												? "opacity-60"
												: token.change > 0
													? "text-plug-green"
													: "text-red-500"
										)}
									>
										<>
											{token.change !== undefined ? (
												<>
													<Counter count={token.change} decimals={2} />%
												</>
											) : (
												"-"
											)}
										</>
									</div>
								</div>
							</div>
						</div>
					)}
				</Accordion>
			</motion.div>

			{token && <TokenFrame id={id} symbol={token.symbol} />}
		</>
	)
}
