import { SignInResponse } from "next-auth/react"

import { atom } from "jotai"
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
