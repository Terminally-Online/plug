import React, { useEffect, useState } from "react"

import { isWebAndroid, isWebIOS } from "@/lib/utils/platform"

declare global {
	interface WindowEventMap {
		beforeinstallprompt: Event
	}
}

export const PwaPrompt: React.FC = () => {
	const [showPrompt, setShowPrompt] = useState(false)
	const [deferredPrompt, setDeferredPrompt] = useState<any>(null)

	useEffect(() => {
		const handleBeforeInstallPrompt = (e: Event) => {
			e.preventDefault()
			setDeferredPrompt(e)
			setShowPrompt(true)
		}

		window.addEventListener("beforeinstallprompt", handleBeforeInstallPrompt as EventListener)

		return () => {
			window.removeEventListener("beforeinstallprompt", handleBeforeInstallPrompt)
		}
	}, [])

	const handleInstall = () => {
		if (deferredPrompt) {
			deferredPrompt.prompt()
			deferredPrompt.userChoice.then((choiceResult: { outcome: string }) => {
				if (choiceResult.outcome === "accepted") {
					console.log("User accepted the install prompt")
				} else {
					console.log("User dismissed the install prompt")
				}
				setDeferredPrompt(null)
			})
		}
	}

	if (!showPrompt) return null

	if (isWebIOS) {
		return (
			<div>
				<p>To install the app, tap the share button and then &quot;Add to Home Screen&quot;</p>
			</div>
		)
	}

	if (isWebAndroid) {
		return (
			<div>
				<p>Install Plug for a better experience</p>
				<button onClick={handleInstall}>Install</button>
			</div>
		)
	}

	return null
}
