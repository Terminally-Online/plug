import { FC, useEffect, useRef, useState } from "react"

import {
	animate,
	motion,
	MotionProps,
	useMotionValue,
	useTransform
} from "framer-motion"
import { GitFork, Zap } from "lucide-react"

import { cn, colors, tagColors } from "@/lib"

type Props = {
	size?: "md" | "lg"
	color?: keyof typeof colors
	glow?: boolean
	title: string
	forks?: number
	runs?: number
	className?: string
} & MotionProps

const sizes: Record<NonNullable<Props["size"]>, string> = {
	md: "text-lg font-bold",
	lg: "text-md lg:text-xl font-bold min-h-[140px] lg:min-h-[200px]"
}

export const LandingActionCard: FC<Props> = ({
	size = "md",
	color = "blue",
	glow = false,
	title,
	forks,
	runs,
	className,
	...props
}) => {
	const base = `rounded-lg p-4 text-white text-left flex flex-col justify-end`

	const ref = useRef<HTMLButtonElement>(null)

	const [isInView, setInView] = useState(false)

	const animatedForks = useMotionValue(0)
	const forksRounded = useTransform(animatedForks, latest =>
		Math.round(latest)
	)
	const animatedRuns = useMotionValue(0)
	const runsRounded = useTransform(animatedRuns, latest => Math.round(latest))

	useEffect(() => {
		const node = ref.current
		if (!node) return

		const observer = new IntersectionObserver(
			entries => {
				entries.forEach(entry => {
					if (entry.isIntersecting) {
						setInView(true)
					}
				})
			},
			{ threshold: 0.1 }
		)

		observer.observe(node)

		return () => {
			observer.unobserve(node)
		}
	}, [])

	useEffect(() => {
		if (!isInView) return
		const controls = animate(
			animatedForks,
			forks || Math.floor(Math.random() * 300),
			{
				duration: 1,
				delay: 0.5
			}
		)
		return controls.stop
	}, [animatedForks, forks, isInView])

	useEffect(() => {
		if (!isInView) return
		const controls = animate(
			animatedRuns,
			runs || Math.floor(Math.random() * 1200) + 200,
			{
				duration: 1,
				delay: 0.5
			}
		)
		return controls.stop
	}, [animatedRuns, runs, isInView])

	return (
		<motion.button
			ref={ref}
			className={cn(base, sizes[size], className)}
			style={{
				backgroundColor: colors[color],
				boxShadow: glow ? `0 0 20px ${colors[color]}` : "none"
			}}
			{...props}
		>
			<div className="flex flex-row gap-2 text-xs">
				<div
					className="flex flex-row items-center gap-1 rounded-full px-2 py-1"
					style={{ backgroundColor: tagColors[color] }}
				>
					<GitFork size={16} className="opacity-40" />
					<motion.span className="tabular-nums">
						{forksRounded}
					</motion.span>{" "}
					Forks
				</div>
				<div
					className="flex flex-row items-center gap-1 rounded-full px-2 py-1"
					style={{ backgroundColor: tagColors[color] }}
				>
					<Zap size={16} className="opacity-40" />
					<motion.span className="tabular-nums">
						{runsRounded}
					</motion.span>{" "}
					Runs
				</div>
			</div>
			<span className="mt-auto w-[65%] 2xl:w-[90%]">{title}</span>
		</motion.button>
	)
}

export default LandingActionCard
