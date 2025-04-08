import { ChainImage } from "@/components/app/sockets/chains/chain.image"
import { ChainId, cn, getChainName } from "@/lib"
import { useColumnActions } from "@/state/columns"
import { FC } from "react"
import { useBytecode } from "wagmi"

type ColumnSettingsDeploymentItemProps = {
	index: number
	chainId: ChainId
	factory: string | null
	address: string
}
export const ColumnSettingsDeploymentItem: FC<ColumnSettingsDeploymentItemProps> = ({ index, chainId, factory, address }) => {
	const { frame } = useColumnActions(index)

	const { data: factoryBytecode } = useBytecode({
		chainId,
		address: factory as `0x${string}`
	})

	const { data: deploymentBytecode } = useBytecode({
		chainId,
		address: address as `0x${string}`
	})

	const handleDeploy = () => {
		if (!factoryBytecode || deploymentBytecode) return

		frame(`${chainId}-deploy`)
	}

	if (!factory) return null

	return (
		<p className="flex flex-row items-center justify-between gap-2 font-bold">
			<ChainImage chainId={chainId} size="xs" />
			<span className="opacity-40">{getChainName(chainId)}</span>{" "}
			<span
				className={cn(
					"group ml-auto flex flex-row items-center gap-1",
					factoryBytecode && !deploymentBytecode && "cursor-pointer"
				)}
				onClick={handleDeploy}
			>
				{factoryBytecode ? deploymentBytecode ? "Published" : "Ready to Publish" : "In Development"}
			</span>
		</p>
	)
}
