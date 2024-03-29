import type { FC, PropsWithChildren } from "react"

import { Selector } from "."

import { useTabs } from "@/contexts"

import { Vault, Vaults, Wallet } from "./actions/manage"

export const Panel: FC<PropsWithChildren> = () => {
	const { pane } = useTabs()

	return pane === "wallet" ? (
		<Wallet />
	) : pane === "vaults" ? (
		<>
			<Selector />
			<Vaults />
		</>
	) : (
		<>
			<Selector />
			<Vault />
		</>
	)
}

export default Panel
