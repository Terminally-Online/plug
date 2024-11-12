import { FC, useRef } from "react"

import { motion, useSpring } from "framer-motion"

import { LandingContainer } from "@/components"
import { GTM_EVENTS, routes, useAnalytics } from "@/lib"
import { useScroll, useTransform } from "framer-motion"

export const LandingFooter: FC = () => {
	const footerRef = useRef<HTMLDivElement>(null)
	const { scrollYProgress } = useScroll({
		target: footerRef,
		offset: ["start end", "end end"]
	})
	const springyProgress = useSpring(scrollYProgress, {
		stiffness: 100,
		damping: 30,
		restDelta: 0.001
	})
	const y = useTransform(
		springyProgress,
		[0, 1],
		["-100%", "0%"],
	)

	const handleCallToAction = useAnalytics(GTM_EVENTS.CTA_CLICKED)

	return (
		<div ref={footerRef} className="relative z-[1] overflow-hidden bg-white pt-32 lg:gap-4 w-full h-full">
			<div className="absolute top-0 h-[2px] w-full bg-gradient-to-r from-plug-green to-plug-yellow" />
			<motion.div
				className="pointer-events-none absolute inset-0 mb-4 w-full opacity-[4%]"
				style={{ 
					y: y,
					backgroundImage: "url(/plug-word-green.svg)", 
					backgroundSize: "cover", 
					backgroundPosition: "center", 
					backgroundRepeat: "no-repeat" 
				}}
			/>

			<LandingContainer className="mb-32 flex-col gap-2">
				<div className="grid grid-cols-1 font-bold text-plug-green/40 lg:grid-cols-12">
					<div className="mb-8 lg:col-span-4 lg:mb-0">
						<p className="lg:max-w-[320px]">
							Automate your transactions on every popular Ethereum based blockchain so that you can log
							off and have everything run without you doing a thing.
						</p>
					</div>
					<div className="mb-2 flex flex-col items-start gap-2 lg:col-span-2">
						<button
							className="transition-all duration-200 ease-in-out hover:text-plug-green"
							onClick={() => handleCallToAction(routes.comingSoon)}
						>
							Smart Contracts
						</button>
						<button
							className="transition-all duration-200 ease-in-out hover:text-plug-green"
							onClick={() => handleCallToAction(routes.github)}
						>
							Code
						</button>
						<button
							className="transition-all duration-200 ease-in-out hover:text-plug-green"
							onClick={() => handleCallToAction(routes.comingSoon)}
						>
							Audits
						</button>
					</div>
					<div className="mb-2 flex flex-col items-start gap-2 lg:col-span-2">
						<button
							className="transition-all duration-200 ease-in-out hover:text-plug-green"
							onClick={() => handleCallToAction(routes.documentation)}
						>
							Documentation
						</button>
						<button
							className="transition-all duration-200 ease-in-out hover:text-plug-green"
							onClick={() => handleCallToAction(routes.comingSoon)}
						>
							Guides
						</button>
					</div>
					<div className="mb-2 flex flex-col gap-2 lg:col-span-2"></div>
					<div className="mb-2 flex flex-col items-start gap-2 lg:col-span-2 lg:items-end lg:text-right">
						<button
							className="transition-all duration-200 ease-in-out hover:text-plug-green"
							onClick={() => handleCallToAction("mailto:hello@onplug.io")}
						>
							hello@onplug.io
						</button>
						<button
							className="transition-all duration-200 ease-in-out hover:text-plug-green"
							onClick={() => handleCallToAction(routes.twitter)}
						>
							Twitter
						</button>
						<button
							className="transition-all duration-200 ease-in-out hover:text-plug-green"
							onClick={() => handleCallToAction(routes.comingSoon)}
						>
							Telegram
						</button>
					</div>
				</div>
			</LandingContainer>

			<div className="border-t-[2px] border-plug-green/10">
				<LandingContainer className="flex flex-col items-start gap-2 py-4 font-bold text-plug-green/40 lg:flex-row lg:gap-8">
					<p>Copyright Terminally Online, Inc. 2024</p>
					<button
						className="transition-all duration-200 ease-in-out hover:text-plug-green lg:ml-auto"
						onClick={() => handleCallToAction(routes.comingSoon)}
					>
						Terms of Service
					</button>
					<button
						className="transition-all duration-200 ease-in-out hover:text-plug-green"
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
