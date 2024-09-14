import Head from "next/head"
import { FC, PropsWithChildren } from "react"

import { LandingFooter, Navbar } from "@/components"

export const StaticLayout: FC<PropsWithChildren & { title: string }> = ({ title, children }) => (
	<>
		<Head>
			<title>{title} | Plug</title>
		</Head>

		<Navbar />
		{children}
		<LandingFooter />
	</>
)
