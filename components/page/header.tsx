import { FC } from "react"

import Image from "next/image"

import { useSession } from "next-auth/react"

import BlockiesSvg from "blockies-react-svg"
import { ChevronLeft, Ellipsis, GitFork, Plus, Share } from "lucide-react"

import { ActionView, Button, Container, Header } from "@/components"
import { useFrame, usePage, usePlugs, useSockets } from "@/contexts"
import {
	cardColors,
	cn,
	formatAddress,
	formatTimeSince,
	formatTitle
} from "@/lib"

const HomePageHeader = () => {
	const { page, handlePage } = usePage()
	const { handleFrame } = useFrame({ id: "global" })
	const { address, ensAvatar } = useSockets()
	const { handle } = usePlugs()

	return (
		<Header
			size="lg"
			label={
				<>
					{address ? (
						<button
							className="flex flex-row items-center gap-2"
							onClick={() => handleFrame("auth")}
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
							page.key !== "home"
								? "opacity-40 hover:opacity-100"
								: ""
						)}
						onClick={() => handlePage({ key: "home" })}
					>
						Home
					</button>

					<button
						className={cn(
							"mr-auto text-lg font-bold transition-all duration-200 ease-in-out",
							page.key !== "activity"
								? "opacity-40 hover:opacity-100"
								: ""
						)}
						onClick={() => handlePage({ key: "activity" })}
					>
						Activity
					</button>
				</>
			}
			nextOnClick={() => handle.plug.add("home")}
			nextLabel={<Plus size={14} />}
		/>
	)
}

const PlugHeader = () => {
	const { data: session } = useSession()
	const { page } = usePage()
	const { handleFrame } = useFrame({ id: "global" })
	const { plug, handle } = usePlugs(page.id)

	const own =
		plug !== undefined && session && session.address === plug.userAddress

	if (!plug) return null

	return (
		<div className="flex min-h-[calc(100vh-80px)] flex-col">
			<Header
				size="lg"
				back={page.from ?? "mine"}
				icon={
					<div
						className="h-6 w-6 min-w-6 rounded-md"
						style={{
							backgroundImage: cardColors[plug.color]
						}}
					/>
				}
				label={plug.name === "" ? "Untitled Plug" : plug.name}
				nextOnClick={own ? () => handleFrame("manage") : () => {}}
				nextLabel={
					own ? (
						<Ellipsis size={14} />
					) : (
						<div className="flex flex-row items-center gap-2">
							<BlockiesSvg
								address={plug.userAddress}
								className="h-5 w-5 rounded-md"
							/>
							<p className="text-sm font-bold opacity-40">
								{formatAddress(plug.userAddress)}
							</p>
						</div>
					)
				}
				nextEmpty={own === false}
			/>

			<div className="mb-4 flex flex-row items-center gap-4">
				<div className="font-bold opacity-40">
					Last updated {formatTimeSince(plug.updatedAt)}
				</div>

				<Button
					variant="secondary"
					className="group ml-auto p-1"
					onClick={() =>
						handle.plug.fork({
							id: plug.id,
							from: page.from ?? undefined
						})
					}
				>
					<GitFork size={14} />
				</Button>

				<Button
					variant="secondary"
					className="group p-1"
					onClick={() => handleFrame("share")}
				>
					<Share size={14} />
				</Button>
			</div>

			<ActionView />
		</div>
	)
}

const DynamicPageHeader = () => {
	const { page, handlePage } = usePage()
	const { handle } = usePlugs()

	return (
		<Header
			size="lg"
			label={
				<>
					<Button
						variant="secondary"
						className="rounded-sm p-1"
						onClick={() => handlePage({ key: "home" })}
					>
						<ChevronLeft size={14} />
					</Button>

					<button
						className={
							"mr-auto text-lg font-bold transition-all duration-200 ease-in-out"
						}
						onClick={() => handlePage({ key: "activity" })}
					>
						{formatTitle(page.key)}
					</button>
				</>
			}
			nextOnClick={() => handle.plug.add(page.key)}
			nextLabel={<Plus size={14} />}
		/>
	)
}

export const PageHeader = () => {
	const { page } = usePage()

	return (
		<Container>
			{["home", "activity"].includes(page.key) ? (
				<HomePageHeader />
			) : page.key === "plug" ? (
				<PlugHeader />
			) : (
				<DynamicPageHeader />
			)}
		</Container>
	)
}
