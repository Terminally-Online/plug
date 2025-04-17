import { SignInResponse } from "next-auth/react"

import { atom, useAtomValue } from "jotai"
import qrcode from "qrcode-generator"

import { WalletConnectProvider } from "@/lib"
import { RouterOutputs } from "@/server/client"

import { atomWithStorage } from "jotai/utils"

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

export const INITIAL_SOCKET: RouterOutputs["socket"]["get"] = {
	id: "anonymous-static",
	createdAt: new Date(),
	updatedAt: new Date(),
	deploymentFactory: null,
	deploymentNonce: 1738,
	deploymentDelegate: null,
	deploymentImplementation: null,
	deploymentSalt: null,
	admin: false,
	socketAddress: "",
	identity: {
		createdAt: new Date(),
		updatedAt: new Date(),
		onboardingAt: null,
		onboardingColor: null,
		onboardingCount: 0,
		socketId: "anonymous-static",
		farcasterId: null,
		ens: null,
		farcaster: null,
		referralCode: "",
		requestedAt: null,
		approvedAt: null,
		onboardedAt: null,
		referrerId: null
	}
}
export const socketModelAtom = atomWithStorage<RouterOutputs["socket"]["get"]>("plug.socketModel", INITIAL_SOCKET)
export const socketAtom = atom(get => {
	const socket = get(socketModelAtom)

	const isDemo = socket.id.startsWith("demo") || false
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
