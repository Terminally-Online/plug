import Head from "next/head"

import { Blocks, CallToAction, Hero, LandingFooter, Light, Transactions, Vision } from "@/components"

const Page = () => (
	<>
		<Head>
			<title>Plug</title>
		</Head>

		<div className="overflow-x-hidden">
			<Hero />
			<Light />
			<Transactions />
			<CallToAction
				text="Operate at a level your friends only dream of."
				description="If you can't beat them, join them. Embrace the future of onchain activity and supercharge your potential with only a couple of clicks and a few minutes."
				button="Get Started"
			/>
			<Vision />
			<Blocks />
			<LandingFooter />
		</div>
	</>
)

export default Page
