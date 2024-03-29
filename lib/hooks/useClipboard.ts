import { useCallback, useState } from "react"

export const useClipboard = () => {
	const [isCopied, setIsCopied] = useState<boolean>(false)

	const copy = useCallback(async (text: string): Promise<void> => {
		try {
			await navigator.clipboard.writeText(text)
			setIsCopied(true)
			setTimeout(() => setIsCopied(false), 2000)
		} catch (err) {
			console.error("Failed to copy: ", err)
			setIsCopied(false)
		}
	}, [])

	return { copy, isCopied }
}

export default useClipboard
