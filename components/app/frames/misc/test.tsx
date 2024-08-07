import { useSession } from "next-auth/react"

import { User } from "lucide-react"

import { AuthButton, Frame } from "@/components"
import { useFrame } from "@/contexts"

export const TestFrame = () => {
	const { frameVisible } = useFrame()

	const isFrame = frameVisible === "test"

	return (
		<Frame
			icon={<User size={18} />}
			label="Test Frame"
			visible={isFrame}
			hasOverlay={true}
			hasChildrenPadding={false}
		>
			<div className="flex flex-col gap-4 px-6 pb-4">
				<p>Hello gorgeous. What are you doing here?</p>
			</div>
		</Frame>
	)
}
