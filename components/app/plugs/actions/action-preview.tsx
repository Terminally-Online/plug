import Image from "next/image"

import { Sentence } from "@/components"
import { usePlugs } from "@/contexts"
import { categories } from "@/lib/constants"

export const ActionPreview = () => {
	const { actions } = usePlugs()

	return (
		<div className="mb-4 flex flex-col gap-2">
			{actions.map((action, index) => (
				<div key={index} className="relative">
					{index < actions.length - 1 && (
						<div className="absolute bottom-[-12px] top-2 z-[3] ml-[11px] w-[2px] bg-grayscale-100" />
					)}

					<div className="relative z-[4] flex flex-col gap-1">
						<div className="flex flex-row items-center gap-4">
							<Image
								className="h-6 w-6 rounded-md"
								src={categories[action.categoryName].image}
								alt={`Icon for ${action.categoryName}`}
								width={24}
								height={24}
							/>

							<Sentence index={index} preview={true} />
						</div>

						<p className="ml-10 text-sm opacity-60">Ready</p>
					</div>
				</div>
			))}
		</div>
	)
}
