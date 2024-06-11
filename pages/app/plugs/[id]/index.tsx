import { useEffect, useRef, useState } from "react"

import Image from "next/image"
import { useSearchParams } from "next/navigation"
import { useRouter } from "next/router"

import BlockiesSvg from "blockies-react-svg"
import {
	Badge,
	Blocks,
	Ellipsis,
	Eye,
	GitFork,
	Globe,
	Link,
	LoaderCircle,
	PencilLine,
	Play,
	Redo,
	SearchIcon,
	Send,
	Settings,
	Share,
	TestTubeDiagonal,
	Twitter,
	Undo,
	Users
} from "lucide-react"

import { Container, Header } from "@/components/app"
import { Frame } from "@/components/app/frames/base"
import { ActionList } from "@/components/app/plugs/actions/action-list"
import { ActionPreview } from "@/components/app/plugs/actions/action-preview"
import { ActionView } from "@/components/app/plugs/actions/action-view"
import { Button } from "@/components/buttons"
import { Checkbox, Search } from "@/components/inputs"
import { usePlugs } from "@/contexts/PlugProvider"
import { colors, routes } from "@/lib/constants"
import { NextPageWithLayout } from "@/lib/types"
import { cn } from "@/lib/utils"

const Page: NextPageWithLayout = () => {
	const searchRef = useRef<HTMLInputElement>(null)

	const router = useRouter()
	const id = router.query.id as string
	const searchParams = useSearchParams()
	const from = searchParams.get("from")

	const {
		plug,
		version,
		handleSelect,
		handleEdit,
		handleFork,
		handleVersionChange
	} = usePlugs()

	const [search, setSearch] = useState("")

	const [manageVisible, setManageVisible] = useState(false)

	const [playStateVisible, setPlayStateVisible] = useState<
		undefined | "socket" | "chain" | "run" | "executing"
	>(undefined)
	const [chainsSelected, setChainsSelected] = useState<string[]>([])

	const [actionsVisible, setActionsVisible] = useState(false)
	const [shareVisible, setShareVisible] = useState(false)

	const [copied, setCopied] = useState(false)

	const handleChainSelect = (chain: string) => {
		setChainsSelected(prev =>
			prev.includes(chain)
				? prev.filter(c => c !== chain)
				: [...prev, chain]
		)
	}

	useEffect(() => handleSelect(id), [handleSelect, id])

	useEffect(() => {
		if (copied) {
			navigator.clipboard.writeText(
				`${window.location.origin}/app/plugs/${plug?.id || ""}`
			)

			const timeout = setTimeout(() => setCopied(false), 2000)
			return () => clearTimeout(timeout)
		}
	}, [copied, plug])

	if (!plug) return null

	return (
		<div className="relative">
			<Header
				size="lg"
				back={from ?? routes.app.plugs.mine}
				icon={
					<div
						className="h-6 w-6 rounded-md"
						style={{
							backgroundColor:
								colors[plug.color as keyof typeof colors]
						}}
					/>
				}
				label={plug.name === "" ? "Untitled Plug" : plug.name}
				nextOnClick={() => setManageVisible(!manageVisible)}
				nextLabel={<Ellipsis size={14} className="opacity-60" />}
			/>

			{plug.versions.length > 0 &&
			plug.versions[plug.versions.length - version] &&
			plug.versions[plug.versions.length - version].actions.length > 0 ? (
				<ActionView />
			) : (
				<div className="mx-auto my-20 flex max-w-[80%] flex-col gap-2 text-center">
					<p className="text-lg font-bold">
						No actions have been added yet.
					</p>
					<p className="opacity-60">
						Create a Plug to actions that you want to do on a
						regular basis and when all the conditions have been met.
					</p>
					<Button
						className="mx-auto mt-4 w-max"
						onClick={() => setActionsVisible(!actionsVisible)}
					>
						Add Action
					</Button>
				</div>
			)}

			<Container className="fixed bottom-0 left-0 right-0 flex flex-col overflow-y-visible">
				<div className="relative overflow-y-visible pt-16">
					<Search
						className="z-[4]"
						icon={<SearchIcon size={14} className="opacity-60" />}
						placeholder="Search protocols and actions"
						handleOnClick={() =>
							setActionsVisible(previousVisible => {
								// TODO: This isn't working for some reason? Potentially happening because of a re-render, but I am honestly sure. Right now though, if we click inside of a Search it selects that component to type in, but we don't want this one to be typed in ever. It's a fake display-only component.

								// NOTE: Focus the input when we click into it. We check for when it is not yet visible because it will be and thus we should be focusing it.
								if (!previousVisible) searchRef.current?.focus()

								return !previousVisible
							})
						}
					>
						<div className="absolute bottom-[60px] left-0 right-0 top-0 z-[-1] bg-gradient-to-t from-white to-white/0" />
						<div className="absolute bottom-0 left-0 right-0 z-[-1] h-[60px] bg-white" />
					</Search>
				</div>

				<div className="flex flex-row justify-between gap-2 bg-white pb-2 pt-2">
					<button
						className={cn(
							"cursor-pointer p-4 text-black/65",
							{ "hover:text-black": version !== 1 },
							{ "opacity-40": version === 1 }
						)}
						onClick={() => handleVersionChange(-1)}
						disabled={version === 1}
					>
						<Undo size={14} />
					</button>
					<button
						className={cn(
							"cursor-pointer p-4 text-black/65",
							{
								"hover:text-black":
									version !== plug.versions.length
							},
							{ "opacity-40": version === plug.versions.length }
						)}
						onClick={() => handleVersionChange(1)}
						disabled={version === plug.versions.length}
					>
						<Redo size={14} />
					</button>
					<button
						className="cursor-pointer p-4 text-black/65 hover:text-black"
						onClick={() => setPlayStateVisible("socket")}
					>
						<Play size={14} />
					</button>
					<button
						className="cursor-pointer p-4 text-black/65 hover:text-black"
						onClick={() =>
							handleFork({ id: plug.id, from: from ?? undefined })
						}
					>
						<GitFork size={14} />
					</button>
					<button
						className="cursor-pointer p-4 text-black/65 hover:text-black"
						onClick={() => setShareVisible(!shareVisible)}
					>
						<Share size={14} />
					</button>
				</div>
			</Container>

			<Frame
				className="z-[2]"
				icon={<Users size={18} className="opacity-60" />}
				label="Choose Socket"
				visible={playStateVisible === "socket"}
				handleVisibleToggle={() => setPlayStateVisible(undefined)}
			>
				<div className="flex flex-col gap-4">
					<div className="flex flex-row items-center gap-4">
						<Checkbox checked={true} handleChange={() => {}} />
						<BlockiesSvg
							address={
								"0x62180042606624f02d8a130da8a3171e9b33894d"
							}
							className="h-6 w-6 rounded-md"
						/>
						<p className="mr-auto font-bold">Defi</p>

						<div className="flex flex-row">
							<Image
								className="mr-[-10px]"
								src="/blockchain/ethereum.png"
								alt="Optimism"
								width={24}
								height={24}
							/>
							<Image
								className="mr-[-10px]"
								src="/blockchain/optimism.png"
								alt="Optimism"
								width={24}
								height={24}
							/>
							<Image
								className="mr-[-10px]"
								src="/blockchain/arbitrum.png"
								alt="Optimism"
								width={24}
								height={24}
							/>
							<Image
								src="/blockchain/base.png"
								alt="Optimism"
								width={24}
								height={24}
							/>
						</div>
					</div>

					<div className="flex flex-row items-center gap-4">
						<Checkbox checked={false} handleChange={() => {}} />
						<BlockiesSvg
							address={
								"0x63130042606624f02d8a130da8a3171e9b33894d"
							}
							className="h-6 w-6 rounded-md"
						/>
						<p className="mr-auto font-bold">Noun Sniper</p>

						<div className="flex flex-row">
							<Image
								src="/blockchain/ethereum.png"
								alt="Optimism"
								width={24}
								height={24}
							/>
						</div>
					</div>

					<div className="flex flex-row gap-4">
						<Button
							variant="secondary"
							className="w-max"
							onClick={() => setPlayStateVisible("chain")}
						>
							Deploy
						</Button>

						<Button
							className="w-full"
							onClick={() => setPlayStateVisible("chain")}
						>
							Continue
						</Button>
					</div>
				</div>
			</Frame>

			<Frame
				className="z-[2]"
				handleBack={() => setPlayStateVisible("socket")}
				icon={<Globe size={18} className="opacity-60" />}
				label="Choose Chain"
				visible={playStateVisible === "chain"}
				handleVisibleToggle={() => setPlayStateVisible(undefined)}
				hasOverlay={true}
			>
				<div className="flex flex-col gap-4">
					<div className="flex flex-row items-center gap-4">
						<Checkbox
							checked={chainsSelected.includes("ethereum")}
							handleChange={() => handleChainSelect("ethereum")}
						/>

						<div className="mr-auto flex flex-row gap-2">
							<Image
								src="/blockchain/ethereum.png"
								alt="Optimism"
								width={24}
								height={24}
							/>
							<p className="font-bold">Ethereum</p>
						</div>

						<p className="opacity-60">0.4 ETH</p>
					</div>
					<div className="flex flex-row items-center gap-4">
						<Checkbox
							checked={chainsSelected.includes("optimism")}
							handleChange={() => handleChainSelect("optimism")}
						/>

						<div className="mr-auto flex flex-row gap-2">
							<Image
								src="/blockchain/optimism.png"
								alt="Optimism"
								width={24}
								height={24}
							/>
							<p className="font-bold">Optimism</p>
						</div>

						<p className="opacity-60">0.2 ETH</p>
					</div>
					<div className="flex flex-row items-center gap-4">
						<Checkbox
							checked={chainsSelected.includes("arbitrum")}
							handleChange={() => handleChainSelect("arbitrum")}
						/>

						<div className="mr-auto flex flex-row gap-2">
							<Image
								src="/blockchain/arbitrum.png"
								alt="Optimism"
								width={24}
								height={24}
							/>
							<p className="font-bold">Arbitrum</p>
						</div>

						<p className="opacity-60">0.8 ETH</p>
					</div>
					<div className="flex flex-row items-center gap-4">
						<Checkbox
							checked={chainsSelected.includes("base")}
							handleChange={() => handleChainSelect("base")}
						/>

						<div className="mr-auto flex flex-row gap-2">
							<Image
								src="/blockchain/base.png"
								alt="Optimism"
								width={24}
								height={24}
							/>
							<p className="font-bold">Base</p>
						</div>

						<p className="opacity-60">0.12 ETH</p>
					</div>
					<div className="flex flex-row gap-4">
						<Button
							variant="secondary"
							className="w-max"
							onClick={() => setPlayStateVisible("run")}
						>
							Execute
						</Button>

						<Button
							className="w-full"
							onClick={() => setPlayStateVisible("chain")}
						>
							Schedule
						</Button>
					</div>
				</div>
			</Frame>

			<Frame
				className="z-[2]"
				handleBack={() => setPlayStateVisible("chain")}
				icon={<Eye size={18} className="opacity-60" />}
				label="Transaction Preview"
				visible={playStateVisible === "run"}
				handleVisibleToggle={() => setPlayStateVisible(undefined)}
				hasOverlay={true}
			>
				<div className="flex flex-col gap-4">
					<p className="font-bold opacity-60">Actions</p>
					<ActionPreview />
					<p className="flex font-bold">
						<span className="mr-auto opacity-60">Use Socket</span>
						<div className="flex flex-row items-center gap-2">
							<BlockiesSvg
								address="0x581BEf12967f06f2eBfcabb7504fA61f0326CD9A"
								className="h-5 w-5 rounded-md"
							/>
							<p className="mr-auto font-bold">Noun Sniper</p>
						</div>
					</p>
					<p className="flex font-bold">
						<span className="mr-auto opacity-60">Run On</span>
						{chainsSelected.map(chain => (
							<>
								<Image
									className="ml-[-10px] h-6 w-6"
									src={`/blockchain/${chain}.png`}
									alt={chain}
									width={24}
									height={24}
								/>
							</>
						))}
					</p>
					<p className="flex font-bold">
						<span className="mr-auto opacity-60">Total</span>
						<div className="flex flex-row gap-2">
							<span className="opacity-40">0.0011 ETH</span>
							<span>$4.19</span>
						</div>
					</p>

					<Button
						className="mt-4 w-full"
						onClick={() => setPlayStateVisible("executing")}
					>
						Run Transaction
					</Button>
				</div>
			</Frame>

			<Frame
				className="z-[2]"
				icon={
					<LoaderCircle
						size={18}
						className="animate-spin opacity-60"
					/>
				}
				label="Execution Processing"
				visible={playStateVisible === "executing"}
				handleVisibleToggle={() => setPlayStateVisible(undefined)}
			>
				<div className="flex flex-col gap-8">
					<p className="leading-6">
						<span className="opacity-60">Your</span>
						<span
							className="rounded-lg bg-gradient-to-tr px-2 py-1 font-bold"
							style={{
								background: `linear-gradient(to right, rgba(0,239,54,0.1), rgba(147,223,0,0.1))`,
								color: `#00E100`
							}}
						>
							{plug.name}
						</span>{" "}
						<span className="opacity-60">
							execution is currently processing.
						</span>
					</p>

					<p className="opacity-60">
						You can go ahead and schedule execution in the future
						while you wait for your transaction to complete.
					</p>
				</div>
			</Frame>

			<Frame
				className="z-[2]"
				icon={<Settings size={18} className="opacity-60" />}
				label="Manage Plug"
				visible={manageVisible}
				handleVisibleToggle={() => setManageVisible(!manageVisible)}
			>
				<div className="flex flex-col gap-4">
					<Search
						icon={<PencilLine size={14} className="opacity-60" />}
						placeholder="Plug name"
						search={plug.name}
						handleSearch={(name: string) =>
							handleEdit({
								...plug,
								name
							})
						}
					/>

					<div className="flex flex-row items-center gap-2">
						<p className="mr-auto font-bold">Private</p>

						<Checkbox
							checked={plug.isPrivate}
							handleChange={(checked: boolean) =>
								handleEdit({
									...plug,
									isPrivate: checked
								})
							}
						/>
					</div>

					<div className="flex flex-row items-center gap-2">
						<p className="font-bold">Color</p>

						<div className="ml-auto flex flex-wrap items-center gap-1">
							{Object.keys(colors).map(color => (
								<div
									key={color}
									className="group flex h-6 w-6 cursor-pointer items-center justify-center rounded-full border-[2px]"
									style={{
										borderColor:
											plug.color === color
												? colors[
														color as keyof typeof colors
													]
												: "transparent"
									}}
									onClick={() =>
										handleEdit({
											...plug,
											color
										})
									}
								>
									<div
										className="h-full w-full rounded-full border-[2px] border-white transition-all duration-200 ease-in-out"
										style={{
											backgroundColor:
												colors[
													color as keyof typeof colors
												]
										}}
									/>
								</div>
							))}
						</div>
					</div>
				</div>

				<div className="mt-[20px] flex flex-row gap-2">
					<Button
						variant="destructive"
						className="w-full"
						onClick={() => {}}
					>
						Delete
					</Button>
				</div>
			</Frame>

			<Frame
				className="scrollbar-hide z-[1] h-[calc(100vh-80px)] overflow-y-auto"
				icon={<Blocks size={18} className="opacity-60" />}
				label="Add Action"
				visible={actionsVisible}
				handleVisibleToggle={() => setActionsVisible(!actionsVisible)}
			>
				<div className="flex flex-col gap-8">
					<Search
						ref={searchRef}
						icon={<SearchIcon size={14} className="opacity-60" />}
						placeholder="Search protocols and actions"
						search={search}
						handleSearch={setSearch}
					/>

					<ActionList
						handleNestedToggle={() =>
							setActionsVisible(!actionsVisible)
						}
					/>
				</div>
			</Frame>

			<Frame
				className="z-[2]"
				icon={<Badge size={18} className="opacity-60" />}
				label="Share Plug"
				visible={shareVisible}
				handleVisibleToggle={() => setShareVisible(!shareVisible)}
			>
				<div className="flex flex-col gap-4">
					<div className="flex flex-row items-center gap-2">
						<Link size={14} className="opacity-60" />
						<p className="font-bold">Direct Link</p>
						<Button
							variant="secondary"
							sizing="sm"
							className="ml-auto"
							onClick={() => setCopied(true)}
						>
							{copied ? "Copied" : "Copy"}
						</Button>
					</div>

					<div className="flex flex-row items-center gap-2">
						<Twitter size={14} className="opacity-60" />
						<p className="font-bold">Twitter</p>
						<a
							className="ml-auto"
							href={`https://twitter.com/intent/tweet?text=${plug.name} using @onplug_io:%0A%0A${window.location.origin}${routes.app.plugs.index}/${plug.id}`}
							target="_blank"
							rel="noopener noreferrer"
						>
							<button className="rounded-full bg-gradient-to-tr from-[#0085CE] to-[#00A2FB] px-[24px] py-[8px] text-xs font-bold text-white hover:from-[#0085CE]/90 hover:to-[#00A2FB]/90">
								Tweet
							</button>
						</a>
					</div>

					<div className="flex flex-row items-center gap-2">
						<Send size={14} className="opacity-60" />
						<p className="font-bold">Telegram</p>
						<a
							className="ml-auto"
							href={`https://t.me/share/url?url=${window.location.origin}${routes.app.plugs.index}/${plug.id}&text=${plug.name} using @onplug_io`}
							target="_blank"
							rel="noopener noreferrer"
						>
							<button className="rounded-full bg-gradient-to-tr from-[#00A2E3] to-[#67D4FF] px-[24px] py-[8px] text-xs font-bold text-white hover:from-[#00A2E3]/90 hover:to-[#67D4FF]/90">
								Share
							</button>
						</a>
					</div>

					<div className="flex flex-row items-center gap-2">
						<Image
							className="opacity-60"
							src="/icons/farcaster.svg"
							alt="Farcaster"
							width={14}
							height={14}
						/>
						<p className="font-bold">Warpcast</p>
						<a
							className="ml-auto"
							href={`https://warpcast.com/~/compose?text=https://twitter.com/intent/tweet?text=${plug.name}%20using%20@onplug_io&embeds[]=${window.location.origin}${routes.app.plugs.index}/${plug.id}`}
							target="_blank"
							rel="noopener noreferrer"
						>
							<button className="rounded-full bg-[#472A91] px-[24px] py-[8px] text-xs font-bold text-white hover:bg-[#472A91]/90">
								Cast
							</button>
						</a>
					</div>
				</div>
			</Frame>
		</div>
	)
}

Page.getLayout = page => <Container>{page}</Container>

export default Page
