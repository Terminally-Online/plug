import { useSession } from "next-auth/react"
import { FC, useCallback, useMemo, useState } from "react"

import { motion } from "framer-motion"
import { AlertTriangle, Calendar, CircleDollarSign, Eye, Library, Pause, Play, Send, Waypoints } from "lucide-react"

import { useAtom, useAtomValue } from "jotai"

import { Frame } from "@/components/app/frames/base"
import { ActionPreview } from "@/components/app/plugs/actions/action-preview"
import { ChainImage } from "@/components/app/sockets/chains/chain.image"
import { Image } from "@/components/app/utils/image"
import { Button } from "@/components/shared/buttons/button"
import { Counter } from "@/components/shared/utils/counter"
import { connectedChains } from "@/contexts"
import { ChainId, cn, formatTitle, getChainName } from "@/lib"
import { useActions } from "@/state/actions"
import { useSocket } from "@/state/authentication"
import { columnByIndexAtom, COLUMNS, isFrameAtom, useColumnActions } from "@/state/columns"
import { plugByIdAtom, usePlugActions } from "@/state/plugs"
import { areAllSentencesValidAtom } from "@/state/sentences"

export const RunFrame: FC<{
	index: number
	item: string
}> = ({ index, item }) => {
	const { data: session } = useSession()

	const [column] = useAtom(columnByIndexAtom(index))
	const frameKey = "run"
	const isFrame = useAtomValue(isFrameAtom)(column, frameKey)
	const { frame, navigate } = useColumnActions(index, frameKey)

	const [plug] = useAtom(plugByIdAtom(item))

	const { socket } = useSocket()
	const [solverActions] = useActions()
	const { queue } = usePlugActions()

	// TODO: The functionality for this was not finished because right now our in our environment we
	//       only have one chain that is valid at any given time.
	const [currentChainIndex] = useState(0)

	const supportedChains = useMemo(() => {
		if (!plug || !plug.actions || !solverActions) return []

		return Array.from(
			plug.actions
				.map(action => {
					const protocol = action.protocol
					const protocolSchema = solverActions[protocol]
					const chains = new Set<number>()

					if (protocolSchema?.metadata.chains) {
						const protocolChainIds = protocolSchema.metadata.chains.flatMap(chain => chain.chainIds)
						for (const chain of connectedChains) {
							if (protocolChainIds.includes(chain.id)) chains.add(chain.id)
						}
					}

					return chains
				})
				.reduce((acc, chains) => {
					if (acc.size === 0) return chains
					return new Set([...acc].filter(chainId => chains.has(chainId)))
				}, new Set<number>())
		) as ChainId[]
	}, [plug, solverActions])

	const chain = useMemo(() => {
		if (!supportedChains || supportedChains.length === 0) return null
		if (supportedChains.length === 1) return supportedChains[0]

		return supportedChains[currentChainIndex]
	}, [supportedChains, currentChainIndex])

	const isActionful = useMemo(() => {
		if (!plug || !solverActions) return false

		return plug.actions.some(action => solverActions[action.protocol]?.schema[action.action]?.type === "action")
	}, [plug, solverActions])

	// Using our atom to check if all sentences are valid
	const checkAllSentencesValid = useAtomValue(areAllSentencesValidAtom)

	const isReady = useMemo(() => {
		if (!plug || plug.actions.length === 0) return false
		if (!isActionful) return false

		// Use our atom to check validation state
		// This now properly accounts for coil references
		return checkAllSentencesValid(item)
	}, [isActionful, plug, item, checkAllSentencesValid])

	const handleRun = useCallback(() => {
		if (!column || !column.item || !chain) return

		const intent = {
			plugId: column.item,
			chainId: chain,
			from: socket.socketAddress,
			startAt: column.schedule?.date?.from ?? new Date(),
			endAt: column.schedule?.date?.to,
			frequency: parseInt(column.schedule?.repeats?.value ?? "0"),
			socket: column.index !== COLUMNS.SIDEBAR_INDEX
		}

		queue(intent, {
			onError: data => console.error(data),
			onSuccess: data => {
				navigate({ index, key: COLUMNS.KEYS.ACTIVITY })
				frame(`${data.id}-activity`)
			}
		})
	}, [index, socket.socketAddress, column, chain, queue, navigate, frame])

	if (!column) return null

	return (
		<Frame
			index={index}
			className="z-[2]"
			icon={<Eye size={18} className="opacity-40" />}
			label={
				<>
					<span className="text-lg font-bold">
						<span className="opacity-40">Run:</span> Preview
					</span>
				</>
			}
			visible={(isFrame && session && session.user.anonymous === false) || false}
			hasOverlay={true}
			handleBack={column.schedule ? () => frame("schedule") : undefined}
		>
			<div className="flex flex-col">
				{plug && plug.actions && plug.actions.length > 0 ? (
					<ActionPreview index={index} item={item} />
				) : (
					<div className="flex rounded-lg border-[1px] border-plug-green/10 p-4 py-4 text-center font-bold text-black/40">
						<p className="mx-auto max-w-[380px]">
							No actions added and configured on this Plug yet. Add some actions to run and schedule it.
						</p>
					</div>
				)}

				{isReady && (
					<>
						<div className="mb-2 mt-4 flex flex-row items-center gap-4">
							<p className="font-bold opacity-40">Details</p>
							<div className="h-[1px] w-full bg-plug-green/10" />
						</div>

						{solverActions &&
							plug?.actions &&
							(() => {
								const uniqueProtocols = Array.from(
									new Set(plug.actions?.map(action => action.protocol))
								)

								const shouldScroll = uniqueProtocols.length >= 3

								return (
									<>
										<div className="relative flex flex-row gap-4 font-bold">
											<span className="flex w-max flex-row items-center gap-4">
												<Library size={18} className="opacity-20" />
												<span className="opacity-40">
													Protocol{plug.actions.length > 1 && "s"}
												</span>
											</span>{" "}
											<div className="relative ml-auto flex w-[45%] overflow-hidden">
												{shouldScroll && (
													<>
														<div className="absolute left-0 top-0 z-[1] h-full w-1/4 bg-gradient-to-r from-plug-white to-transparent" />
														<div className="absolute right-0 top-0 z-[1] h-full w-1/4 bg-gradient-to-l from-plug-white to-transparent" />
													</>
												)}

												<motion.div
													className="ml-auto flex flex-row items-center justify-start gap-1 font-bold tabular-nums"
													animate={{
														x: shouldScroll ? ["0%", "-50%"] : 0
													}}
													transition={{
														duration: shouldScroll ? uniqueProtocols.length * 10 : 0,
														ease: "linear",
														repeat: Infinity,
														repeatDelay: 0
													}}
												>
													{[...Array(shouldScroll ? 2 : 1)].map((_, i) => (
														<div key={i} className="flex flex-row items-center gap-4">
															{uniqueProtocols.map(protocol => (
																<div
																	key={protocol}
																	className={cn(
																		"flex w-max flex-row items-center gap-2",
																		shouldScroll && "ml-4"
																	)}
																>
																	<Image
																		src={
																			solverActions[protocol]?.metadata.icon ?? ""
																		}
																		alt={formatTitle(protocol)}
																		width={48}
																		height={48}
																		className="mr-1 h-4 w-4 rounded-[4px]"
																	/>
																	<span className="whitespace-nowrap">
																		{formatTitle(protocol)}
																	</span>
																</div>
															))}
														</div>
													))}
												</motion.div>
											</div>
										</div>
									</>
								)
							})()}

						{chain && (
							<p className="flex flex-row justify-between font-bold">
								<span className="flex w-max flex-row items-center gap-4">
									<Waypoints size={18} className="opacity-20" />
									<span className="opacity-40">Blockchain</span>
								</span>{" "}
								<span className="flex flex-row items-center gap-2 font-bold">
									<ChainImage chainId={chain} size="xs" />
									{getChainName(chain)}
								</span>
							</p>
						)}

						<p className="flex flex-row justify-between font-bold">
							<span className="flex w-full flex-row items-center gap-4">
								<CircleDollarSign size={18} className="opacity-20" />
								<span className="opacity-40">Gas Fee</span>
							</span>{" "}
							<span className="flex flex-row items-center gap-1 font-bold tabular-nums">
								<span className="ml-auto flex flex-row items-center gap-1 pl-2 opacity-40">
									<Counter count={0.0} /> ETH
								</span>
								<span className="ml-2 flex flex-row items-center">Free</span>
							</span>
						</p>
					</>
				)}

				{column.schedule && (
					<>
						<div className="mb-2 mt-4 flex flex-row items-center gap-4">
							<p className="font-bold opacity-40">Schedule</p>
							<div className="h-[1px] w-full bg-plug-green/10" />
						</div>

						<p className="flex flex-row justify-between font-bold">
							<span className="flex w-full flex-row items-center gap-4">
								<Calendar size={18} className="opacity-20" />
								<span className="opacity-40">Frequency</span>
							</span>{" "}
							{column.schedule.repeats.label}
						</p>

						{column.schedule.date && column.schedule.date.from instanceof Date && (
							<p className="flex flex-row justify-between font-bold">
								<span className="flex w-full flex-row items-center gap-4">
									<Play size={18} className="opacity-20" />
									<span className="opacity-40">Start At</span>
								</span>{" "}
								<Counter count={column.schedule.date.from.toLocaleDateString()} />
							</p>
						)}

						{column.schedule.date && column.schedule.date.to instanceof Date && (
							<p className="flex flex-row justify-between font-bold">
								<span className="flex w-full flex-row items-center gap-4">
									<Pause size={18} className="opacity-20" />
									<span className="opacity-40">Stop At</span>
								</span>{" "}
								<Counter count={column.schedule.date.to.toLocaleDateString()} />
							</p>
						)}
					</>
				)}

				<Button
					variant={isReady ? "primary" : "primaryDisabled"}
					className="mt-4 w-full py-4"
					onClick={handleRun}
					disabled={!isReady}
				>
					{isReady ? (
						<span className="flex flex-row items-center justify-center gap-2">
							<Send size={14} className="opacity-60" />
							Submit
						</span>
					) : !plug?.actions || plug?.actions?.length === 0 ? (
						<span className="flex flex-row items-center justify-center gap-2">
							<AlertTriangle size={14} className="opacity-60" />
							No Actions Added
						</span>
					) : !isActionful ? (
						<span className="flex flex-row items-center justify-center gap-2">
							<AlertTriangle size={14} className="opacity-60" />
							Only Constraints Added
						</span>
					) : (
						<span className="flex flex-row items-center justify-center gap-2">
							<AlertTriangle size={14} className="opacity-60" />
							Required Inputs Incomplete
						</span>
					)}
				</Button>
			</div>
		</Frame>
	)
}
