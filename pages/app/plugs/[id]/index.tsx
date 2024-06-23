import {
	Ellipsis,
	GitFork,
	Play,
	Redo,
	SearchIcon,
	Share,
	Undo
} from "lucide-react"

import { Container, Header } from "@/components/app"
import { ActionsFrame } from "@/components/app/frames/plugs/[id]/actions"
import { ExecuteFrame } from "@/components/app/frames/plugs/[id]/execute"
import { ManageFrame } from "@/components/app/frames/plugs/[id]/manage"
import { ShareFrame } from "@/components/app/frames/plugs/[id]/share"
import { ActionView } from "@/components/app/plugs/actions/action-view"
import { Button } from "@/components/buttons"
import { Search } from "@/components/inputs"
import { useFrame } from "@/contexts"
import { usePlugs } from "@/contexts/PlugProvider"
import { colors, routes } from "@/lib/constants"
import { useNavigation } from "@/lib/hooks/useNavigation"
import { NextPageWithLayout } from "@/lib/types"

const Page: NextPageWithLayout = () => {
	const { id, from } = useNavigation()
	const { plug, handle } = usePlugs(id)
	const { handleFrameVisible } = useFrame()

	if (!plug) return null

	return (
		<Container>
			<div className="flex min-h-[calc(100vh-80px)] flex-col">
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
					nextOnClick={() => handleFrameVisible("manage")}
					nextLabel={<Ellipsis size={14} className="opacity-60" />}
				/>

				<ActionView />
			</div>

			<Container className="fixed bottom-0 left-0 right-0 flex flex-col overflow-y-visible">
				<div className="relative overflow-y-visible pt-16">
					<Search
						className="z-[4]"
						icon={<SearchIcon size={14} className="opacity-60" />}
						placeholder="Search protocols and actions"
						handleOnClick={() => handleFrameVisible("actions")}
					>
						<div className="absolute bottom-[60px] left-0 right-0 top-0 z-[-1] bg-gradient-to-t from-white to-white/0" />
						<div className="absolute bottom-0 left-0 right-0 z-[-1] h-[60px] bg-white" />
					</Search>
				</div>

				<Button
					className="mb-4 mt-2"
					onClick={() => handleFrameVisible("socket")}
				>
					Run
				</Button>

				{/* <div className="flex flex-row justify-between gap-2 bg-white pb-2 pt-2">
					<button
						className="cursor-pointer p-4 text-black/65 hover:text-black"
						onClick={() =>
							handle.plug.fork({
								id: plug.id,
								from: from ?? undefined
							})
						}
					>
						<GitFork size={14} />
					</button>
					<button
						className="cursor-pointer p-4 text-black/65 hover:text-black"
						onClick={() => handleFrameVisible("share")}
					>
						<Share size={14} />
					</button>
				</div> */}
			</Container>

			<>
				<ExecuteFrame />
				<ManageFrame />
				<ActionsFrame />
				<ShareFrame />
			</>
		</Container>
	)
}

export default Page
