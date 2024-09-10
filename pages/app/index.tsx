import { useEffect } from "react"

import { NextPageContext } from "next"

import { Session } from "next-auth"
import { getSession, signIn } from "next-auth/react"

import { ConsolePage } from "@/components/pages/console"
import { FrameProvider, PlugProvider, RootProvider, SocketProvider, WalletProvider } from "@/contexts"

Page.getInitialProps = async (ctx: NextPageContext) => {
	const session = await getSession(ctx)
	return { session }
}

function Page({ session }: { session: Session | null }) {
	useEffect(() => {
		if (session !== null) return

		signIn("credentials", {
			message: "0x0",
			signature: "0x0",
			redirect: false
		})
	}, [session])

	return (
		<RootProvider session={session}>
			<WalletProvider>
				<SocketProvider>
					<FrameProvider>
						<PlugProvider>
							<ConsolePage />
						</PlugProvider>
					</FrameProvider>
				</SocketProvider>
			</WalletProvider>
		</RootProvider>
	)
}

export default Page
