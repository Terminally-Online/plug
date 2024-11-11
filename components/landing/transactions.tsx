import { Execution, LandingContainer, Recurring, Scheduled } from "@/components"

export const Transactions = () => {
	return (
		<div className="relative z-[0] mb-[80px] h-full">
			<LandingContainer className="grid grid-cols-2 gap-8 xl:grid-cols-6 xl:grid-rows-2">
				<Scheduled />
				<Execution />
				<Recurring />
			</LandingContainer>
		</div>
	)
}
