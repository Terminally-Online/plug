import { useSearchParams } from "next/navigation"
import { FC, useEffect, useState } from "react"

import { Asterisk } from "lucide-react"

import { Search } from "@/components/app/inputs/search"
import { Button } from "@/components/shared/buttons/button"
import { cn } from "@/lib"
import { api } from "@/server/client"
import { useAccount } from "@/lib/hooks/account/useAccount"

const TWEET_TEMPLATES = [
	`Just discovered @onplug_io - a game-changing platform for automated trading. Can't wait to get access!`,
	`Excited about @onplug_io's smart automation tools for trading. Looking forward to joining the community!`,
	`This new automated trading platform looks incredible. Would love to be part of the early access group for @onplug_io!`,
	`The future of trading is here with @onplug_io. Ready to experience next-level automation!`,
	`Heard amazing things about @onplug_io's trading automation. Hope to get access soon!`
]

const UUID_REGEX = /^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/i

export const ReferralRequired: FC = () => {
	const searchParams = useSearchParams()
	const account = useAccount()

	const { mutate, error, isPending, isError, isSuccess } = api.socket.referral.submit.useMutation({
		onSuccess: () => {
			// Refetch socket data which will trigger reactive updates
			// socketQuery.refetch()
		}
	})
	const requestAccess = api.socket.referral.request.useMutation()

	const [referralCode, setReferralAddress] = useState("")

	const handleRequestAccess = async () => {
		const randomTweet = TWEET_TEMPLATES[Math.floor(Math.random() * TWEET_TEMPLATES.length)]
		const tweetText = encodeURIComponent(randomTweet + `\n\nMy address: ${account.address}`)
		window.open(`https://twitter.com/intent/tweet?text=${tweetText}`, "_blank")
		await requestAccess.mutateAsync()
	}

	useEffect(() => {
		const rfid = searchParams.get("rfid")

		if (rfid === null) return

		setReferralAddress(rfid)
	}, [searchParams])

	useEffect(() => {
		if (UUID_REGEX.test(referralCode) === false) return

		mutate(referralCode)
	}, [referralCode, mutate])

	return (
		<div className="flex w-full flex-col items-start justify-between bg-white p-2">
			<div className="flex h-full w-full items-center justify-center md:rounded-lg md:border-[1px] md:border-plug-green/10">
				<div className="flex max-w-[480px] flex-col items-center">
					<h1 className="text-2xl font-bold">Get Access to Plug.</h1>
					<p className="mb-8 text-center font-bold text-black/40">
						Enter the referral code you received to get started or request one by tagging{" "}
						<button onClick={handleRequestAccess} className="text-plug-green hover:opacity-80">
							@onplug_io
						</button>{" "}
						on Twitter.
					</p>

					<div className="w-full space-y-6">
						{isSuccess && (
							<div className="w-full rounded-lg border border-plug-yellow bg-plug-yellow p-4 text-sm text-plug-yellow">
								Successfully approved! If you are not redirected in a few seconds, please refresh the
								page.
							</div>
						)}

						<div className="flex flex-col gap-2">
							<Search
								icon={<Asterisk size={14} className="opacity-60" />}
								placeholder="Referral Code"
								search={referralCode}
								handleSearch={setReferralAddress}
								clear
							/>
							<div className="flex flex-row items-center gap-2">
								<Button
									variant="secondary"
									className="w-max py-4"
									onClick={handleRequestAccess}
									disabled={requestAccess.isPending}
								>
									{requestAccess.isPending ? "Requesting..." : "Get Access"}
								</Button>

								<button
									className={cn(
										"w-full rounded-lg py-4 font-bold",
										referralCode && isPending === false
											? "cursor-pointer bg-plug-yellow text-plug-green"
											: "border-[1px] border-plug-green bg-white text-plug-green"
									)}
									onClick={() => referralCode && mutate(referralCode)}
									disabled={isPending || !referralCode}
								>
									{isPending ? "Submitting..." : "Submit Code"}
								</button>
							</div>

							{isError && error && (
								<p className="mx-auto mt-2 max-w-[360px] text-center text-sm font-bold text-red-500">
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
