import { signOut } from "next-auth/react"
import { useCallback, useMemo } from "react"

import { UseDisconnectReturnType, useDisconnect as useDisconnectWagmi } from "wagmi"

export function useDisconnect(out: boolean = false): UseDisconnectReturnType {
	const { connectors, disconnect, ...rest } = useDisconnectWagmi()

	const disconnectAll = useCallback(() => {
		connectors.forEach(connector => {
			disconnect({ connector })
		})

		if (out) signOut()
	}, [connectors, out, disconnect])

	return useMemo(() => ({ ...rest, disconnect: disconnectAll, connectors }), [disconnectAll, connectors, rest])
}
