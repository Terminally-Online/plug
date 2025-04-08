import { FC } from "react"

import { useSwitchChain } from "wagmi"

import { Button } from "./button"
import { useAccount } from "@/lib/hooks/account/useAccount"
import { ChainId } from "@/lib"

type Props = Parameters<typeof Button>[0] & { chainId: ChainId }

export const ChainSpecificButton: FC<Props> = ({ chainId, ...rest }) => {
	const account = useAccount()

	const { switchChainAsync } = useSwitchChain()

	const handleSwitch = async () => {
		await switchChainAsync({ chainId: chainId })
	}

	if (chainId !== account.chainId)
		return (
			<Button {...rest} href={undefined} onClick={handleSwitch}>
				Switch Chain
			</Button>
		)

	return <Button {...rest} />
}
