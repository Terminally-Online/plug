import Head from "next/head"

import { Blocks, Curve3D, Hero, LandingFooter, Navbar, Platform, Transactions, Vision } from "@/components"

const Page = () => (
	<>
		<Head>
			<title>Plug</title>
		</Head>

		<div className="overflow-x-hidden">
			<Navbar />
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
