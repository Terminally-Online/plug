import { Sentence } from "@/components/app/sentences"
import { Button } from "@/components/buttons"
import { useFrame } from "@/contexts"
import { useActions } from "@/contexts/ActionProvider"

export const ActionView = () => {
	const { handleFrameVisible } = useFrame()
	const { actions } = useActions()

	return (
		<>
			{actions && actions.length > 0 ? (
				<div className="mb-72">
					<div className="flex flex-col">
						{actions.map((action, index) => (
							<>
								<Sentence action={action} />

								{index < actions.length - 1 && (
									<div className="mx-auto h-4 w-[2px] bg-grayscale-100" />
								)}
							</>
						))}
					</div>

					{/*<div className="mt-12">
						<h4 className="font-bold opacity-40">
							Next Action Suggestions
						</h4>
					</div> */}
				</div>
			) : (
				<div className="mx-auto my-20 flex max-w-[80%] flex-col gap-2 text-center">
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
