import type { AppProps, AppType } from "next/app"

import { Session } from "next-auth"

import localFont from "next/font/local"
// import { getSession, SessionProvider } from "next-auth/react"
import { GoogleTagManager } from "@next/third-parties/google"
import { SpeedInsights } from "@vercel/speed-insights/next"

// import { api } from "@/lib/api"
import { type NextPageWithLayout } from "@/lib/types"

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

type AppPropsWithLayout = AppProps & {
	Component: NextPageWithLayout
}

const PlugApp: AppType<{
	session: Session | null
}> = ({ Component, pageProps }: AppPropsWithLayout) => {
	const getLayout = Component.getLayout ?? (page => page)

	return (
		<>
			<style jsx global>{`
				* {
					font-family: ${satoshi.style.fontFamily}, sans-serif;
				}
			`}</style>

			<GoogleTagManager gtmId="GTM-PT3JT2P9" />
			{getLayout(<Component {...pageProps} />)}
			<SpeedInsights />
		</>
	)
}

export default PlugApp

// PlugApp.getInitialProps = async ({ ctx }) => {
// 	return {
// 		session: await getSession(ctx)
// 	}
// }
//
// export default api.withTRPC(PlugApp)
