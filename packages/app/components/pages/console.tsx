import { signIn, useSession } from "next-auth/react"
import { useRouter } from "next/router"
import { useEffect, useRef } from "react"

import { LoaderCircle } from "lucide-react"

import { useMediaQuery } from "@/lib"
import { useSocket } from "@/state/authentication"
import { COLUMNS, useColumnStore } from "@/state/columns"
import { usePlugStore } from "@/state/plugs"
import { useSubscriptions } from "@/state/subscriptions"

import { DesktopConsole } from "./desktop"
import { MobileConsole } from "./mobile"

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

	const router = useRouter()
	const { md } = useMediaQuery()
	const { socket } = useSocket()
	const { columns, handle } = useColumnStore()
	const { plugs } = usePlugStore()

	console.log("[ConsolePage] Render", {
		socket: !!socket,
		md: md,
		columnsLength: columns.length,
		hasPlugs: plugs.length > 0
	})

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

			{md ? <DesktopConsole /> : <MobileConsole />}
		</>
	)
}
