import { useSearchParams } from "next/navigation"
import { FC, useEffect, useState } from "react"

import { Asterisk } from "lucide-react"

import { Button, Search } from "@/components"
import { cn, greenGradientStyle, useConnect } from "@/lib"
import { api } from "@/server/client"
import { useSocket } from "@/state"

const TWEET_TEMPLATES = [
	`Yo @onplug_io! Wen automation? Ready to ape into this Plug life fam 🦍`,
	`@onplug_io ser pls gib access, need to automate my degen plays 🚀`,
	`gm @onplug_io! Time to stop being poor and start being automated 💸`,
	`@onplug_io wen moon? Need that alpha automation access rn fr fr 🌙`,
	`ayoo @onplug_io! I'm ready to become an automation maxi, LFG 🔥`
]

export const ReferralRequired: FC = () => {
	const { account } = useConnect()
	const { socket } = useSocket()

	const { mutate, error, isLoading, isError, isSuccess } = api.socket.referral.submit.useMutation()
	const requestAccess = api.socket.referral.request.useMutation()

	const searchParams = useSearchParams()
	const [referrerAddress, setReferralAddress] = useState("")

	useEffect(() => {
		const rfid = searchParams.get("rfid")
		if (rfid) {
			setReferralAddress(rfid)
		}
	}, [searchParams])

	const isVisible = Boolean(account.isAuthenticated && socket && !socket.identity?.approvedAt)

	const handleRequestAccess = async () => {
		const randomTweet = TWEET_TEMPLATES[Math.floor(Math.random() * TWEET_TEMPLATES.length)]
		const tweetText = encodeURIComponent(randomTweet + `\n\nMy address: ${account.address}`)
		window.open(`https://twitter.com/intent/tweet?text=${tweetText}`, "_blank")
		await requestAccess.mutateAsync()
	}

	if (!isVisible) return null

	return (
		<div className="flex w-full flex-col items-start justify-between bg-white p-2">
			<div className="flex h-full w-full items-center justify-center rounded-lg border-[1px] border-grayscale-100">
				<div className="flex max-w-[480px] flex-col items-center">
					<h1 className="text-2xl font-bold">Get Access to Plug.</h1>
					<p className="mb-8 text-center font-bold text-black/40">
						Enter the referral code you received to get started or request one by tagging{" "}
						<button
							onClick={handleRequestAccess}
							style={{
								...greenGradientStyle
							}}
						>
							@onplug_io
						</button>{" "}
						on Twitter.
					</p>

					<div className="w-full space-y-6">
						{isSuccess && (
							<div className="w-full rounded-lg border border-green-200 bg-green-50 p-4 text-sm text-green-600">
								Successfully approved! If you are not redirected in a few seconds, please refresh the
								page.
							</div>
						)}

						<div className="flex flex-col gap-2">
							<Search
								icon={<Asterisk size={14} className="opacity-60" />}
								placeholder="Referral Code"
								search={referrerAddress}
								handleSearch={setReferralAddress}
								clear
							/>
							<div className="flex flex-row items-center gap-2">
								<Button
									variant="secondary"
									className="w-max py-4"
									onClick={handleRequestAccess}
									disabled={requestAccess.isLoading}
								>
									{requestAccess.isLoading ? "Requesting..." : "Get Access"}
								</Button>

								<button
									className={cn(
										"w-full rounded-lg py-4 font-bold",
										referrerAddress && isLoading === false
											? "cursor-pointer bg-gradient-to-tr from-plug-green to-plug-yellow text-white"
											: "border-[1px] border-plug-green bg-white text-plug-green"
									)}
									onClick={() => referrerAddress && mutate({ referrerAddress })}
									disabled={isLoading || !referrerAddress}
								>
									{isLoading ? "Submitting..." : "Submit Code"}
								</button>
							</div>

							{isError && error && (
								<p className="mx-auto max-w-[320px] text-center font-bold text-red-500">
									{error.message}
								</p>
							)}
						</div>
					</div>
				</div>
			</div>
		</div>
	)
}
