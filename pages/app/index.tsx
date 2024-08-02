import { PageContent, PageHeader } from "@/components/app"
import { NextPageWithLayout } from "@/lib"

const Page: NextPageWithLayout = () => {
	return (
		<>
			<PageHeader />
			<PageContent />
		</>
	)
}

export default Page
