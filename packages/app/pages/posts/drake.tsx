import { motion } from "framer-motion"
import { useState } from "react"

const Page = () => {
	const [isAnimating, setIsAnimating] = useState(false)

	const handleReveal = () => {
		console.log("Reveal button clicked");
		setIsAnimating(true); // Trigger the animation
	};

	const handleReset = () => {
		console.log("Reset button clicked");
		setIsAnimating(false); // Reset the animation
	};

	return (
		<div className="absolute inset-0">
			<h1 className="text-6xl font-bold">
				Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello
				World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World
				Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello
				World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World
				Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello
				World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World
				Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello
				World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World
				Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello
				World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World
				Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello
				World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World
				Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello
				World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World
				Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello
				World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World
				Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello
				World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World
				Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello World Hello
				World Hello World Hello World Hello World Hello World Hello World
			</h1>
			{/* 
			<motion.div
				className="absolute inset-0 bg-plug-white"
				initial={{ filter: "blur(0px)" }}
				animate={{ filter: ["blur(0px)", "blur(100px)"] }}
				transition={{ duration: 1 }}
			/> */}

			<motion.div
				className="absolute inset-0 z-[999] overflow-hidden"
				animate={{
					background: [
						"radial-gradient(circle at center, transparent 0%, rgba(0,0,0,0.95) 2%, black 5%)",
						"radial-gradient(circle at center, transparent 100%, rgba(0,0,0,0.95) 98%, black 100%)"
					],
					scale: [1, 1.2],
					filter: ["perspective(400px) distort(0deg)", "perspective(400px) distort(20deg)"]
				}}
				style={{
					transformOrigin: "center center",
					backfaceVisibility: "hidden"
				}}
				transition={{ duration: 3, ease: "easeInOut" }}
			/>

			{/* Buttons for Reveal and Reset */}
			<div className="absolute bottom-8 left-1/2 -translate-x-1/2 z-10 flex space-x-4">
				<button
					onClick={handleReveal}
					className="px-6 py-2 bg-black text-white rounded-full"
				>
					Reveal
				</button>
				<button
					onClick={handleReset}
					className="px-6 py-2 bg-red-500 text-white rounded-full"
				>
					Reset
				</button>
			</div>
		</div>
	)
}

export default Page
