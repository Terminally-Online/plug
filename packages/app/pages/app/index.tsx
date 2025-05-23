import { NextPageContext } from "next"
import { Session } from "next-auth"
import { getSession, SessionProvider, signIn } from "next-auth/react"

import { ConsolePage } from "@/components/pages/console"
import { ActivityProvider, BeforeInstallProvider, DataProvider, WalletProvider } from "@/contexts"
import { api } from "@/server/client"
import { ConnectionProvider } from "@/lib/hooks/account/useConnect"

export const getInitialProps = async (context: NextPageContext) => {
	let session = await getSession(context)

	// TODO: This is not functioning correctly because our JWT cookie is reliant on tokens.
	if (!session) {
		try {
			await signIn("credentials", {
				message: "0x0",
				signature: "0x0",
				chainId: 0,
				redirect: true
			})

			session = await getSession(context)
		} catch (error) {
			console.error("Auto-authentication failed:", error)
		}
	}

	return { props: { session } }
}

const Page = ({ session }: { session: Session | null }) => {
	return (
		<SessionProvider session={session}>
			<BeforeInstallProvider>
				<WalletProvider>
					<ConnectionProvider>
						<DataProvider session={session}>
							<ActivityProvider>
								<ConsolePage />
							</ActivityProvider>
						</DataProvider>
					</ConnectionProvider>
				</WalletProvider>
			</BeforeInstallProvider>
		</SessionProvider>
	)
}

export default api.withTRPC(Page)
