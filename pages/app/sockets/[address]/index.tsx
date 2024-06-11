import { useEffect, useState } from "react"

import { useRouter } from "next/router"

import BlockiesSvg from "blockies-react-svg"
import { Ellipsis, PencilLine, SearchIcon } from "lucide-react"

import { Container, Header } from "@/components/app"
import { Frame } from "@/components/app/frames/base"
import { SocketActivity } from "@/components/app/sockets/activity"
import { SocketAssetList } from "@/components/app/sockets/asset-list"
import { Button } from "@/components/buttons"
import { Search } from "@/components/inputs/search"
import { useBalances, useSockets } from "@/contexts"
import { routes } from "@/lib/constants"
import { NextPageWithLayout } from "@/lib/types"
import { cn } from "@/lib/utils"

const Page: NextPageWithLayout = () => {
	const router = useRouter()
	const address = router.query.address as string

	const { socket, handleSelect, handleRename } = useSockets()
	const { search, balances, handleSearch } = useBalances({
		address: socket?.socketAddress || ""
	})

	const [tab, setTab] = useState<"assets" | "activity" | "search">("activity")
	const [manageVisible, setManageVisible] = useState(false)

	const [copied, setCopied] = useState(false)

	useEffect(() => handleSelect(address), [handleSelect, address])

	useEffect(() => {
		if (copied) {
			navigator.clipboard.writeText(socket?.socketAddress ?? "")
			setTimeout(() => setCopied(false), 2000)
		}
	}, [copied, socket])

	if (!socket) return null

	return (
		<>
			<Header
				size="lg"
				back={routes.app.index}
				icon={
					<BlockiesSvg
						address={socket.socketAddress}
						className="h-6 w-6 rounded-lg"
					/>
				}
				label={socket.name}
				nextOnClick={() => setManageVisible(!manageVisible)}
				nextLabel={<Ellipsis size={14} />}
			/>

			<div className="flex flex-col gap-4">
				<Search
					icon={<SearchIcon size={14} className="opacity-60" />}
					placeholder="Search activity and assets"
					search={search}
					handleSearch={(search: string) => handleSearch(search)}
				/>

				<div className="flex flex-row gap-2 text-lg font-bold">
					<h3
						className={cn(
							tab !== "activity" && "opacity-40",
							"cursor-pointer transition-all duration-200 ease-in-out hover:opacity-100"
						)}
						onClick={() => setTab("activity")}
					>
						Activity
					</h3>
					<h3
						className={cn(
							tab !== "assets" && "opacity-40",
							"cursor-pointer transition-all duration-200 ease-in-out hover:opacity-100"
						)}
						onClick={() => setTab("assets")}
					>
						Assets
					</h3>
				</div>
			</div>

			<div className="mt-[20px]">
				{tab === "assets" ? (
					<SocketAssetList balances={balances} />
				) : (
					<SocketActivity />
				)}
			</div>

			<Frame
				className="z-[2]"
				label="Manage Socket"
				visible={manageVisible}
				handleVisibleToggle={() => setManageVisible(!manageVisible)}
			>
				<Search
					icon={<PencilLine size={14} className="opacity-60" />}
					placeholder="Socket name"
					search={socket.name}
					handleSearch={handleRename}
				/>

				<div className="mt-[20px] flex flex-row gap-2">
					<Button
						variant="secondary"
						className="w-max"
						onClick={() => setCopied(true)}
					>
						{copied ? "Copied" : "Copy"}
					</Button>
					<Button
						variant="primary"
						className="w-full"
						onClick={() => {}}
					>
						Deploy Onchain
					</Button>
				</div>
			</Frame>
		</>
	)
}

Page.getLayout = page => <Container>{page}</Container>

export default Page
