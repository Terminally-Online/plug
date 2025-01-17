import { FC } from "react"
import { ChevronLeft } from "lucide-react" 
import { cn } from "@/lib/utils"

export const Header: FC<{
    size?: "sm" | "md" | "lg"
    label: string | JSX.Element
    nextLabel?: string | JSX.Element 
    nextOnClick?: () => void
    nextEmpty?: boolean
    nextPadded?: boolean
    onBack?: () => void
    className?: string
    icon?: JSX.Element
}> = ({
    size = "md",
    label,
    nextLabel,
    nextOnClick,
    nextEmpty,
    nextPadded = true,
    onBack,
    className,
    icon
}) => {
    return (
        <div className={cn("flex w-full flex-row items-center justify-between gap-4", 
            nextPadded && "px-4 py-4",
            className
        )}>
            <div className="flex flex-row items-center gap-4">
                {onBack && (
                    <button 
                        onClick={onBack}
                        className="flex h-8 w-8 items-center justify-center rounded-md hover:bg-plug-green/5"
                    >
                        <ChevronLeft size={18} />
                    </button>
                )}
                
                <div className="flex items-center gap-2">
                    {icon}
                    <div className={cn(
                        "font-bold",
                        size === "sm" && "text-sm", 
                        size === "md" && "text-base",
                        size === "lg" && "text-lg"
                    )}>
                        {label}
                    </div>
                </div>
            </div>

            {nextLabel && (
                <button 
                    type="button"
                    onClick={(e) => {
                        e.stopPropagation()
                        nextOnClick?.()
                    }}
                    className={cn(
                        "flex h-8 items-center justify-center rounded-md px-2 cursor-pointer",
                        !nextEmpty && "bg-plug-green/5"
                    )}
                >
                    {nextLabel}
                </button>
            )}
        </div>
    )
}