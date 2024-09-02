import { CalendarClock } from "lucide-react"

import { InfoCard } from "@/components"

export const Execution = () => {
	return (
		<InfoCard
			icon={<CalendarClock size={24} className="opacity-40" />}
			text='"If this, then that” driven execution.'
			description="Stop missing opportunities and have your transactions execute no matter where you are or what you're doing. If you want it done, it will be delivered on a silver platter."
			className="col-span-2 h-[540px] xl:col-span-4 xl:row-span-2 xl:h-full"
		/>
	)
}
