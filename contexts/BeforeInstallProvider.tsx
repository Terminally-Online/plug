import Image from "next/image"
import { createContext, useContext, useSyncExternalStore } from "react"
import { isAndroid, isBrowser, isChrome, isIOS, isMacOs, isSafari, isWindows, osName } from "react-device-detect"

import { ChevronRight } from "lucide-react"

type Instructions = Record<
	"macOS" | "iOS" | "android" | "linux" | "windows" | "safari" | "chrome" | "arc",
	(string | JSX.Element)[]
>

const instructions = {
	android: ["Open this page in Chrome"],
	windows: ["Open this page in Edge"],
	linux: ["Open this page in Chromium or Chrome"],
	iOS: [
		"Open this page in Safari",
		<span key="ios-share" className="flex items-center">
			Click the Share button
			<Image
				className="mx-1 inline-block h-4 w-4 opacity-60"
				src="/misc/ios-share-icon.svg"
				alt="apple-share-icon"
				width={48}
				height={48}
			/>
			in the Safari toolbar, then choose Add to home screen
		</span>,
		"Type the name that you want to use for the web app, then click Add."
	],
	macOS: [
		"Open this page in Safari or Chrome. Arc does not support native apps on macOS.",
		<span key="macos-share" className="items-center">
			From the menu bar, choose File <ChevronRight size={14} className="mb-1 inline-block h-4 w-4" /> Add to Dock.
			Or click the Share button
			<Image
				className="mx-1 mb-1 inline-block h-4 w-4 opacity-60"
				src="/misc/ios-share-icon.svg"
				alt="apple-share-icon"
				width={48}
				height={48}
			/>
			in the Safari toolbar, then choose Add to Dock
		</span>,
		"Type the name that you want to use for the web app, then click Add"
	],
	arc: ["Open this page in Safari or Chrome. Arc does not support native apps on macOS."],
	safari: [
		<span key="safari-share" className="items-center">
			Click the Share button
			<Image
				className="mx-1 mb-1 inline-block h-4 w-4 opacity-60"
				src="/misc/ios-share-icon.svg"
				alt="apple-share-icon"
				width={48}
				height={48}
			/>
			in the Safari toolbar, then choose Add to Dock
		</span>,
		"Type the name that you want to use for the web app, then click Add"
	],
	chrome: [
		"Click the three-dot menu in the top right corner to open the options menu.",
		"Select 'More tools' > 'Create shortcut'. to open the shortcut creation menu.",
		"Finish by checking 'Open as window' and clicking 'Create' to install the app."
	]
} satisfies Instructions

const getInstructions = (isArcBrowser: boolean) => {
	if (isMacOs) {
		if (isArcBrowser) return instructions.arc
		if (isSafari) return instructions.safari
		if (isChrome) return instructions.chrome
		return instructions.macOS
	}

	if (isIOS) return instructions.iOS
	if (isAndroid) return instructions.android
	if (osName === "Linux") return instructions.linux
	if (isWindows) return instructions.windows

	return []
}

interface BeforeInstallPromptEvent extends Event {
	prompt(): Promise<UserChoice>
}

type UserChoice = { outcome: "accepted" | "dismissed"; platform: string }

type AppInstallManager = {
	prompt?: (() => Promise<UserChoice>) | null
	isNativePromptAvailable?: boolean
	isArcBrowser: boolean
	instructions: Array<string | JSX.Element>
}

const appInstallManagerStore: AppInstallManager = {
	prompt: null,
	isNativePromptAvailable: false,
	isArcBrowser: false,
	instructions: []
}

const AppInstallManagerContext = createContext<AppInstallManager | null>(null)

const checkIsArcBrowser = () => {
	return window.getComputedStyle(document.documentElement).getPropertyValue("--arc-palette-title") ? true : false
}

const subscribeToLoad = (callback: () => void) => {
	async function delayCallback() {
		// delay the callback to be sure everything is loaded
		await new Promise(resolve => setTimeout(resolve, 1000))
		callback()
	}

	window.addEventListener("load", delayCallback)

	return () => {
		window.removeEventListener("load", delayCallback)
	}
}

function subscribeToBeforeInstallPrompt(callback: () => void) {
	function saveInstallPrompt(event: Event) {
		event.preventDefault()

		appInstallManagerStore.prompt = (event as BeforeInstallPromptEvent).prompt.bind(event)

		appInstallManagerStore.isNativePromptAvailable = true

		callback()
	}

	window.addEventListener("beforeinstallprompt", saveInstallPrompt)

	return () => {
		window.removeEventListener("beforeinstallprompt", saveInstallPrompt)
	}
}

/**
 * Use it in root.tsx
 * Wrap `<Outlet />` with `<AppInstallManagerProvider>`
 *
 */
export const BeforeInstallProvider = ({ children }: { children: React.ReactElement }) => {
	const appInstallManager = useSyncExternalStore(
		subscribeToBeforeInstallPrompt,
		() => appInstallManagerStore,
		() => null
	)

	// 🚨 Some chrome based browsers don't support prompt to install even if they support the event.
	// [21/10/2023] : ❌ Arc Browser
	const isArcBrowser = useSyncExternalStore(
		subscribeToLoad,
		() => checkIsArcBrowser(),
		() => false
	)

	if (appInstallManager) {
		if (isArcBrowser) {
			appInstallManager.prompt = null
			appInstallManager.isNativePromptAvailable = false
		}
	}

	const instructions = getInstructions(isArcBrowser)

	return (
		<AppInstallManagerContext.Provider value={{ ...appInstallManager, isArcBrowser, instructions }}>
			{children}
		</AppInstallManagerContext.Provider>
	)
}

/**
 * Use `BeforeInstallPromptEvent.prompt` to prompt the user to install the PWA.
 * If the PWA is already installed by the current browser, `available` will always be false and `prompt` will always be null.
 *
 * [21/10/2023]
 *
 * ❌ On Safari and Firefox, `available` will always be false and `prompt` will always be null.
 * These the browser does not support prompt to install, `beforeinstallprompt` event is not fired.
 * https://developer.mozilla.org/en-US/docs/Web/API/BeforeInstallPromptEvent#browser_compatibility
 *
 * 🤷‍♂️ Arc Browser, even if it's based on Chromium, doesn't support prompt to install.
 * `prompt` never moves from pending to resolved.
 *
 * @returns the BeforeInstallPromptEvent if available
 */
export const useBeforeInstall = () => {
	const context = useContext(AppInstallManagerContext)

	if (isBrowser && context === undefined) {
		throw new Error(`useAppInstallManager must be used within a XProvider.`)
	}

	return context ?? appInstallManagerStore
}
