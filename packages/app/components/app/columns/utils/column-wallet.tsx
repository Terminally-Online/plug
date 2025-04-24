import { FC } from "react"

import { SocketAssets } from "@/components/app/sockets/assets"
import { useSocket } from "@/state/authentication"

export const ColumnWallet: FC<{ index: number }> = () => {
	const { socket } = useSocket()

	return (
		<div className="flex h-full flex-col gap-4 overflow-y-scroll p-4 text-center">
			<SocketAssets index={-2} address={socket.id} hasTokens hasCollectibles />
		</div>
	)
}
