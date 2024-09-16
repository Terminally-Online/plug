import { useSession } from "next-auth/react"
import { createContext, FC, PropsWithChildren, useContext, useMemo, useState } from "react"

import { mainnet } from "viem/chains"
import { GetEnsAvatarReturnType, GetEnsNameReturnType, normalize } from "viem/ens"
import { useEnsAvatar, useEnsName } from "wagmi"

import { UserSocketModel } from "@/prisma/types"
import { api, RouterOutputs } from "@/server/client"

const TIME = 5 * 60 * 1000

export const SocketContext = createContext<{
	address: string | undefined
	name: GetEnsNameReturnType | undefined
	avatar: GetEnsAvatarReturnType | undefined
	socket: UserSocketModel | undefined
	collectibles: RouterOutputs["socket"]["balances"]["collectibles"]
	positions: RouterOutputs["socket"]["balances"]["positions"]
	isAnonymous: boolean
	isDemo: boolean
}>({
	address: undefined,
	name: undefined,
	avatar: undefined,
	socket: undefined,
	collectibles: [],
	positions: {
		tokens: [],
		protocols: []
	},
	isAnonymous: false,
	isDemo: false
})

export const SocketProvider: FC<PropsWithChildren> = ({ children }) => {
	const { data: session } = useSession()

	const { data: ensName } = useEnsName({
		chainId: mainnet.id,
		address: session?.address as `0x${string}`
	})
	const { data: ensAvatar } = useEnsAvatar({
		chainId: mainnet.id,
		name: normalize(ensName ?? "") || undefined
	})

	const { data: socketData } = api.socket.get.useQuery(
		{
			name: ensName,
			avatar: ensAvatar
		},
		{
			onSettled: data => setSocket(data)
		}
	)

	const isDemo = socketData?.id.startsWith("demo") || false
	const isAnonymous = isDemo || socketData?.id.startsWith("anonymous") || false
	const enabled =
		socketData !== undefined &&
		socketData.socketAddress !== undefined &&
		socketData.id.startsWith("anonymous") === false

	const { data: collectibles = [] } = api.socket.balances.collectibles.useQuery(socketData?.socketAddress, {
		enabled,
		staleTime: TIME,
		refetchInterval: TIME
	})
	const { data: positions = { tokens: [], protocols: [] } } = api.socket.balances.positions.useQuery(
		socketData?.socketAddress,
		{
			enabled,
			staleTime: TIME,
			refetchInterval: TIME
		}
	)

	const [socket, setSocket] = useState<UserSocketModel | undefined>(socketData)

	// useEffect(() => {
	// 	setSocket(socketData)
	// }, [socketData])

	return (
		<SocketContext.Provider
			value={{
				address: session?.address,
				name: socket?.identity?.ens?.name || ensName || undefined,
				avatar: socket?.identity?.ens?.avatar || ensAvatar || undefined,
				socket,
				collectibles,
				positions,
				isAnonymous,
				isDemo
			}}
		>
			{children}
		</SocketContext.Provider>
	)
}

export const useSockets = () => useContext(SocketContext)

// 	(id ?: string) => {
// 	const { socket, collectibles, positions, ...context } = useContext(SocketContext)
// 	const { id: socketId, columns } = socket ?? {}

// 	const column = useMemo(() => (columns ? columns.find(column => column.id === id) : undefined), [columns, id])
// 	const isExternal = useMemo(() => {
// 		return column !== undefined && column.viewAs !== null && column.viewAs.id !== socketId
// 	}, [column, socketId])
// 	const socketAddress = useMemo(() => {
// 		return column?.viewAs?.socketAddress
// 	}, [column])

// 	const { data: columnCollectibles = undefined } = api.socket.balances.collectibles.useQuery(socketAddress, {
// 		enabled: isExternal,
// 		staleTime: TIME,
// 		refetchInterval: TIME
// 	})
// 	const { data: columnPositions = undefined } = api.socket.balances.positions.useQuery(socketAddress, {
// 		enabled: isExternal,
// 		staleTime: TIME,
// 		refetchInterval: TIME
// 	})

// 	return {
// 		...context,
// 		socket,
// 		column,
// 		isExternal,
// 		collectibles: columnCollectibles ?? collectibles,
// 		positions: columnPositions ?? positions
// 	}
// }
