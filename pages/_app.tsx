import WalletProvider from '@/contexts/WalletProvider'

import { Session } from 'next-auth'

import Layout from '@/components/auth/layout'
import { api } from '@/lib/api'
import { type NextPageWithLayout } from '@/lib/types'

import './styles.css'
import { getSession, SessionProvider } from 'next-auth/react'
import type { AppProps, AppType } from 'next/app'

type AppPropsWithLayout = AppProps & {
	Component: NextPageWithLayout
}

// Use of the <SessionProvider> is mandatory to allow components that call
// `useSession()` anywhere in your application to access the `session` object.
const MyApp: AppType<{
	session: Session | null
}> = ({ Component, pageProps }: AppPropsWithLayout) => {
	const getLayout = Component.getLayout ?? (page => page)

	return (
		<SessionProvider session={pageProps.session}>
			<WalletProvider>
				{getLayout(<Component {...pageProps} />)}
			</WalletProvider>
		</SessionProvider>
	)
}

MyApp.getInitialProps = async ({ ctx }) => {
	return {
		session: await getSession(ctx)
	}
}

export default api.withTRPC(MyApp)
