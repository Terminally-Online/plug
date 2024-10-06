import type { AppProps } from "next/app"
import localFont from "next/font/local"
import React, { FC } from "react"
import { useEffect } from "react"

import { api } from "@/server/client"

import { GoogleTagManager } from "@next/third-parties/google"

// import PwaPrompt from "@/components/shared/pwa"
import { GTM_ID, NextPageWithLayout } from "@/lib"
import { registerServiceWorker } from "@/lib/pwa"

import "./styles.css"

const satoshi = localFont({
	src: [
		{ path: "../assets/Satoshi-Light.ttf", weight: "300" },
		{
			path: "../assets/Satoshi-LightItalic.ttf",
			weight: "300",
			style: "italic"
		},
		{ path: "../assets/Satoshi-Regular.ttf", weight: "400" },
		{ path: "../assets/Satoshi-Bold.ttf", weight: "700" },
		{
			path: "../assets/Satoshi-BoldItalic.ttf",
			weight: "700",
			style: "italic"
		},
		{ path: "../assets/Satoshi-Black.ttf", weight: "900" },
		{
			path: "../assets/Satoshi-BlackItalic.ttf",
			weight: "900",
			style: "italic"
		}
	],
	variable: "--font-satoshi"
})

const PlugApp: FC<
	AppProps & {
		Component: NextPageWithLayout
	}
> = ({ Component, pageProps }) => {
	const getLayout = Component.getLayout ?? (page => page)

	useEffect(() => {
		registerServiceWorker()
	}, [])

	return (
		<React.Fragment>
			<style jsx global>
				{`
					* {
						font-family: ${satoshi.style.fontFamily}, sans-serif;
					}
				`}
			</style>

			<GoogleTagManager gtmId={GTM_ID} />

			{getLayout(<Component {...pageProps} />)}
			{/* <PwaPrompt /> */}
		</React.Fragment>
	)
}

export default api.withTRPC(PlugApp)
