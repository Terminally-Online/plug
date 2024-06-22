import { useCallback, useEffect, useState } from "react"

export const useClipboard = (text: string) => {
	const [copied, setCopied] = useState<boolean>(false)

	useEffect(() => {
		let timeout: NodeJS.Timeout | null = null

		if (copied) {
			navigator.clipboard.writeText(text)
			timeout = setTimeout(() => setCopied(false), 2000)
		}

		return () => (timeout ? clearTimeout(timeout) : undefined)
	}, [text, copied])

	return { copied, handleCopied: useCallback(() => setCopied(true), []) }
}

export default useClipboard
