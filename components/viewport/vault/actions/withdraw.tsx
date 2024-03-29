import { FC, PropsWithChildren } from "react"

import { Balance } from "@/components/viewport/vault/actions"

export const Withdraw: FC<PropsWithChildren> = () => {
	return <Balance direction={-1} action="Withdrawal" />
}

export default Withdraw
