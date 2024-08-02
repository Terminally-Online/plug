import {
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
		<div className="min-w-screen flex w-full flex-row overflow-visible">
			<ConsoleSidebar />
			<ConsoleColumnRow />
		</div>
	)
}

const Page = () => {
	const { isMobile } = useMediaQuery()

	return isMobile ? <MobilePage /> : <DesktopPage />
}

export default Page
