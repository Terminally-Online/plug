import React, { FC, useMemo, useState } from "react"

import { ExternalLink, EyeOff } from "lucide-react"

import { useAtom, useAtomValue } from "jotai"

import { Frame } from "@/components/app/frames/base"
import { SocketTokenPriceChart } from "@/components/app/sockets/tokens/token-chart"
import { TokenImage } from "@/components/app/sockets/tokens/token-image"
import { Image } from "@/components/app/utils/image"
import { Counter } from "@/components/shared/utils/counter"
import { chains, cn, formatTitle, getChainId, getTextColor } from "@/lib"
import { RouterOutputs } from "@/server/client"
import { columnByIndexAtom, isFrameAtom } from "@/state/columns"

export const PositionFrame: FC<{
	index: number
	protocol: RouterOutputs["socket"]["balances"]["positions"]["protocols"][number]
}> = ({ index, protocol }) => {
	const { positions } = protocol

	const [column] = useAtom(columnByIndexAtom(index))
	const frameKey = `${protocol.name}-position`
	const isFrame = useAtomValue(isFrameAtom)(column, frameKey)

	const [color, setColor] = useState("")
	const [colors, setColors] = useState<Record<string, string>>({})
	const [excludedKeys, setExcludedKeys] = useState<Array<string>>([])

	const textColor = getTextColor(color ?? "#ffffff")

	const groupedPositions = useMemo(() => {
		const grouped: Record<
			string,
			Array<RouterOutputs["socket"]["balances"]["positions"]["protocols"][number]["positions"][number]>
		> = {}

		positions.forEach(position => {
			if (grouped[position.type] === undefined) {
				grouped[position.type] = []
			}

			grouped[position.type].push(position)
		})

		return grouped
	}, [positions])

	const keys = useMemo(() => {
		return positions
			.filter(position => Boolean(position.fungible))
			.map(position => {
				const contract =
					position.fungible.implementations.find(implementation => implementation.chain === position.chain)
						?.contract ?? ""

				return {
					chain: position.chain,
					contract
				}
			})
			.filter(position => position.contract !== "")
			.map(position => `${position.chain}:${position.contract}`)
			.filter(key => !excludedKeys.includes(key))
	}, [positions, excludedKeys])

	return (
		<Frame
			index={index}
			className="max-h-[85vh] overflow-y-auto overflow-x-hidden"
			icon={<TokenImage logo={protocol?.icon ?? ""} symbol={protocol.name} size="sm" handleColor={setColor} />}
			label={protocol.name}
			visible={isFrame}
			hasChildrenPadding={false}
			hasOverlay
		>
			<React.Fragment>
				{keys.length > 0 && (
					<SocketTokenPriceChart
						enabled={isFrame}
						keys={keys}
						colors={colors}
						// handleHeader={setHeader}
						// handleTooltip={setTooltipData}
					/>
				)}

				<div className={cn("relative mb-4 flex flex-col gap-4 px-6 font-bold", keys.length > 0 && "pt-4")}>
					<a
						className="flex w-full items-center justify-center gap-2 rounded-lg bg-plug-green/10 py-4 font-bold transition-all duration-200 ease-in-out hover:opacity-90"
						style={{
							backgroundColor: color ?? "",
							color: textColor
						}}
						href={protocol.url}
						target="_blank"
						rel="noreferrer"
					>
						<ExternalLink size={14} className="opacity-60" />
						Manage
					</a>

					{Object.keys(groupedPositions).map(type => (
						<div key={type}>
							<div className="mb-2 flex flex-row items-center gap-2">
								<p className="font-bold opacity-40">{formatTitle(type)}</p>
								<div className="h-[2px] w-full bg-plug-green/10" />
								<p className="font-bold opacity-40">Value</p>
							</div>

							<div className="flex flex-col gap-2">
								{groupedPositions[type].map((position, positionIndex) => {
									if (!position.fungible) return null

									const implementation = position.fungible.implementations.find(
										implementation => implementation.chain === position.chain
									)
									const key = `${position.chain}:${implementation?.contract}`

									return (
										<div key={`${type}-${positionIndex}`}>
											<div className="flex flex-row items-center gap-4">
												<button
													className="relative"
													onClick={() =>
														setExcludedKeys(prev =>
															prev.includes(key)
																? prev.filter(excludedKey => excludedKey !== key)
																: [...prev, key]
														)
													}
												>
													<TokenImage
														logo={position.fungible.icon ?? ""}
														symbol={position.fungible.symbol}
														blur={false}
														handleColor={
															Object.keys(colors).includes(key)
																? undefined
																: color =>
																		setColors(prev =>
																			prev
																				? {
																						...prev,
																						[key]: color
																					}
																				: {
																						[key]: color
																					}
																		)
														}
													/>

													<div
														className={cn(
															"absolute left-1/2 top-1/2 z-[20] flex h-10 w-10 -translate-x-1/2 -translate-y-1/2 animate-fade-in items-center justify-center rounded-full transition-all duration-200 ease-in-out",
															excludedKeys.includes(key)
																? "bg-white/90 hover:bg-white/80"
																: "hover:bg-white/20"
														)}
													>
														{excludedKeys.includes(key) && (
															<EyeOff size={14} className="opacity-40" />
														)}
													</div>
												</button>

												<div className="flex w-full flex-col gap-0">
													<div className="flex flex-row items-center justify-between gap-2">
														<p>{position.fungible.name}</p>
														<div className="flex flex-row">
															<p>$</p>
															<Counter count={position.value ?? 0} />
														</div>
													</div>

													<div className="relative flex flex-row items-center justify-between gap-2 text-sm text-black/40">
														<Image
															src={chains[getChainId(position.chain)].logo}
															alt={position.chain}
															className="h-4 w-4 rounded-full"
															width={24}
															height={24}
														/>
														<div>
															<Counter count={position.balance ?? 0} />
														</div>
														<div
															className={cn(
																"ml-auto flex flex-row items-center text-sm",
																position.change === undefined
																	? "opacity-60"
																	: position.change > 0
																		? "text-plug-green"
																		: "text-red-500"
															)}
														>
															<>
																{position.change !== undefined ? (
																	<>
																		<Counter count={position.change} decimals={2} />
																		<p>%</p>
																	</>
																) : (
																	"-"
																)}
															</>
														</div>
													</div>
												</div>
											</div>
										</div>
									)
								})}
							</div>
						</div>
					))}
				</div>
			</React.Fragment>
		</Frame>
	)
}
