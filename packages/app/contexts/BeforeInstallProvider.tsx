import Image from "next/image"
import React, {
    createContext,
    FC,
    ReactElement,
    useContext,
    useMemo,
    useSyncExternalStore,
    type JSX,
} from "react";
import { isAndroid, isBrowser, isChrome, isIOS, isMacOs, isSafari, isWindows, osName } from "react-device-detect"

type OSType = "macOS" | "iOS" | "android" | "linux" | "windows" | "safari" | "chrome" | "arc"
type Instructions = Record<OSType, (string | JSX.Element)[]>

const SHARE_ICON_PATH = "/misc/ios-share-icon.svg"

const instructions: Instructions = {
	android: [
		"Open this page in Google Chrome. Click the 'Continue' button if you already are.",
		"Tap the three-dot menu in the top right corner and select 'Add to Home screen'."
	],
	windows: [
		"Open this page in Microsoft Edge. Click the 'Continue' button if you already are.",
		"Click the '...' menu in the top right corner of your browser.",
		"Finally, click select 'Apps' > 'Install this site as an app'."
	],
	linux: [
		"Open this page in Google Chrome. Click the 'Continue' button if you already are.",
		"Click the three-dot menu in the top right corner of your browser.",
		"Finally, select 'More tools' > 'Create shortcut'. Check 'Open as window' and click 'Create'."
	],
	iOS: [
		"Open this page in Safari.",
		<span key="ios-share" className="flex items-center">
			Tap the Share button
			<Image
				className="mx-1 inline-block h-4 w-4 opacity-60"
				src={SHARE_ICON_PATH}
				alt="apple-share-icon"
				width={48}
				height={48}
			/>
			at the bottom of the screen, then select 'Add to Home Screen'.
		</span>
	],
	macOS: [
		"Open this page in Safari or Google Chrome.",
		<span key="macos-share" className="items-center">
			Click the Share button
			<Image
				className="mx-1 mb-1 inline-block h-4 w-4 opacity-60"
				src={SHARE_ICON_PATH}
				alt="apple-share-icon"
				width={48}
				height={48}
			/>
			in the toolbar and select 'Add to Dock'.
		</span>
	],
	arc: [
		"Arc browser does not support installing web apps. Please use Safari or Chrome on macOS instead.",
		"We have submit a feature request to the Arc team to have this functionality enabled."
	],
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
			in the toolbar and select 'Add to Dock'.
		</span>
	],
	chrome: [
		"Click the three-dot menu in the top right corner of your browser",
		"Click 'More tools' > 'Create shortcut'. Check 'Open as window' and click 'Create' to install the app."
	]
} satisfies Instructions

const getInstructions = (isArcBrowser: boolean): (string | JSX.Element)[] => {
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

interface AppInstallManager {
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

const checkIsArcBrowser = (): boolean => {
	return !!window.getComputedStyle(document.documentElement).getPropertyValue("--arc-palette-title")
}

const subscribeToLoad = (callback: () => void): (() => void) => {
	const delayCallback = async (): Promise<void> => {
		await new Promise(resolve => setTimeout(resolve, 1000))
		callback()
	}

	window.addEventListener("load", delayCallback)

	return () => {
		window.removeEventListener("load", delayCallback)
	}
}

const subscribeToBeforeInstallPrompt = (callback: () => void): (() => void) => {
	const saveInstallPrompt = (event: Event): void => {
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

export const BeforeInstallProvider: FC<{ children: ReactElement<any> }> = ({ children }) => {
	const appInstallManager = useSyncExternalStore(
		subscribeToBeforeInstallPrompt,
		() => appInstallManagerStore,
		() => null
	)

	const isArcBrowser = useSyncExternalStore(subscribeToLoad, checkIsArcBrowser, () => false)

	const memoizedAppInstallManager = useMemo(() => {
		if (appInstallManager && isArcBrowser) {
			return {
				...appInstallManager,
				prompt: null,
				isNativePromptAvailable: false
			}
		}
		return appInstallManager
	}, [appInstallManager, isArcBrowser])

	const instructions = useMemo(() => getInstructions(isArcBrowser), [getInstructions, isArcBrowser])

	const contextValue = useMemo(
		() => ({
			...memoizedAppInstallManager,
			isArcBrowser,
			instructions
		}),
		[memoizedAppInstallManager, isArcBrowser, instructions]
	)

	return <AppInstallManagerContext.Provider value={contextValue}>{children}</AppInstallManagerContext.Provider>
}

export const useBeforeInstall = (): AppInstallManager => {
	const context = useContext(AppInstallManagerContext)

	if (isBrowser && context === undefined) {
		throw new Error("useBeforeInstall must be used within a BeforeInstallProvider.")
	}

	return context ?? appInstallManagerStore
}
