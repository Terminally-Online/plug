import { FC, useCallback } from "react"

import { formatEther } from "viem"
import { useBalance, useBytecode, useSendTransaction } from "wagmi"

import { CircleDollarSign, Diameter, Router, Waypoints } from "lucide-react"

import { useAtom, useAtomValue } from "jotai"

import { ScrollingError } from "@/components/app/frames/assets/scrolling-error"
import { Frame } from "@/components/app/frames/base"
import { ChainImage } from "@/components/app/sockets/chains/chain.image"
import { ChainSpecificButton } from "@/components/shared/buttons/authenticate"
import { Button } from "@/components/shared/buttons/button"
import { ChainId, formatTitle, getBlockExplorerAddress, getChainName, routes } from "@/lib"
import { api } from "@/server/client/api"
import { useSocket } from "@/state/authentication"
import { columnByIndexAtom, isFrameAtom } from "@/state/columns"

export const SocketDeployFrame: FC<{ index: number; chainId: ChainId }> = ({ index, chainId }) => {
	const [column] = useAtom(columnByIndexAtom(index))
	const frameKey = `${chainId}-deploy`
	const isFrame = useAtomValue(isFrameAtom)(column, frameKey)

	const { socket } = useSocket()

	const { data: bytecode } = useBytecode({
		chainId,
		address: socket.socketAddress as `0x${string}`
	})

	const { data: balanceData } = useBalance({
		chainId,
		address: socket.socketAddress as `0x${string}`
	})

	const { data: intent, error: intentError } = api.solver.actions.intent.useQuery(
		{
			chainId,
			from: socket.id,
			inputs: [
				{
					protocol: "plug",
					action: "deploy",
					nonce: socket.deploymentNonce ?? 1738,
					admin: socket.id ?? "",
					delegate: socket.deploymentDelegate ?? "",
					implementation: socket.deploymentImplementation ?? ""
				}
			],
			options: {
				isEOA: true,
				simulate: true
			}
		},
		{
			enabled: !bytecode && socket && socket.id.startsWith("0x")
		}
	)
	const { error, sendTransaction, isPending } = useSendTransaction()

	const handleDeploy = useCallback(() => {
		if (bytecode || !intent) return

		sendTransaction(
			{
				to: intent.transactions[0].to,
				data: intent.transactions[0].data,
				value: intent.transactions[0].value
			},
			{ onError: error => console.error(error) }
		)
	}, [bytecode, intent])

	return (
		<Frame
			index={index}
			icon={<Router size={18} className="opacity-40" />}
			label="Publish Socket"
			visible={isFrame}
			hasOverlay
		>
			<div className="flex flex-col gap-4">
				{intentError || error ? (
					<ScrollingError error={intentError?.message ?? error?.message ?? ""} />
				) : (
					<p className="mx-auto max-w-[400px] text-center font-bold opacity-40">
						Only use this for emergencies as it consumes gas not needed. You can read more in{" "}
						<a href={routes.documentation}>our documentation</a>.
					</p>
				)}

				<div className="flex flex-col gap-2">
					<div className="flex flex-row items-center gap-4">
						<p className="font-bold opacity-40">Details</p>
						<div className="h-[2px] w-full bg-plug-green/10" />
					</div>
					<p className="flex w-full flex-row items-center gap-4 font-bold">
						<Waypoints size={18} className="opacity-20" />
						<span className="mr-auto opacity-40">Chain</span>
						<span className="flex flex-row items-center gap-2">
							<ChainImage chainId={chainId} size="xs" />
							{formatTitle(getChainName(chainId))}
						</span>
					</p>
					<div className="w-full font-bold">
						<p className="flex w-full flex-row items-center gap-4">
							<Diameter size={18} className="opacity-20" />
							<span className="mr-auto opacity-40">Code Length</span>
							{bytecode?.length ?? 0}
						</p>
					</div>
					<div className="w-full font-bold">
						<p className="flex w-full flex-row items-center gap-4">
							<CircleDollarSign size={18} className="opacity-20" />
							<span className="mr-auto opacity-40">Native Balance</span>
							{balanceData ? Number(formatEther(balanceData.value)).toFixed(6) : "0.00"} ETH
						</p>
					</div>
				</div>

				<div className="flex flex-row gap-2">
					<Button
						variant="secondary"
						className="w-max"
						href={getBlockExplorerAddress(chainId, socket.socketAddress)}
					>
						View
					</Button>
					<ChainSpecificButton
						className="w-full py-4"
						chainId={chainId}
						onClick={handleDeploy}
						disabled={Boolean(bytecode)}
					>
						{isPending ? "Publishing..." : bytecode ? "Already Published" : "Publish"}
					</ChainSpecificButton>
				</div>
			</div>
		</Frame>
	)
}
