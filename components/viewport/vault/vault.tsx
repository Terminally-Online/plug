import type { FC, PropsWithChildren } from "react"

import { Panel } from "@/components/viewport/vault"
import { useTabs } from "@/contexts"

export const Vault: FC<PropsWithChildren> = () => {
	const { expanded } = useTabs()

	return expanded ? (
		<div className="flex h-screen flex-col bg-stone-900 text-white">
			<Panel />
		</div>
	) : null
}

export default Vault
