import { useSession } from "next-auth/react"

import { User } from "lucide-react"

import { AuthButton, Frame } from "@/components"
import { useColumns } from "@/state"

export const AuthFrame = () => {
	const index = -1
	const { data: session } = useSession()
	const { isFrame } = useColumns(index, "auth")

	return (
		<Frame
			index={index}
			icon={<User size={18} />}
			label={session?.address ? "Account" : "Login"}
			visible={isFrame}
			hasOverlay={true}
			hasChildrenPadding={false}
		>
			<div className="flex flex-col gap-4 px-6 pb-4">
				<AuthButton />
			</div>
		</Frame>
	)
}
