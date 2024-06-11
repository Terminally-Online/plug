import { FC, PropsWithChildren } from "react"

import Head from "next/head"

import { Footer } from "@/components/layouts/static/footer"
import { Navbar } from "@/components/layouts/static/navbar"

export const StaticLayout: FC<PropsWithChildren & { title: string }> = ({
	title,
	children
}) => (
	<>
		<Head>
			<title>{title} | Plug</title>
		</Head>

		<Navbar />
		{children}
		<Footer />
	</>
)
