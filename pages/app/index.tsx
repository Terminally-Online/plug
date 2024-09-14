import { NextPageContext } from "next"
import { Session } from "next-auth"
import { getSession } from "next-auth/react"

import { ConsolePage } from "@/components/pages/console"
import { FrameProvider, PlugProvider, RootProvider, SocketProvider, WalletProvider } from "@/contexts"
import { ConnectionProvider } from "@/lib"

Page.getInitialProps = async (ctx: NextPageContext) => {
	const session = await getSession(ctx)
	return { session }
}

function Page({ session }: { session: Session | null }) {
	return (
		<RootProvider session={session}>
			<WalletProvider>
				<ConnectionProvider>
					<SocketProvider>
						<FrameProvider>
							<PlugProvider>
								<ConsolePage />
							</PlugProvider>
						</FrameProvider>
					</SocketProvider>
				</ConnectionProvider>
			</WalletProvider>
		</RootProvider>
	)
}

export default Page
