import Image from "next/image"

import BlockiesSvg from "blockies-react-svg"
import { Cable, Plus, UsersRound } from "lucide-react"

import { Container, Header } from "@/components/app"
import { AccountFrame } from "@/components/app/frames/account"
import { PlugGrid } from "@/components/app/plugs/grid"
import { SocketList } from "@/components/app/sockets/socket-list"
import { AuthButton } from "@/components/buttons"
import { useFrame, usePlugs, useSockets } from "@/contexts"
import { routes } from "@/lib/constants"
import { formatAddress } from "@/lib/functions"
import { NextPageWithLayout } from "@/lib/types"
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
							className="flex w-max flex-row items-center rounded-[10px] p-2"
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
								className="rounded-md"
							/>
						) : (
							<BlockiesSvg
								className="h-6 w-6 rounded-md"
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

			{address === undefined ? (
				<div className="my-32 flex flex-col gap-[30px]">
					<p className="mx-auto w-[80%] max-w-[360px] text-center text-lg opacity-60">
						Step into Plug and get started by connecting your wallet
						to manage your Sockets and Plug in one place.
					</p>

					<AuthButton />
				</div>
			) : (
				<SocketList />
			)}

			<AccountFrame />
		</>
	)
}

Page.getLayout = page => (
	<>
		{/* <Header
			size="lg"
			label={
				<>
					<div
						className="mr-2 flex w-max flex-row items-center gap-2 rounded-md p-2"
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
			className="z-[20] bg-white"
		/> */}
		<Container>{page}</Container>
	</>
)

export default Page
