import { FC, useEffect, useMemo, useState } from "react"

import { Bell, Waypoints } from "lucide-react"

import { Frame, Image, TokenImage } from "@/components"
import { chains, cn, formatTitle, getChainId, getTextColor } from "@/lib"
import { RouterOutputs } from "@/server/client"
import { useColumnStore } from "@/state/columns"

type Token =
	| NonNullable<RouterOutputs["socket"]["balances"]["positions"]>["tokens"][number]
	| NonNullable<RouterOutputs["solver"]["tokens"]["get"]>[number]

type SwapConfirmFrameProps = {
	index: number
	tokenIn: Token
	tokenOut: Token
}

export const SwapConfirmFrame: FC<SwapConfirmFrameProps> = ({ index, tokenIn, tokenOut }) => {
	const {
		isFrame,
		handle: { frame }
	} = useColumnStore(index, `${tokenOut.symbol}-${tokenIn.symbol}-swap-confirm`)

	const [tokenOutColor, setTokenOutColor] = useState("#000000")
	const [tokenInColor, setTokenInColor] = useState("#000000")

	return (
		<Frame
			index={index}
			icon={
				<div className="-gap-2 relative flex flex-row items-center">
					<div className="relative h-8 w-10">
						<TokenImage
							logo={
								tokenOut?.icon ||
								`https://token-icons.llamao.fi/icons/tokens/${getChainId(tokenOut.implementations[0].chain)}/${tokenOut.implementations[0].contract}?h=240&w=240`
							}
							symbol={tokenOut.symbol}
							size="sm"
							handleColor={setTokenOutColor}
						/>
					</div>
					<div className="relative -ml-4 h-8 w-10">
						<TokenImage
							logo={
								tokenIn?.icon ||
								`https://token-icons.llamao.fi/icons/tokens/${getChainId(tokenIn.implementations[0].chain)}/${tokenIn.implementations[0].contract}?h=240&w=240`
							}
							symbol={tokenIn.symbol}
							size="sm"
							handleColor={setTokenInColor}
						/>
					</div>
				</div>
			}
			label="Execution Queued"
			visible={isFrame}
			handleBack={() => frame(`${tokenOut.symbol}-token`)}
			hasOverlay
		>
			<div className="flex flex-col">
				<p className="font-bold">
					<span className="opacity-40">The swap of</span>{" "}
					<span className="relative inline-flex items-center gap-1 rounded-sm px-2">
						<span
							className="absolute inset-0 -z-[1] rounded-sm opacity-10"
							style={{
								background: `linear-gradient(to right, ${tokenOutColor}, ${tokenInColor})`,
								opacity: 0.1
							}}
						/>
						{tokenOut.symbol}
						<span className="opacity-40">→</span>
						{tokenIn.symbol}
					</span>{" "}
					<span className="opacity-40">
						has been successfully queued. Your intent will now be simulated and run as soon as it can.
					</span>
				</p>

				<div className="mb-2 mt-4 flex flex-row items-center gap-4">
					<p className="font-bold opacity-40">Transaction</p>
					<div className="h-[2px] w-full bg-plug-green/10" />
				</div>

				<div className="flex flex-col">
					<p className="flex flex-row justify-between font-bold">
						<span className="flex w-full flex-row items-center gap-4">
							<Bell size={18} className="opacity-20" />
							<span className="opacity-40">Status</span>
						</span>{" "}
						{formatTitle("Running")}
					</p>

					<p className="flex w-full flex-row items-center gap-4 font-bold">
						<Waypoints size={18} className="opacity-20" />
						<span className="mr-auto opacity-40">Chain</span>
						<span className="flex flex-row items-center gap-2">
							<Image className="h-4 w-4" src={chains[1].logo} alt="ethereum" width={24} height={24} />
							Ethereum
						</span>
					</p>
				</div>

				<button
					className={cn(
						"mt-4 flex w-full items-center justify-center gap-2 rounded-lg border-[1px] px-12 py-4 font-bold transition-all duration-200 ease-in-out hover:opacity-90 hover:brightness-105"
					)}
					style={{
						backgroundColor: tokenInColor,
						color: getTextColor(tokenInColor),
						borderColor: tokenInColor
					}}
					onClick={() => frame()}
				>
					Done
				</button>
			</div>
		</Frame>
	)
}
