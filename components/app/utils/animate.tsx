import { FC, PropsWithChildren } from "react"

import { motion } from "framer-motion"

const Base = () => {}

const List: FC<PropsWithChildren> = ({ children }) => (
	<motion.div
		className="flex flex-col gap-2"
		initial="hidden"
		animate="visible"
		variants={{
			hidden: { opacity: 0 },
			visible: {
				opacity: 1,
				transition: {
					staggerChildren: 0.05
				}
			}
		}}
	>
		{children}
	</motion.div>
)

const ListItem: FC<PropsWithChildren> = ({ children }) => (
	<motion.div
		variants={{
			hidden: { opacity: 0, y: 10 },
			visible: {
				opacity: 1,
				y: 0,
				transition: {
					type: "spring",
					stiffness: 100,
					damping: 10
				}
			}
		}}
	>
		{children}
	</motion.div>
)

export const Animate = Object.assign(Base, { List, ListItem })
