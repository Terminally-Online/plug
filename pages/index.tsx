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
import { StaticLayout } from "@/components/layouts/static"
import { NextPageWithLayout } from "@/lib/types"

const Page: NextPageWithLayout = () => (
	<>
		<Hero />
		<Steps />
		<Examples />
		<Value />
		<Templates />
		<CallToAction
			text="Never miss another opportunity."
			description="Instead of being trapped inside waiting to click the buttons you can go outside and live the life you want without worrying about missing every opportunity. Your capital can finally manage itself."
			button="Get Early Access"
		/>
		<Vision />
		<Letter />
		<FrequentlyAskedQuestions />
		<CallToAction
			text="Level up your onchain activity."
			description="As soon as your conditions are met and expected outcomes can be delivered your transactions will automatically run without you doing a thing."
			button="Get Started"
		/>
	</>
)

Page.getLayout = page => <StaticLayout title="Home">{page}</StaticLayout>

export default Page
