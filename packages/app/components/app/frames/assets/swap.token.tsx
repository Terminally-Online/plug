import { FC } from "react"

import { BadgeCheck, SearchIcon } from "lucide-react"

import { Accordion, Frame, Search, TokenImage } from "@/components"
import { getChainId, useDebounce } from "@/lib"
import { api, RouterOutputs } from "@/server/client"
import { useColumnStore } from "@/state"

type Token =
	| NonNullable<RouterOutputs["socket"]["balances"]["positions"]>["tokens"][number]
	| NonNullable<RouterOutputs["solver"]["tokens"]["get"]>[number]

type SwapTokenFrameProps = {
	index: number
	tokenOut: Token
	handleTokenIn: (token: Token) => void
}

export const SwapTokenFrame: FC<SwapTokenFrameProps> = ({ index, tokenOut, handleTokenIn }) => {
	const {
		isFrame,
		handle: { frame }
	} = useColumnStore(index, `${tokenOut.symbol}-swap-token`)

	const [search, debouncedSearch, handleSearch] = useDebounce("")

	const { data: tokens } = api.solver.tokens.get.useQuery(debouncedSearch, {
		keepPreviousData: true
	})

	return (
		<Frame
			index={index}
			className="min-h-[480px]"
			icon={
				<div className="relative h-8 w-10">
					<TokenImage
						logo={
							tokenOut?.icon ||
							`https://token-icons.llamao.fi/icons/tokens/${getChainId(tokenOut.implementations[0].chain)}/${tokenOut.implementations[0].contract}?h=240&w=240`
						}
						symbol={tokenOut.symbol}
						size="sm"
					/>
				</div>
			}
			label={`Swap ${tokenOut.symbol}`}
			visible={isFrame}
			handleBack={() => frame(`${tokenOut.symbol}-token`)}
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

				{tokens
					?.filter(token => token?.symbol !== tokenOut?.symbol)
					.map((token, tokenIndex) => (
						<Accordion
							key={tokenIndex}
							onExpand={() => {
								handleTokenIn(token)
								frame(`${tokenOut.symbol}-${token.symbol}-swap-amount`)
							}}
						>
							<div className="flex flex-row items-center gap-4">
								{token.implementations && token.implementations.length > 0 && (
									<TokenImage
										logo={
											token?.icon ||
											`https://token-icons.llamao.fi/icons/tokens/${getChainId(token.implementations[0].chain)}/${token.implementations[0].contract}?h=240&w=240`
										}
										symbol={token?.symbol}
									/>
								)}
								<div className="flex flex-col text-left">
									<p className="flex flex-row items-center gap-2 font-bold">
										{token.name}
										{token.flags.verified && (
											<div className="group rounded-full bg-plug-green text-plug-yellow">
												<BadgeCheck size={14} className="opacity-60" />
											</div>
										)}
									</p>
									<p className="text-sm font-bold opacity-60">{token.symbol}</p>
								</div>
							</div>
						</Accordion>
					))}
			</div>
		</Frame>
	)
}