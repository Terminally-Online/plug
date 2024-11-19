import { signIn, useSession } from "next-auth/react"
import { useEffect } from "react"

import { ConsoleColumnRow, ConsoleSidebar } from "@/components"
import { socket } from "@/server/api/routers/socket"

const Page = () => {
	const { data: session } = useSession()

	useEffect(() => {
		if (session?.user.id) return
	
		signIn("credentials", {
			message: "0x0",
			signature: "0x0",
			chainId: 0,
			redirect: false
		})
	}, [session?.user.id])

	return (
		<div className="min-w-screen flex w-full flex-row overflow-visible">
			<ConsoleSidebar />
			<ConsoleColumnRow />
		</div>
	)
}

export default Page
