import { FC } from "react"

import { Button, Ecosystem, LandingContainer } from "@/components"
import { greenGradientStyle, routes } from "@/lib"

const EARLY_ACCESS =
	process.env.NEXT_PUBLIC_EARLY_ACCESS === "false" ? false : true

export const Hero: FC = () => (
	<div className="relative flex h-[1050px] w-full overflow-hidden lg:h-[900px] lg:items-center">
		<LandingContainer>
			<div className="mt-16 flex flex-col gap-[15px] lg:mt-0 lg:max-w-[70%] lg:gap-[30px]">
				<h1 className="text-[42px] font-bold sm:text-[52px] md:text-[72px] 2xl:text-[96px]">
					Automate your onchain activity with an{" "}
					<span style={{ ...greenGradientStyle }}>
						“if this, then that”
					</span>{" "}
					interface.
				</h1>
				<p className="text-[18px] font-light text-black/40 sm:text-[24px] lg:max-w-[85%]">
					Use Plug to build your own transaction workflows or choose
					from community generated strategies. Let our bots execute
					your transactions and never worry about missing an
					opportunity again. No code needed.
				</p>
				<Button
					href={EARLY_ACCESS ? routes.earlyAccess : routes.app.index}
					className="mt-[30px] w-max"
				>
					{EARLY_ACCESS ? "Get Early Access" : "Enter the App"}
				</Button>
			</div>
		</LandingContainer>

		<Ecosystem />
	</div>
)
