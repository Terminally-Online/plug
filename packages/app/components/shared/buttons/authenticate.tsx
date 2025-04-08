import { FC } from "react"

import { Chain } from "viem"
import { useSwitchChain } from "wagmi"

import { Button } from "./button"
import { useAccount } from "@/lib/hooks/account/useAccount"

type Props = Parameters<typeof Button>[0] & { chain: Chain }

export const ChainSpecificButton: FC<Props> = ({ chain, ...rest }) => {
	const { chainId } = useAccount()

	const { switchChainAsync } = useSwitchChain()

	const handleSwitch = async () => {
		await switchChainAsync({ chainId: chain.id })
	}

	if (chainId !== chain.id)
		return (
			<Button {...rest} href={undefined} onClick={handleSwitch}>
				Switch Chain
			</Button>
		)

	return <Button {...rest} />
}
