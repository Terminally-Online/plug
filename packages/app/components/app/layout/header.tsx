import { FC } from "react"
import { ChevronLeft, LogOut } from "lucide-react" 
import { cn } from "@/lib/utils"
import { useRouter } from "next/router"
import { useDisconnect } from "wagmi"
import { signOut } from "next-auth/react"

type HeaderProps = {
    size?: "sm" | "md" | "lg"
    label: string | JSX.Element
    nextLabel?: string | JSX.Element 
    nextOnClick?: () => void | Promise<void>
    nextEmpty?: boolean
    onBack?: () => void
    className?: string
}

export const Header = ({
    size = "md",
    label,
    nextLabel,
    nextOnClick,
    nextEmpty,
    onBack,
    ...props
}: HeaderProps) => {
    return (
        <div className={cn("flex w-full flex-row items-center justify-between gap-4 px-4 py-4", props.className)}>
            <div className="flex flex-row items-center gap-4">
                {onBack && (
                    <button 
                        onClick={onBack}
                        className="flex h-8 w-8 items-center justify-center rounded-md hover:bg-plug-green/5"
                    >
                        <ChevronLeft size={18} />
                    </button>
                )}
                
                <div className={cn(
                    "font-bold",
                    size === "sm" && "text-sm", 
                    size === "md" && "text-base",
                    size === "lg" && "text-lg"
                )}>
                    {label}
                </div>
            </div>

            {nextLabel && (
                <button 
                    type="button"
                    onClick={(e) => {
                        e.stopPropagation()
                        console.log("Header: Next button clicked")
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
