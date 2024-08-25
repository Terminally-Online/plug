import { useAccount, useSignMessage } from "wagmi"

import { AuthButton } from "@/components/shared"

export const ConsoleAuthenticate = () => {
	const { isConnected, isConnecting } = useAccount()
	const { signMessageAsync, isLoading: isSigning } = useSignMessage()

	return (
		<div className="flex h-full flex-col items-center justify-center px-4 text-center">
			{isConnecting === false && isConnected === false && (
				<div className="flex flex-col gap-2">
					<p className="font-bold">Connect your wallet to continue.</p>
					<p className="max-w-[320px] text-sm opacity-40">
						You can view everything here anonymously, but to perform actions, you will need to authenticate.
					</p>
				</div>
			)}

			{isConnecting && isConnected === false && (
				<div className="flex flex-col gap-2">
					<p className="font-bold">Confirm connection in your wallet.</p>
					<p className="max-w-[320px] text-sm opacity-40">
						You can view everything here anonymously, but to perform actions, you will need to authenticate.
					</p>
				</div>
			)}

			{isConnected && isSigning === false && (
				<div className="flex flex-col gap-2">
					<p className="font-bold">Sign the message to confirm.</p>
					<p className="max-w-[320px] text-sm opacity-40">
						We ask you to sign a message to prove the ownership of the wallet you are connected so that no one else can
						impersonate you.
					</p>
				</div>
			)}

			<AuthButton className="mt-8 w-max" />
		</div>
	)
}
