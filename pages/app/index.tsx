import { FC, useEffect, useRef, useState } from "react"

import { signIn, signOut, useSession } from "next-auth/react"

import { AuthFrame, ConsoleColumnRow, ConsoleSidebar, PageContent, PageHeader } from "@/components"
import { useSockets } from "@/contexts"
import { useMediaQuery } from "@/lib"

const MobilePage = () => {
	const { page } = useSockets()

	if (!page) return null

	return (
		<>
			<PageHeader />
			<PageContent />
			<AuthFrame id={page.id} />
		</>
	)
}

const DesktopPage = () => {
	return (
		<div className="min-w-screen flex h-screen w-full flex-row overflow-y-hidden overflow-x-visible">
			<ConsoleSidebar />
			<ConsoleColumnRow />
		</div>
	)
}

const Page = () => {
	const { data: session } = useSession()

	const { md } = useMediaQuery()

	useEffect(() => {
		if (session !== null) return

		signIn("credentials", {
			message: "0x0",
			signature: "0x0",
			redirect: false
		})
	}, [session])

	return <>{md ? <DesktopPage /> : <MobilePage />}</>
}

export default Page
