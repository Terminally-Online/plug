import { useEffect } from "react"

import { signIn, useSession } from "next-auth/react"

import { AuthFrame, ConsoleColumnRow, ConsoleSidebar, PageContent, PageHeader } from "@/components"
import { FrameProvider, PlugProvider, SocketProvider, useSockets, WalletProvider } from "@/contexts"
import { useMediaQuery } from "@/lib"

const MobilePage = () => {
	const { page } = useSockets()

	if (!page) return null

	return (
		<>
			<PageHeader />
			<PageContent />
			<AuthFrame id={page.id} />
		</>
	)
}

const DesktopPage = () => {
	return (
		<div className="min-w-screen flex h-screen w-full flex-row overflow-y-hidden overflow-x-visible">
			<ConsoleSidebar />
			<ConsoleColumnRow />
		</div>
	)
}

const Page = () => {
	const { data: session } = useSession()
	const { md } = useMediaQuery()

	useEffect(() => {
		if (session !== null) return

		signIn("credentials", {
			message: "0x0",
			signature: "0x0",
			redirect: false
		})
	}, [session])

	return (
		<WalletProvider>
			<SocketProvider>
				<FrameProvider>
					<PlugProvider>{md ? <DesktopPage /> : <MobilePage />}</PlugProvider>
				</FrameProvider>
			</SocketProvider>
		</WalletProvider>
	)
}

export default Page
