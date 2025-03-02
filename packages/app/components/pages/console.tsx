import { signIn, useSession } from "next-auth/react"
import { useRouter } from "next/router"
import { useEffect, useRef } from "react"

import { LoaderCircle } from "lucide-react"

import { useMediaQuery } from "@/lib"
import { useSocket } from "@/state/authentication"
import { COLUMNS, primaryColumnsAtom, useColumnActions } from "@/state/columns"
import { plugsAtom } from "@/state/plugs"

import { DesktopConsole } from "./desktop"
import { MobileConsole } from "./mobile"
import { useAtomValue } from "jotai"

export const ConsolePage = () => {
	const hasHandledInitialUrl = useRef(false)

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

	const columns = useAtomValue(primaryColumnsAtom)
	const { add } = useColumnActions()
	const plugs = useAtomValue(plugsAtom)

	// useEffect(() => {
	// 	if (!socket || socket.id === address) return
	//
	// 	disconnect()
	// }, [socket, address, disconnect])

	useEffect(() => {
		if (!socket || !socket.identity) return

		if (socket.identity.referrerId && socket.identity.referralCode && !router.query.rfid) {
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

		const { plug, ...restQuery } = router.query
		router.replace(
			{
				pathname: router.pathname,
				query: restQuery
			},
			undefined,
			{ shallow: true }
		)

		add({
			index: columns[columns.length - 1]?.index + 1 || 0,
			key: COLUMNS.KEYS.PLUG,
			item: plugId,
			from: COLUMNS.KEYS.MY_PLUGS
		})
	}, [router, router.query, columns, plugs, add])

	if (!socket)
		return (
			<div className="absolute bottom-0 left-0 right-0 top-0 flex h-screen w-screen items-center justify-center">
				<LoaderCircle size={24} className="animate-spin opacity-40" />
			</div>
		)

	return md ? <DesktopConsole /> : <MobileConsole />
}
