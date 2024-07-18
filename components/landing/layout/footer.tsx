import { FC } from "react"

import Image from "next/image"
import Link from "next/link"

import { LandingContainer } from "@/components"
import { routes } from "@/lib"

export const LandingFooter: FC = () => (
	<>
		<LandingContainer className="mb-16 mt-24 flex-col gap-2 lg:gap-4">
			<Image src="/black-logo.svg" alt="Logo" width={96} height={64} />

			<div className="grid grid-cols-1 lg:grid-cols-12">
				<div className="mb-8 lg:col-span-4 lg:mb-0">
					<p className="text-black/65 lg:max-w-[320px]">
						Automate your transactions on every popular Ethereum
						based blockchain so that you can log off and have
						everything run without you doing a thing.
					</p>
				</div>
				<div className="mb-2 flex flex-col gap-2 lg:col-span-2">
					<a
						href={routes.comingSoon}
						className="opacity-60 transition-opacity duration-200 hover:opacity-100"
					>
						Smart Contracts
					</a>
					<a
						href={routes.github}
						className="opacity-60 transition-opacity duration-200 hover:opacity-100"
					>
						Code
					</a>
					<a
						href={routes.comingSoon}
						className="opacity-60 transition-opacity duration-200 hover:opacity-100"
					>
						Audits
					</a>
				</div>
				<div className="mb-2 flex flex-col gap-2 lg:col-span-2">
					<a
						href={routes.documentation}
						className="opacity-60 transition-opacity duration-200 hover:opacity-100"
					>
						Documentation
					</a>
					<a
						href={routes.comingSoon}
						className="opacity-60 transition-opacity duration-200 hover:opacity-100"
					>
						Guides
					</a>
				</div>
				<div className="mb-2 flex flex-col gap-2 lg:col-span-2">
					<a
						href={routes.comingSoon}
						className="opacity-60 transition-opacity duration-200 hover:opacity-100"
					>
						Brand Kit
					</a>
					<a
						href={routes.comingSoon}
						className="opacity-60 transition-opacity duration-200 hover:opacity-100"
					>
						Investor Memo
					</a>
				</div>
				<div className="mb-2 flex flex-col gap-2 lg:col-span-2 lg:text-right">
					<a
						href={routes.comingSoon}
						className="opacity-60 transition-opacity duration-200 hover:opacity-100"
					>
						hello@onplug.io
					</a>
					<a
						href={routes.twitter}
						target="_blank"
						rel="noreferrer"
						className="opacity-60 transition-opacity duration-200 hover:opacity-100"
					>
						Twitter
					</a>
					<Link
						href={routes.comingSoon}
						className="opacity-60 transition-opacity duration-200 hover:opacity-100"
					>
						Telegram
					</Link>
				</div>
			</div>
		</LandingContainer>

		<div className="bg-[#d9d9d9]/10">
			<LandingContainer className="flex flex-col gap-2 py-4 lg:flex-row lg:gap-8">
				<p>Copyright Terminally Online, Inc. 2024</p>
				<Link
					href="/coming-soon/"
					className="opacity-60 transition-opacity duration-200 hover:opacity-100 lg:ml-auto"
				>
					Terms of Service
				</Link>
				<Link
					href="/coming-soon/"
					className="opacity-60 transition-opacity duration-200 hover:opacity-100"
				>
					Privacy Policy
				</Link>
			</LandingContainer>
		</div>
	</>
)

export default LandingFooter
