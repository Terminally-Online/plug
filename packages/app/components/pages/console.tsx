import { signIn, useSession } from "next-auth/react"
import { useRouter } from "next/router"
import { useEffect, useRef } from "react"

import { useAtomValue } from "jotai"

import { useConnect, useMediaQuery } from "@/lib"
import { useDisconnect } from "@/lib/hooks/wallet/useDisconnect"
import { useSocket } from "@/state/authentication"
import { COLUMNS, primaryColumnsAtom, useColumnActions } from "@/state/columns"
import { plugsAtom } from "@/state/plugs"

import { DesktopConsole } from "./desktop"
import { MobileConsole } from "./mobile"

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

	const {
		account: { address }
	} = useConnect()
	const { disconnect } = useDisconnect(true)
	const { socket } = useSocket()

	const columns = useAtomValue(primaryColumnsAtom)
	const { add } = useColumnActions()
	const plugs = useAtomValue(plugsAtom)

	useEffect(() => {
		const isAnonymous = !socket.id.startsWith("0x")
		const isPotentiallyExpired = !address || socket.id === address
		if (isAnonymous || isPotentiallyExpired) return

		disconnect()
	}, [socket, address, disconnect])

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

	return md ? <DesktopConsole /> : <MobileConsole />
}
