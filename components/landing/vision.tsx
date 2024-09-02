import { FC } from "react"

import { Blob, BookProfit, Cardiogram, LandingContainer } from "@/components"

import { Underperforming } from "./underperforming"

export const Vision: FC = () => (
	<div className="relative z-[0] mt-[80px]">
		<LandingContainer className="mb-[80px] grid grid-cols-2 gap-8 xl:grid-cols-6 xl:grid-rows-2">
			<BookProfit />
			<Underperforming />
			<Cardiogram />
		</LandingContainer>

		<Blob className="blur-[220px]" left={"60%"} top={"30%"} width={"400"} height={"300"} />
	</div>
)
