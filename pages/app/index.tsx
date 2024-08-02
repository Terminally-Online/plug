import { FC } from "react"

import Image from "next/image"

import BlockiesSvg from "blockies-react-svg"
import { ChevronLeft, Plus } from "lucide-react"

import {
	AccountFrame,
	Button,
	Container,
	Header,
	PlugsDiscover,
	SocketActivity,
	SocketAssets
} from "@/components"
import { Plugs } from "@/components/shared/framework/plugs"
import { useFrame, usePage, usePlugs, useSockets } from "@/contexts"
import { cn, formatTitle, NextPageWithLayout, routes } from "@/lib"

const PageHeader: FC<{ page: string }> = ({ page }) => {
	const { handlePage } = usePage()
	const { handleFrameVisible } = useFrame()
	const { address, ensAvatar } = useSockets()

	const isHomePage = page === "home" || page === "activity"

	return (
		<>
			{isHomePage ? (
				<>
					{address ? (
						<button
							className="flex flex-row items-center gap-2"
							onClick={() => handleFrameVisible("account")}
						>
							{ensAvatar ? (
								<Image
									src={ensAvatar}
									alt="ENS Avatar"
									width={24}
									height={24}
									className="h-6 w-6 rounded-sm"
								/>
							) : (
								<BlockiesSvg
									className="h-6 w-6 rounded-sm"
									address={address}
								/>
							)}
						</button>
					) : (
						<div
							className="flex h-6 w-6 flex-row items-center justify-center rounded-sm"
							style={{
								backgroundImage:
									"linear-gradient(30deg, #00E100, #A3F700)"
							}}
						>
							<Image
								src="/white-icon.svg"
								alt="Logo"
								width={662}
								height={616}
								className="h-3 w-auto"
							/>
						</div>
					)}

					<button
						className={cn(
							"text-lg font-bold transition-all duration-200 ease-in-out",
							page !== "home"
								? "opacity-40 hover:opacity-100"
								: ""
						)}
						onClick={() => handlePage("home")}
					>
						Home
					</button>

					<button
						className={cn(
							"mr-auto text-lg font-bold transition-all duration-200 ease-in-out",
							page !== "activity"
								? "opacity-40 hover:opacity-100"
								: ""
						)}
						onClick={() => handlePage("activity")}
					>
						Activity
					</button>
				</>
			) : (
				<>
					<Button
						variant="secondary"
						className="rounded-sm p-1"
						onClick={() => handlePage("home")}
					>
						<ChevronLeft size={14} />
					</Button>

					<button
						className={
							"mr-auto text-lg font-bold transition-all duration-200 ease-in-out"
						}
						onClick={() => handlePage("activity")}
					>
						{formatTitle(page)}
					</button>
				</>
			)}
		</>
	)
}

const PageContent: FC<{ page: string }> = ({ page }) => {
	switch (page) {
		case "home":
			return (
				<Container>
					<Plugs hideEmpty={true} />
					<SocketAssets />
				</Container>
			)
		case "activity":
			return (
				<Container>
					<SocketActivity />
				</Container>
			)
		case "discover":
			return <PlugsDiscover />
		case
		// case "create":
		// 	return <PlugsColumn />
		// case "plug":
		// 	return <PlugsColumn />
		default:
			return <></>
	}
}

const Page: NextPageWithLayout = () => {
	const { page } = usePage()
	const { handle } = usePlugs()

	return (
		<>
			<Container>
				<Header
					size="lg"
					label={<PageHeader page={page} />}
					nextOnClick={() => handle.plug.add(routes.app.index)}
					nextLabel={<Plus size={14} />}
				/>
			</Container>

			<PageContent page={page} />

			<AccountFrame />
		</>
	)
}

export default Page
