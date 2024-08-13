import { FC, PropsWithChildren } from "react"

import { Session } from "next-auth"
import { SessionProvider } from "next-auth/react"

import {
	BalancesProvider,
	FrameProvider,
	PlugProvider,
	SocketProvider,
	WalletProvider
} from "@/contexts"

export const RootProvider: FC<
	PropsWithChildren & {
		session: Session | null
	}
> = ({ session, children }) => (
	<SessionProvider session={session}>
		<WalletProvider>
			<SocketProvider>
				<FrameProvider>
					<PlugProvider>
						<BalancesProvider>{children}</BalancesProvider>
					</PlugProvider>
				</FrameProvider>
			</SocketProvider>
		</WalletProvider>
	</SessionProvider>
)
