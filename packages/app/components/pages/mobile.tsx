import { memo } from "react"

import { AuthFrame } from "@/components/app/frames/misc/auth"
import { PageContent } from "@/components/page/content"
import { PageNavbar } from "@/components/page/navbar"

export const MobileConsole = memo(() => {
	return (
		<>
			<PageContent />
			<PageNavbar />
			<AuthFrame />
		</>
	)
})

MobileConsole.displayName = "MobileConsole"
