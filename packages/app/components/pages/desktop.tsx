import { useSession } from "next-auth/react"
import { memo } from "react"

import { ConsoleColumnRow } from "@/components/app/columns/column-row"
import { ConsoleSidebar } from "@/components/app/sidebar"
import { LoginRequired } from "@/components/app/utils/login-required"
import { ReferralRequired } from "@/components/app/utils/referral-required"
import { useSocket } from "@/state/authentication"

export const DesktopConsole = memo(() => {
	const { data: session } = useSession()
	const { socket } = useSocket()

	const isAuthenticated = session?.user.id?.startsWith("0x")
	const isReferred = Boolean(socket && socket.identity?.referrerId)

	return (
		<div className="min-w-screen flex h-screen w-full flex-row overflow-y-hidden overflow-x-visible">
			<ConsoleSidebar />

			{!isAuthenticated ? <LoginRequired /> : !isReferred ? <ReferralRequired /> : <ConsoleColumnRow />}
		</div>
	)
})

DesktopConsole.displayName = "DesktopConsole"
