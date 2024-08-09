import { FC } from "react"

import { useSession } from "next-auth/react"

import { User } from "lucide-react"

import { AuthButton, Frame } from "@/components"
import { useFrame } from "@/contexts"

export const AuthFrame: FC<{ id: string }> = ({ id }) => {
	const { isFrame } = useFrame({ id: "global", key: "auth" })
	const { data: session } = useSession()

	return (
		<Frame
			id={id}
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
