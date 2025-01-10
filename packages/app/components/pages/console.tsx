import { signIn, useSession } from "next-auth/react"
import { useRouter } from "next/router"
import { memo, useEffect, useRef } from "react"

import { useAccount } from "wagmi"

import { LoaderCircle } from "lucide-react"

import { AuthFrame, ConsoleColumnRow, ConsoleSidebar, PageContent, PageNavbar } from "@/components"
import { LoginRequired } from "@/components/app/utils/login-required"
import { ReferralRequired } from "@/components/app/utils/referral-required"
import { useConnect, useMediaQuery } from "@/lib"
import { useRenderTracking } from "@/lib/hooks/useRenderTracking"
import { COLUMNS, useColumnStore, usePlugStore, useSocket, useSubscriptions } from "@/state"

const MobilePage = () => {
	return (
		<>
			<PageContent />
			<PageNavbar />
			<AuthFrame />
		</>
	)
}

const DesktopPage = memo(() => {
	const { data: session } = useSession()
	const { socket } = useSocket()

	const needsReferral = Boolean(session?.user.id?.startsWith("0x") && socket && !socket.identity?.referrerId)
	// const renderCount = useRenderTracking("DesktopPage")

	return (
		<div className="min-w-screen flex h-screen w-full flex-row overflow-y-hidden overflow-x-visible">
			{/* <div className="border-plug-blue group pointer-events-none absolute left-0 top-0 z-[9999] h-full w-full rounded-lg border-[8px]">
				<p className="bg-plug-blue absolute bottom-0 w-max rounded-tr-lg px-4 py-2 font-bold text-plug-white">
					{renderCount.current}
				</p>
			</div> */}

			<ConsoleSidebar />

			{!session?.user.id?.startsWith("0x") ? (
				<LoginRequired />
			) : needsReferral ? (
				<ReferralRequired />
			) : (
				<ConsoleColumnRow />
			)}
		</div>
	)
})

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
