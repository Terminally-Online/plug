import React, { FC, useMemo, useState } from "react"

import { ExternalLink } from "lucide-react"

import { useAtom, useAtomValue } from "jotai"

import { Frame } from "@/components/app/frames/base"
import { TokenImage } from "@/components/app/sockets/tokens/token-image"
import { Accordion } from "@/components/shared/utils/accordion"
import { Counter } from "@/components/shared/utils/counter"
import { cn, getTextColor, getZerionTokenIconUrl } from "@/lib"
import { RouterOutputs } from "@/server/client"
import { columnByIndexAtom, COLUMNS, isFrameAtom } from "@/state/columns"

export const PositionFrame: FC<{
	index: number
	protocols?: NonNullable<RouterOutputs["service"]["zerion"]["wallet"]["positions"]>["data"]
}> = ({ index, protocols }) => {
	const protocol = protocols?.at(0)
	const [column] = useAtom(columnByIndexAtom(index))
	const frameKey = `${protocol?.relationships.dapp?.data.id ?? ""}-position`
	const isFrame = useAtomValue(isFrameAtom)(column, frameKey)

	const [color, setColor] = useState("#ffffff")
	const textColor = getTextColor(color)

	const positions = protocols

	const groupedPositions = useMemo(() => {
		if (!positions) return {}
		const grouped: Record<string, typeof positions> = {}
		positions.forEach(position => {
			const type = position.attributes?.position_type || "unknown"
			if (!grouped[type]) grouped[type] = []
			grouped[type].push(position)
		})
		return grouped
	}, [positions])

	return (
		<Frame
			index={index}
			className="max-h-[85vh] overflow-y-auto overflow-x-hidden"
			label={protocol?.attributes?.application_metadata?.name ?? ""}
			icon={
				<TokenImage
					logo={protocol?.attributes?.application_metadata?.icon?.url ?? ""}
					symbol={protocol?.attributes?.application_metadata?.name ?? protocol?.id ?? ""}
					size="sm"
					handleColor={setColor}
				/>
			}
			visible={isFrame}
			hasChildrenPadding={false}
			hasOverlay
		>
			<div className="py-4">
				<div className={cn("relative flex flex-col gap-4 px-6 font-bold")}>
					{index === COLUMNS.SIDEBAR_INDEX && (
						<a
							className="flex w-full items-center justify-center gap-2 rounded-lg bg-plug-green/10 py-4 font-bold transition-all duration-200 ease-in-out hover:opacity-90"
							style={{
								backgroundColor: color ?? "",
								color: textColor
							}}
							href={protocol?.attributes?.application_metadata?.url ?? undefined}
							target="_blank"
							rel="noreferrer"
						>
							<ExternalLink size={14} className="opacity-60" />
							Manage
						</a>
					)}

					{Object.entries(groupedPositions).map(([type, group]) => (
						<div key={type}>
							<div className="mb-2 flex flex-row items-center gap-2">
								<p className="font-bold opacity-40">{type.charAt(0).toUpperCase() + type.slice(1)}</p>
								<div className="h-[1px] w-full bg-plug-green/10" />
							</div>
							<div className="relative flex flex-col gap-2">
								{group.map((position, idx) => {
									const change = position.attributes?.changes?.percent_1d
									return (
										<Accordion key={position.id + idx}>
											<div className="flex w-full flex-row items-center gap-4">
												<div className="relative h-10 min-w-10">
													<TokenImage
														logo={getZerionTokenIconUrl(position)}
														symbol={position.attributes.fungible_info.symbol}
													/>
												</div>
												<div className="flex w-full flex-col items-center truncate overflow-ellipsis tabular-nums">
													<div className="flex w-full flex-row font-bold">
														<p className="truncate whitespace-nowrap font-bold">
															{position.attributes.fungible_info.name}
														</p>
														<div className="ml-auto flex flex-row items-center">
															$
															<Counter
																count={(
																	position.attributes.value ??
																	position.attributes.price ??
																	0
																).toLocaleString("en-US", {
																	minimumFractionDigits: 2,
																	maximumFractionDigits: 2
																})}
																decimals={2}
															/>
														</div>
													</div>
													<div className="flex w-full flex-row gap-4 font-bold">
														<div className="flex flex-row items-center gap-2 truncate overflow-ellipsis">
															<div className="flex flex-row items-center gap-1 truncate text-sm opacity-40">
																<Counter
																	count={position.attributes.quantity.float ?? 0}
																/>
																<p className="whitespace-nowrap">
																	{position.attributes.fungible_info.symbol?.toUpperCase()}
																</p>
															</div>
														</div>
														<div
															className={cn(
																"ml-auto flex flex-row items-center text-sm",
																change === undefined
																	? "opacity-60"
																	: change >= 0
																		? "text-chart-green"
																		: "text-plug-red"
															)}
														>
															<>
																{typeof change === "number" ? (
																	<>
																		<Counter count={change} decimals={2} />%
																	</>
																) : (
																	"-"
																)}
															</>
														</div>
													</div>
												</div>
											</div>
										</Accordion>
									)
								})}
							</div>
						</div>
					))}
				</div>
			</div>
		</Frame>
	)
}
