import { FC, HTMLAttributes } from "react"

import { X } from "lucide-react"

import { Accordion, Button, Fragments, Image } from "@/components"
import { usePlugs } from "@/contexts"
import { Action, ActionSchema, cn } from "@/lib"
import { useActions } from "@/state"

export const Sentence: FC<
	HTMLAttributes<HTMLButtonElement> & {
		index: number
		item: string
		action: Action
		actionIndex: number
		preview?: boolean
	}
> = ({ index, item, action, actionIndex, preview = false, className, ...props }) => {
	const { own, actions, handle } = usePlugs(item)
	const [solverActions] = useActions()

	const actionSchema = solverActions[action.protocol]

	return (
		<>
			<Accordion
				className={cn("hover:cursor-auto hover:border-grayscale-100 hover:bg-white", className)}
				{...props}
			>
				<div className={cn("flex flex-row items-center font-bold")}>
					<p className="flex w-full flex-wrap items-center gap-[4px]">
						<div className="relative h-6 w-10">
							<Image
								className="absolute mr-2 h-6 w-6 rounded-sm blur-xl filter"
								src={actionSchema.metadata.icon}
								alt={`Icon for ${action.protocol}`}
								width={64}
								height={64}
							/>
							<Image
								className="absolute mr-2 h-6 w-6 rounded-sm"
								src={actionSchema.metadata.icon}
								alt={`Icon for ${action.protocol}`}
								width={64}
								height={64}
							/>
						</div>

						<Fragments
							index={index}
							item={item}
							action={action}
							actionIndex={actionIndex}
							preview={preview}
						/>
					</p>

					{preview === false && own && (
						<Button
							variant="secondary"
							className="mb-auto ml-4 mt-[4px] rounded-sm p-1"
							onClick={() =>
								handle.action.edit({
									id: item,
									actions: JSON.stringify(actions.filter((_, i) => i !== actionIndex))
								})
							}
						>
							<X size={14} className="opacity-60" />
						</Button>
					)}
				</div>
			</Accordion>

			{preview === false && actionIndex < actions.length - 1 && (
				<div className="mx-auto h-2 w-[2px] bg-grayscale-0" />
			)}
		</>
	)
}
