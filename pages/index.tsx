import { Blocks, CallToAction, Demo, Hero, LandingFooter, Light, Transactions, Vision } from "@/components"

const Page = () => (
	<div className="overflow-x-hidden">
		<Hero />
		<Demo />
		<Light />
		<Transactions />
		<CallToAction
			text="Operate at a level your friends only dream of."
			description="Experience extraordinary results, elevate your onchain performance, leave the rest of the market behind, and change the scale of your goals."
			button="Get Started"
		/>
		<Vision />
		<Blocks />
		<LandingFooter />
	</div>
)

export default Page
