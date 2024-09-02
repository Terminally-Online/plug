import { CalendarClock } from "lucide-react"

import { InfoCard } from "@/components"

export const Scheduled = () => {
	return (
		<InfoCard
			icon={<CalendarClock size={24} className="opacity-40" />}
			text="Transactions on a schedule."
			description="Not automating your transactions means you're leaving money on the table. The market is evolving - time to upgrade."
			className="col-span-2 h-[280px] sm:h-[320px] 2xl:h-[300px]"
		/>
	)
}
