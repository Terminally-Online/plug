import type { FC } from "react"

import { Button } from "@/components/buttons"
import { Container } from "@/components/landing/container"
import { Ecosystem } from "@/components/landing/ecosystem"
import { greenGradientStyle, routes } from "@/lib/constants"

const EARLY_ACCESS =
	process.env.NEXT_PUBLIC_EARLY_ACCESS === "false" ? false : true

export const Hero: FC = () => (
	<div className="relative flex h-[1050px] w-full overflow-hidden lg:h-[900px] lg:items-center">
		<Container>
			<div className="mt-4 flex flex-col gap-[15px] lg:mt-0 lg:max-w-[70%] lg:gap-[30px]">
				<h1 className="text-[42px] font-bold lg:text-[72px] 2xl:text-[96px]">
					Automate your onchain activity with an{" "}
					<span style={{ ...greenGradientStyle }}>
						“if this, then that”
					</span>{" "}
					interface.
				</h1>
				<p className="text-[18px] font-light text-black/40 lg:max-w-[85%] lg:text-[24px]">
					Use Plug to build your own transaction workflows or choose
					from community generated strategies. Let our bots execute
					your transactions and never worry about missing an
					opportunity again. No code needed.
				</p>
				<Button
					href={EARLY_ACCESS ? routes.earlyAccess : routes.app}
					className="mt-[30px] w-max"
				>
					{EARLY_ACCESS ? "Get Early Access" : "Enter the App"}
				</Button>
			</div>
		</Container>

		<Ecosystem />
	</div>
)
