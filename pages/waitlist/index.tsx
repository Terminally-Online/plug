import { useEffect } from "react"

import { useRouter } from "next/router"

import { Button, Callout, StaticLayout } from "@/components"
import { GTM_EVENTS, useAnalytics } from "@/lib"

const WAITLIST_FORM =
	"https://docs.google.com/forms/d/e/1FAIpQLSf4ttqF5PizhP_F2jHBGTuaH-q6YunG4PkUcaK8JRhljXg5oQ/viewform"

const Page = () => {
	const router = useRouter()
	const handleCallToAction = useAnalytics(GTM_EVENTS.CTA_CLICKED, WAITLIST_FORM, true, "/")

	useEffect(() => {
		const timeout = setTimeout(() => {
			handleCallToAction()
		}, 250)

		return () => clearTimeout(timeout)
	}, [router, handleCallToAction])

	return (
		<StaticLayout title="Waitlist">
			<div className="flex h-screen w-screen items-center justify-center">
				<Callout
					title="You are being redirected to the waitlist. One moment please..."
					description="If you are not automatically redirect, please click the button below."
				>
					<Button variant="secondary" sizing="sm" href="/">
						Home
					</Button>
					<Button sizing="sm" onClick={() => handleCallToAction()}>
						Join Waitlist
					</Button>
				</Callout>
			</div>
		</StaticLayout>
	)
}

export default Page
