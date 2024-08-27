import { FC, HTMLAttributes, useState } from "react"

import { CircleDollarSign, ImageIcon } from "lucide-react"

import { Header, SocketCollectionList, SocketPositionList, SocketTokenList } from "@/components"
import { useBalances, useSockets } from "@/contexts"
import { cn } from "@/lib"

export const SocketAssets: FC<HTMLAttributes<HTMLDivElement> & { id: string }> = ({ id, className, ...props }) => {
	const { anonymous } = useSockets()
	const { positions } = useBalances()
	const { tokens } = positions

	const [tokensExpanded, setTokensExpanded] = useState(false)
	const [positionsExpanded, setPositionsExpanded] = useState(false)

	return (
		<div className={cn("flex h-full flex-col gap-2", className)} {...props}>
			{anonymous && (
				<div className="flex h-full flex-col items-center justify-center text-center font-bold">
					<p>You are anonymous.</p>
					<p className="max-w-[320px] opacity-40">
						To view the collectibles you are holding you must authenticate a wallet.
					</p>
				</div>
			)}

			{anonymous === false && (
				<>
					<Header
						size="sm"
						icon={<CircleDollarSign size={14} className="opacity-40" />}
						label="Tokens"
						nextLabel={tokens.length < 5 ? undefined : tokensExpanded ? "Collapse" : "See All"}
						nextOnClick={() => setTokensExpanded(!tokensExpanded)}
					/>
					<SocketTokenList id={id} expanded={tokensExpanded} />

					<Header
						size="sm"
						icon={<CircleDollarSign size={14} className="opacity-40" />}
						label="Positions"
						nextLabel={
							positions.protocols.length < 3 ? undefined : positionsExpanded ? "Collapse" : "See All"
						}
						nextOnClick={() => setPositionsExpanded(!positionsExpanded)}
					/>
					<SocketPositionList id={id} expanded={positionsExpanded} />

					<Header size="sm" icon={<ImageIcon size={14} className="opacity-40" />} label="Collectibles" />
					<SocketCollectionList id={id} />
				</>
			)}
		</div>
	)
}
