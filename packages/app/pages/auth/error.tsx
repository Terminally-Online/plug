import { useRouter } from "next/router"
import { useEffect, useState } from "react"

import { Home, Loader } from "lucide-react"

import { Button } from "@/components/shared/buttons/button"

export default function AuthError() {
	const {
		push,
		query: { error }
	} = useRouter()

	const errorMessages: Record<string, string> = {
		Configuration: "There is a problem with the server configuration.",
		AccessDenied: "Access denied. You don't have permission to view this resource.",
		Verification: "The verification link may have expired or already been used.",
		SignIn: "Try signing in with a different account.",
		OAuthSignin: "Error in constructing an authorization URL.",
		OAuthCallback: "Error in handling the response from an OAuth provider.",
		OAuthCreateAccount: "Error in creating a user in the database.",
		EmailCreateAccount: "Error in creating a user in the database.",
		Callback: "Error in the OAuth callback handler.",
		Default: "An authentication error occurred. Please try signing in again."
	}

	const message = error ? errorMessages[error as string] || errorMessages.Default : errorMessages.Default

	return (
		<main className="flex min-h-screen flex-col items-center justify-center p-4">
			<div className="w-full max-w-md rounded-lg border-[2px] border-plug-green/10 p-6 shadow-sm">
				<div className="mb-8 space-y-2 text-center">
					<h1 className="text-xl font-semibold tracking-tight">
						You've encounted an: <span className="opacity-60">Authentication Error</span>. This has been
						relayed to our team.
					</h1>
					<p className="mx-auto max-w-[280px] text-sm font-bold opacity-60">{message}</p>
				</div>
				<div className="flex flex-row items-center justify-center gap-2">
					<Button
						variant="primary"
						onClick={() => push("/app")}
						className="flex w-full flex-row items-center justify-center gap-2"
					>
						<Home size={16} className="opacity-60" />
						Go Back
					</Button>
				</div>
			</div>
		</main>
	)
}
