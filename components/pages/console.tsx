import { signIn, useSession } from "next-auth/react"

import { LoaderCircle } from "lucide-react"

import { AuthFrame, ConsoleColumnRow, ConsoleSidebar, PageContent, PageNavbar } from "@/components"
import { LoginRequired } from "@/components/app/utils/login-required"
import { ReferralRequired } from "@/components/app/utils/referral-required"
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
	const { data: session } = useSession()
	const { socket } = useSocket()

	const isAuthenticated = session?.user.id?.startsWith("0x")
	const isApproved = socket?.identity?.approvedAt !== null

	return (
		<div className="min-w-screen flex h-screen w-full flex-row overflow-y-hidden overflow-x-visible">
			<ConsoleSidebar />
			{!isAuthenticated ? (
				<LoginRequired />
			) : !isApproved ? (
				<ReferralRequired />
			) : (
				<ConsoleColumnRow />
			)}
		</div>
	)
}

export const ConsolePage = () => {
	const { md } = useMediaQuery()
	// NOTE: This makes the session required for the console page. When the user does
	// not have a session, they will be automatically logged into an anonymous account.
	// This enables users to maintain their session through reloads, and on log out,
	// automatically roll over back to an anonymous account maintaining the local
	// console state they already have in their `localStorage`.
	// ...
	// New user → Anonymous Session → Phantom Socket
	// Wallet user → Connect Wallet → Sign Message → Authenticated Session → Socket
	// Existing user → Existing Session → Socket
	// ...
	// We have to include the socket in the loading state because the socket is a dependency
	// to render the console page. Therefore, until both the session and a socket, whether
	// anonymous or authenticated, are available, we will show a loading state.
	const { data: session } = useSession({
		required: true,
		onUnauthenticated: () =>
			signIn("credentials", {
				message: "0x0",
				signature: "0x0",
				chainId: 0,
				redirect: true
			})
	})
	const { socket } = useSocket()

	if (!session?.user.id || !socket)
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
