import { useSession } from "next-auth/react"
import { useAccount as useAccountWagmi } from "wagmi"

export const useAccount = () => {
	const account = useAccountWagmi()

	const { data: session } = useSession()

	const isConnected = account.status === "connected"
	const isAuthenticated = session?.user.id?.startsWith("0x") || false

	return { ...account, ...session, isConnected, isAuthenticated }
}
