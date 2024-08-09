import { User } from "lucide-react"

import { Frame } from "@/components"
import { useFrame } from "@/contexts"

export const TestFrame = () => {
	const { id, isFrame } = useFrame({ id: "global", key: "test" })

	return (
		<Frame
			id={id}
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
