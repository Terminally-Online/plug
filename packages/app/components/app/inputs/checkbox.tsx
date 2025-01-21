import type { FC } from "react"

import { motion } from "framer-motion"
import { Check } from "lucide-react"

import { cn } from "@/lib"

type Props = {
	checked: boolean
	handleChange: (checked: boolean) => void
	disabled?: boolean
}

export const Checkbox: FC<Props> = ({ checked, handleChange, disabled = false }) => (
	<button
		className={cn(
			"h-min w-min rounded-[4px] border-2",
			disabled === false
				? "cursor-pointer border-plug-green"
				: "cursor-not-allowed border-plug-green/10"
		)}
		onClick={() => disabled === false && handleChange(!checked)}
	>
		<motion.div
			className="rounded-4"
			initial={{ background: "transparent" }}
			animate={{
				background: checked ? "#385842" : "transparent"
			}}
			transition={{ duration: 0.2, ease: "easeInOut" }}
		>
			<Check
				size={14}
				strokeWidth={3}
				className={cn(
					checked ? "text-plug-yellow" : "text-transparent"
				)}
			/>
		</motion.div>
	</button>
)
