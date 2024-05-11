import Head from "next/head"

import { Footer, Navbar } from "@/components/base"
import {
	CallToAction,
	Examples,
	FrequentlyAskedQuestions,
	Hero,
	Letter,
	Steps,
	Templates,
	Value,
	Vision
} from "@/components/landing"

const Page = () => (
	<>
		<Head>
			<title>Home | Plug</title>
		</Head>
		<Navbar />
		<Hero />
		<Steps />
		<Examples />
		<Value />
		<CallToAction
			text="Get what you want from every transaction."
			description="Simultaneous settlement ensures transactions only execute when the conditions and expected outcomes you set in your intent can be met. No fees are paid and tokens move unless everything happens as expected."
			button="Get Started"
		/>
		<Templates />
		<Vision />
		<CallToAction
			text="Blockchains were not built for humans."
			description="Instead of being trapped inside waiting to click the buttons you can go outside and live the life you want without worrying about missing every opportunity. Your capital can finally manage itself."
			button="Get Early Access"
		/>
		<Letter />
		<FrequentlyAskedQuestions />
		<Footer />
	</>
)

export default Page
