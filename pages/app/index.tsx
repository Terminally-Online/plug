import { NextPageContext } from "next"
import { Session } from "next-auth"
import { getSession, SessionProvider } from "next-auth/react"

import { ConsolePage } from "@/components/pages/console"
import { DataProvider, PlugProvider, WalletProvider } from "@/contexts"
import { ConnectionProvider } from "@/lib"

Page.getInitialProps = async (ctx: NextPageContext) => {
	const session = await getSession(ctx)
	return { session }
}

function Page({ session }: { session: Session | null }) {
	return (
		<SessionProvider session={session}>
			<WalletProvider>
				<ConnectionProvider>
					<DataProvider>
						<PlugProvider>
							<ConsolePage />
						</PlugProvider>
					</DataProvider>
				</ConnectionProvider>
			</WalletProvider>
		</SessionProvider>
	)
}

export default Page
