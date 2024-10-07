import { Session } from "next-auth"
import { signIn, useSession } from "next-auth/react"
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

export const ConsolePage = ({ session }: { session: Session | null }) => {
	const { md } = useMediaQuery()

	useEffect(() => {
		if (session?.user.id) return

		signIn("credentials", {
			message: "0x0",
			signature: "0x0",
			chainId: 0,
			redirect: true
		})
	}, [session])

	if (!session?.user.id)
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

			{md ? <DesktopPage /> : <MobilePage />}
		</>
	)
}
