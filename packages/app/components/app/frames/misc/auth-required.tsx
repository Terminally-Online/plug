import { useSession } from "next-auth/react"
import { FC, useMemo } from "react"

import { User } from "lucide-react"

import { Button, Frame } from "@/components"
import { useColumnData, useSidebar } from "@/state"

const FRAMES_REQUIRED_AUTH = ["schedule", "run"]

export const AuthRequiredFrame: FC<{ index: number }> = ({ index }) => {
	const { data: session } = useSession()
	const { is, handleActivePane } = useSidebar()
	const { column } = useColumnData(index)

	const isFrame = useMemo(
		() =>
			(session &&
				column &&
				session.user.anonymous === true &&
				FRAMES_REQUIRED_AUTH.includes(column?.frame || "")) ||
			false,
		[session, column]
	)

	return (
		<Frame
			index={index}
			icon={<User size={18} className="opacity-40" />}
			label="Authentication Required"
			visible={isFrame}
			hasOverlay={true}
			hasChildrenPadding={false}
		>
			<div className="flex flex-col gap-4 px-6 pb-4">
				<p className="text-sm font-bold opacity-40">
					Before running or scheduling this Plug, you must authenticate by connecting a wallet so that we can
					route everything through the appropriate onchain contracts.
				</p>
				<Button
					className="w-full"
					onClick={() => (is.authenticating ? undefined : handleActivePane("authenticating"))}
				>
					{is.authenticating ? "Logging in..." : "Login"}
				</Button>
			</div>
		</Frame>
	)
}
