import { signOut } from "next-auth/react"
import { useCallback, useMemo } from "react"

import { UseDisconnectReturnType, useDisconnect as useDisconnectWagmi } from "wagmi"

import { useSetAtom } from "jotai"

import { INITIAL_SOCKET, socketModelAtom } from "@/state/authentication"

export function useDisconnect(out: boolean = false): UseDisconnectReturnType {
	const { connectors, disconnect, ...rest } = useDisconnectWagmi()

	const setSocket = useSetAtom(socketModelAtom)

	const disconnectAll = useCallback(() => {
		connectors.forEach(connector => {
			disconnect({ connector })
		})

		if (out) signOut()

		setSocket(INITIAL_SOCKET)
	}, [connectors, out, disconnect, setSocket])

	return useMemo(() => ({ ...rest, disconnect: disconnectAll, connectors }), [disconnectAll, connectors, rest])
}
