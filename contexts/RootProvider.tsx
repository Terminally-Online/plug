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

import { PageProvider } from "./PagesProvider"

export const RootProvider: FC<
	PropsWithChildren & {
		session: Session | null
	}
> = ({ session, children }) => (
	<SessionProvider session={session}>
		<PageProvider>
			<FrameProvider>
				<WalletProvider>
					<SocketProvider>
						<PlugProvider>
							<BalancesProvider>{children}</BalancesProvider>
						</PlugProvider>
					</SocketProvider>
				</WalletProvider>
			</FrameProvider>
		</PageProvider>
	</SessionProvider>
)
