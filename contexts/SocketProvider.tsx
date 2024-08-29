import { createContext, FC, PropsWithChildren, useContext, useMemo, useState } from "react"

import { useSession } from "next-auth/react"

import { useEnsAvatar, useEnsName } from "wagmi"

import { ConsoleColumnModel, UserSocketModel } from "@/prisma/types"
import { api, RouterOutputs } from "@/server/client"

import { mainnet } from "viem/chains"
import { GetEnsAvatarReturnType, GetEnsNameReturnType, normalize } from "viem/ens"

export const SocketContext = createContext<{
	address: string | undefined
	name: GetEnsNameReturnType | undefined
	avatar: GetEnsAvatarReturnType | undefined
	socket: UserSocketModel | undefined
	collectibles: RouterOutputs["socket"]["balances"]["collectibles"]
	positions: RouterOutputs["socket"]["balances"]["positions"]
	page: ConsoleColumnModel | undefined
	isAnonymous: boolean
	handle: {
		columns: {
			add: (data: { key: string; id?: string; index?: number; item?: string }) => void
			navigate: (data: { id?: string; key: string; item?: string; from?: string }) => void
			remove: (id: string) => void
			move: (data: { from: number; to: number }) => void
			as: (data: { id: string; as: string }) => void
		}
	}
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
	page: undefined,
	isAnonymous: false,
	handle: {
		columns: {
			add: () => {},
			navigate: () => {},
			remove: () => {},
			move: () => {},
			as: () => {}
		}
	}
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

	const { data: socketData } = api.socket.get.useQuery({
		name: ensName,
		avatar: ensAvatar
	})

	const isAnonymous = !socketData || socketData.id.startsWith("anonymous")
	const enabled =
		socketData !== undefined &&
		socketData.socketAddress !== undefined &&
		socketData.id.startsWith("anonymous") === false

	const { data: collectibles = [] } = api.socket.balances.collectibles.useQuery(socketData?.socketAddress, {
		enabled
	})

	const { data: positions = { tokens: [], protocols: [] } } = api.socket.balances.positions.useQuery(
		socketData?.socketAddress,
		{
			enabled,
			refetchInterval: 5 * 60 * 1000
		}
	)

	const [socket, setSocket] = useState<UserSocketModel | undefined>(socketData)

	const page = useMemo(() => socket?.columns.find(column => column.index === -1), [socket])

	const handle = {
		columns: {
			add: api.socket.columns.add.useMutation({
				onSuccess: data => setSocket(data)
			}),
			navigate: api.socket.columns.navigate.useMutation({
				onMutate: data => {
					setSocket(
						prev =>
							prev && {
								...prev,
								columns: prev.columns.map(column =>
									column.id === data.id
										? {
												...column,
												key: data.key,
												item: data.item ?? null,
												from: data.from ?? null
											}
										: column
								)
							}
					)
				}
			}),
			remove: api.socket.columns.remove.useMutation({
				onMutate: data => {
					const previousSocket = socket

					setSocket(
						previousSocket && {
							...previousSocket,
							columns: previousSocket.columns
								.filter(column => column.id !== data)
								.sort((a, b) => a.index - b.index)
								.map((column, index) => ({
									...column,
									// Subtract 1 to account for the app column.
									index: index - 1
								}))
						}
					)

					return previousSocket
				},
				onError: (_, __, context) => setSocket(context)
			}),
			move: api.socket.columns.move.useMutation(),
			as: api.socket.columns.as.useMutation({
				onSuccess: data => setSocket(data)
			})
		}
	}

	return (
		<SocketContext.Provider
			value={{
				address: session?.address,
				name: socket?.identity?.ens?.name || ensName || undefined,
				avatar: socket?.identity?.ens?.avatar || ensAvatar || undefined,
				socket,
				collectibles,
				positions,
				page,
				isAnonymous,
				handle: {
					columns: {
						add: data => handle.columns.add.mutate(data),
						navigate: data => handle.columns.navigate.mutate(data),
						remove: data => handle.columns.remove.mutate(data),
						move: data => handle.columns.move.mutate(data),
						as: data => handle.columns.as.mutate(data)
					}
				}
			}}
		>
			{children}
		</SocketContext.Provider>
	)
}

export const useSockets = () => useContext(SocketContext)
