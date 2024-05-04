import { Footer } from "@/components/base"
import Navbar from "@/components/base/navbar"
import { MainButton } from "@/components/buttons"
import { greenGradientStyle } from "@/lib/constants"

const Page = () => {
	return (
		<>
			<Navbar />

			<div className="flex min-h-[800px] flex-col items-center justify-center gap-2">
				<h1
					className="text-[72px] font-bold"
					style={{ ...greenGradientStyle }}
				>
					Coming Soon
				</h1>

				<p className="max-w-[320px] text-center text-black/65">
					You've found a section that is still under construction.
					Check back soon for updates.
				</p>

				<div className="mt-8 flex flex-row gap-2">
					<MainButton variant="secondary" text="Back Home" />
					<MainButton variant="primary" text="Get Early Access" />
				</div>
			</div>

			<Footer />
		</>
	)
}

export default Page
