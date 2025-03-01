import { FC, useCallback } from "react"

import { Bell, CircleDollarSign, Clock10, MessageCircleQuestionIcon, Share } from "lucide-react"

import { Frame } from "@/components/app/frames/base"
import { ActionPreview } from "@/components/app/plugs/actions/action-preview"
import { ActivityIcon } from "@/components/app/sockets/activity/activity-item"
import { Button } from "@/components/shared/buttons/button"
import { Counter } from "@/components/shared/utils/counter"
import { DateSince } from "@/components/shared/utils/date-since"
import { formatTitle } from "@/lib"
import { RouterOutputs } from "@/server/client"
import { useActions } from "@/state/actions"
import { useColumnActions } from "@/state/columns"

export const SimulationFrame: FC<{
	index: number
	activity: RouterOutputs["plugs"]["activity"]["get"][number] | undefined
	simulationId: string | undefined
}> = ({ index, activity, simulationId }) => {
	const [solverActions] = useActions()

	const { frame } = useColumnActions(index, `${activity?.id}-activity`)

	const simulation = activity?.runs.find(sim => sim.id === simulationId)
	const actions = activity?.inputs

	const handleShare = useCallback(() => {
		if (!activity) return

		try {
			const cleanedActions = activity.inputs.map(action => ({
				protocol: action.protocol?.toLowerCase?.(),
				sentence: solverActions[action.protocol]["schema"][action.action].sentence
			}))
			console.log("cleanedActions", cleanedActions)
			const params = new URLSearchParams({
				name: activity.plug.name.slice(0, 100), // Reasonable name length
				protocols: cleanedActions.map(a => a.protocol).join(","),
				sentences: cleanedActions.map(a => a.sentence).join(",")
			})
			const url = `/api/canvas/opengraph?${params.toString()}`

			console.log("Debug - OpenGraph URL:", url)
			window.open(url, "_blank")
		} catch (e) {
			console.error("Share generation failed:", e, {
				plugId: activity.plug.id,
				actions: activity.inputs
			})
		}
	}, [activity, solverActions])

	if (!activity || !simulation) return null

	return (
		<Frame
			index={index}
			icon={<ActivityIcon status={simulation.status} />}
			label={`Simulation ${formatTitle(simulation.status)}`}
			visible={simulation !== undefined}
			handleBack={() => frame()}
			hasOverlay={true}
		>
			<ActionPreview index={index} item={activity.plug.id} actions={actions} errors={simulation.errors ?? []} />

			{simulation.error && (
				<p className="mx-auto mt-4 px-8 text-center text-sm font-bold text-plug-red">
					Warning: {simulation.error}
				</p>
			)}

			{simulation.status === "success" ? (
				<Button
					className="mt-4 flex w-full flex-row items-center justify-center gap-2 py-4"
					onClick={handleShare}
				>
					<Share size={18} className="opacity-60" />
					Share
				</Button>
			) : (
				<Button
					className="mt-4 flex w-full flex-row items-center justify-center gap-2 py-4"
					onClick={() => {
						const message = [`Socket: ${activity.plug.socketId}`, `Simulation ${simulation.id}`].join(" - ")

						const encodedMessage = encodeURIComponent(message)
						window.open(`https://t.me/evmchance?text=${encodedMessage}`, "_blank")
					}}
				>
					<MessageCircleQuestionIcon size={18} className="opacity-60" />
					Get Help
				</Button>
			)}

			<div className="mb-2 mt-4 flex flex-row items-center gap-4">
				<p className="font-bold opacity-40">Details</p>
				<div className="h-[2px] w-full bg-plug-green/10" />
			</div>

			<p className="flex flex-row items-center justify-between gap-4 font-bold">
				<Bell size={18} className="opacity-20" />
				<span className="opacity-40">Status</span>{" "}
				<span className="ml-auto">{formatTitle(simulation.status)}</span>
			</p>

			<p className="flex flex-row items-center justify-between gap-4 font-bold">
				<Clock10 size={18} className="opacity-20" />
				<span className="opacity-40">Simulated</span>{" "}
				<span className="ml-auto">
					<DateSince date={new Date(simulation.createdAt)} />
				</span>
			</p>

			{simulation.gasEstimate && (
				<p className="flex flex-row items-center justify-between gap-4 font-bold">
					<CircleDollarSign size={18} className="opacity-20" />
					<span className="opacity-40">Fee</span>{" "}
					<span className="group ml-auto flex cursor-pointer flex-row items-center gap-4">
						<Counter count={simulation.gasEstimate} />
					</span>
				</p>
			)}
		</Frame>
	)
}
