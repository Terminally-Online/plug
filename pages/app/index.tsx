import { Cable, Cross, Plus, UsersRound, X } from "lucide-react"

import { Container, Header } from "@/components/app"
import { PlugGrid } from "@/components/app/plugs/grid"
import { SocketList } from "@/components/app/sockets/socket-list"
import { AuthButton, Button } from "@/components/buttons"
import { usePlugs, useSockets } from "@/contexts"
import { routes } from "@/lib/constants"
import { NextPageWithLayout } from "@/lib/types"

const Page: NextPageWithLayout = () => {
	const { address, sockets, handleAdd: handleSocketAdd } = useSockets()
	const { handle } = usePlugs()

	const hasSockets = sockets && sockets.length > 0

	return (
		<>
			<Header
				size="lg"
				label="Plug"
				nextOnClick={() => handle.plug.add(routes.app.index)}
				nextLabel={<Plus size={14} className="opacity-60" />}
			/>

			<div className="flex flex-col gap-2 rounded-md bg-grayscale-100 p-4">
				<div className="flex w-full items-center font-bold">
					<h3>Hey, Plug is in private testing mode</h3>
					<Button
						variant="secondary"
						className="ml-auto p-1"
						onClick={() => {}}
					>
						<X size={14} className="opacity-60" />
					</Button>
				</div>
				<p className="mr-8 text-sm">
					Nothing is final and everything is subject to change. If you
					stumble upon any bugs or have feedback please let us know
					for a little treat as a thank you!
				</p>
				<p className="text-sm font-bold underline">
					Submit Feedback Now
				</p>
			</div>

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
