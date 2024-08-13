import {
	AuthFrame,
	ConsoleColumnRow,
	ConsoleSidebar,
	PageContent,
	PageHeader
} from "@/components"
import { useSockets } from "@/contexts"
import { useMediaQuery } from "@/lib"

const MobilePage = () => {
	const { page } = useSockets()

	if (!page) return null

	return (
		<>
			<PageHeader />
			<PageContent />
			<AuthFrame id={page.id} />
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

	return <>{md ? <DesktopPage /> : <MobilePage />}</>
}

export default Page
