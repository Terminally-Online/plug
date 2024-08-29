import { FC, PropsWithChildren } from "react"

import { Session } from "next-auth"
import { SessionProvider } from "next-auth/react"

import { FrameProvider, PlugProvider, SocketProvider, WalletProvider } from "@/contexts"

export const RootProvider: FC<
	PropsWithChildren & {
		session: Session | null
	}
> = ({ session, children }) => {
	return (
		<SessionProvider session={session}>
			<WalletProvider>
				<SocketProvider>
					<FrameProvider>
						<PlugProvider>{children}</PlugProvider>
					</FrameProvider>
				</SocketProvider>
			</WalletProvider>
		</SessionProvider>
	)
}
