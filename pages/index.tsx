import Head from "next/head"

import { Blocks, Curve3D, Hero, LandingFooter, Platform, Transactions, Vision } from "@/components"

const Page = () => (
	<>
		<Head>
			<title>
				A single interface to manage, compose, schedule, and execute all your transactions and holdings in one
				place.
			</title>
		</Head>

		<div className="overflow-x-hidden">
			<Hero />
			<Platform />
			<Curve3D />
			<Transactions />
			<Blocks />
			<Vision />
			<LandingFooter />
		</div>
	</>
)

export default Page
