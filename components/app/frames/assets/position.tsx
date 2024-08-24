import { FC, useMemo } from "react"

import Image from "next/image"

import { ExternalLink } from "lucide-react"

import { Counter } from "@/components/shared"
import { useFrame } from "@/contexts"
import { cn, formatTitle } from "@/lib"
import { RouterOutputs } from "@/server/client"

import { Frame } from "../base"

export const PositionFrame: FC<{
	id: string
	protocol: RouterOutputs["socket"]["balances"]["positions"]["protocols"][number]
}> = ({ id, protocol }) => {
	const { isFrame } = useFrame({
		id,
		key: `position-${protocol.name}`
	})

	// Group positions by type.
	const groupedPositions = useMemo(() => {
		const grouped: Record<
			string,
			Array<
				RouterOutputs["socket"]["balances"]["positions"]["protocols"][number]["positions"][number]
			>
		> = {}

		protocol.positions.forEach(position => {
			if (grouped[position.type] === undefined) {
				grouped[position.type] = []
			}

			grouped[position.type].push(position)
		})

		return grouped
	}, [protocol])

	return (
		<Frame
			id={id}
			icon={
				<div className="relative h-10 w-10">
					<Image
						className="absolute left-1/2 top-1/2 -translate-x-1/2 -translate-y-1/2 animate-fade-in rounded-full blur-2xl filter transition-all duration-200 ease-in-out"
						src={protocol.icon}
						alt={protocol.name}
						style={{
							width: `2rem`,
							minWidth: `2rem`,
							height: `2rem`
						}}
						width={240}
						height={240}
					/>
					<div
						className="absolute left-1/2 top-1/2 -translate-x-1/2 -translate-y-1/2 animate-fade-in rounded-full bg-grayscale-100"
						style={{
							backgroundImage: `url(${protocol.icon})`,
							backgroundSize: "cover",
							backgroundPosition: "center",
							backgroundRepeat: "no-repeat",
							width: "2rem",
							minWidth: "2rem",
							height: "2rem"
						}}
					/>
				</div>
			}
			label={protocol.name}
			visible={isFrame}
			hasChildrenPadding={false}
			hasOverlay
		>
			<div className="relative mb-4 flex flex-col gap-4 px-6 font-bold">
				{Object.keys(groupedPositions).map(type => (
					<div key={type}>
						<div className="mb-2 flex flex-row items-center gap-2">
							<p className="font-bold opacity-40">
								{formatTitle(type)}
							</p>

							<div className="h-[2px] w-full bg-grayscale-100" />
						</div>

						<div className="flex flex-col gap-2">
							{groupedPositions[type].map((position, index) => (
								<div key={`${type}-${index}`}>
									<div className="flex flex-row items-center gap-4">
										<Image
											className="h-8 w-8 rounded-full"
											src={position.fungible.icon ?? ""}
											alt=""
											width={48}
											height={48}
										/>
										<div className="flex w-full flex-col gap-0">
											<div className="flex flex-row items-center justify-between gap-2">
												<p>{position.fungible.name}</p>
												<p className="flex flex-row">
													$
													<Counter
														count={
															position.value ?? 0
														}
													/>
												</p>
											</div>

											<div className="flex flex-row items-center justify-between gap-2 text-sm text-black/40">
												<p>
													<Counter
														count={
															position.balance ??
															0
														}
													/>
												</p>
												<p
													className={cn(
														"ml-auto text-sm",
														position.change ===
															undefined
															? "opacity-60"
															: position.change >
																  0
																? "text-plug-green"
																: "text-red-500"
													)}
												>
													<span className="ml-auto flex flex-row items-center">
														{position.change !==
														undefined ? (
															<>
																<Counter
																	count={
																		position.change
																	}
																	decimals={2}
																/>
																%
															</>
														) : (
															"-"
														)}
													</span>
												</p>
											</div>
										</div>
									</div>
								</div>
							))}
						</div>
					</div>
				))}

				<a
					className="flex w-full items-center justify-center gap-2 rounded-lg bg-grayscale-100 py-4 font-bold transition-all duration-200 ease-in-out hover:opacity-90"
					// style={{
					// 	backgroundColor: metadata?.color ?? "",
					// 	color: textColor
					// }}
					href={protocol.url}
					target="_blank"
					rel="noreferrer"
				>
					<ExternalLink size={14} className="opacity-60" />
					Manage
				</a>
			</div>
		</Frame>
	)
}
