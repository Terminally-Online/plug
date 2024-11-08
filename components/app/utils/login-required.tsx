import { useSession } from "next-auth/react"
import { FC } from "react"

import { Button } from "@/components/shared"
import { useSidebar } from "@/state"

export const LoginRequired: FC = () => {
  const { data: session } = useSession()
  const { handleSidebar } = useSidebar()

  const isVisible = !session?.user.id?.startsWith("0x")

  if (!isVisible) return null

  const handleLogin = () => {
    handleSidebar("authenticating")
  }

  return (
    <div className="h-full w-full flex items-center justify-center bg-white">
      <div className="flex flex-col items-center gap-8 max-w-md w-full px-4">
        <h1 className="text-2xl font-bold">Welcome to Plug.</h1>
        
        <div className="w-full space-y-6">
          <div className="flex flex-col gap-2">
            <p className="text-sm font-medium text-center opacity-60">
              Connect your wallet to get started
            </p>

            <Button
              className="w-full"
              onClick={handleLogin}
              variant="primary"
            >
              Connect
            </Button>
          </div>
        </div>
      </div>
    </div>
  )
}