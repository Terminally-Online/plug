import { NextPageContext } from "next"
import { Session } from "next-auth"
import { getSession, SessionProvider } from "next-auth/react"

import { ConsolePage } from "@/components/pages/console"
import { ActivityProvider, BeforeInstallProvider, DataProvider, WalletProvider } from "@/contexts"
import { ConnectionProvider } from "@/lib"

export const getServerSideProps = async (context: NextPageContext) => {
	const session = await getSession(context)
	return { props: { session } }
}

export default function Page({ session }: { session: Session | null }) {
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
