import { memo } from "react"
import { useSession } from "next-auth/react"
import { AuthFrame } from "@/components/app/frames/misc/auth"
import { InfoFrame } from "@/components/app/frames/misc/info"
import { PageContent } from "@/components/page/content"
import { PageNavbar } from "@/components/page/navbar"
import { PageHeader } from "@/components/page/header"
import { COLUMNS, useColumnData } from "@/state/columns"
import { useSocket } from "@/state/authentication"

export const MobileConsole = memo(() => {
    const { data: session } = useSession()
    const { socket } = useSocket()
    const { column } = useColumnData(COLUMNS.MOBILE_INDEX)
    
    // Only show navbar if not on a Plug page
    const showNavbar = column?.key !== COLUMNS.KEYS.PLUG
    
    // Check if user is referred
    const isReferred = Boolean(socket && socket.identity?.referrerId)
    const isAuthenticated = session?.user.id?.startsWith("0x")
    
    // Hide UI elements if authenticated but not referred
    const showUI = !isAuthenticated || isReferred

    return (
        <>
            {showUI && <PageHeader />}
            <PageContent />
            {showUI && showNavbar && <PageNavbar />}
            <AuthFrame />
            <InfoFrame />
        </>
    )
})

MobileConsole.displayName = "MobileConsole"
