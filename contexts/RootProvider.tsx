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

import { ConsoleProvider } from "./ConsoleProvider"

type Props = PropsWithChildren & {
	session: Session | null
}

export const RootProvider: FC<Props> = ({ session, children }) => (
	<SessionProvider session={session}>
		<FrameProvider>
			<WalletProvider>
				<SocketProvider>
					<PlugProvider>
						<BalancesProvider>
							<ConsoleProvider>{children}</ConsoleProvider>
						</BalancesProvider>
					</PlugProvider>
				</SocketProvider>
			</WalletProvider>
		</FrameProvider>
	</SessionProvider>
)
