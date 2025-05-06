import { FC } from "react"

import { SearchIcon } from "lucide-react"

import { useAtom, useAtomValue } from "jotai"

import { Frame } from "@/components/app/frames/base"
import { Search } from "@/components/app/inputs/search"
import { TokenImage } from "@/components/app/sockets/tokens/token-image"
import { Accordion } from "@/components/shared/utils/accordion"
import { Counter } from "@/components/shared/utils/counter"
import { cn, getZerionTokenIconUrl, useDebounce, ZerionFungible, ZerionPosition } from "@/lib"
import { api } from "@/server/client"
import { columnByIndexAtom, isFrameAtom, useColumnActions } from "@/state/columns"

type SwapTokenFrameProps = {
	index: number
	tokenOut: ZerionPosition
	handleTokenIn: (token: ZerionFungible) => void
}

export const SwapTokenFrame: FC<SwapTokenFrameProps> = ({ index, tokenOut, handleTokenIn }) => {
	const [column] = useAtom(columnByIndexAtom(index))
	const frameKey = `${tokenOut.attributes.fungible_info.symbol}-swap-token`
	const isFrame = useAtomValue(isFrameAtom)(column, frameKey)
	const { frame } = useColumnActions(index, frameKey)

	const [search, debouncedSearch, handleSearch] = useDebounce("", 100)

	const { data } = api.service.zerion.fungibles.list.useQuery(
		{
			query: {
				filter: {
					implementationChainId: "base",
					searchQuery: debouncedSearch || undefined
				}
			}
		},
		{ placeholderData: prev => prev }
	)
	const tokens = data?.data || []

	return (
		<Frame
			index={index}
			icon={
				<div className="relative h-8 w-10">
					<TokenImage
						logo={getZerionTokenIconUrl(tokenOut)}
						symbol={tokenOut.attributes.fungible_info.symbol}
						size="sm"
					/>
				</div>
			}
			label={`Swap ${tokenOut.attributes.fungible_info.symbol}`}
			visible={isFrame}
			handleBack={() => frame(`${tokenOut.attributes.fungible_info.symbol}-token`)}
			hasChildrenPadding={false}
			hasOverlay
		>
			<div className="relative flex h-full flex-col gap-2 px-6 pb-4">
				<Search
					icon={<SearchIcon size={14} className="opacity-60" />}
					placeholder="Search tokens"
					search={search}
					handleSearch={handleSearch}
					clear
				/>

				{tokens.map((token, tokenIndex) => {
					if (tokenOut.relationships.fungible.data.id === token.id) return null

					return (
						<Accordion
							key={tokenIndex}
							onExpand={() => {
								handleTokenIn(token)
								frame(
									`${tokenOut.attributes.fungible_info.symbol}-${token.attributes.symbol}-swap-amount`
								)
							}}
						>
							<div className="flex flex-row items-center gap-4">
								{token.attributes.implementations && token.attributes.implementations.length > 0 && (
									<TokenImage
										logo={getZerionTokenIconUrl(token.attributes?.icon?.url)}
										symbol={token.attributes.symbol}
									/>
								)}

								<div className="flex w-full flex-col items-center justify-between">
									<div className="flex w-full flex-row font-bold">
										<p className="truncate whitespace-nowrap font-bold">{token.attributes.name}</p>
										<div className="ml-auto flex flex-row items-center">
											$
											<Counter
												count={(token.attributes?.market_data?.price ?? 0).toLocaleString(
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

									<div className="flex w-full flex-row font-bold">
										<div className="flex flex-row items-center gap-2 truncate overflow-ellipsis">
											<p className="flex flex-row items-center gap-1 truncate whitespace-nowrap text-sm opacity-40">
												{token.attributes.symbol?.toUpperCase()}
											</p>
										</div>

										<div
											className={cn(
												"ml-auto flex flex-row items-center text-sm",
												token.attributes?.market_data?.changes?.percent_1d
													? token.attributes?.market_data?.changes?.percent_1d >= 0
														? "text-chart-green"
														: "text-plug-red"
													: "opacity-60"
											)}
										>
											<>
												{token.attributes?.market_data?.changes?.percent_1d ? (
													<>
														<Counter
															count={token.attributes?.market_data?.changes?.percent_1d}
															decimals={2}
														/>
														%
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
		</Frame>
	)
}
