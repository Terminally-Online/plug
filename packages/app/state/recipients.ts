import { useCallback, useMemo } from "react"

import { useAtom } from "jotai"

import { atomWithStorage } from "jotai/utils"
import { useAccount } from "@/lib/hooks/account/useAccount"

const recipientAtom = atomWithStorage<string[]>("plug.recipients", [])

export const useRecipients = (recipient: string) => {
	const { address } = useAccount()

	const [recipients, setRecipients] = useAtom(recipientAtom)

	const handleRecent = useCallback(
		(recipient: string) => {
			if (recipient === "") return
			setRecipients(prev => {
				const newRecipients = [recipient, ...prev.filter(r => r !== recipient && r !== address)]
				return newRecipients.slice(0, 10)
			})
		},
		[setRecipients]
	)

	const minimizedRecipients = useMemo(() => recipients.filter(r => r !== recipient), [recipients, recipient])

	return { recipients: minimizedRecipients, handleRecent }
}
