import type { FC } from "react"

import { ArrowDownWideNarrow, FilePen, SlidersHorizontal } from "lucide-react"

import { StepCard } from "@/components/cards"
import { Container } from "@/components/landing/container"

export const Steps: FC = () => (
	<Container className="grid gap-8 lg:grid-cols-3">
		<StepCard
			index={1}
			title="Set Rules"
			description="Choose a set of conditions to determine when your transaction can be executed."
			initial={{ opacity: 0, y: 20 }}
			whileInView={{ opacity: 1, y: 0 }}
			transition={{ duration: 0.2 }}
		>
			<SlidersHorizontal size={24} />
		</StepCard>
		<StepCard
			index={2}
			title="Define Actions"
			description="Bundle the actions that will automatically execute once all of your rules are satisfied."
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
	</Container>
)
