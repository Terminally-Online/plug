import type { AppProps, AppType } from "next/app"

import { Session } from "next-auth"
// import { getSession, SessionProvider } from "next-auth/react"

import { GoogleTagManager } from "@next/third-parties/google"
import { SpeedInsights } from "@vercel/speed-insights/next"

// import { api } from "@/lib/api"
import { type NextPageWithLayout } from "@/lib/types"

import "./styles.css"

type AppPropsWithLayout = AppProps & {
	Component: NextPageWithLayout
}

const PlugApp: AppType<{
	session: Session | null
}> = ({ Component, pageProps }: AppPropsWithLayout) => {
	const getLayout = Component.getLayout ?? (page => page)

	return (
		<>
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
