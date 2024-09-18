import { signIn } from "next-auth/react"
import { useEffect } from "react"

import { LoaderCircle } from "lucide-react"

import { AuthFrame, ConsoleColumnRow, ConsoleSidebar, PageContent, PageHeader, PageNavbar } from "@/components"
import { useMediaQuery } from "@/lib"
import { useColumns, useSocket } from "@/state"

const MobilePage = () => {
	return (
		<>
			<PageHeader />
			<PageContent />
			<PageNavbar />

			<AuthFrame />
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

export const ConsolePage = () => {
	const { md } = useMediaQuery()
	const { socket } = useSocket()
	const { columns } = useColumns()

	useEffect(() => {
		if (socket) return

		signIn("credentials", {
			message: "0x0",
			signature: "0x0",
			chainId: 0,
			redirect: false
		})
	}, [socket])

	if (!socket)
		return (
			<div className="absolute bottom-0 left-0 right-0 top-0 flex h-screen w-screen items-center justify-center">
				<LoaderCircle size={24} className="animate-spin opacity-60" />
				Socket has not been loaded.
			</div>
		)

	return (
		<>
			{/*
	            <FeatureRequestFrame />
	            <DeletedFrame />
	        */}

			{md ? <DesktopPage /> : <MobilePage />}
		</>
	)
}
