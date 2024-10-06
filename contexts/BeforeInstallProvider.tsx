import { createContext, useContext, useSyncExternalStore } from "react"
import { isBrowser } from "react-device-detect"

/**
 * This interface is experimental.
 *
 * https://developer.mozilla.org/en-US/docs/Web/API/BeforeInstallPromptEvent/BeforeInstallPromptEvent
 *
 */
interface BeforeInstallPromptEvent extends Event {
	/**
	 * Allows a developer to show the install prompt at a time of their own choosing.
	 * This method returns a Promise.
	 */
	prompt(): Promise<UserChoice>
}

type UserChoice = { outcome: "accepted" | "dismissed"; platform: string }

type AppInstallManager =
	| {
			prompt: null
			isNativePromptAvailable: false
	  }
	| {
			prompt: () => Promise<UserChoice>
			isNativePromptAvailable: true
	  }

/**
 * This is the most reliable way (I found) to work with the `BeforeInstallPromptEvent` on the browser.
 *
 * We will implement what I call the 'external store pattern'.
 */
const appInstallManagerStore: AppInstallManager = {
	prompt: null,
	isNativePromptAvailable: false
}

const AppInstallManagerContext = createContext<AppInstallManager | null>(null)

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

	return context
}

function checkIsArcBrowser() {
	// https://webmasters.stackexchange.com/a/142231/138612
	return window.getComputedStyle(document.documentElement).getPropertyValue("--arc-palette-title") ? true : false
}

function subscribeToLoad(callback: () => void) {
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

	return <AppInstallManagerContext.Provider value={appInstallManager}>{children}</AppInstallManagerContext.Provider>
}
