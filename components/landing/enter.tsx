import Link from "next/link"

import { motion } from "framer-motion"
import { ArrowRight } from "lucide-react"

export default function Enter() {
	return (
		<Link href="/canvas">
			<motion.button
				type="submit"
				className="flex w-full flex-row items-center justify-center gap-4 bg-white p-4 px-8 text-black transition-all duration-200 ease-in-out hover:bg-stone-950 hover:text-white disabled:cursor-not-allowed disabled:opacity-50"
			>
				ENTER APP
				<ArrowRight className="opacity-40" size={18} />
			</motion.button>
		</Link>
	)
}
