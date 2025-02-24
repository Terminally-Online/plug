import { FC, HTMLAttributes, PropsWithChildren } from "react"
import { motion, Variants } from "framer-motion"
import { cn } from "@/lib";

interface MarqueeProps extends HTMLAttributes<HTMLDivElement>, PropsWithChildren {
	height?: string;
	duration?: number;
	fontSize?: string;
}

const marqueeVariants = (duration: number): Variants => ({
	animate: {
		x: ["0%", "-50%"],
		transition: {
			x: {
				repeat: Infinity,
				repeatType: "loop",
				duration,
				ease: "linear",
			},
		},
	},
})

const calculateDuration = (children: React.ReactNode): number => {
	if (typeof children === 'string') {
		return Math.max(5, Math.min(80, children.length * 1.2));
	}
	return 15;
}

export const Marquee: FC<MarqueeProps> = ({
	children,
	duration: userDuration,
	fontSize = "text-md",
	className,
	...props
}) => {
	const calculatedDuration = userDuration ?? calculateDuration(children);
	
	return (
		<div
			className={cn(`relative w-full max-w-full`, className)}
			{...props}
		>
			<motion.div
				className="absolute whitespace-nowrap"
				variants={marqueeVariants(calculatedDuration)}
				animate="animate"
				style={{
					transition: `all ${calculatedDuration}s linear infinite`,
				}}
			>
				<h1
					className={`${fontSize}`}
				>
					{children} {children} {children} {children} {children} {children}
				</h1>
			</motion.div>
		</div>
	)
}
