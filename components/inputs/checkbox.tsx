import type { FC } from "react"

import { motion } from "framer-motion"
import { Check } from "lucide-react"

type Props = { checked: boolean; handleChange: (checked: boolean) => void }

export const Checkbox: FC<Props> = ({ checked, handleChange }) => (
	<button
		className="h-min w-min cursor-pointer rounded-[6px] bg-gradient-to-tr from-[#00E100] to-[#A3F700] p-[2px]"
		onClick={() => handleChange(!checked)}
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
