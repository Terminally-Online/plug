import { FC } from "react"

import { Image } from "@/components/app/utils/image"
import { Accordion } from "@/components/shared/utils/accordion"
import { Counter } from "@/components/shared/utils/counter"
import { cn } from "@/lib"
import { RouterOutputs } from "@/server/client"
import { useColumnActions } from "@/state/columns"

export const SocketPositionItem: FC<{
	index: number
	protocol?: RouterOutputs["socket"]["balances"]["positions"]["protocols"][number]
}> = ({ index, protocol }) => {
	const { frame } = useColumnActions(index, `${protocol?.name ?? ""}-position`)

	const { positions } = protocol ?? {}

	const change = positions
		? positions.reduce((acc, position) => acc + (position.change ?? 0), 0) /
			positions.filter(position => position.change !== undefined).length
		: 0

	return (
		<Accordion loading={positions === undefined} onExpand={() => frame()}>
			{protocol === undefined ? (
				<div className="invisible">
					<p>.</p>
					<p>.</p>
				</div>
			) : (
				<div className="flex w-full flex-row items-center gap-4">
					<div className="relative h-10 min-w-10">
						<Image
							className="absolute left-1/2 -translate-x-1/2 animate-fade-in rounded-full blur-2xl filter transition-all duration-200 ease-in-out"
							src={protocol.icon}
							alt=""
							width={240}
							height={240}
						/>
						<Image
							className="absolute left-1/2 top-1/2 h-10 w-10 min-w-10 -translate-x-1/2 -translate-y-1/2 animate-fade-in rounded-full bg-plug-green/10"
							src={protocol.icon}
							alt=""
							width={240}
							height={240}
						/>
					</div>

					<div className="relative flex w-full flex-col">
						<div className="flex w-full flex-row justify-between">
							<p className="mr-auto font-bold">{protocol.name}</p>
							<div className="ml-auto flex flex-row font-bold">
								<p>$</p>
								<Counter
									count={protocol.positions.reduce((acc, position) => acc + (position.value ?? 0), 0)}
								/>
							</div>
						</div>
						<div className="flex w-full flex-row items-center gap-2 text-sm font-bold">
							<p className="opacity-40">{protocol.positions.length} Positions</p>
							<div
								className={cn(
									"ml-auto flex flex-row items-center text-sm",
									change === undefined
										? "opacity-60"
										: change > 0
											? "text-plug-green"
											: "text-red-500"
								)}
							>
								<>
									{change !== undefined ? (
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
			)}
		</Accordion>
	)
}
