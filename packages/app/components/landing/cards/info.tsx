import Link from "next/link"
import { useRouter } from "next/router"
import { FC, HTMLAttributes, PropsWithChildren, ReactNode } from "react"

import { motion, MotionProps } from "framer-motion"
import { ArrowRight } from "lucide-react"

import { cn } from "@/lib/utils"
import Image from "next/image"

export const InfoCard: FC<
	HTMLAttributes<HTMLDivElement> &
	MotionProps &
	PropsWithChildren<{
		icon?: ReactNode
		text: string | React.ReactNode
		description?: string
		author?: string
		href?: string
	}>
> = ({ children, icon, text, description, author, href, className, ...props }) => {
	const router = useRouter()

	const handleClick = () => {
		if (href) {
			router.push(href)
		}
	}

	return (
		<motion.div
			className={cn(
				"relative flex min-h-[240px] flex-row items-end gap-8 rounded-xl border-[1px] border-plug-green/10 bg-plug-white p-8 text-plug-green",
				className
			)}
			initial={{ transform: "translateY(20px)", opacity: 0 }}
			whileInView={{
				transform: ["translateY(20px)", "translateY(0px)"],
				opacity: [0, 1]
			}}
			transition={{ duration: 0.3 }}
			onClick={handleClick}
			{...props}
		>
			<div className="absolute bottom-0 left-0 right-0 top-0 overflow-hidden rounded-xl">{children}</div>

			<div className="flex-rows flex items-center gap-4">
				{icon && <div className="mb-auto mt-1">{icon}</div>}
				<div className="z-[10] flex flex-col gap-2 font-bold">
					<h2 className="flex items-center text-lg lg:text-2xl">{text}</h2>
					{description && <p className="max-w-[480px] text-plug-green/40">{description}</p>}
					<div className="flex flex-row justify-between items-center">
						{href && (
							<Link href={href} className="mt-2 flex flex-row items-center gap-2 text-plug-green">
								Read More
								<span className="opacity-40 transition-opacity duration-300 group-hover:opacity-100">
									<ArrowRight size={14} />
								</span>
							</Link>
						)}
						{author && <div className="flex flex-row gap-2 items-center">
							<Image
								src={`/users/${author}.png`}
								alt={author ?? ""}
								width={24}
								height={24}
								className="rounded-full w-6 h-6"
							/>

							<p>{author}</p>
						</div>}
					</div>
				</div>
			</div>
		</motion.div>
	)
}

export default InfoCard
