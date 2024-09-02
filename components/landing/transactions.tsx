import { Blob, Execution, LandingContainer, Recurring, Scheduled } from "@/components"

export const Transactions = () => {
	return (
		<div className="relative z-[0] mb-[80px] h-full">
			<LandingContainer className="grid grid-cols-2 gap-8 xl:grid-cols-6 xl:grid-rows-2">
				<Scheduled />
				<Execution />
				<Recurring />
			</LandingContainer>

			<Blob className="blur-[220px]" left={"60%"} top={"30%"} width={"400"} height={"300"} />
		</div>
	)
}
