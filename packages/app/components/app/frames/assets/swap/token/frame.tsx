import { FC } from "react"

import { BadgeCheck, SearchIcon } from "lucide-react"

import { useAtom, useAtomValue } from "jotai"

import { Frame } from "@/components/app/frames/base"
import { Search } from "@/components/app/inputs/search"
import { TokenImage } from "@/components/app/sockets/tokens/token-image"
import { Accordion } from "@/components/shared/utils/accordion"
import { getChainId, greenGradientStyle, useDebounce } from "@/lib"
import { api, RouterOutputs } from "@/server/client"
import { columnByIndexAtom, isFrameAtom, useColumnActions } from "@/state/columns"

type Token =
	| NonNullable<RouterOutputs["socket"]["balances"]["positions"]>["tokens"][number]
	| NonNullable<RouterOutputs["solver"]["tokens"]["get"]>[number]

type SwapTokenFrameProps = {
	index: number
	tokenOut: Token
	handleTokenIn: (token: Token) => void
}

export const SwapTokenFrame: FC<SwapTokenFrameProps> = ({ index, tokenOut, handleTokenIn }) => {
	const [column] = useAtom(columnByIndexAtom(index))
	const frameKey = `${tokenOut.symbol}-swap-token`
	const isFrame = useAtomValue(isFrameAtom)(column, frameKey)
	const { frame } = useColumnActions(index, frameKey)

	const [search, debouncedSearch, handleSearch] = useDebounce("")

	const { data: tokens } = api.solver.tokens.get.useQuery(debouncedSearch, {
		placeholderData: prev => prev
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

								<div className="w-full flex flex-row justify-between gap-2 items-center">
									<div className="flex flex-col text-left w-full">
										<p className="font-bold">{token.name}</p>
										<p className="text-sm font-bold opacity-60">{token.symbol}</p>
									</div>

									{token.flags.verified && (
										<p className="font-bold text-sm" style={{ ...greenGradientStyle }}>
											Verified
										</p>
									)}
								</div>
							</div>
						</Accordion>
					))}
			</div>
		</Frame>
	)
}
