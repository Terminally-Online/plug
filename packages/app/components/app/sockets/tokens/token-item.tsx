import React, { FC, memo, useState } from "react"

import { SwapFrame } from "@/components/app/frames/assets/swap"
import { TokenFrame } from "@/components/app/frames/assets/token"
import { TransferFrame } from "@/components/app/frames/assets/transfer"
import { TokenImage } from "@/components/app/sockets/tokens/token-image"
import { Accordion } from "@/components/shared/utils/accordion"
import { Counter } from "@/components/shared/utils/counter"
import { cn, getChainId, getTextColor } from "@/lib"
import { RouterOutputs } from "@/server/client"
import { useColumnStore } from "@/state/columns"

import { ChainImage } from "../chains/chain.image"

const DEFAULT_TOKEN_COLOR = "#ffffff"

type SocketTokenItemProps = {
	index: number
	token?: NonNullable<RouterOutputs["socket"]["balances"]["positions"]>["tokens"][number]
	isListToken?: boolean
}

export const SocketTokenItem: FC<SocketTokenItemProps> = memo(({ index, token, isListToken }) => {
	const { handle } = useColumnStore(index, `${token?.symbol}-token`)
	const [color, setColor] = useState(DEFAULT_TOKEN_COLOR)
	const textColor = getTextColor(color)

	return (
		<>
			<Accordion loading={token === undefined} onExpand={token === undefined ? () => {} : () => handle.frame()}>
				{token === undefined ? (
					<div className="invisible">
						<p>.</p>
						<p>.</p>
					</div>
				) : (
					<div className="flex w-full flex-row items-center gap-4">
						<div className="relative h-10 min-w-10">
							{token.implementations && token.implementations.length > 0 && (
								<TokenImage
									logo={
										// @ts-ignore
										token.icon?.url ||
										token.icon ||
										`https://token-icons.llamao.fi/icons/tokens/${getChainId(token.implementations[0].chain)}/${token.implementations[0].contract}?h=240&w=240`
									}
									symbol={token.symbol}
									handleColor={setColor}
								/>
							)}
						</div>

						<div className="flex w-full flex-col items-center truncate overflow-ellipsis tabular-nums">
							<div className="flex w-full flex-row font-bold">
								<p className="truncate whitespace-nowrap font-bold">{token.name}</p>
								<p className="ml-auto flex flex-row items-center">
									$
									<Counter
										count={(token.value ?? token.price ?? 0).toLocaleString("en-US", {
											minimumFractionDigits: 2,
											maximumFractionDigits: 2
										})}
										decimals={2}
									/>
								</p>
							</div>

							<div className="flex w-full flex-row gap-4 font-bold">
								<div className="flex flex-row items-center gap-2 truncate overflow-ellipsis">
									<div className="flex flex-row items-center">
										{token.implementations?.map((implementation, implementationIndex) => (
											<ChainImage
												key={implementationIndex}
												chainId={getChainId(implementation.chain)}
												size="xs"
											/>
										))}
									</div>
									<div className="flex flex-row items-center gap-1 truncate text-sm opacity-40">
										<Counter count={token.balance ?? 0} />
										<p className="whitespace-nowrap">{token.symbol?.toUpperCase()}</p>
									</div>
								</div>

								<p
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
										{token.change ? (
											<>
												<Counter count={token.change} decimals={2} />%
											</>
										) : (
											"-"
										)}
									</>
								</p>

								{/* <pre className="text-left text-xs">
									{JSON.stringify({ ...token, description: "" }, null, 2)}
								</pre> */}
							</div>
						</div>
					</div>
				)}
			</Accordion>

			{token && (
				<>
					<TokenFrame index={index} token={token} color={color} textColor={textColor} />
					<TransferFrame index={index} token={token} color={color} textColor={textColor} />
					<SwapFrame index={index} tokenOut={token} color={color} textColor={textColor} />
				</>
			)}
		</>
	)
})

SocketTokenItem.displayName = "SocketTokenItem"
