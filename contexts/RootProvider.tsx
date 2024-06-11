import type { FC, PropsWithChildren } from "react"

import { Session } from "next-auth"
import { SessionProvider } from "next-auth/react"

import { WalletProvider } from "@/components/auth/connector"
import { BalancesProvider, SocketProvider } from "@/contexts"

import { PlugProvider } from "./PlugProvider"

type Props = PropsWithChildren & {
	session: Session | null
}

export const RootProvider: FC<Props> = ({ session, children }) => (
	<SessionProvider session={session}>
		<SocketProvider>
			<PlugProvider>
				<BalancesProvider>
					<WalletProvider>{children}</WalletProvider>
				</BalancesProvider>
			</PlugProvider>
		</SocketProvider>
	</SessionProvider>
)
