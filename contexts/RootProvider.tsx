import { Session } from "next-auth"
import { SessionProvider } from "next-auth/react"
import { FC, PropsWithChildren } from "react"

export const RootProvider: FC<
	PropsWithChildren & {
		session: Session | null
	}
> = ({ session, children }) => {
	return <SessionProvider session={session}>{children}</SessionProvider>
}
