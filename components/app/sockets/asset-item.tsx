import { type FC, useCallback, useState } from "react"

import Image from "next/image"

import { motion } from "framer-motion"
import { ChevronDown } from "lucide-react"

import { Button } from "@/components/buttons"
import { useBalances } from "@/contexts"
import { getChainImage } from "@/lib/functions"

import { AssetPercentages } from "./asset-percentages"

type Props = {
	token: NonNullable<ReturnType<typeof useBalances>["balances"]>[number]
	handleSelect?: (
		token: NonNullable<ReturnType<typeof useBalances>["balances"]>[number]
	) => void
}

export const SocketAssetItem: FC<Props> = ({ token, handleSelect }) => {
	const [expanded, setExpanded] = useState(false)

	const handleClick = useCallback(() => {
		if (handleSelect !== undefined) handleSelect(token)
		else setExpanded(!expanded)
	}, [token, expanded, handleSelect])

	return (
		<div className="flex cursor-pointer flex-col gap-4">
			<button
				className="group flex h-min w-full cursor-pointer flex-row items-center gap-2"
				onClick={handleClick}
			>
				<Image
					src={token.logoURI ?? ""}
					alt={token.symbol}
					className="h-6 w-6 rounded-full bg-white/20"
					width={48}
					height={48}
				/>

				<span className="flex flex-col font-bold">{token.symbol}</span>

				<span className="ml-auto tabular-nums opacity-60">
					{token.balanceFormatted.toLocaleString()}
				</span>
				<AssetPercentages chains={token.chains} />
				<Button
					variant="secondary"
					className="p-1 group-hover:bg-grayscale-100"
					onClick={handleClick}
				>
					<motion.div
						style={{
							rotate: handleSelect !== undefined ? -90 : 0
						}}
						animate={{
							rotate:
								handleSelect !== undefined
									? -90
									: expanded
										? 180
										: 0
						}}
						transition={{
							duration: 0.2
						}}
					>
						<ChevronDown size={16} className="opacity-60" />
					</motion.div>
				</Button>
			</button>

			{handleSelect === undefined && expanded && (
				<div className="ml-[4px] flex flex-col gap-4">
					{token.chains.map((chain, index) => (
						<div
							key={index}
							className="flex flex-row items-center gap-2"
						>
							<Image
								src={getChainImage(chain.chainId)}
								alt="Ethereum"
								className="h-4 w-4 rounded-full"
								width={16}
								height={16}
							/>
							<p className="mr-auto font-bold opacity-40">
								{chain.chainName}
							</p>

							<p className="tabular-nums opacity-60">
								{chain.balanceFormatted.toLocaleString()}
							</p>
							<p className="w-[56px] text-right tabular-nums">
								{chain.percentage}%
							</p>
						</div>
					))}
				</div>
			)}
		</div>
	)
}
