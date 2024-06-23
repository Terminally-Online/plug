import { Sentence } from "@/components/app/sentences/sentence"
import { Button } from "@/components/buttons"
import { useFrame, usePlugs } from "@/contexts"

export const ActionView = () => {
	const { handleFrameVisible } = useFrame()
	const { actions } = usePlugs()

	return (
		<>
			{actions && actions.length > 0 ? (
				<div className="mb-72 flex flex-col">
					{Array.from({ length: actions.length }).map((_, index) => (
						<Sentence key={index} index={index} />
					))}

					<div className="mt-12">
						<h4 className="font-bold opacity-40">
							Next Action Suggestions
						</h4>
					</div>
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
