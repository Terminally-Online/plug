// import { signOut } from "next-auth/react"
// import { sendAnalyticsEvent } from "uniswap/src/features/telemetry/send"
// import { logger } from "utilities/src/logger/logger"
// import { getCurrentPageFromLocation } from "utils/urlRoutes"
// import { useRouter } from "next/router"
import { createContext, PropsWithChildren, useContext } from "react"

import { UserRejectedRequestError } from "viem"
import {
	ResolvedRegister,
	useAccount,
	UseAccountReturnType,
	UseConnectReturnType,
	useConnect as useConnectWagmi,
	useSignMessage,
	UseSignMessageReturnType
} from "wagmi"

// import { useAccountDrawer } from "components/AccountDrawer/MiniPortfolio/hooks"
// import { walletTypeToAmplitudeWalletType } from "components/Web3Provider/walletConnect"
import { useDisconnect } from "@/lib/hooks/wallet/useDisconnect"

const ConnectionContext = createContext<
	| {
			connection: UseConnectReturnType<ResolvedRegister["config"]>
			account: UseAccountReturnType<ResolvedRegister["config"]>
			sign: UseSignMessageReturnType
	  }
	| undefined
>(undefined)

export function ConnectionProvider({ children }: PropsWithChildren) {
	const connection = useConnectWagmi({
		mutation: {
			onError(error) {
				if (error instanceof UserRejectedRequestError) connection.reset()
			}
		}
	})

	const account = useAccount()
	const sign = useSignMessage()
	const { disconnect } = useDisconnect()

	// useEffect(() => {
	// 	if (!accountDrawer.isOpen && connection.isPending) {
	// 		connection.reset()
	// 		disconnect()
	// 	}
	// }, [connection, accountDrawer.isOpen, disconnect])

	return <ConnectionContext.Provider value={{ connection, account, sign }}>{children}</ConnectionContext.Provider>
}

/**
 * Wraps wagmi.useConnect in a singleton provider to provide the same connect state to all callers.
 * @see {@link https://wagmi.sh/react/api/hooks/useConnect}
 * @see {@link https://wagmi.sh/react/api/hooks/useAccount}
 * @see {@link https://wagmi.sh/react/api/hooks/useSignMessage}
 */
export function useConnect() {
	const value = useContext(ConnectionContext)
	if (!value) {
		throw new Error("useConnect must be used within a ConnectionProvider")
	}
	return value
}
