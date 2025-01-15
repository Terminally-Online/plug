import Head from "next/head"

import { Blocks } from "@/components/landing/blocks"
import { Curve3D } from "@/components/landing/curve-3d"
import { Hero } from "@/components/landing/hero"
import { LandingFooter } from "@/components/landing/layout/footer"
import { Navbar } from "@/components/landing/layout/navbar"
import { Platform } from "@/components/landing/platform"
import { Transactions } from "@/components/landing/transactions"
import { Vision } from "@/components/landing/vision"

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
