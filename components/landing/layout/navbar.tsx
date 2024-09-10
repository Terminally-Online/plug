import { FC } from "react"

import Image from "next/image"

import { Book, Twitter } from "lucide-react"

import { Button, LandingContainer } from "@/components"
import { GTM_EVENTS, routes, useAnalytics } from "@/lib"

export const Navbar: FC = () => {
	const handleCallToAction = useAnalytics(
		GTM_EVENTS.CTA_CLICKED,
		process.env.NEXT_PUBLIC_EARLY_ACCESS === "false" ? routes.app : routes.earlyAccess
	)

	return (
		<LandingContainer className="fixed z-[10] w-full items-center gap-8 bg-gradient-to-b from-white to-white/0 py-8">
			<div className="flex w-full flex-row items-center gap-4">
				<button className="mr-8" onClick={() => handleCallToAction(routes.index)}>
					<Image src="/black-icon.svg" alt="Logo" width={24} height={24} />
				</button>
				<button className="mr-4" onClick={() => handleCallToAction(routes.documentation)}>
					<Book size={18} className="opacity-40 transition-opacity duration-200 hover:opacity-100" />
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
