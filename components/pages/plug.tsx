import { useSession } from "next-auth/react"

import { SearchIcon } from "lucide-react"

import {
	ActionsFrame,
	Button,
	Container,
	ExecuteFrame,
	ManagePlugFrame,
	Search,
	ShareFrame
} from "@/components"
import { useFrame, usePage, usePlugs } from "@/contexts"

export const PagePlug = () => {
	const { page } = usePage()
	const { data: session } = useSession()
	const { handleFrame } = useFrame({ id: "global" })
	const { plug } = usePlugs(page.id)

	const own =
		plug !== undefined && session && session.address === plug.userAddress

	if (!plug) return null

	return (
		<Container>
			<Container className="fixed bottom-0 left-0 right-0 flex flex-col overflow-y-visible">
				<div className="absolute bottom-[150px] left-0 right-0 top-0 z-[-1] bg-gradient-to-t from-white to-white/0" />
				<div className="absolute bottom-0 left-0 right-0 z-[-1] h-[150px] bg-white" />

				{own && (
					<Search
						className="z-[4] pt-16"
						icon={<SearchIcon size={14} className="opacity-60" />}
						placeholder="Search protocols and actions"
						handleOnClick={() => handleFrame("actions")}
					/>
				)}

				<div className="mb-4 mt-2 flex flex-row gap-2">
					<Button
						variant="secondary"
						className="w-max"
						onClick={() => handleFrame("socket-run")}
					>
						Run
					</Button>

					<Button
						className="w-full"
						onClick={() => handleFrame("socket-schedule")}
					>
						Schedule
					</Button>
				</div>
			</Container>

			<ExecuteFrame />
			<ManagePlugFrame />
			<ActionsFrame />
			<ShareFrame />
		</Container>
	)
}
