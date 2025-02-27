import { FC } from "react"

import { ShieldX } from "lucide-react"

import { Frame } from "@/components/app/frames/base"
import { useColumnStore } from "@/state/columns"

export const ErrorFrame: FC<{ index: number }> = ({ index }) => {
    const { column, isFrame } = useColumnStore(index, "error")

    if (!column) return null

    return (
        <Frame
            index={index}
            className="z-[2]"
            icon={<ShieldX size={18} className="opacity-40" />}
            label="Error Encountered"
            visible={isFrame}
        >
            <div className="font-bold flex flex-col gap-4">
                <p className="opacity-40 max-w-[380px] mx-auto">We encountered an error while processing your request. Our team has been automatically notified.</p>
            </div>
        </Frame>
    )
}
