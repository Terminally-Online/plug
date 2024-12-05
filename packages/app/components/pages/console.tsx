import { signIn, useSession } from "next-auth/react"
import { useRouter } from "next/router"
import { useEffect, useRef } from "react"

import { LoaderCircle } from "lucide-react"

import { AuthFrame, ConsoleColumnRow, ConsoleSidebar, PageContent, PageNavbar } from "@/components"
import { LoginRequired } from "@/components/app/utils/login-required"
import { ReferralRequired } from "@/components/app/utils/referral-required"
import { useConnect, useMediaQuery } from "@/lib"
import { COLUMNS, useColumnStore, usePlugStore, usePlugSubscriptions, useSocket, useSubscriptions } from "@/state"

const MobilePage = () => {
	// Add URL parameter handling for mobile

	return (
		<>
			<PageContent />
			<PageNavbar />
			<AuthFrame />
		</>
	)
}

const DesktopPage = () => {
	const { account } = useConnect()
	const { socket } = useSocket()

	const needsReferral = Boolean(account.isAuthenticated && socket && !socket.identity?.referrerId)

	return (
		<div className="min-w-screen flex h-screen w-full flex-row overflow-y-hidden overflow-x-visible">
			<ConsoleSidebar />
			<ConsoleColumnRow />
			{/* {!account.isAuthenticated ? <LoginRequired /> : needsReferral ? <ReferralRequired /> : <ConsoleColumnRow />} */}
		</div>
	)
}

export const ConsolePage = () => {
	useSubscriptions()

	const { md } = useMediaQuery()
	const router = useRouter()
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

	const { columns, handle } = useColumnStore()
	const { plugs } = usePlugStore()

	const hasHandledInitialUrl = useRef(false)

	useEffect(() => {
		if (!socket || !socket.identity) return

		if (socket.identity.approvedAt && socket.identity.referralCode && !router.query.rfid) {
			router.replace(
				{
					query: { ...router.query, rfid: socket.identity.referralCode }
				},
				undefined,
				{ shallow: true }
			)
		}
	}, [socket, router])

	useEffect(() => {
		const plugId = router.query.plug as string

		if (!plugId || !plugs.length || hasHandledInitialUrl.current) return

		hasHandledInitialUrl.current = true

		// Clear the plug param from URL while preserving other params
		const { plug, ...restQuery } = router.query
		router.replace(
			{
				pathname: router.pathname,
				query: restQuery
			},
			undefined,
			{ shallow: true }
		)

		handle.add({
			index: columns[columns.length - 1]?.index + 1 || 0,
			key: COLUMNS.KEYS.PLUG,
			item: plugId,
			from: COLUMNS.KEYS.MY_PLUGS
		})
	}, [router, router.query, columns, plugs, handle])

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
