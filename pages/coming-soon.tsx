import Head from "next/head"
import Link from "next/link"

import { Footer, Navbar } from "@/components/base"
import { Button } from "@/components/buttons"
import { greenGradientStyle } from "@/lib/constants"
import { routes } from "@/lib/constants/routes"

const Page = () => {
	return (
		<>
			<Head>
				<title>Coming Soon | Plug</title>
			</Head>

			<Navbar />

			<div className="flex min-h-[800px] flex-col items-center justify-center gap-2">
				<h1
					className="text-[48px] font-bold lg:text-[72px]"
					style={{ ...greenGradientStyle }}
				>
					Coming Soon
				</h1>

				<p className="max-w-[320px] text-center text-black/65">
					You have found a section that is still under construction.
					Check back soon for updates.
				</p>

				<div className="mt-8 flex flex-row gap-2">
					<Button variant="secondary" href={routes.index}>
						Back Home
					</Button>
					<Button variant="primary" href={routes.earlyAccess}>
						Get Early Access
					</Button>
				</div>
			</div>

			<Footer />
		</>
	)
}

export default Page
