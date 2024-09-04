import { Blocks, CallToAction, Demo, Hero, LandingFooter, Light, Transactions, Vision } from "@/components"

const Page = () => (
	<div className="overflow-x-hidden">
		<Hero />
		<Demo />
		<Light />
		<Transactions />
		<CallToAction
			text="Operate at a level your friends only dream of."
			description="If you can’t beat them, join them. Embrace the future of onchain activity and 
			supercharge your potential. So easy you’ll be able to teach your friends."
			button="Get Started"
		/>
		<Vision />
		<Blocks />
		<LandingFooter />
	</div>
)

export default Page
