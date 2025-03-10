import { memo } from "react"

import { ConsoleColumnRow } from "@/components/app/columns/column-row"
import { ConsoleSidebar } from "@/components/app/sidebar"
import { ConsoleOnboarding } from "@/components/pages/onboard"
import { useSocket } from "@/state/authentication"
import { useConnect } from "@/lib"

export const DesktopConsole = memo(() => {
	const { account: { isAuthenticated } } = useConnect()
	const { socket } = useSocket()

	return (
		<div className="min-w-screen flex h-screen w-full flex-row overflow-y-hidden overflow-x-visible">
			<ConsoleSidebar />

			{isAuthenticated && socket?.identity?.onboardingAt ? (
				<ConsoleColumnRow />
			) : (
				<ConsoleOnboarding />
			)}
		</div>
	)
})

DesktopConsole.displayName = "DesktopConsole"
