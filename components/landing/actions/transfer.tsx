import { CalendarClock } from "lucide-react"

import { InfoCard } from "../cards"

export const ActionTransfer = () => {
	return (
		<InfoCard
			icon={<CalendarClock size={24} className="opacity-40" />}
			text="Transfer."
			description="Send your assets anywhere in the world."
			className="relative z-[99999] h-[320px] sm:h-[320px] 2xl:h-[300px]"
		>
			<div className="absolute bottom-[50%] left-0 right-0 top-0 bg-gradient-to-b from-plug-white/0 to-plug-white" />
			<div className="absolute bottom-0 left-0 right-0 top-[50%] bg-plug-white" />
		</InfoCard>
	)
}
