import { FC } from "react"

import { Book, Twitter } from "lucide-react"

import { Button, Image, LandingContainer } from "@/components"
import { env } from "@/env"
import { GTM_EVENTS, routes, useAnalytics } from "@/lib"

export const Navbar: FC = () => {
	const handleCallToAction = useAnalytics(
		GTM_EVENTS.CTA_CLICKED,
		env.NEXT_PUBLIC_EARLY_ACCESS ? routes.earlyAccess : routes.app
	)

	return (
		<LandingContainer className="fixed z-[10] w-full items-center gap-8 bg-gradient-to-b from-white to-white/0 py-8">
			<div className="flex w-full flex-row items-center gap-4">
				<button
					className="mr-8 flex flex-row items-center gap-8"
					onClick={() => handleCallToAction(routes.index)}
				>
					<Image src="/plug-logo-green.svg" alt="Logo" width={32} height={32} />
					<Image src="/plug-word-green.svg" alt="Logo" width={64} height={32} />
				</button>
				<button onClick={() => handleCallToAction(routes.twitter)}>
					<Twitter size={18} className="opacity-40 transition-opacity duration-200 hover:opacity-100" />
				</button>

				<Button
					variant="secondary"
					className="ml-auto px-4 py-2"
					sizing="sm"
					onClick={() => handleCallToAction()}
				>
					Enter App
				</Button>
			</div>
		</LandingContainer>
	)
}

export default Navbar
