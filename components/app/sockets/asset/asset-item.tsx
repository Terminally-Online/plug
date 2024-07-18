import { FC, useCallback, useMemo, useState } from "react"

import Image from "next/image"

import { AssetPercentages } from "@/components/app"
import { AccordionContent, Counter } from "@/components/utils"
import { useBalances } from "@/contexts"
import { cn } from "@/lib"
import { getChainImage } from "@/lib/functions"
import { PriceData } from "@/lib/functions/llama/price"

type Props = {
	token: NonNullable<ReturnType<typeof useBalances>["balances"]>[number]
	priceData: undefined | PriceData
	handleSelect?: (
		token: NonNullable<ReturnType<typeof useBalances>["balances"]>[number]
	) => void
}

export const SocketAssetItem: FC<Props> = ({
	token,
	priceData,
	handleSelect
}) => {
	const [expanded, setExpanded] = useState(false)

	const priceChange = priceData
		? priceData[
				`${token.chains[0].chainName.toLowerCase()}:${token.chains[0].address}`
			]?.change
		: undefined

	const totalValue = useMemo(() => {
		if (priceData === undefined) return undefined

		return token.chains
			.map(token => {
				const coinKey =
					`${token.chainName.toLowerCase()}:${token.address}` as const

				const price =
					token.balanceFormatted * priceData[coinKey]?.price ?? 0

				return price
			})
			.reduce((acc, price) => acc + price, 0)
	}, [priceData, token])

	const handleClick = useCallback(() => {
		if (handleSelect !== undefined) handleSelect(token)
		else setExpanded(!expanded)
	}, [token, expanded, handleSelect])

	return (
		<div className="flex cursor-pointer flex-col gap-4">
			<button
				className={cn(
					"group group flex h-min w-full cursor-pointer flex-col items-center overflow-hidden rounded-[16px] border-[1px] border-grayscale-0 p-4 transition-all duration-200 ease-in-out",
					expanded
						? "bg-grayscale-0 hover:bg-white"
						: "bg-white hover:border-white hover:bg-grayscale-0"
				)}
				onClick={handleClick}
			>
				<span className="flex w-full flex-row items-center gap-4 tabular-nums">
					<span className="relative h-10 w-10">
						<Image
							src={token.logoURI ?? ""}
							alt={token.symbol}
							className="animate-fade-in absolute left-1/2 top-1/2 h-48 w-48 -translate-x-1/2 -translate-y-1/2 rounded-full blur-2xl filter transition-all duration-200 ease-in-out"
							width={72}
							height={72}
						/>
						<Image
							src={token.logoURI ?? ""}
							alt={token.symbol}
							className="animate-fade-in absolute left-1/2 top-1/2 h-10 w-10 min-w-10 -translate-x-1/2 -translate-y-1/2 rounded-full bg-grayscale-100"
							width={48}
							height={48}
						/>
					</span>

					<span className="flex flex-col">
						<span className="mr-auto font-bold">{token.name}</span>
						<span className="mr-auto flex flex-row items-center gap-2">
							<AssetPercentages chains={token.chains} />
							<span className="flex flex-row items-center gap-2 text-sm opacity-60">
								<Counter count={token.balanceFormatted} />
								{token.symbol}
							</span>
						</span>
					</span>

					<span className="ml-auto flex flex-col text-right">
						<span className="ml-auto">
							<span className="ml-auto flex flex-row items-center gap-2 text-right font-bold">
								<span className="ml-auto flex flex-row items-center">
									{totalValue ? (
										<>
											$
											<Counter count={totalValue} />
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
						</span>
					</span>
				</span>

				<AccordionContent expanded={expanded}>
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
										count={expanded ? chain.percentage : 0}
									/>
									%
								</p>
							</div>
						))}
					</span>
				</AccordionContent>
			</button>
		</div>
	)
}
