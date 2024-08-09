import { CheckCircle } from "lucide-react"

import { Button, Frame } from "@/components"
import { useFrame } from "@/contexts"

export const FeatureRequestFrame = () => {
	const {
		id,
		isFrame,
		prevFrame: from,
		handleFrame
	} = useFrame({ id: "global", key: "featureRequestSubmit", seperator: "-" })

	return (
		<Frame
			id={id}
			icon={<CheckCircle size={18} />}
			label="Feature Request Submit"
			visible={isFrame}
		>
			<div className="flex flex-col items-center gap-2">
				<p className="w-full opacity-60">
					Your feedback has been submitted. Thank you for your
					feedback. Our team will look into it soon!
				</p>

				{from && (
					<Button
						variant="primary"
						className="mt-4 w-full"
						onClick={() => handleFrame(from)}
					>
						Continue
					</Button>
				)}
			</div>
		</Frame>
	)
}
