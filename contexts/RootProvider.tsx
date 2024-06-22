import type { FC, PropsWithChildren } from "react"

import { Session } from "next-auth"
import { SessionProvider } from "next-auth/react"

import { WalletProvider } from "@/components/auth/connector"
import {
	BalancesProvider,
	FrameProvider,
	PlugProvider,
	SocketProvider
} from "@/contexts"

import { ActionProvider } from "./ActionProvider"

type Props = PropsWithChildren & {
	session: Session | null
}

export const RootProvider: FC<Props> = ({ session, children }) => (
	<SessionProvider session={session}>
		<FrameProvider>
			<SocketProvider>
				<PlugProvider>
					<ActionProvider>
						<BalancesProvider>
							<WalletProvider>{children}</WalletProvider>
						</BalancesProvider>
					</ActionProvider>
				</PlugProvider>
			</SocketProvider>
		</FrameProvider>
	</SessionProvider>
)
