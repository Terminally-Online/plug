import { useSession } from "next-auth/react"

import { useAccount as useAccountWagmi } from "wagmi"

export const useAccount = () => {
	const account = useAccountWagmi()

	const { data: session } = useSession()

	const isConnected = account.status === "connected"
	const isAuthenticated = session?.user.id?.startsWith("0x") || false

	// TODO: We have an issue caused by spreading the session into the same thing.
	return {
		id: session?.address ?? null,
		user: session?.user ?? null,
		expires: session?.expires ?? null,
		...account,
		isConnected,
		isAuthenticated
	}
}
