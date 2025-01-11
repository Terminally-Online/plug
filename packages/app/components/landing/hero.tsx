import { motion } from "framer-motion"

import { LandingContainer } from "@/components/landing/layout/container"
import { Button } from "@/components/shared/buttons/button"
import { env } from "@/env"
import { GTM_EVENTS, routes, useAnalytics } from "@/lib"

export const Hero = () => {
	const handleCallToAction = useAnalytics(
		GTM_EVENTS.CTA_CLICKED,
		env.NEXT_PUBLIC_EARLY_ACCESS ? routes.earlyAccess : routes.app
	)

	return (
		<div className="relative z-[2] flex h-full w-screen">
			<div className="z-2 relative w-full">
				<LandingContainer className="flex h-full flex-col py-8 text-plug-green">
					<div className="relative my-auto flex items-center pb-12">
						<div className="flex flex-col gap-4 md:my-auto xl:gap-8">
							<motion.h1
								className="text-[48px] font-black leading-tight md:text-[72px] lg:max-w-[840px] lg:text-[82px] xl:max-w-[1240px] xl:text-[96px]"
								initial={{ transform: "translateY(-20px)", opacity: 0 }}
								whileInView={{
									transform: ["translateY(-20px)", "translateY(0px)"],
									opacity: [0, 1]
								}}
								transition={{ duration: 0.3 }}
							>
								Your all-in-one app for onchain activity.
							</motion.h1>

							<motion.p
								className="max-w-[480px] text-xl font-bold text-plug-green/40 md:max-w-[520px] lg:max-w-[620px] lg:text-[24px] xl:max-w-[620px] 2xl:max-w-[720px]"
								initial={{ transform: "translateY(20px)", opacity: 0 }}
								whileInView={{
									transform: ["translateY(20px)", "translateY(0px)"],
									opacity: [0, 1]
								}}
								transition={{
									duration: 0.3,
									delay: 0.15
								}}
							>
								Plug gives you a single interface to manage, compose, schedule, and execute all your
								onchain activity in one place.
							</motion.p>

							<Button
								variant="none"
								className="mt-8 w-max rounded-md border-[1px] border-plug-yellow/30 bg-plug-yellow px-8 py-3 text-center font-black text-plug-green filter backdrop-blur-xl transition-all duration-200 ease-in-out hover:bg-plug-yellow/50"
								onClick={() => handleCallToAction()}
							>
								Enter App
							</Button>
						</div>
					</div>
				</LandingContainer>
			</div>
		</div>
	)
}
