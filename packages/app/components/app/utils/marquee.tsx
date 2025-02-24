import { FC, HTMLAttributes, PropsWithChildren } from "react"
import { motion, Variants } from "framer-motion"
import { cn } from "@/lib";

interface MarqueeProps extends HTMLAttributes<HTMLDivElement>, PropsWithChildren {
    height?: string;
    duration?: number;
    strokeWidth?: number;
    strokeColor?: string;
    fontSize?: string;
}

const marqueeVariants: Variants = {
    animate: {
        x: ["0%", "-50%"],
        transition: {
            x: {
                repeat: Infinity,
                repeatType: "loop",
                duration: 15,
                ease: "linear",
            },
        },
    },
}

export const Marquee: FC<MarqueeProps> = ({ 
    children, 
    duration = 15,
    strokeWidth = 2,
    strokeColor = "#f4955c",
    fontSize = "text-md",
	className,
    ...props 
}) => {
    return (
        <div 
            className={cn(`relative w-full max-w-full`, className)} 
            {...props}
        >
            <motion.div
                className="absolute whitespace-nowrap"
                variants={marqueeVariants}
                animate="animate"
                style={{
                    transition: `all ${duration}s linear infinite`,
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
