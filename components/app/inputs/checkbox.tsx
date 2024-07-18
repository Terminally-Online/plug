import type { FC } from "react"

import { motion } from "framer-motion"
import { Check } from "lucide-react"

import { cn } from "@/lib"

type Props = {
	checked: boolean
	handleChange: (checked: boolean) => void
	disabled?: boolean
}

export const Checkbox: FC<Props> = ({
	checked,
	handleChange,
	disabled = false
}) => (
	<button
		className={cn(
			"h-min w-min rounded-[6px] p-[2px]",
			disabled === false
				? "cursor-pointer bg-gradient-to-tr from-[#00E100] to-[#A3F700]"
				: "cursor-not-allowed bg-grayscale-100"
		)}
		onClick={() => disabled === false && handleChange(!checked)}
	>
		<motion.div
			className="rounded-[4px] p-[2px]"
			initial={{ background: "#FFFFFF " }}
			animate={{
				background: checked
					? "linear-gradient(to top right, #00E100, #A3F700)"
					: "linear-gradient(to top right, #FFFFFF, #FFFFFF)"
			}}
			transition={{ duration: 0.2, ease: "easeInOut" }}
		>
			<Check
				size={14}
				style={{
					color: checked ? "#FFFFFF" : "transparent"
				}}
			/>
		</motion.div>
	</button>
)
