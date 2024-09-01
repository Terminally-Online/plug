import { FC } from "react"

import { BookProfit, Cardiogram, LandingContainer } from "@/components"

import { Underperforming } from "./underperforming"

export const Vision: FC = () => (
	<LandingContainer className="mb-[80px] grid grid-cols-2 gap-8 xl:grid-cols-6 xl:grid-rows-2">
		<BookProfit />
		<Underperforming />
		<Cardiogram />
	</LandingContainer>
)
