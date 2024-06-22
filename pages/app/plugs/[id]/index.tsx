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
import { Search } from "@/components/inputs"
import { useFrame } from "@/contexts"
import { usePlugs } from "@/contexts/PlugProvider"
import { colors, routes } from "@/lib/constants"
import { useNavigation } from "@/lib/hooks/useNavigation"
import { NextPageWithLayout } from "@/lib/types"
import { cn } from "@/lib/utils"

const Page: NextPageWithLayout = () => {
	const { id, from } = useNavigation()
	const { plug, version, actions } = usePlugs(id)
	const { handleFrameVisible } = useFrame()

	if (!plug) return null

	return (
		<Container className="relative">
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

				<div className="flex flex-row justify-between gap-2 bg-white pb-2 pt-2">
					<button
						className={cn(
							"cursor-pointer p-4 text-black/65",
							{ "hover:text-black": version !== 1 },
							{ "opacity-40": version === 1 }
						)}
						onClick={() => actions.plug.handleVersion(-1)}
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
						onClick={() => actions.plug.handleVersion(1)}
						disabled={version === plug.versions.length}
					>
						<Redo size={14} />
					</button>
					<button
						className="cursor-pointer p-4 text-black/65 hover:text-black"
						onClick={() => handleFrameVisible("socket")}
					>
						<Play size={14} />
					</button>
					<button
						className="cursor-pointer p-4 text-black/65 hover:text-black"
						onClick={() =>
							actions.plug.handleFork({
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
				</div>
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
