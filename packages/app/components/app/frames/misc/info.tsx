import { HelpCircle } from "lucide-react"
import { Frame } from "@/components/app/frames/base"
import { useColumnStore, COLUMNS } from "@/state/columns"

export const InfoFrame = () => {
    const { isFrame } = useColumnStore(COLUMNS.MOBILE_INDEX, "info-frame")

    return (
        <Frame
            index={COLUMNS.MOBILE_INDEX}
            icon={<HelpCircle size={18} />}
            label="Information"
            visible={isFrame}
            hasOverlay={true}
        >
            <div className="flex flex-col gap-4">
                {/* Content will be added later */}
            </div>
        </Frame>
    )
}