import { Cable, Plus, UsersRound } from "lucide-react"

import { Container, Header } from "@/components/app"
import { PlugGrid } from "@/components/app/plugs/grid"
import { SocketList } from "@/components/app/sockets/socket-list"
import { AuthButton } from "@/components/buttons"
import { usePlugs, useSockets } from "@/contexts"
import { routes } from "@/lib/constants"
import { NextPageWithLayout } from "@/lib/types"

const Page: NextPageWithLayout = () => {
	const { address, sockets, handleAdd: handleSocketAdd } = useSockets()
	const { handleAdd: handlePlugAdd } = usePlugs()

	const hasSockets = sockets && sockets.length > 0

	return (
		<>
			<Header
				size="lg"
				label="Plug"
				nextOnClick={() => handlePlugAdd(routes.app.index)}
				nextLabel={<Plus size={14} className="opacity-60" />}
			/>

			<Header
				size="md"
				icon={<Cable size={14} className="opacity-60" />}
				label="Plugs"
				nextHref={routes.app.plugs.index}
				nextLabel="See All"
			/>
			<PlugGrid from={routes.app.index} count={8} all={true} />

			<Header
				size="md"
				icon={<UsersRound size={14} className="opacity-60" />}
				label="Sockets"
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
		</>
	)
}

Page.getLayout = page => <Container>{page}</Container>

export default Page
