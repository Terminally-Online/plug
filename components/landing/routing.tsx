import { Wallet } from "lucide-react"

import { InfoCard } from "@/components"

export const Routing = () => {

	return (
		<InfoCard
			icon={<Wallet size={24} className="opacity-40" />}
			text="Only the best results."
			description="Every transaction is meticulously planned and routed to ensure the best results."
			className="col-span-2 h-[280px] sm:h-[320px] 2xl:h-[300px]"
		>
			<div className="absolute bottom-[40%] left-0 right-0 top-0 bg-gradient-to-b from-plug-white/0 to-plug-white" />
			<div className="absolute bottom-0 left-0 right-0 top-[60%] bg-plug-white" />
		</InfoCard>
	)
}

export default Routing
