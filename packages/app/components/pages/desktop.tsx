import { ConsoleColumnRow } from "@/components/app/columns/column-row"
import { ConsoleSidebar } from "@/components/app/sidebar"
import { ConsoleOnboarding } from "@/components/pages/onboard"
import { useSocket } from "@/state/authentication"

export const DesktopConsole = () => {
	const { socket } = useSocket()

	return (
		<div className="min-w-screen flex h-screen w-full flex-row overflow-y-hidden overflow-x-visible">
			<ConsoleSidebar />

			{socket?.identity?.onboardingAt ? <ConsoleColumnRow /> : <ConsoleOnboarding />}
		</div>
	)
}
