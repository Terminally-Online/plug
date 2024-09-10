import { FC, PropsWithChildren } from "react"

import { Session } from "next-auth"
import { SessionProvider } from "next-auth/react"

export const RootProvider: FC<
	PropsWithChildren & {
		session: Session | null
	}
> = ({ session, children }) => {
	return <SessionProvider session={session}>{children}</SessionProvider>
}
