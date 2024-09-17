import { FC, HTMLAttributes, useState } from "react"

import { CircleDollarSign, ImageIcon } from "lucide-react"

import { Callout, Header, SocketCollectionList, SocketPositionList, SocketTokenList } from "@/components"
import { cn } from "@/lib"
import { useColumns, useHoldings, useSocket } from "@/state"

export const SocketAssets: FC<HTMLAttributes<HTMLDivElement> & { index?: number }> = ({
	index = -1,
	className,
	...props
}) => {
	const { isAnonymous } = useSocket()
	const { column, isExternal } = useColumns(index)
	const { collectibles, tokens, protocols } = useHoldings(column?.viewAs?.socketAddress)

	const [expanded, setExpanded] = useState<Array<string>>([])

	return (
		<div className={cn("flex h-full flex-col gap-2", className)} {...props}>
			<Callout.Anonymous index={index} viewing="assets" />

			{(isAnonymous === false || isExternal) && (
				<>
					{tokens.length > 0 && (
						<>
							<Header
								size="sm"
								icon={<CircleDollarSign size={14} className="opacity-40" />}
								label="Tokens"
								nextLabel={
									tokens.length < 5 ? undefined : expanded.includes("tokens") ? "Collapse" : "See All"
								}
								nextOnClick={() =>
									setExpanded(prev =>
										prev.includes("tokens")
											? prev.filter(key => key !== "tokens")
											: [...prev, "tokens"]
									)
								}
							/>
							<SocketTokenList index={index} expanded={expanded.includes("tokens")} isColumn={false} />
						</>
					)}

					{protocols.length > 0 && (
						<>
							<Header
								size="sm"
								icon={<CircleDollarSign size={14} className="opacity-40" />}
								label="Positions"
								nextLabel={
									protocols.length < 3
										? undefined
										: expanded.includes("positions")
											? "Collapse"
											: "See All"
								}
								nextOnClick={() =>
									setExpanded(prev =>
										prev.includes("positions")
											? prev.filter(key => key !== "positions")
											: [...prev, "positions"]
									)
								}
							/>
							<SocketPositionList
								index={index}
								expanded={expanded.includes("positions")}
								isColumn={false}
							/>
						</>
					)}

					{collectibles.length > 0 && (
						<>
							<Header
								size="sm"
								icon={<ImageIcon size={14} className="opacity-40" />}
								label="Collectibles"
							/>
							<SocketCollectionList index={index} expanded={true} isColumn={false} />
						</>
					)}
				</>
			)}
		</div>
	)
}
