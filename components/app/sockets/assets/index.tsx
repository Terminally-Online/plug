import { FC, HTMLAttributes, useState } from "react"

import exp from "constants"
import { CircleDollarSign, ImageIcon } from "lucide-react"

import { Callout, Header, SocketCollectionList, SocketPositionList, SocketTokenList } from "@/components"
import { useSockets } from "@/contexts"
import { cn } from "@/lib"

export const SocketAssets: FC<HTMLAttributes<HTMLDivElement> & { id: string }> = ({ id, className, ...props }) => {
	const { isAnonymous: anonymous, positions } = useSockets()

	const [expanded, setExpanded] = useState<Array<string>>([])

	return (
		<div className={cn("flex h-full flex-col gap-2", className)} {...props}>
			<Callout.Anonymous viewing="assets" />

			{anonymous === false && (
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
								prev.includes("tokens") ? prev.filter(key => key !== "tokens") : [...prev, "tokens"]
							)
						}
					/>
					<SocketTokenList id={id} expanded={expanded.includes("tokens")} />

					<Header
						size="sm"
						icon={<CircleDollarSign size={14} className="opacity-40" />}
						label="Positions"
						nextLabel={
							positions.protocols.length < 3
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
					<SocketPositionList id={id} expanded={expanded.includes("positions")} />

					<Header size="sm" icon={<ImageIcon size={14} className="opacity-40" />} label="Collectibles" />
					<SocketCollectionList id={id} />
				</>
			)}
		</div>
	)
}
