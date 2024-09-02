import { Blocks, CallToAction, Demo, Hero, LandingFooter, Light, Transactions, Vision } from "@/components"

const Page = () => (
	<div className="overflow-x-hidden">
		<Hero />
		<Demo />
		<Light />
		<Transactions />
		<CallToAction
			text="Operate at a level your friends only dream of."
			description="With best in class execution and access to top strategies your capital can work harder than ever before. Compose top protocols and automate your portfolio's growth."
			button="Get Started"
		/>
		<Vision />
		<Blocks />
		<LandingFooter />
	</div>
)

export default Page
