import { useSession } from "next-auth/react"
import { memo } from "react"

import { AuthFrame } from "@/components/app/frames/misc/auth"
import { PageContent } from "@/components/page/content"
import { PageHeader } from "@/components/page/header"
import { PageNavbar } from "@/components/page/navbar"
import { useSocket } from "@/state/authentication"
import { COLUMNS, useColumnData } from "@/state/columns"

import { ColumnAuthenticate } from "../app/columns/utils/column-authenticate"
import Container from "../app/layout/container"
import { ReferralRequired } from "../app/utils/referral-required"

export const MobileConsole = memo(() => {
	const { data: session } = useSession()
	const { socket } = useSocket()
	const { column } = useColumnData(COLUMNS.MOBILE_INDEX)

	const showNavbar = column?.key !== COLUMNS.KEYS.PLUG

	const isAuthenticated = session?.user.id?.startsWith("0x")
	const isReferred = Boolean(socket && socket.identity?.referrerId)

	const showUI = !isAuthenticated || isReferred

	return (
		<>
			{showUI && <PageHeader />}
			{!isAuthenticated ? (
				<Container>
					<ColumnAuthenticate index={COLUMNS.MOBILE_INDEX} />
				</Container>
			) : !isReferred ? (
				<Container>
					<ReferralRequired />
				</Container>
			) : (
				<PageContent />
			)}
			{showUI && showNavbar && <PageNavbar />}

			<AuthFrame />
		</>
	)
})

MobileConsole.displayName = "MobileConsole"
