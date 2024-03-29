import { FC, PropsWithChildren } from "react"

import { Balance } from "@/components/viewport/vault/actions"

export const Deposit: FC<PropsWithChildren> = () => {
	return <Balance direction={1} action="Deposit" />
}

export default Deposit
