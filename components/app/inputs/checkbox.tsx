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
			"h-min w-min rounded-[6px] p-[2px]",
			disabled === false
				? "cursor-pointer bg-gradient-to-tr from-plug-green to-plug-yellow"
				: "cursor-not-allowed bg-grayscale-100"
		)}
		onClick={() => disabled === false && handleChange(!checked)}
	>
		<motion.div
			className="rounded-[4px]"
			initial={{ background: "#FFFFFF " }}
			animate={{
				background: checked
					? "linear-gradient(to top right, #385842, #D2F38A)"
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
