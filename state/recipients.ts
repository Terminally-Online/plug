import { useCallback, useMemo } from "react"

import { useAtom } from "jotai"

import { atomWithStorage } from "jotai/utils"

const recipientAtom = atomWithStorage<string[]>("recipients", [])

export const useRecipients = (recipient: string) => {
	const [recipients, setRecipients] = useAtom(recipientAtom)

	const handleRecent = useCallback(
		(recipient: string) => {
			setRecipients(prev => {
				const newRecipients = [recipient, ...prev.filter(r => r !== recipient)]
				return newRecipients.slice(0, 6)
			})
		},
		[setRecipients]
	)

	const minimizedRecipients = useMemo(() => recipients.filter(r => r !== recipient), [recipients, recipient])

	return { recipients: minimizedRecipients, handleRecent }
}
