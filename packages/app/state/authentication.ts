import { SignInResponse } from "next-auth/react"

import { UserSocketModel } from "@/prisma/types"

import { atom, useAtomValue } from "jotai"
import qrcode from "qrcode-generator"

import { WalletConnectProvider } from "@/lib"

export const authenticationResponseAtom = atom<SignInResponse | undefined>(undefined)
export const authenticationLoadingAtom = atom(false)
export const authenticationAtom = atom(get => ({
	isError: !!get(authenticationResponseAtom)?.error,
	error: get(authenticationResponseAtom)?.error ?? undefined,
	isLoading: get(authenticationLoadingAtom) && get(authenticationResponseAtom) === undefined,
	isSuccess: get(authenticationResponseAtom)?.ok === true,
	status: get(authenticationResponseAtom)?.status,
	url: get(authenticationResponseAtom)?.url
}))

export const walletConnectProviderAtom = atom<WalletConnectProvider | undefined>(undefined)
export const walletConnectURIAtom = atom<string | undefined>(undefined)
export const walletConnectURIMatrixAtom = atom<
	{ moduleCount: number; getModule: (row: number, col: number) => boolean } | undefined
>(get => {
	const uri = get(walletConnectURIAtom)

	if (uri === undefined) return

	const qr = qrcode(0, "L")
	qr.addData(uri)
	qr.make()

	return {
		moduleCount: qr.getModuleCount(),
		getModule: (row: number, col: number) => qr.isDark(row, col)
	}
})

export const socketModelAtom = atom<UserSocketModel | undefined>(undefined)
export const socketAtom = atom(get => {
	const socket = get(socketModelAtom)

	const isDemo = socket?.id.startsWith("demo") || false
	const isAnonymous = socket === undefined || isDemo || socket?.id.startsWith("anonymous") || false
	const isApproved = (socket && Boolean(socket.identity?.approvedAt)) || false

	const name = socket?.identity?.ens?.name ?? undefined
	const avatar = socket?.identity?.ens?.avatar ?? undefined

	return {
		isDemo,
		isAnonymous,
		isApproved,
		name,
		avatar,
		socket
	}
})
export const useSocket = () => useAtomValue(socketAtom)
