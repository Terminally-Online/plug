import { FC, memo, useState } from "react"

import { TokenImage } from "@/components/app/sockets/tokens/token-image"
import { Accordion } from "@/components/shared/utils/accordion"
import { Counter } from "@/components/shared/utils/counter"
import { cn, getTextColor, getZerionTokenIconUrl, ZerionPosition } from "@/lib"
import { useColumnActions } from "@/state/columns"

import { SwapFrame } from "../../frames/assets/swap/frame"
import { TokenFrame } from "../../frames/assets/token/frame"
import { TransferAmountFrame } from "../../frames/assets/transfer/amount/frame"
import { TransferRecipientFrame } from "../../frames/assets/transfer/recipient/frame"

const DEFAULT_TOKEN_COLOR = "#ffffff"

type SocketTokenItemProps = {
	index: number
	token: ZerionPosition | undefined
}&
	React.HTMLAttributes<HTMLDivElement>

export const SocketTokenItem: FC<SocketTokenItemProps> = memo(({ index, token, ...props }) => {
	const { frame } = useColumnActions(index, `${token?.attributes.fungible_info.symbol}-token`)

	const [color, setColor] = useState(DEFAULT_TOKEN_COLOR)
	const textColor = getTextColor(color)

	const isPlaceholder = token === undefined

	return (
		<>
			<Accordion loading={isPlaceholder} onExpand={isPlaceholder ? () => {} : () => frame()} {...props}>
				{token && (
					<div className="flex w-full flex-row items-center gap-4">
						<div className="relative h-10 min-w-10">
							{token.attributes.fungible_info.implementations &&
								token.attributes.fungible_info.implementations.length > 0 && (
									<TokenImage
										logo={getZerionTokenIconUrl(token)}
										symbol={token.attributes.fungible_info.symbol}
										handleColor={setColor}
									/>
								)}
						</div>

						<div className="flex w-full flex-col items-center truncate overflow-ellipsis tabular-nums">
							<div className="flex w-full flex-row font-bold">
								<p className="truncate whitespace-nowrap font-bold">
									{token.attributes.fungible_info.name}
								</p>
								<div className="ml-auto flex flex-row items-center">
									$
									<Counter
										count={(token.attributes.value ?? token.attributes.price ?? 0).toLocaleString(
											"en-US",
											{
												minimumFractionDigits: 2,
												maximumFractionDigits: 2
											}
										)}
										decimals={2}
									/>
								</div>
							</div>

							<div className="flex w-full flex-row gap-4 font-bold">
								<div className="flex flex-row items-center gap-2 truncate overflow-ellipsis">
									<div className="flex flex-row items-center gap-1 truncate text-sm opacity-40">
										<Counter count={token.attributes.quantity.float ?? 0} />
										<p className="whitespace-nowrap">
											{token.attributes.fungible_info.symbol?.toUpperCase()}
										</p>
									</div>
								</div>

								<div
									className={cn(
										"ml-auto flex flex-row items-center text-sm",
										token.attributes.changes?.percent_1d === undefined
											? "opacity-60"
											: token.attributes.changes?.percent_1d >= 0
												? "text-chart-green"
												: "text-plug-red"
									)}
								>
									<>
										{token.attributes.changes?.percent_1d ? (
											<>
												<Counter count={token.attributes.changes?.percent_1d} decimals={2} />%
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

			{token && (
				<>
					<TokenFrame index={index} token={token} color={color} textColor={textColor} />

					<TransferRecipientFrame index={index} token={token} />
					<TransferAmountFrame index={index} token={token} color={color} textColor={textColor} />

					<SwapFrame index={index} tokenOut={token} color={color} textColor={textColor} />
				</>
			)}
		</>
	)
})

SocketTokenItem.displayName = "SocketTokenItem"
