import { FC } from "react"

import Image from "next/image"

import { LandingContainer } from "@/components"
import { GTM_EVENTS, routes, useAnalytics } from "@/lib"

export const LandingFooter: FC = () => {
	const handleCallToAction = useAnalytics(GTM_EVENTS.CTA_CLICKED)

	return (
		<div className="relative z-[12] bg-white pt-16 lg:gap-4">
			<div className="absolute top-[-4px] h-[8px] w-full bg-gradient-to-r from-plug-green to-plug-yellow blur-xl filter" />
			<div className="absolute top-0 h-[2px] w-full bg-gradient-to-r from-plug-green to-plug-yellow" />

			<LandingContainer className="mb-16 flex-col gap-2">
				<Image className="mb-4" src="/black-logo.svg" alt="Logo" width={96} height={64} />

				<div className="grid grid-cols-1 font-bold text-black/40 lg:grid-cols-12">
					<div className="mb-8 lg:col-span-4 lg:mb-0">
						<p className="lg:max-w-[320px]">
							Automate your transactions on every popular Ethereum based blockchain so that you can log
							off and have everything run without you doing a thing.
						</p>
					</div>
					<div className="mb-2 flex flex-col items-start gap-2 lg:col-span-2">
						<button
							className="transition-all duration-200 ease-in-out hover:text-black"
							onClick={() => handleCallToAction(routes.comingSoon)}
						>
							Smart Contracts
						</button>
						<button
							className="transition-all duration-200 ease-in-out hover:text-black"
							onClick={() => handleCallToAction(routes.github)}
						>
							Code
						</button>
						<button
							className="transition-all duration-200 ease-in-out hover:text-black"
							onClick={() => handleCallToAction(routes.comingSoon)}
						>
							Audits
						</button>
					</div>
					<div className="mb-2 flex flex-col items-start gap-2 lg:col-span-2">
						<button
							className="transition-all duration-200 ease-in-out hover:text-black"
							onClick={() => handleCallToAction(routes.documentation)}
						>
							Documentation
						</button>
						<button
							className="transition-all duration-200 ease-in-out hover:text-black"
							onClick={() => handleCallToAction(routes.comingSoon)}
						>
							Guides
						</button>
					</div>
					<div className="mb-2 flex flex-col gap-2 lg:col-span-2"></div>
					<div className="mb-2 flex flex-col items-start gap-2 lg:col-span-2 lg:items-end lg:text-right">
						<button
							className="transition-all duration-200 ease-in-out hover:text-black"
							onClick={() => handleCallToAction("mailto:hello@onplug.io")}
						>
							hello@onplug.io
						</button>
						<button
							className="transition-all duration-200 ease-in-out hover:text-black"
							onClick={() => handleCallToAction(routes.twitter)}
						>
							Twitter
						</button>
						<button
							className="transition-all duration-200 ease-in-out hover:text-black"
							onClick={() => handleCallToAction(routes.comingSoon)}
						>
							Telegram
						</button>
					</div>
				</div>
			</LandingContainer>

			<div className="border-t-[2px] border-grayscale-100">
				<LandingContainer className="flex flex-col gap-2 py-4 font-bold text-black/40 lg:flex-row lg:gap-8">
					<p>Copyright Terminally Online, Inc. 2024</p>
					<button
						className="transition-all duration-200 ease-in-out hover:text-black lg:ml-auto"
						onClick={() => handleCallToAction(routes.comingSoon)}
					>
						Terms of Service
					</button>
					<button
						className="transition-all duration-200 ease-in-out hover:text-black"
						onClick={() => handleCallToAction(routes.comingSoon)}
					>
						Privacy Policy
					</button>
				</LandingContainer>
			</div>
		</div>
	)
}

export default LandingFooter
