import Head from "next/head"
import { FC, PropsWithChildren } from "react"

import { LandingFooter } from "@/components/landing/layout/footer"
import { Navbar } from "@/components/landing/layout/navbar"

type StaticLayoutProps = PropsWithChildren & { 
	title: string,
	ogTitle?: string,
	ogDescription?: string,
	ogImage?: string
}

export const StaticLayout: FC<StaticLayoutProps> = ({ 
	title,
	ogTitle, 
	ogDescription,
	ogImage,
	children 
}) => (
	<>
		<Head>
			<title>{title} | Plug</title>
			{/* Open Graph tags for better link previews */}
			<meta property="og:title" content={ogTitle || `${title} | Plug`} />
			{ogDescription && <meta property="og:description" content={ogDescription} />}
			{ogImage && <meta property="og:image" content={ogImage} />}
			<meta property="og:type" content="website" />
			<meta name="twitter:card" content="summary_large_image" />
		</Head>

		<Navbar />
		{children}
		<LandingFooter />
	</>
)
