import { FC } from "react"

import { Bell, Waypoints } from "lucide-react"

import { useAtom, useAtomValue } from "jotai"

import { Frame } from "@/components/app/frames/base"
import { ChainImage } from "@/components/app/sockets/chains/chain.image"
import { TokenImage } from "@/components/app/sockets/tokens/token-image"
import { Counter } from "@/components/shared/utils/counter"
import { cn, getChainId } from "@/lib"
import { RouterOutputs } from "@/server/client"
import { columnByIndexAtom, isFrameAtom, useColumnActions } from "@/state/columns"

type Token =
	| NonNullable<RouterOutputs["socket"]["balances"]["positions"]>["tokens"][number]
	| NonNullable<RouterOutputs["solver"]["tokens"]["get"]>[number]

type TransferSuccessFrame = {
	index: number
	token: Token
	color: string
	textColor: string
}

export const TransferSuccessFrame: FC<TransferSuccessFrame> = ({ index, token, color, textColor }) => {
	const [column] = useAtom(columnByIndexAtom(index))
	const frameKey = `${token.symbol}-transfer-success`
	const isFrame = useAtomValue(isFrameAtom)(column, frameKey)
	const { frame } = useColumnActions(index, frameKey)

	return (
		<Frame
			index={index}
			icon={
				<div className="relative h-8 w-10">
					<TokenImage
						logo={
							token?.icon ||
							`https://token-icons.llamao.fi/icons/tokens/${getChainId(token.implementations[0].chain)}/${token.implementations[0].contract}?h=240&w=240`
						}
						symbol={token.symbol}
						size="sm"
					/>
				</div>
			}
			label="Transfer Confirmed"
			visible={isFrame}
			handleBack={() => frame(`${token.symbol}-token`)}
			hasOverlay
		>
			<div className="flex flex-col">
				<span className="text-left font-bold">
					<span className="opacity-40">Your transfer of</span>{" "}
					<span className="mr-1 inline-flex w-min tabular-nums">
						<Counter count={column?.transfer?.precise ?? ""} />{" "}
					</span>
					<div
						className="inline-flex items-center rounded-xs px-2"
						style={{ backgroundColor: color, color: textColor }}
					>
						{token?.symbol}
					</div>{" "}
					<span className="opacity-40">
						has been confirmed! Any updates to your token balances or positions will be reflected shortly.
					</span>
				</span>

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
						Success
					</p>

					<p className="flex w-full flex-row items-center gap-4 font-bold">
						<Waypoints size={18} className="opacity-20" />
						<span className="mr-auto opacity-40">Chain</span>
						<span className="flex flex-row items-center gap-2">
							<ChainImage chainId={getChainId("base")} size="xs" />
							Base
						</span>
					</p>
				</div>

				<button
					className={cn(
						"mt-4 flex w-full items-center justify-center gap-2 rounded-lg border-[1px] px-12 py-4 font-bold transition-all duration-200 ease-in-out hover:opacity-90 hover:brightness-105"
					)}
					style={{
						backgroundColor: color,
						color: textColor,
						borderColor: color
					}}
					onClick={() => frame()}
				>
					Done
				</button>
			</div>
		</Frame>
	)
}
