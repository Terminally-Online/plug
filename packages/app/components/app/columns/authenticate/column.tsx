import { useSession } from "next-auth/react"
import { FC } from "react"


import { useAtomValue } from "jotai"

import { Callout } from "@/components/app/utils/callout"
import { Button } from "@/components/shared/buttons/button"
import {
	formatAddress,
} from "@/lib"
import { useAccount } from "@/lib/hooks/account/useAccount"
import { useAuthenticate } from "@/lib/hooks/account/useAuthenticate"
import { authenticationAtom } from "@/state/authentication"
import { columnByIndexAtom, useColumnActions } from "@/state/columns"
import { ConnectorList } from "@/components/app/columns/authenticate/connector/list"

export const ColumnAuthenticate: FC<{ index: number }> = ({ index }) => {
	const { data: session } = useSession()
	const account = useAccount()

	const { authenticate, failureReason, isLoading } = useAuthenticate()
	const { navigate } = useColumnActions()

	const column = useAtomValue(columnByIndexAtom(index))
	const authentication = useAtomValue(authenticationAtom)

	const handleAuthenticate = () => {
		authenticate(undefined, {
			onSuccess: () => navigate({ index, from: column?.from }),
			onError: error => console.error(error)
		})
	}

	return (
		<div className="flex h-full flex-col items-center justify-center text-center">
			{authentication.isLoading && (
				<Callout
					title="Authentication loading."
					description="We are loading all the state of your account. One moment please."
				/>
			)}

			{session?.user.id !== account.address &&
				account.address &&
				isLoading === false &&
				authentication.isLoading === false && (
					<Callout
						title={failureReason ? "Signature error." : "Prove ownership."}
						description={
							failureReason
								? "An internal error was received while signing the message. " +
								failureReason.message.split("Details:")[1].split("Details:")[0].trim()
								: `Please sign the message to prove your ownership of ${formatAddress(account.address)}.`
						}
					>
						<Button className="mt-2" sizing="sm" onClick={handleAuthenticate}>
							Sign Message
						</Button>
					</Callout>
				)}

			{account.address && isLoading && (
				<Callout
					title="Proving ownership."
					description={`Completing the signing process to prove ownership of ${formatAddress(account.address)}`}
				>
					<Button className="mt-2" sizing="sm" disabled>
						Signing...
					</Button>
				</Callout>
			)}

			{account.address === undefined && <ConnectorList index={index} from={column?.from} />}
		</div>
	)
}
