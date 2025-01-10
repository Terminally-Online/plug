import { memo } from "react"

import { AuthFrame } from "../app"
import { PageContent } from "../page/content"
import { PageNavbar } from "../page/navbar"

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
