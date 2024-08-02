import { ConsoleColumnRow, ConsoleSidebar } from "@/components"

const Page = () => {
	return (
		<div className="min-w-screen flex w-full flex-row overflow-visible">
			<ConsoleSidebar />
			<ConsoleColumnRow />
		</div>
	)
}

export default Page
