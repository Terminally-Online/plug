import { memo } from "react"

import { AuthFrame } from "@/components/app/frames/misc/auth"
import { PageContent } from "@/components/page/content"
import { PageNavbar } from "@/components/page/navbar"
import { PageHeader } from "@/components/page/header"
import { COLUMNS, useColumnData } from "@/state/columns"

export const MobileConsole = memo(() => {
    const { column } = useColumnData(COLUMNS.MOBILE_INDEX)
    
    // Only show navbar if not on a Plug page
    const showNavbar = column?.key !== COLUMNS.KEYS.PLUG

    console.log("[MobileConsole] Render", {
        hasColumn: !!column,
        columnKey: column?.key,
        showNavbar
    })
    
    return (
        <>
            <PageHeader />
            <PageContent />
            {showNavbar && <PageNavbar />}
            <AuthFrame />
        </>
    )
})

MobileConsole.displayName = "MobileConsole"
