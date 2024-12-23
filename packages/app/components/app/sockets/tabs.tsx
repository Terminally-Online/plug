import { createRef, FC, useEffect, useRef, useState } from "react"

import { motion } from "framer-motion"
import { Activity, Landmark, PiggyBank } from "lucide-react"

import { cn, formatTitle } from "@/lib"

const tabs = [
	{
		name: "Activity",
		icon: <Activity size={14} />
	},
	{
		name: "Assets",
		icon: <PiggyBank size={14} />
	},
	{
		name: "Positions",
		icon: <Landmark size={14} />
	}
]

export const SocketTabs: FC<{
	selected: number
	onSelect: (index: number) => void
}> = ({ selected, onSelect }) => {
	const textRefs = useRef(tabs.map(() => createRef<HTMLButtonElement>()))
	const [underlineStyle, setUnderlineStyle] = useState({ width: 0, x: 0 })

	useEffect(() => {
		const currentRef = textRefs.current[selected].current

		if (currentRef) {
			const { offsetWidth, offsetLeft } = currentRef

			setUnderlineStyle({
				width: offsetWidth,
				x: offsetLeft
			})
		}
	}, [selected])

	return (
		<div className="border-b-[1px] border-plug-green/10 px-4">
			<ul className="relative flex list-none gap-4 pb-2">
				{tabs.map((tab, index) => (
					<button
						key={tab.name}
						ref={textRefs.current[index]}
						onClick={() => onSelect(index)}
						className={cn(
							"flex cursor-pointer flex-row items-center gap-2 border-none bg-transparent font-bold outline-none transition-all duration-200 ease-in-out",
							selected === index ? "opacity-100" : "opacity-40 hover:opacity-80",
							index === 0 && "mr-auto"
						)}
					>
						<span className="opacity-40">{tab.icon}</span>
						{formatTitle(tab.name)}
					</button>
				))}

				<motion.div
					layout
					initial={false}
					animate={underlineStyle}
					transition={{ type: "spring", stiffness: 300, damping: 30 }}
					className="absolute bottom-[-1px] h-[1px] bg-black"
				/>
			</ul>
		</div>
	)
}
