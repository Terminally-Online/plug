import { FC } from "react"

import Image from "next/image"
import Link from "next/link"

import { LandingContainer } from "@/components"
import { routes } from "@/lib"

export const LandingFooter: FC = () => (
	<div className="relative z-[12] bg-white pt-16 lg:gap-4">
		<div className="absolute top-[-4px] h-[8px] w-full bg-gradient-to-r from-plug-green to-plug-yellow blur-xl filter" />
		<div className="absolute top-0 h-[2px] w-full bg-gradient-to-r from-plug-green to-plug-yellow" />

		<LandingContainer className="mb-16 flex-col gap-2">
			<Image src="/black-logo.svg" alt="Logo" width={96} height={64} />

			<div className="grid grid-cols-1 font-bold text-black/40 lg:grid-cols-12">
				<div className="mb-8 lg:col-span-4 lg:mb-0">
					<p className="lg:max-w-[320px]">
						Automate your transactions on every popular Ethereum based blockchain so that you can log off
						and have everything run without you doing a thing.
					</p>
				</div>
				<div className="mb-2 flex flex-col gap-2 lg:col-span-2">
					<a href={routes.comingSoon} className="transition-all duration-200 ease-in-out hover:text-black">
						Smart Contracts
					</a>
					<a href={routes.github} className="transition-all duration-200 ease-in-out hover:text-black">
						Code
					</a>
					<a href={routes.comingSoon} className="transition-all duration-200 ease-in-out hover:text-black">
						Audits
					</a>
				</div>
				<div className="mb-2 flex flex-col gap-2 lg:col-span-2">
					<a href={routes.documentation} className="transition-all duration-200 ease-in-out hover:text-black">
						Documentation
					</a>
					<a href={routes.comingSoon} className="transition-all duration-200 ease-in-out hover:text-black">
						Guides
					</a>
				</div>
				<div className="mb-2 flex flex-col gap-2 lg:col-span-2">
					<a href={routes.comingSoon} className="transition-all duration-200 ease-in-out hover:text-black">
						Brand Kit
					</a>
					<a href={routes.comingSoon} className="transition-all duration-200 ease-in-out hover:text-black">
						Investor Memo
					</a>
				</div>
				<div className="mb-2 flex flex-col gap-2 lg:col-span-2 lg:text-right">
					<a href={routes.comingSoon} className="transition-all duration-200 ease-in-out hover:text-black">
						hello@onplug.io
					</a>
					<a
						href={routes.twitter}
						target="_blank"
						rel="noreferrer"
						className="transition-all duration-200 ease-in-out hover:text-black"
					>
						Twitter
					</a>
					<Link href={routes.comingSoon} className="transition-all duration-200 ease-in-out hover:text-black">
						Telegram
					</Link>
				</div>
			</div>
		</LandingContainer>

		<div className="border-t-[1px] border-grayscale-100">
			<LandingContainer className="flex flex-col gap-2 py-4 font-bold text-black/40 lg:flex-row lg:gap-8">
				<p>Copyright Terminally Online, Inc. 2024</p>
				<Link
					href="/coming-soon/"
					className="transition-all duration-200 ease-in-out hover:text-black lg:ml-auto"
				>
					Terms of Service
				</Link>
				<Link href="/coming-soon/" className="transition-all duration-200 ease-in-out hover:text-black">
					Privacy Policy
				</Link>
			</LandingContainer>
		</div>
	</div>
)

export default LandingFooter
