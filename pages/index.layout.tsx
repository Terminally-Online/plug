import type { FC, PropsWithChildren } from "react"

import Image from "next/image"
import Link from "next/link"

import { Activity, Book, Github, Twitter } from "lucide-react"

import logoWhite from "@/assets/logo-white.svg"
import EnterApp from "@/components/landing/enter"
import Marquee from "@/components/landing/marquee"
import Theme from "@/components/landing/theme"

const PageLayout: FC<PropsWithChildren> = ({ children }) => (
	<div className="flex h-full min-h-screen flex-col">
		<Marquee />

		<div className="lg:hidden">
			<EnterApp />
		</div>

		<div className="text-shadow-blur flex flex-col px-4 pb-4 shadow-white dark:shadow-black lg:px-8">
			<div className="mb-auto mt-4 flex flex-row items-center justify-center gap-4">
				<Link className="relative mr-12 text-xl" href="/">
					<div className="flex flex-row items-center gap-2">
						<Image
							src={logoWhite}
							alt="Plug"
							width={24}
							height={24}
							className="rounded-full bg-stone-900 p-1"
						/>
						PLUG
					</div>
					<span className="absolute left-full top-0 ml-2 h-min rounded-md bg-stone-900 p-1 text-[8px] font-bold leading-none text-white dark:bg-white dark:text-black">
						ALPHA
					</span>
				</Link>

				<a
					href="https://docs.onplug.io"
					target="_blank"
					rel="noopener noreferrer"
				>
					<Book
						className="text-black/60 transition-all duration-200 ease-in-out hover:text-black/80 dark:text-white/40 dark:hover:text-white/80"
						size={18}
					/>
				</a>

				<a
					href="https://status.onplug.io"
					target="_blank"
					rel="noopener noreferrer"
				>
					<Activity
						className="text-black/60 transition-all duration-200 ease-in-out hover:text-black/80 dark:text-white/40 dark:hover:text-white/80"
						size={18}
					/>
				</a>

				<Theme />

				<a
					href="https://twitter.com/onplug_io"
					target="_blank"
					rel="noopener noreferrer"
				>
					<Twitter
						className="text-black/60 transition-all duration-200 ease-in-out hover:text-black/80 dark:text-white/40 dark:hover:text-white/80"
						size={18}
					/>
				</a>

				<a
					href="https://github.com/nftchance/plug"
					target="_blank"
					rel="noopener noreferrer"
				>
					<Github
						className="text-black/60 transition-all duration-200 ease-in-out hover:text-black/80 dark:text-white/40 dark:hover:text-white/80"
						size={18}
					/>
				</a>
			</div>
		</div>

		{children}

		<div className="hidden lg:block">
			<EnterApp />
		</div>
	</div>
)

export default PageLayout
