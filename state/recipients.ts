import { useCallback, useMemo } from "react"

import { useAtom } from "jotai"

import { useConnect } from "@/lib"

import { atomWithStorage } from "jotai/utils"

const recipientAtom = atomWithStorage<string[]>("recipients", [])

export const useRecipients = (recipient: string) => {
	const { account } = useConnect()

	const [recipients, setRecipients] = useAtom(recipientAtom)

	const handleRecent = useCallback(
		(recipient: string) => {
			if (recipient === "") return
			setRecipients(prev => {
				const newRecipients = [recipient, ...prev.filter(r => r !== recipient && r !== account.address)]
				return newRecipients.slice(0, 10)
			})
		},
		[account, setRecipients]
	)

	const minimizedRecipients = useMemo(() => recipients.filter(r => r !== recipient), [recipients, recipient])

	return { recipients: minimizedRecipients, handleRecent }
}
