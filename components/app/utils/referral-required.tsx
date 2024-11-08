import { useSession } from "next-auth/react"
import { FC, useState } from "react"

import { Button } from "@/components/shared"
import { cn, greenGradientStyle, useClipboard } from "@/lib"
import { api } from "@/server/client"
import { useSocket } from "@/state"
import { Search } from "../inputs"
import { SearchIcon } from "lucide-react"

export const ReferralRequired: FC = () => {
	const { data: session } = useSession()
	const { socket } = useSocket()
	const [referralAddress, setReferralAddress] = useState("")
	const [error, setError] = useState<string>()
	const [success, setSuccess] = useState(false)

	const submitReferral = api.socket.submitReferral.useMutation({
		onError: error => setError(error.message),
		onSuccess: () => {
			setSuccess(true)
			setError(undefined)
			// Add a small delay before refreshing to show the success message
			setTimeout(() => {
				window.location.reload()
			}, 500)
		}
	})

	const requestAccess = api.socket.requestAccess.useMutation()

	const isAuthenticated = session?.user.id?.startsWith("0x")
	const isVisible = Boolean(isAuthenticated && socket && !socket.identity?.approvedAt)

	if (!isVisible) return null

	const handleSubmitReferral = () => {
		if (!referralAddress) return
		submitReferral.mutate({ referrerAddress: referralAddress })
	}

	const handleRequestAccess = async () => {
		const tweetTemplates = [
			`Hey @onplug_io! I'm ready to automate my onchain life with Plug 🔌\n\nMy address: ${session?.user.id}`,
			`Excited to try @onplug_io - the best way to automate in crypto! 🤖\n\nMy address: ${session?.user.id}`,
			`Hey @onplug_io - I'm looking to make Plug my new home for onchain activity ⚡\n\nMy address: ${session?.user.id}`,
			`Need that @onplug_io access to start automating my onchain activities 🎯\n\nMy address: ${session?.user.id}`,
			`Hey @onplug_io! Let's streamline my crypto life with automated onchain actions ⚡\n\nMy address: ${session?.user.id}`
		]

		const randomTweet = tweetTemplates[Math.floor(Math.random() * tweetTemplates.length)]
		const tweetText = encodeURIComponent(randomTweet)
		window.open(`https://twitter.com/intent/tweet?text=${tweetText}`, "_blank")
		await requestAccess.mutateAsync()
	}

	return (
		<div className="flex h-full w-full flex-col items-start justify-between bg-white">
			<div className="flex w-full justify-center pt-[30vh]">
				<div className="flex w-full max-w-md flex-col px-4">
					<div className="flex flex-col items-center gap-8">
						<h1 className="text-2xl font-bold">Start Automating with Plug.</h1>

						<div className="w-full space-y-6">
							{error && (
								<div className="w-full rounded-lg border border-red-200 bg-red-50 p-4 text-sm text-red-600">
									{error}
								</div>
							)}

							{success && (
								<div className="w-full rounded-lg border border-green-200 bg-green-50 p-4 text-sm text-green-600">
									Successfully approved! Please refresh the page.
								</div>
							)}

							<div className="flex flex-col gap-2">
								<p className="text-center text-sm font-medium opacity-60">
									Enter an address from an existing user to get started.
								</p>
								<input
									type="text"
									placeholder="0x... or ENS name"
									value={referralAddress}
									onChange={e => setReferralAddress(e.target.value)}
									className="w-full rounded-lg border-[1px] border-white bg-grayscale-0 px-4 py-2 font-mono text-sm outline-none transition-colors duration-200 ease-in-out hover:border-grayscale-100 hover:bg-white"
								/>
								<Button
									className="w-full"
									onClick={handleSubmitReferral}
									disabled={submitReferral.isLoading || !referralAddress}
									variant={submitReferral.isLoading || !referralAddress ? "disabled" : "primary"}
								>
									{submitReferral.isLoading ? "Submitting..." : "Submit Referral"}
								</Button>
							</div>
						</div>
					</div>
				</div>
			</div>

			<div className="flex w-full justify-center pb-8">
				<div className="flex w-full max-w-md flex-col gap-2 px-4">
					<p className="text-center text-sm font-medium opacity-60">No invite code? No Problem.</p>
					<Button
						className="w-full"
						onClick={handleRequestAccess}
						disabled={requestAccess.isLoading}
						variant="primary"
					>
						{requestAccess.isLoading ? "Requesting..." : "Request Access"}
					</Button>
				</div>
			</div>
		</div>
	)
}

