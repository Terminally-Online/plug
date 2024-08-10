import { useEffect, useState } from "react"

import {
	CallToAction,
	Examples,
	FrequentlyAskedQuestions,
	Hero,
	LandingFooter,
	Templates,
	Value,
	Vision
} from "@/components"

const Page = () => {
	const [expanded, setExpanded] = useState(false)

	const handleScroll = () => {
		const element = document.getElementById("below-the-fold")
		if (element) element.scrollIntoView({ behavior: "smooth" })
	}

	const handleExpand = () => {
		setExpanded(true)
		handleScroll()
	}

	useEffect(() => {
		handleScroll()
	}, [expanded])

	return (
		<>
			<Hero handleExpand={handleExpand} />

			{expanded && (
				<div id="below-the-fold">
					<Examples />
					<Value />
					<Templates />
					<CallToAction
						text="Never miss another opportunity."
						description="Instead of being trapped inside waiting to click the buttons you can go outside and live the life you want without worrying about missing every opportunity. Your capital can finally manage itself."
						button="Get Early Access"
					/>
					<Vision />
					<FrequentlyAskedQuestions />
					<CallToAction
						text="Level up your onchain activity."
						description="As soon as all of the constraints on your intent are met and expected outcomes can be delivered your transaction will automatically run without you doing a thing."
						button="Get Started"
					/>
					<LandingFooter />
				</div>
			)}
		</>
	)
}

export default Page
