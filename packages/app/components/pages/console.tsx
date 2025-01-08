import { signIn, useSession } from "next-auth/react"
import { useRouter } from "next/router"
import { useEffect, useRef } from "react"
import { LoaderCircle } from "lucide-react"

import { AuthFrame, ConsoleColumnRow, ConsoleSidebar, PageContent, PageNavbar } from "@/components"
import { ColumnAuthenticate } from "@/components/app/columns/utils/column-authenticate"
import { LoginRequired } from "@/components/app/utils/login-required"
import { ReferralRequired } from "@/components/app/utils/referral-required"
import { useConnect, useMediaQuery } from "@/lib"
import { COLUMNS, useColumnStore, usePlugStore, usePlugSubscriptions, useSocket, useSubscriptions } from "@/state"

const MobilePage = () => {
	const { data: session } = useSession()
	const { account } = useConnect()
	const { socket } = useSocket()

	// Show auth column when not authenticated
	if (!account.isAuthenticated) {
		return <ColumnAuthenticate index={COLUMNS.MOBILE_INDEX} />
	}

	// Show referral required when not approved
	const needsReferral = !socket?.identity?.referrerId

	return (
		<div className="min-w-screen flex h-screen w-full flex-row overflow-y-hidden overflow-x-visible">
			<ConsoleSidebar />
			{!account.isAuthenticated ? <LoginRequired /> : needsReferral ? <ReferralRequired /> : <ConsoleColumnRow />}
		</div>
	)
}

export const ConsolePage = () => {
	const hasHandledInitialUrl = useRef(false)

	useSubscriptions()
	useSession({
		required: true,
		onUnauthenticated: () =>
			signIn("credentials", {
				message: "0x0",
				signature: "0x0",
				chainId: 0,
				redirect: true
			})
	})

	const { md } = useMediaQuery()
	const router = useRouter()

	const { socket } = useSocket()
	const { columns, handle } = useColumnStore()
	const { plugs } = usePlugStore()

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

			{md ? <DesktopPage /> : <MobilePage />}
		</>
	)
}
