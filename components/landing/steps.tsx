import { FC } from "react"

import { ArrowDownWideNarrow, FilePen, SlidersHorizontal } from "lucide-react"

import { LandingContainer, StepCard } from "@/components"

export const Steps: FC = () => (
	<LandingContainer className="grid gap-8 xl:grid-cols-3">
		<StepCard
			index={1}
			title="Set Constraints"
			description="Set the constraints that will define the outcomes you expect from your transaction."
			initial={{ opacity: 0, y: 20 }}
			whileInView={{ opacity: 1, y: 0 }}
			transition={{ duration: 0.2 }}
		>
			<SlidersHorizontal size={24} />
		</StepCard>
		<StepCard
			index={2}
			title="Define Actions"
			description="Declare the actions that will be executed once all of your constraints are satisfied."
			initial={{ opacity: 0, y: 20 }}
			whileInView={{ opacity: 1, y: 0 }}
			transition={{ duration: 0.2, delay: 0.2 }}
		>
			<ArrowDownWideNarrow size={24} />
		</StepCard>
		<StepCard
			index={3}
			title="Declare Intent"
			description="Sign a gasless signature to signal your intent of execution then sit back and relax."
			initial={{ opacity: 0, y: 20 }}
			whileInView={{ opacity: 1, y: 0 }}
			transition={{ duration: 0.4, delay: 0.4 }}
		>
			<FilePen size={24} />
		</StepCard>
	</LandingContainer>
)
