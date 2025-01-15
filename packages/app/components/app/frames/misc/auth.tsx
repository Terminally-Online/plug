import { useSession } from "next-auth/react"

import { User } from "lucide-react"

import { Frame } from "@/components/app/frames/base"
import { AuthButton } from "@/components/shared/buttons/auth"
import { useColumnStore } from "@/state/columns"

export const AuthFrame = () => {
	const index = -1
	const { data: session } = useSession()
	const { isFrame } = useColumnStore(index, "auth")

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
