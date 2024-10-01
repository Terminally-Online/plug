import { FC, useState } from "react"

import { Pen, Sparkle } from "lucide-react"

import { api } from "@/server/client"

import { Button, Frame, Search } from "@/components"
import { useColumns } from "@/state"

export const FeatureRequestFrame: FC<{ index: number }> = ({ index }) => {
	const { isFrame } = useColumns(index, "freatureRequest")

	const [message, setMessage] = useState("")

	const handleFeatureRequest = api.misc.featureRequest.useMutation({
		onMutate: () => {
			setMessage("")
			// TODO: We un-implemented this when beginning to store frames on columns.
			// frame(`featureRequestSubmit-${from}`)
		}
	})

	return (
		<Frame
			index={index}
			icon={<Sparkle size={18} />}
			label="Feature Request"
			visible={isFrame}
			// handleBack={from ? () => handleFrame(from) : undefined}
			hasOverlay={true}
		>
			<div className="flex flex-col items-center gap-2">
				<p className="mb-4 w-full opacity-60">
					Have a piece of feedback or feature request to submit? Feel free to share more details about your
					request.
				</p>

				<p className="flex w-full font-bold">Message</p>

				<Search
					className="w-full"
					icon={<Pen size={14} />}
					placeholder="Have specific details to share? Go ahead!"
					search={message}
					handleSearch={setMessage}
					textArea={true}
				/>

				{message && message.length > 360 && (
					<p className="w-full text-sm opacity-60">
						Woah! You wrote us an essay. Thank you. Your time and effort is very much appreciated.
					</p>
				)}

				<Button
					variant={message && message.length > 4 ? "primary" : "disabled"}
					className="mt-4 w-full"
					onClick={() =>
						handleFeatureRequest.mutate({
							context: "",
							message
						})
					}
					disabled={!message || message.length < 4}
				>
					{message && message.length > 4 ? "Submit Request" : "Write Feedback"}
				</Button>
			</div>
		</Frame>
	)
}
