import { FC, HTMLAttributes, useState } from "react"

import { CircleDollarSign, ImageIcon } from "lucide-react"

import { Callout, Header, SocketCollectionList, SocketPositionList, SocketTokenList } from "@/components"
import { useSockets } from "@/contexts"
import { cn } from "@/lib"

export const SocketAssets: FC<HTMLAttributes<HTMLDivElement> & { id: string }> = ({ id, className, ...props }) => {
	const { isAnonymous, isExternal, collectibles, positions } = useSockets(id)
	const { tokens, protocols } = positions

	const [expanded, setExpanded] = useState<Array<string>>([])

	return (
		<div className={cn("flex h-full flex-col gap-2", className)} {...props}>
			<Callout.Anonymous id={id} viewing="assets" />

			{(isAnonymous === false || isExternal) && (
				<>
					{tokens.length > 0 && (
						<>
							<Header
								size="sm"
								icon={<CircleDollarSign size={14} className="opacity-40" />}
								label="Tokens"
								nextLabel={
									positions.tokens.length < 5
										? undefined
										: expanded.includes("tokens")
											? "Collapse"
											: "See All"
								}
								nextOnClick={() =>
									setExpanded(prev =>
										prev.includes("tokens")
											? prev.filter(key => key !== "tokens")
											: [...prev, "tokens"]
									)
								}
							/>
							<SocketTokenList id={id} expanded={expanded.includes("tokens")} isColumn={false} />
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
							<SocketPositionList id={id} expanded={expanded.includes("positions")} isColumn={false} />
						</>
					)}

					{collectibles.length > 0 && (
						<>
							<Header
								size="sm"
								icon={<ImageIcon size={14} className="opacity-40" />}
								label="Collectibles"
							/>
							<SocketCollectionList id={id} expanded={true} isColumn={false} />
						</>
					)}
				</>
			)}
		</div>
	)
}
