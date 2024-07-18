import { useMemo } from "react"

import Image from "next/image"

import { useSession } from "next-auth/react"

import { Plus } from "lucide-react"

import { Button, Sentence } from "@/components"
import { useFrame, usePlugs } from "@/contexts"
import { actions, categories, formatTitle, getValues } from "@/lib"

const baseSuggestions = Object.entries(actions).flatMap(
	([categoryName, actions]) =>
		Object.keys(actions).map(actionName => ({
			categoryName,
			actionName
		}))
)

const getProtocolFrequency = (
	actions: Array<{
		categoryName: string
		actionName: string
	}>
): Record<string, number> => {
	const protocolFrequency: Record<string, number> = {}

	actions.forEach(action => {
		protocolFrequency[action.categoryName] =
			(protocolFrequency[action.categoryName] || 0) + 1
	})

	return protocolFrequency
}

export const ActionView = () => {
	const { data: session } = useSession()
	const { handleFrameVisible } = useFrame()
	const { id, plug, actions, handle } = usePlugs()

	const own = plug && session && session.address === plug.userAddress

	const suggestions = useMemo(() => {
		const protocolFrequency = getProtocolFrequency(actions)
		const selectedActions = new Set(
			actions.map(action => `${action.categoryName}-${action.actionName}`)
		)

		// Handle case where there are no actions
		if (actions.length === 0) {
			return baseSuggestions.slice(0, 5)
		}

		// Get the most recent action
		const mostRecentAction = actions[actions.length - 1]

		return baseSuggestions
			.filter(
				suggestion =>
					!selectedActions.has(
						`${suggestion.categoryName}-${suggestion.actionName}`
					)
			)
			.sort((a, b) => {
				// Prioritize suggestions related to the most recent action
				if (a.categoryName === mostRecentAction.categoryName) return -1
				if (b.categoryName === mostRecentAction.categoryName) return 1
				return (
					(protocolFrequency[b.categoryName] || 0) -
					(protocolFrequency[a.categoryName] || 0)
				)
			})
			.slice(0, 3)
	}, [actions])

	return (
		<>
			{actions && actions.length > 0 ? (
				<div className="mb-72 flex flex-col">
					{actions.map((_, index) => (
						<Sentence key={index} index={index} />
					))}

					{own && (
						<div className="mt-12">
							<h4 className="mb-2 font-bold opacity-40">
								Next Action Suggestions
							</h4>
							<div className="flex flex-col gap-2">
								{suggestions.map((suggestion, idx) => (
									<button
										key={idx}
										className="flex items-center gap-4 rounded-lg bg-grayscale-0 p-4 font-bold"
										onClick={() =>
											handle.action.edit({
												id,
												actions: JSON.stringify([
													...actions,
													{
														...suggestion,
														values: getValues(
															suggestion.categoryName,
															suggestion.actionName
														)
													}
												])
											})
										}
									>
										<Image
											src={
												categories[
													suggestion.categoryName
												].image
											}
											alt={suggestion.categoryName}
											width={24}
											height={24}
											className="rounded-md"
										/>
										<p className="flex w-full flex-wrap items-center gap-[8px] truncate">
											{formatTitle(suggestion.actionName)}
										</p>
										<Button
											variant="secondary"
											className="group p-1"
											onClick={() =>
												handle.action.edit({
													id,
													actions: JSON.stringify([
														...actions,
														{
															...suggestion,
															values: getValues(
																suggestion.categoryName,
																suggestion.actionName
															)
														}
													])
												})
											}
										>
											<Plus size={14} />
										</Button>
									</button>
								))}
							</div>
						</div>
					)}
				</div>
			) : (
				<div className="mx-auto my-auto flex h-full max-w-[80%] flex-col gap-2 text-center">
					<p className="text-lg font-bold">
						No actions have been added yet.
					</p>
					<p className="opacity-60">
						Create a Plug to actions that you want to do on a
						regular basis and when all the conditions have been met.
					</p>
					<Button
						className="mx-auto mt-4 w-max"
						onClick={() => handleFrameVisible("actions")}
					>
						Add Action
					</Button>
				</div>
			)}
		</>
	)
}
