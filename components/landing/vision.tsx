import { FC } from "react"

import { Blob, BookProfit, Cardiogram, LandingContainer } from "@/components"

import { Underperforming } from "./underperforming"

export const Vision: FC = () => (
	<div className="relative z-[0]">
		<LandingContainer className="mb-[80px] grid grid-cols-2 gap-8 xl:grid-cols-6 xl:grid-rows-2">
			<BookProfit />
			<Underperforming />
			<Cardiogram />
		</LandingContainer>

		<Blob left={"60%"} top={"100%"} width={"1000"} height={"500"} />
	</div>
)
