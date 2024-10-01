import { FC } from "react"

import { RouterOutputs } from "@/server/client"

import { Accordion, Counter, SocketTokenPercentages, TokenFrame, TokenImage } from "@/components"
import { cn, getChainId } from "@/lib"
import { useFrame } from "@/state"

export const SocketTokenItem: FC<{
	index: number
	tokenIndex: number
	token?: NonNullable<RouterOutputs["socket"]["balances"]["positions"]>["tokens"][number]
}> = ({ index, tokenIndex, token }) => {
	const { key, handleFrame } = useFrame({
		index,
		key: `${index}-${tokenIndex}-token`
	})

	return (
		<>
			<Accordion
				loading={token === undefined}
				onExpand={
					token === undefined
						? () => {}
						: () => {
								handleFrame()
								console.log("clicked", key)
							}
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
							logo={
								token?.icon ||
								`https://token-icons.llamao.fi/icons/tokens/${getChainId(token.implementations[0].chain)}/${token.implementations[0].contract}?h=240&w=240`
							}
							symbol={token?.symbol}
						/>

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

							<div className="flex w-full flex-row gap-4 font-bold">
								<div className="flex flex-row items-center gap-2 truncate overflow-ellipsis">
									<SocketTokenPercentages implementations={token.implementations} />
									<div className="flex flex-row items-center gap-1 truncate text-sm opacity-40">
										<Counter count={token.balance} />
										<p className="whitespace-nowrap">{token.symbol.toUpperCase()}</p>
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

			{token !== undefined && <TokenFrame index={index} tokenIndex={tokenIndex} token={token} />}
		</>
	)
}
