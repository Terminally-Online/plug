import { SignInResponse } from "next-auth/react"

import { atom } from "jotai"

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
