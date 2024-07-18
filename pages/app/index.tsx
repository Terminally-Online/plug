import Image from "next/image"

import BlockiesSvg from "blockies-react-svg"
import { Cable, Plus, UsersRound } from "lucide-react"

import {
	AccountFrame,
	Container,
	Header,
	PlugGrid,
	SocketList
} from "@/components"
import { useFrame, usePlugs, useSockets } from "@/contexts"
import { formatAddress, NextPageWithLayout, routes } from "@/lib"
import { api } from "@/server/client"

const Page: NextPageWithLayout = () => {
	const { handleFrameVisible } = useFrame()
	const {
		address,
		ensName,
		ensAvatar,
		sockets,
		handleAdd: handleSocketAdd
	} = useSockets()
	const { handle } = usePlugs()

	const { data: plugs } = api.plug.all.useQuery({ target: "mine", limit: 8 })

	const hasSockets = sockets && sockets.length > 0

	return (
		<>
			<Header
				size="lg"
				label={
					<>
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
						<p className="mr-auto text-xl font-bold">Plug</p>
					</>
				}
				nextOnClick={() => handle.plug.add(routes.app.index)}
				nextLabel={<Plus size={14} />}
			>
				{address && (
					<button
						className="ml-auto flex flex-row items-center gap-2"
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
						<p className="font-bold opacity-40">
							{ensName ?? formatAddress(address)}
						</p>
					</button>
				)}
			</Header>

			<Header
				size="md"
				icon={<Cable size={14} className="opacity-40" />}
				label="Plugs"
				nextHref={routes.app.plugs.index}
				nextLabel="See All"
			/>
			<PlugGrid from={routes.app.index} count={8} plugs={plugs} />

			<Header
				size="md"
				icon={<UsersRound size={14} className="opacity-40" />}
				label="Accounts"
				nextOnClick={handleSocketAdd}
				nextLabel={hasSockets ? "Create New" : undefined}
			/>
			<SocketList />

			<AccountFrame />
		</>
	)
}

Page.getLayout = page => <Container>{page}</Container>

export default Page
