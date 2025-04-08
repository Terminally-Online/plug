import { ChainId, formatAddress, formatTitle, getChainName } from "@/lib";
import { FC, useCallback } from "react";
import { Frame } from "@/components/app/frames/base";
import { useAtom, useAtomValue } from "jotai";
import { columnByIndexAtom, isFrameAtom } from "@/state/columns";
import { useBytecode, useSendTransaction, useBalance } from "wagmi";
import { useSocket } from "@/state/authentication";
import { ChainSpecificButton } from "@/components/shared/buttons/authenticate";
import { CircleDollarSign, Diameter, Hash, Router, Waypoints } from "lucide-react";
import { api } from "@/server/client/api";
import { formatEther } from "viem";
import { ScrollingError } from "@/components/app/frames/assets/scrolling-error";
import { ChainImage } from "@/components/app/sockets/chains/chain.image";

export const SocketDeployFrame: FC<{ index: number, chainId: ChainId }> = ({ index, chainId }) => {
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

    const { data: intent, error: intentError, isLoading } = api.solver.actions.intent.useQuery({
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
    }, {
        enabled: !bytecode && socket && socket.id.startsWith("0x"),
    })
    const { error, sendTransaction, isPending } = useSendTransaction()

    const handleDeploy = useCallback(() => {
        if (bytecode || !intent) return

        sendTransaction({
            to: intent.transactions[0].to,
            data: intent.transactions[0].data,
            value: intent.transactions[0].value,
        }, { onError: error => console.error(error) })
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
                {error ? <ScrollingError error={error?.message ?? ""} /> : <p className="text-center opacity-40 font-bold max-w-[400px] mx-auto">Please only use this for dire emergencies as it consumes gas not needed.</p>}

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
                            {balanceData ? Number(formatEther(balanceData.value)).toFixed(6) : '0.00'} ETH
                        </p>
                    </div>
                </div>

                <ChainSpecificButton
                    className="py-4 w-full"
                    chainId={chainId}
                    onClick={handleDeploy}
                    disabled={Boolean(bytecode)}
                >
                    {isPending ? "Publishing..." : bytecode ? "Already Published" : "Publish"}
                </ChainSpecificButton>
            </div>
        </Frame>
    )
}
