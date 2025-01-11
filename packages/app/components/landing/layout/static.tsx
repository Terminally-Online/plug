import Head from "next/head"
import { FC, PropsWithChildren } from "react"

import { LandingFooter } from "@/components/landing/layout/footer"
import { Navbar } from "@/components/landing/layout/navbar"

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
