import { FC } from "react"

import { Sentence } from "@/components"
import { usePlugs } from "@/contexts"

export const ActionPreview: FC<{ index: number; item: string; review?: boolean }> = ({
	index,
	item,
	review = false
}) => {
	const { actions } = usePlugs(item)

	return (
		<div className="mb-4 flex flex-col">
			{actions.map((_, actionIndex) => (
				<div key={actionIndex} className="relative">
					<div className="relative z-[4] mb-2 flex flex-col">
						<Sentence index={index} item={item} actionIndex={actionIndex} preview={true} />

						{review === true && <p className="text-sm font-bold opacity-40">Ready</p>}
					</div>
				</div>
			))}
		</div>
	)
}
