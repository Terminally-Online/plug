import { FC } from "react"

import { Activity, Book, Twitter } from "lucide-react"

import { Image } from "@/components/app/utils/image"
import { LandingContainer } from "@/components/landing/layout/container"
import { Button } from "@/components/shared/buttons/button"
import { env } from "@/env"
import { GTM_EVENTS, routes, useAnalytics } from "@/lib"

export const Navbar: FC = () => {
	const handleCallToAction = useAnalytics(
		GTM_EVENTS.CTA_CLICKED,
		env.NEXT_PUBLIC_EARLY_ACCESS ? routes.earlyAccess : routes.app
	)

	return (
		<LandingContainer className="flex h-full flex-col py-8 text-plug-green">
			<div className="flex flex-row items-center justify-between xl:justify-start">
				<button
					className="flex flex-row items-center gap-4 md:gap-8 xl:mr-16"
					onClick={() => handleCallToAction(routes.index)}
				>
					<Image src="/plug-logo-green.svg" alt="Logo" width={24} height={24} />
					<Image src="/plug-word-green.svg" alt="Logo" width={64} height={24} />
				</button>

				<div className="ml-auto flex items-center gap-6 md:gap-8 xl:ml-0">
					<button onClick={() => handleCallToAction(routes.status)}>
						<Activity size={18} className="opacity-80 transition-opacity duration-200 hover:opacity-100" />
					</button>
					<button onClick={() => handleCallToAction(routes.twitter)}>
						<Twitter size={18} className="opacity-80 transition-opacity duration-200 hover:opacity-100" />
					</button>
					{/* <button onClick={() => handleCallToAction(routes.posts)}>
						<Book size={18} className="opacity-80 transition-opacity duration-200 hover:opacity-100" />
					</button> */}
				</div>

				<div className="mx-8 hidden h-[2px] w-full bg-plug-green/10 xl:block" />

				<Button
					variant="none"
					className="ml-6 w-max min-w-[110px] rounded-md border-[1px] border-plug-yellow/20 bg-plug-yellow px-4 py-2 text-center text-sm font-black text-plug-green filter backdrop-blur-xl transition-all duration-200 ease-in-out hover:bg-plug-yellow/50 md:ml-8"
					onClick={() => handleCallToAction()}
				>
					Enter App
				</Button>
			</div>
		</LandingContainer>
	)
}
