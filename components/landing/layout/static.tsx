import { FC, PropsWithChildren } from "react"

import Head from "next/head"

import { LandingFooter, Navbar } from "@/components"

export const StaticLayout: FC<PropsWithChildren & { title: string }> = ({
	title,
	children
}) => (
	<>
		<Head>
			<title>{title} | PLUG</title>
		</Head>

		<Navbar />
		{children}
		<LandingFooter />
	</>
)
