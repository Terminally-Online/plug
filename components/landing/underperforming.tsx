import { Code } from "lucide-react"

import { InfoCard } from "@/components"

export const Underperforming = () => {
	return (
		<InfoCard
			icon={<Code size={24} className="opacity-40" />}
			text="You're underperforming."
			description="You don't need to be a rocket scientist to outcompete. Top performers no longer execute transactions manually."
			className="col-span-2 h-[280px] sm:h-[320px] 2xl:h-[300px]"
		/>
	)
}
