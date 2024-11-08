import { useSession } from "next-auth/react"
import { FC, useState } from "react"

import { Button } from "@/components/shared"
import { useClipboard } from "@/lib"
import { api } from "@/server/client"
import { useSocket } from "@/state"

export const ReferralRequired: FC = () => {
  const { data: session } = useSession()
  const { socket } = useSocket()
  const { copied, handleCopied } = useClipboard(session?.user.id ?? "")
  const [referralAddress, setReferralAddress] = useState("")
  const [error, setError] = useState<string>()
  const [success, setSuccess] = useState(false)

  const submitReferral = api.socket.submitReferral.useMutation({
    onError: (error) => setError(error.message),
    onSuccess: () => {
      setSuccess(true)
      setError(undefined)
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
    const tweetText = encodeURIComponent(
      `Hey @plug_hq! I'd love to try out Plug.\n\nMy address: ${session?.user.id}`
    )
    window.open(`https://twitter.com/intent/tweet?text=${tweetText}`, '_blank')
    await requestAccess.mutateAsync()
  }

  return (
    <div className="flex h-full w-full items-center justify-center bg-white">
      <div className="flex max-w-md w-full flex-col items-center gap-8 px-4">
        <h1 className="text-2xl font-bold">You need access to Plug.</h1>
        
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
              Enter a referral address from an existing user
            </p>
            <input
              type="text"
              placeholder="0x... or ENS name"
              value={referralAddress}
              onChange={(e) => setReferralAddress(e.target.value)}
              className="w-full rounded-lg border border-grayscale-100 bg-white px-4 py-2 font-mono text-sm"
            />
            <Button
              className="w-full"
              onClick={handleSubmitReferral}
              disabled={submitReferral.isLoading || !referralAddress}
            >
              {submitReferral.isLoading ? "Submitting..." : "Submit Referral"}
            </Button>
          </div>

          <div className="flex flex-col gap-2">
            <p className="text-center text-sm font-medium opacity-60">
              Or share your address to get referred
            </p>
            <Button 
              className="w-full font-mono"
              onClick={handleCopied}
              variant="secondary"
            >
              {copied ? "Copied!" : session?.user.id}
            </Button>

            <Button
              className="w-full"
              onClick={handleRequestAccess}
              disabled={requestAccess.isLoading}
              variant="primary"
            >
              {requestAccess.isLoading ? "Requesting..." : "Request Access on Twitter"}
            </Button>
          </div>
        </div>
      </div>
    </div>
  )
}