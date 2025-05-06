import { FC } from "react"

import { Image } from "@/components/app/utils/image"
import { Accordion } from "@/components/shared/utils/accordion"
import { Counter } from "@/components/shared/utils/counter"
import { cn } from "@/lib"
import { RouterOutputs } from "@/server/client"
import { useColumnActions } from "@/state/columns"

import { PositionFrame } from "../../frames/assets/position/frame"

export const SocketPositionItem: FC<{
	index: number
	protocols?: NonNullable<RouterOutputs["service"]["zerion"]["wallet"]["positions"]>["data"]
}> = ({ index, protocols }) => {
	const protocol = protocols?.[0]
	const { frame } = useColumnActions(index, `${protocol?.relationships.dapp?.data.id ?? ""}-position`)

	const count = protocols?.length ?? 0
	const totalValue = protocols?.reduce((sum, p) => sum + (p.attributes?.value ?? 0), 0) ?? 0
	const validChanges = protocols
		?.map(p => p.attributes?.changes?.percent_1d)
		.filter(c => typeof c === "number") as number[]
	const change =
		validChanges && validChanges.length > 0 ? validChanges.reduce((sum, c) => sum + c, 0) / validChanges.length : 0

	return (
		<>
			<Accordion loading={protocol === undefined} onExpand={() => frame()}>
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
								src={protocol.attributes?.application_metadata?.icon?.url ?? ""}
								alt=""
								width={240}
								height={240}
							/>
							<Image
								className="absolute left-1/2 top-1/2 h-10 w-10 min-w-10 -translate-x-1/2 -translate-y-1/2 animate-fade-in rounded-full bg-plug-green/10"
								src={protocol.attributes?.application_metadata?.icon?.url ?? ""}
								alt=""
								width={240}
								height={240}
							/>
						</div>

						<div className="relative flex w-full flex-col">
							<div className="flex w-full flex-row justify-between">
								<p className="mr-auto font-bold">{protocol.attributes?.application_metadata?.name}</p>
								<div className="ml-auto flex flex-row font-bold">
									<p>$</p>
									<Counter
										count={totalValue.toLocaleString("en-US", {
											minimumFractionDigits: 2,
											maximumFractionDigits: 2
										})}
									/>
								</div>
							</div>
							<div className="flex w-full flex-row items-center gap-2 text-sm font-bold">
								<p className="opacity-40">
									{count} Position{count === 1 ? "" : "s"}
								</p>
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

			<PositionFrame index={index} protocols={protocols} />
		</>
	)
}
