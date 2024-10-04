import { signIn } from "next-auth/react"
import { useEffect } from "react"

import { LoaderCircle } from "lucide-react"

import { AuthFrame, ConsoleColumnRow, ConsoleSidebar, PageContent, PageNavbar } from "@/components"
import { useMediaQuery } from "@/lib"
import { useSocket } from "@/state"

const MobilePage = () => {
	return (
		<>
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

	// useEffect(() => {
	// 	if (socket) return
	//
	// 	signIn("credentials", {
	// 		message: "0x0",
	// 		signature: "0x0",
	// 		chainId: 0,
	// 		redirect: false
	// 	})
	// }, [socket])

	if (!socket)
		return (
			<div className="absolute bottom-0 left-0 right-0 top-0 flex h-screen w-screen items-center justify-center">
				<LoaderCircle size={24} className="animate-spin opacity-60" />
			</div>
		)

	return (
		<>
			{/*
	            <FeatureRequestFrame />
	            <DeletedFrame />
	        */}

			{/*
			<p>{session?.address}</p>
			<p>{socket.id}</p>
			<p>{socket.socketAddress}</p>
			<p>{JSON.stringify(columns, null, 2)}</p>
			*/}

			{md ? <DesktopPage /> : <MobilePage />}
		</>
	)
}
