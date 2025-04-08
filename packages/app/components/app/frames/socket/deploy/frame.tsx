import { ChainId } from "@/lib";
import { FC, useCallback } from "react";
import { Frame } from "@/components/app/frames/base";
import { useAtom, useAtomValue } from "jotai";
import { columnByIndexAtom, COLUMNS, isFrameAtom } from "@/state/columns";
import { useBytecode, useSendTransaction } from "wagmi";
import { useSocket } from "@/state/authentication";
import { ChainSpecificButton } from "@/components/shared/buttons/authenticate";
import { Router } from "lucide-react";
import { api } from "@/server/client/api";
import { env } from "@/env";
import { ScrollingError } from "../../assets/scrolling-error";

export const SocketDeployFrame: FC<{ index: number, chainId: ChainId }> = ({ index, chainId }) => {
    const [column] = useAtom(columnByIndexAtom(index))
    const frameKey = `${chainId}-deploy`
    const isFrame = useAtomValue(isFrameAtom)(column, frameKey)

    const { socket } = useSocket()


    const { data: bytecode } = useBytecode({
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
                {error ? <ScrollingError error={error?.message ?? ""} /> : <p className="text-center text-sm opacity-40 font-bold max-w-[400px] mx-auto">The first time your Socket sees a transaction it will automatically deploy. If you are new here, you can disregard this. Please only use this for emergencies.</p>}


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
