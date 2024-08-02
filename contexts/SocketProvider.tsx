import {
	createContext,
	FC,
	PropsWithChildren,
	useContext,
	useState
} from "react"

import { useSession } from "next-auth/react"

import { useEnsAvatar, useEnsName } from "wagmi"

import { UserSocketModel } from "@/prisma/types"
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
	socket: UserSocketModel | undefined
	handle: {
		columns: {
			add: (data: { key: string; id?: string }) => void
			remove: (id: string) => void
			resize: (data: { id: string; width: number }) => void
			move: (data: { from: number; to: number }) => void
		}
	}
}>({
	address: undefined,
	ensName: undefined,
	ensAvatar: undefined,
	socket: undefined,
	handle: {
		columns: {
			add: () => {},
			remove: () => {},
			resize: () => {},
			move: () => {}
		}
	}
})

export const SocketProvider: FC<PropsWithChildren> = ({ children }) => {
	const { data: session } = useSession()
	const { data: socketData } = api.socket.get.useQuery()

	const { data: ensName } = useEnsName({
		address: session?.address as `0x${string}`
	})
	const { data: ensAvatar } = useEnsAvatar({
		name: normalize(ensName ?? "") || undefined
	})

	const [socket, setSocket] = useState<UserSocketModel | undefined>(
		socketData
	)

	const handle = {
		columns: {
			add: api.socket.columns.add.useMutation({
				// TODO: Generate and hunt uuids and create it onMutate instead
				//       of waiting on the server.
				onSuccess: data => setSocket(data)
			}),
			remove: api.socket.columns.remove.useMutation({
				onMutate: data => {
					const previousSocket = socket

					setSocket(
						previousSocket && {
							...previousSocket,
							columns: previousSocket.columns
								.filter(column => column.id !== data)
								.map((column, index) => ({
									...column,
									index
								}))
						}
					)

					return previousSocket
				},
				onError: (_, __, context) => setSocket(context)
			}),
			resize: api.socket.columns.resize.useMutation({
				onMutate: data => {
					const previousSocket = socket

					setSocket(
						previousSocket && {
							...previousSocket,
							columns: previousSocket.columns.map(column =>
								column.id === data.id
									? { ...column, width: data.width }
									: column
							)
						}
					)

					return previousSocket
				},
				onError: (_, __, context) => setSocket(context)
			}),
			move: api.socket.columns.move.useMutation({
				onMutate: data => {
					const previousSocket = socket

					const columns = previousSocket?.columns ?? []
					const [removed] = columns.splice(data.from, 1)
					columns.splice(data.to, 0, removed)

					setSocket(
						previousSocket && {
							...previousSocket,
							columns: columns.map((column, index) => ({
								...column,
								index
							}))
						}
					)

					return previousSocket
				},
				onError: (_, __, context) => setSocket(context)
			})
		}
	}

	return (
		<SocketContext.Provider
			value={{
				address: session?.address,
				ensName,
				ensAvatar,
				socket,
				handle: {
					columns: {
						add: data => handle.columns.add.mutate(data),
						remove: data => handle.columns.remove.mutate(data),
						resize: data => handle.columns.resize.mutate(data),
						move: data => handle.columns.move.mutate(data)
					}
				}
			}}
		>
			{children}
		</SocketContext.Provider>
	)
}

export const useSockets = () => useContext(SocketContext)
