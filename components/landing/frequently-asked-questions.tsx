import { FC, useState } from "react"

import { motion } from "framer-motion"
import { ChevronDown } from "lucide-react"

import { LandingContainer } from "@/components"

const FrequentlyAskedQuestion: FC<{ text: string; description: string }> = ({
	text,
	description
}) => {
	const [collapsed, setCollapsed] = useState(true)

	return (
		<div className="flex flex-col gap-2 border-b-[1px] border-[#D9D9D9]/40 pb-4">
			<button
				onClick={() => setCollapsed(!collapsed)}
				className="z-[30] flex w-full items-center text-[24px] font-bold"
			>
				{text}
				<motion.span
					className="ml-auto transform rounded-full bg-[#FBFBFB] p-1 transition-transform"
					initial={{ rotate: 0 }}
					animate={{ rotate: collapsed ? 0 : 180 }}
					transition={{ duration: 0.2 }}
				>
					<ChevronDown size={24} className="opacity-40" />
				</motion.span>
			</button>
			<motion.p
				className="text-black/65 opacity-40 lg:mr-16"
				initial={{ height: 0, opacity: 0 }}
				animate={{
					height: collapsed ? 0 : "auto",
					opacity: collapsed ? 0 : 1
				}}
				transition={{ duration: 0.2 }}
			>
				{description}
			</motion.p>
		</div>
	)
}

export const FrequentlyAskedQuestions: FC = () => (
	<LandingContainer>
		<div className="my-[90px] grid lg:grid-cols-12">
			<h3 className="mb-8 text-[28px] font-bold lg:col-span-4 lg:col-start-2 lg:mb-0 lg:w-[60%] lg:text-[64px] 2xl:w-[50%]">
				Frequently Asked Questions
			</h3>

			<div className="flex flex-col gap-8 lg:col-span-5 lg:col-start-7">
				<FrequentlyAskedQuestion
					text="How does Plug work?"
					description="With a simple plug-and-play interface, you can establish conditions and outcomes across top Ethereum based protocols. Execute immediately, schedule for the future, or set up a recurring transaction. Plug’s bots take care of the hard work."
				/>
				<FrequentlyAskedQuestion
					text="Who is Plug for?"
					description="Plug was made with you in mind. Whether you’re managing on behalf of yourself, a venture fund, a market maker, or a DAO, Plug enables you to operate securely and in ways never before possible without having to write a single line of code. Even if you’re not an expert you can utilize the Plugs curated by the team to get started in seconds."
				/>
				<FrequentlyAskedQuestion
					text="What should I automate?"
					description="All across there are actions that you find boring or difficult and struggle to scale. Automation serves as a way to minimize the impact of each of those downsides and replace them with significant upside and potential. If there’s something you don’t like doing or can’t do, you should automate it."
				/>
				<FrequentlyAskedQuestion
					text="Is Plug really trustless?"
					description="Yes! Plug utilizes intents with embedded transaction conditions that are checked during the time of simulation and onchain execution. You don’t have to trust us or a centralized oracle of any kind. All data is sourced and verified onchain during simulation and execution."
				/>
			</div>
		</div>
	</LandingContainer>
)
