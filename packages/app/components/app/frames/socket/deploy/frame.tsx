import { ChainId } from "@/lib";
import { FC } from "react";
import { Frame } from "@/components/app/frames/base";
import { useAtom, useAtomValue } from "jotai";
import { columnByIndexAtom, COLUMNS, isFrameAtom } from "@/state/columns";
import { useBytecode, useSendTransaction } from "wagmi";
import { useSocket } from "@/state/authentication";
import { ChainSpecificButton } from "@/components/shared/buttons/authenticate";
import { Router } from "lucide-react";
import { api } from "@/server/client/api";

export const SocketDeployFrame: FC<{ index: number, chainId: ChainId }> = ({ index, chainId }) => {
    const [column] = useAtom(columnByIndexAtom(index))
    const frameKey = `${chainId}-deploy`
    const isFrame = useAtomValue(isFrameAtom)(column, frameKey)

    const { socket } = useSocket()

    const { error, sendTransaction, isPending } = useSendTransaction()

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
                nonce: 123,
                admin: "0x62180042606624f02d8a130da8a3171e9b33894d",
                delegate: "0x62180042606624f02d8a130da8a3171e9b33894d",
                implementation: "0x00000000906bb1a5fe6527c051A4C3b1c4595a8a"
            }
        ],
        options: {
            isEOA: true,
            simulate: true
        }
    }, {
        enabled: !bytecode && socket.id.startsWith("0x"),
    })

    const handleDeploy = () => {
        if (bytecode) return

        // send it to the router
        // solver needs to have an action for deployment

        // TODO: Implement this.
    }

    return (
        <Frame
            index={index}
            icon={<Router size={18} className="opacity-40" />}
            label="Publish Socket"
            visible={isFrame}
            hasOverlay
        >
            <div className="flex flex-col gap-4">
                <p className="text-center text-sm opacity-40 font-bold max-w-[400px] mx-auto">The first time your Socket sees a transaction it will automatically deploy. If you are new here, you can disregard this. Please only use this for emergencies.</p>
                <ChainSpecificButton chainId={chainId} className="py-4 w-full" onClick={handleDeploy} disabled={bytecode !== undefined}>
                    {bytecode ? "Already Published" : "Publish"}
                </ChainSpecificButton>
            </div>
        </Frame>
    )
}
