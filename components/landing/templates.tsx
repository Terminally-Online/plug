import type { FC } from "react"

import { LandingActionCard, LandingContainer } from "@/components"
import { colors, routes } from "@/lib"

export const Templates: FC = () => {
	const templates = [
		"Bridge to Optimism, Base, and Bera",
		"Top-Up Gearbox Loan Health Factor",
		"Bid on Noun with Pineapple Hat",
		"Buy Beta When Majors Move",
		"Fill Ethena Liquidty Cap to Limit",
		"Renew ENS Annually at Low Gas",
		"Enter Yearn When Above 65% APY",
		"Withdraw ETH:USDC Liquidity Rewards",
		"Rebalance Portfolio Monthly"
	]

	return (
		<LandingContainer className="my-[90px] flex-col items-center gap-4">
			<h2 className="text-center text-[28px] font-bold lg:w-[60%] lg:text-[64px] 2xl:w-[50%]">
				Start today with best-practice templates
			</h2>
			<p className="text-center text-[18px] font-light opacity-40 lg:w-[45%] lg:text-[24px]">
				No need to start from scratch. In just a few minutes, you can
				deploy a strategy that has been battle-tested by the Plug team
				and industry experts.
			</p>

			<div className="grid w-full grid-cols-1 2xl:grid-cols-12">
				<a
					href={routes.earlyAccess}
					target="_blank"
					rel="noreferrer"
					className="mt-[40px] grid gap-4 md:grid-cols-2 2xl:col-start-4 2xl:col-end-10 2xl:grid-cols-3"
				>
					{templates.map((template, index) => (
						<LandingActionCard
							key={index}
							size="lg"
							color={
								Object.keys(colors)[
									index % Object.keys(colors).length
								] as keyof typeof colors
							}
							glow={false}
							title={template}
							initial={{ opacity: 0, y: 20 }}
							whileInView={{ opacity: 1, y: 0 }}
							transition={{
								duration: 0.2,
								delay: 0.1 * index
							}}
							className={index > 7 ? "hidden 2xl:flex" : ""}
						/>
					))}
				</a>
			</div>
		</LandingContainer>
	)
}
