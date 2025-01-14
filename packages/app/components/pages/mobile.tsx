import { memo } from "react"

import { AuthFrame } from "@/components/app/frames/misc/auth"
import { PageContent } from "@/components/page/content"
import { PageNavbar } from "@/components/page/navbar"
import { PageHeader } from "@/components/page/header"
import { COLUMNS, useColumnData } from "@/state/columns"

export const MobileConsole = memo(() => {
    const { column } = useColumnData(COLUMNS.MOBILE_INDEX)
    
    return (
        <>
            <PageHeader />
            <PageContent />
            <PageNavbar />
            <AuthFrame />
        </>
    )
})

MobileConsole.displayName = "MobileConsole"
