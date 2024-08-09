import {
	AuthFrame,
	ConsoleColumnRow,
	ConsoleSidebar,
	PageContent,
	PageHeader
} from "@/components"
import { useMediaQuery } from "@/lib"

const MobilePage = () => {
	return (
		<>
			<PageHeader />
			<PageContent />
		</>
	)
}

const DesktopPage = () => {
	return (
		<div className="min-w-screen flex h-screen w-full flex-row overflow-y-hidden overflow-x-visible">
			<ConsoleSidebar />
			<ConsoleColumnRow />
		</div>
	)
}

const Page = () => {
	const { md } = useMediaQuery()

	return (
		<>
			{md ? <DesktopPage /> : <MobilePage />}

			<AuthFrame id="global" />
		</>
	)
}

export default Page
