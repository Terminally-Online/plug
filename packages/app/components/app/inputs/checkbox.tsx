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
			"h-min w-min rounded-xs border-2",
			disabled === false ? "cursor-pointer border-plug-green" : "cursor-not-allowed border-plug-green/10",
			checked && "border-plug-yellow bg-plug-yellow text-plug-green"
		)}
		onClick={() => disabled === false && handleChange(!checked)}
	>
		<motion.div
			initial={{ background: "transparent" }}
			animate={{
				background: checked ? "#D2F38A" : "transparent" // Using plug-yellow color when checked
			}}
			transition={{ duration: 0.2, ease: "easeInOut" }}
		>
			<Check
				size={14}
				strokeWidth={3}
				className={cn(
					checked ? "text-[#385842]" : "text-transparent" // Using plug-green color for check
				)}
			/>
		</motion.div>
	</button>
)
