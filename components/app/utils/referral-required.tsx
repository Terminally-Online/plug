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
    onError: (error) => setError(error.message),
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
  const isVisible = Boolean(
    isAuthenticated &&
    socket &&
    !socket.identity?.approvedAt
  )

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
      `Need that @onplug_io access to start automating my onchain activities 🥵\n\nMy address: ${session?.user.id}`,
      `Hey @onplug_io! Let's streamline my crypto life with automated onchain actions ⚡\n\nMy address: ${session?.user.id}`
    ]

    const randomTweet = tweetTemplates[Math.floor(Math.random() * tweetTemplates.length)]
    const tweetText = encodeURIComponent(randomTweet)
    window.open(`https://twitter.com/intent/tweet?text=${tweetText}`, '_blank')
    await requestAccess.mutateAsync()
  }

  return (
    <div className="flex w-full items-start justify-between flex-col bg-white">
      <div className="w-full h-full flex justify-center pt-[30vh] border-[1px] border-grayscale-100 m-2 rounded-lg">
        <div className="flex max-w-md w-full flex-col px-4">
          <div className="flex flex-col items-center">
            <h1 className="text-2xl font-bold">Get Access to Plug.</h1>
            <p className="font-bold text-black/40 mb-8 text-center">Enter the referral code you received to get started or request one by tagging <button
              onClick={handleRequestAccess}
              style={{
                ...greenGradientStyle
              }}>@onplug_io</button> on Twitter.</p>

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
                <Search
                  icon={<SearchIcon size={14} className="opacity-60" />}
                  placeholder="0x... or ENS name"
                  search={referralAddress}
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

                  <Button
                    className={cn("w-full py-4")}
                    // variant={submitReferral.isLoading || !referralAddress ? "disabled" : "primary"}
                    onClick={handleSubmitReferral}
                    disabled={submitReferral.isLoading || !referralAddress}
                  >
                    {submitReferral.isLoading ? "Submitting..." : "Submit Code"}
                  </Button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}
