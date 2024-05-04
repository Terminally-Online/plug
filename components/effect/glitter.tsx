import type { FC, PropsWithChildren } from "react"
import { useEffect, useRef, useState } from "react"

import { motion } from "framer-motion"

const getRandomMaxOpacity = (min = 0.2, max = 0.4) => {
	return Math.random() * (max - min) + min
}

const getRandomDuration = (min = 0.1, max = 0.2) => {
	return Math.random() * (max - min) + min
}

export const Glitter: FC<PropsWithChildren> = ({ children }) => {
	const ref = useRef<HTMLSpanElement>(null)

	const [dots, setDots] = useState(0)

	useEffect(() => {
		if (ref.current) {
			const width = ref.current.offsetWidth
			setDots(Math.floor(width * 3))
		}
	}, [ref])

	return (
		<span className="relative" ref={ref}>
			{children}
			<motion.span
				className="absolute left-0 right-0 top-[15%] hidden flex-wrap lg:flex"
				initial={{ opacity: 0 }}
				whileInView={{ opacity: 1 }}
				transition={{ duration: 1 }}
			>
				{Array.from({ length: dots }).map((_, i) => (
					<motion.div
						key={i}
						className="ml-[-2px] mt-[-2px]"
						style={{
							width: 8,
							height: 8,
							borderRadius: "50%",
							backgroundColor: "white"
						}}
						initial={{ opacity: 0 }}
						animate={{
							opacity: [0, getRandomMaxOpacity()]
						}}
						transition={{
							duration: getRandomDuration(),
							repeat: Infinity,
							repeatType: "reverse"
						}}
					/>
				))}
			</motion.span>
		</span>
	)
}

export default Glitter
