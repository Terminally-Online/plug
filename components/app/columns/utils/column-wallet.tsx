import { useSession } from "next-auth/react"
import { FC } from "react"

import { useSocket } from "@/state"

export const ColumnWallet: FC<{ index: number }> = () => {
	const { data: session } = useSession()
	const { socket } = useSocket()

	if (!socket || !session?.user.id) return null

	return <div className="flex h-full flex-col gap-4 overflow-y-scroll text-center">Wallet</div>
}
