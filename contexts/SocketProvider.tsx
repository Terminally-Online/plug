import { createContext, FC, PropsWithChildren, useContext, useMemo, useState } from "react"

import { useSession } from "next-auth/react"

import { useEnsAvatar, useEnsName } from "wagmi"

import { ConsoleColumnModel, UserSocketModel } from "@/prisma/types"
import { api } from "@/server/client"

import { mainnet } from "viem/chains"
import { GetEnsAvatarReturnType, GetEnsNameReturnType, normalize } from "viem/ens"

export const SocketContext = createContext<{
	address: string | undefined
	name: GetEnsNameReturnType | undefined
	avatar: GetEnsAvatarReturnType | undefined
	socket: UserSocketModel | undefined
	page: ConsoleColumnModel | undefined
	anonymous: boolean
	handle: {
		columns: {
			add: (data: { key: string; id?: string; index?: number; item?: string }) => void
			navigate: (data: { id?: string; key: string; item?: string; from?: string }) => void
			remove: (id: string) => void
			move: (data: { from: number; to: number }) => void
		}
	}
}>({
	address: undefined,
	name: undefined,
	avatar: undefined,
	socket: undefined,
	page: undefined,
	anonymous: false,
	handle: {
		columns: {
			add: () => {},
			navigate: () => {},
			remove: () => {},
			move: () => {}
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

	const anonymous = !socketData || socketData.id.startsWith("anonymous")

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
			move: api.socket.columns.move.useMutation()
		}
	}

	return (
		<SocketContext.Provider
			value={{
				address: session?.address,
				name: socket?.identity?.ens?.name || ensName || undefined,
				avatar: socket?.identity?.ens?.avatar || ensAvatar || undefined,
				socket,
				page,
				anonymous,
				handle: {
					columns: {
						add: data => handle.columns.add.mutate(data),
						navigate: data => handle.columns.navigate.mutate(data),
						remove: data => handle.columns.remove.mutate(data),
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
