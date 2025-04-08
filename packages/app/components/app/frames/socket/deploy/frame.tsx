import { ChainId } from "@/lib";
import { FC } from "react";
import { Frame } from "@/components/app/frames/base";
import { useAtom, useAtomValue } from "jotai";
import { columnByIndexAtom, isFrameAtom } from "@/state/columns";
import { useBytecode, useSendTransaction } from "wagmi";
import { useSocket } from "@/state/authentication";
import { ChainSpecificButton } from "@/components/shared/buttons/authenticate";
import { Router } from "lucide-react";

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

    const handleDeploy = () => {
        if (bytecode) return

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
                <p className="text-center text-sm opacity-40 font-bold">The first time your Socket sees a transaction it will automatically deploy. If you're new here, you can disregard this. Please only use this for emergencies.</p>
                <ChainSpecificButton chainId={chainId} className="py-4 w-full" onClick={handleDeploy} disabled={bytecode !== undefined}>
                    {bytecode ? "Already Published" : "Publish"}
                </ChainSpecificButton>
            </div>
        </Frame>
    )
}
