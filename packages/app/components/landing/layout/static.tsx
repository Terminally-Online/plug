import Head from "next/head"
import { FC, PropsWithChildren } from "react"

import { LandingFooter } from "@/components/landing/layout/footer"
import { Navbar } from "@/components/landing/layout/navbar"

type StaticLayoutProps = PropsWithChildren & {
	title: string
	description?: string
	img?: string
}

export const StaticLayout: FC<StaticLayoutProps> = ({ title, description, img, children }) => (
	<>
		<Head>
			<title>{title} | Plug</title>
			<meta property="og:title" content={`${title} | Plug`} />

			{description && <meta property="og:description" content={description} />}
			{img && <meta property="og:image" content={img} />}
		</Head>

		<Navbar />

		{children}

		<LandingFooter />
	</>
)
