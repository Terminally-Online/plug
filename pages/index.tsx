import { CallToAction, Demo, Hero, LandingFooter, Light, Vision } from "@/components"

const Page = () => (
	<div className="overflow-x-hidden">
		<Hero />
		<Demo />
		<Light />
		<Vision />
		<CallToAction
			text="Level up your onchain activity."
			description="As soon as all of the constraints on your intent are met and expected outcomes can be delivered your transaction will automatically run without you doing a thing."
			button="Get Started"
		/>
		<LandingFooter />
	</div>
)

export default Page
