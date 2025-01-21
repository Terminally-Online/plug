import { useSession } from "next-auth/react"
import { FC, useCallback, useMemo, useState } from "react"

import { anvil, mainnet } from "viem/chains"

import { motion } from "framer-motion"
import {
	AlertTriangle,
	Calendar,
	CircleDollarSign,
	Eye,
	Globe,
	Hash,
	Library,
	Pause,
	Play,
	Send,
	Waypoints
} from "lucide-react"

import { Frame } from "@/components/app/frames/base"
import { ActionPreview } from "@/components/app/plugs/actions/action-preview"
import { ChainImage } from "@/components/app/sockets/chains/chain.image"
import { Image } from "@/components/app/utils/image"
import { Button } from "@/components/shared/buttons/button"
import { Counter } from "@/components/shared/utils/counter"
import { connectedChains } from "@/contexts"
import { ChainId, cn, formatTitle, getChainName } from "@/lib"
import { useActions } from "@/state/actions"
import { COLUMNS, useColumnStore } from "@/state/columns"
import { Flag, useFlags } from "@/state/flags"
import { usePlugStore } from "@/state/plugs"

export const RunFrame: FC<{
	index: number
	item: string
}> = ({ index, item }) => {
	const { data: session } = useSession()
	const { getFlag } = useFlags()
	const {
		column,
		isFrame,
		handle: { frame, navigate }
	} = useColumnStore(index, "run")
	const {
		actions,
		handle: {
			plug: { queue }
		}
	} = usePlugStore(item)
	const [solverActions] = useActions()

	// TODO: The functionality for this was not finished because right now our in our environment we only have
	//       one chain that is valid at any given time.
	const [currentChainIndex, setCurrentChainIndex] = useState(0)

	const supportedChains = useMemo(() => {
		if (!actions || !solverActions) return []

		return Array.from(
			actions
				.map(action => {
					const protocol = action.protocol
					const protocolSchema = solverActions[protocol]
					const chains = new Set<number>()

					if (protocolSchema?.metadata.chains) {
						protocolSchema.metadata.chains.forEach(chainId => {
							if (!connectedChains.some(chain => chain.id === chainId)) return

							const isDev = getFlag(Flag.SHOW_DEVELOPER)
							if ((isDev && chainId === mainnet.id) || (!isDev && chainId === anvil.id)) return

							chains.add(chainId)
						})
					}

					return chains
				})
				.reduce((acc, chains) => {
					if (acc.size === 0) return chains
					return new Set([...acc].filter(chainId => chains.has(chainId)))
				}, new Set<number>())
		) as ChainId[]
	}, [actions, solverActions, getFlag])

	const chain = useMemo(() => {
		if (!supportedChains || supportedChains.length === 0) return null
		if (supportedChains.length === 1) return supportedChains[0]

		return supportedChains[currentChainIndex]
	}, [supportedChains, currentChainIndex])

	const isActionful = useMemo(() => {
		if (!solverActions) return false

		return actions.some(action => solverActions[action.protocol]?.schema[action.action]?.type === "action")
	}, [actions, solverActions])

	const isReady = useMemo(() => {
		if (!actions || actions.length === 0) return false
		if (!isActionful) return false

		const sentences = document.querySelectorAll(`[data-sentence][data-action-preview="${item}"]`)

		return Array.from(sentences).every(sentence => sentence.getAttribute("data-valid") === "true")
	}, [isActionful, actions, item])

	const handleRun = useCallback(() => {
		if (!column || !column.item || !chain) return

		queue(
			{
				workflowId: column.item,
				chainId: chain,
				startAt: column.schedule?.date?.from ?? new Date(),
				endAt: column.schedule?.date?.to,
				frequency: parseInt(column.schedule?.repeats?.value ?? "0")
			},
			{
				onSuccess: data => {
					navigate({ index, key: COLUMNS.KEYS.ACTIVITY })
					frame(`${data.id}-activity`)
				}
			}
		)
	}, [index, column, chain, queue, navigate, frame])

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
				{actions && actions.length > 0 ? (
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
							<div className="h-[2px] w-full bg-plug-green/10" />
						</div>

						{solverActions && (
							<p className="relative flex flex-row gap-4 font-bold">
								<span className="flex w-max flex-row items-center gap-4">
									<Library size={18} className="opacity-20" />
									<span className="opacity-40">Protocols</span>
								</span>{" "}
								<div className="relative ml-auto flex w-[45%] overflow-hidden">
									{actions.length >= 3 && (
										<>
											<div className="absolute left-0 top-0 z-[1] h-full w-1/4 bg-gradient-to-r from-plug-white to-transparent" />
											<div className="absolute right-0 top-0 z-[1] h-full w-1/4 bg-gradient-to-l from-plug-white to-transparent" />
										</>
									)}

									<motion.div
										className="ml-auto flex flex-row items-center justify-start gap-1 font-bold tabular-nums"
										animate={{
											x: actions.length >= 3 ? ["0%", "-50%"] : 0
										}}
										transition={{
											duration: actions.length >= 3 ? actions.length * 10 : 0,
											ease: "linear",
											repeat: Infinity,
											repeatDelay: 0
										}}
									>
										{[...Array(actions.length >= 3 ? 6 : 1)].map((_, i) => (
											<div key={i} className="flex flex-row items-center gap-4">
												{Array.from(new Set(actions?.map(action => action.protocol))).map(
													protocol => (
														<div
															key={protocol}
															className={cn(
																"flex w-max flex-row items-center gap-2",
																actions.length >= 3 && "ml-4"
															)}
														>
															<Image
																src={solverActions[protocol]?.metadata.icon ?? ""}
																alt={formatTitle(protocol)}
																width={48}
																height={48}
																className="mr-1 h-4 w-4 rounded-[4px]"
															/>
															<span className="whitespace-nowrap">
																{formatTitle(protocol)}
															</span>
														</div>
													)
												)}
											</div>
										))}
									</motion.div>
								</div>
							</p>
						)}

						<p className="flex flex-row justify-between font-bold">
							<span className="flex w-full flex-row items-center gap-4">
								<Hash size={18} className="opacity-20" />
								<span className="opacity-40">Actions</span>
							</span>{" "}
							<span className="flex flex-row items-center gap-1 font-bold tabular-nums">
								<Counter count={actions?.length ?? 0} />
							</span>
						</p>

						{supportedChains.length !== 1 && (
							<p className="flex flex-row justify-between font-bold">
								<span className="flex w-max flex-row items-center gap-4">
									<Globe size={18} className="opacity-20" />
									<span className="opacity-40">Supported Chains</span>
								</span>{" "}
								<span className="group flex flex-row items-center font-bold">
									{supportedChains.map(chain => (
										<div className="-ml-1 transition-all duration-100 group-hover:ml-1" key={chain}>
											<ChainImage chainId={chain} size="xs" />
										</div>
									))}
								</span>
							</p>
						)}

						{chain && (
							<p className="flex flex-row justify-between font-bold">
								<span className="flex w-max flex-row items-center gap-4">
									<Waypoints size={18} className="opacity-20" />
									<span className="opacity-40">Chain</span>
								</span>{" "}
								<span className="flex flex-row items-center gap-2 font-bold">
									<ChainImage chainId={chain} size="xs" />
									{getChainName(chain)}
								</span>
							</p>
						)}

						{!column.schedule && (
							<p className="flex flex-row justify-between font-bold">
								<span className="flex w-full flex-row items-center gap-4">
									<CircleDollarSign size={18} className="opacity-20" />
									<span className="opacity-40">Fee</span>
								</span>{" "}
								<span className="flex flex-row items-center gap-1 font-bold tabular-nums">
									<span className="ml-auto flex flex-row items-center gap-1 pl-2 opacity-40">
										<Counter count={0.00011} /> ETH
									</span>
									<span className="ml-2 flex flex-row items-center">
										$<Counter count={0.049} />
									</span>
								</span>
							</p>
						)}
					</>
				)}

				{column.schedule && (
					<>
						<div className="mb-2 mt-4 flex flex-row items-center gap-4">
							<p className="font-bold opacity-40">Schedule</p>
							<div className="h-[2px] w-full bg-plug-green/10" />
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
					) : !isActionful ? (
						<span className="flex flex-row items-center justify-center gap-2">
							<AlertTriangle size={14} className="opacity-60" />
							Only Constraints Added
						</span>
					) : actions?.length === 0 ? (
						<span className="flex flex-row items-center justify-center gap-2">
							<AlertTriangle size={14} className="opacity-60" />
							No Actions Added
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
