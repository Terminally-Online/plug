import {
	createContext,
	FC,
	PropsWithChildren,
	useContext,
	useMemo,
	useState
} from "react"

import { useSession } from "next-auth/react"

import { useEnsAvatar, useEnsName } from "wagmi"

import { UserSocket } from "@/server/api/routers/socket"
import { api } from "@/server/client"

import {
	GetEnsAvatarReturnType,
	GetEnsNameReturnType,
	normalize
} from "viem/ens"

export const SocketContext = createContext<{
	address: string | undefined
	ensName: GetEnsNameReturnType | undefined
	ensAvatar: GetEnsAvatarReturnType | undefined
	socket: UserSocket | undefined
	sockets: Array<UserSocket> | undefined
	handleAdd: () => void
	handleSelect: (address: string) => void
	handleRename: (name: string) => void
	handleDeploy: (chainIds: Array<number>, version?: number) => void
}>({
	address: undefined,
	ensName: undefined,
	ensAvatar: undefined,
	socket: undefined,
	sockets: undefined,
	handleAdd: () => {},
	handleSelect: () => {},
	handleRename: () => {},
	handleDeploy: () => {}
})

export const SocketProvider: FC<PropsWithChildren> = ({ children }) => {
	const { data: session } = useSession()
	const { data: apiSockets } = api.socket.all.useQuery()

	const { data: ensName } = useEnsName({
		address: session?.address as `0x${string}`
	})
	const { data: ensAvatar } = useEnsAvatar({
		name: normalize(ensName ?? "") || undefined
	})

	const [socketAddress, setSocketAddress] = useState<string | undefined>(
		undefined
	)
	const [sockets, setSockets] = useState<Array<UserSocket> | undefined>(
		apiSockets
	)

	const socket = useMemo(
		() =>
			sockets &&
			sockets.find(socket => socket.socketAddress === socketAddress),
		[sockets, socketAddress]
	)

	const handleSocketAdd = api.socket.add.useMutation()
	api.socket.onAdd.useSubscription(undefined, {
		onData: (data: UserSocket) =>
			setSockets(prev => (!prev ? [data] : [...prev, data]))
	})

	const handleSocketRename = api.socket.rename.useMutation()
	api.socket.onRename.useSubscription(undefined, {
		onData: (data: UserSocket) =>
			setSockets(prev =>
				!prev
					? [data]
					: prev.map(socket =>
							socket.socketAddress === data.socketAddress
								? data
								: socket
						)
			)
	})

	return (
		<SocketContext.Provider
			value={{
				address: session?.address,
				ensName,
				ensAvatar,
				socket,
				sockets,
				handleAdd: () => handleSocketAdd.mutate(),
				handleSelect: setSocketAddress,
				handleRename: (name: string) =>
					handleSocketRename.mutate({
						address: socketAddress || "",
						name
					}),
				handleDeploy: () => {}
			}}
		>
			{children}
		</SocketContext.Provider>
	)
}

export const useSockets = () => useContext(SocketContext)
