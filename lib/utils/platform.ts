/**
 * Note, this file has one counterpart:
 *
 *   - index.native.ts
 *
 * Be sure to keep it in sync! They should export the exact same set of constants
 * to avoid type mis-matches.
 *
 * Also - we need to check globals exist (document, navigator), because this
 * runs in a service worker for extension.
 *
 */

// Platform
export const isAndroid: boolean = false
export const isIOS: boolean = false
// see: https://stackoverflow.com/a/14301832

export const isWeb: boolean = true
export const isMobileWeb: boolean =
	// https://stackoverflow.com/a/29509267
	typeof navigator !== "undefined" && /iPhone|iPad|iPod|Android|Mobi/i.test(navigator.userAgent)

// Operating System
// via https://stackoverflow.com/questions/9038625/detect-if-device-is-ios
export const isWebIOS: boolean =
	typeof document !== "undefined" &&
	typeof navigator !== "undefined" &&
	(["iPad Simulator", "iPhone Simulator", "iPod Simulator", "iPad", "iPhone"].includes(navigator.platform) ||
		// iPad on iOS 13 detection
		(navigator.userAgent.includes("Mac") && "ontouchend" in document))

// via https://stackoverflow.com/questions/6031412/detect-android-phone-via-javascript-jquery
export const isWebAndroid: boolean =
	typeof navigator !== "undefined" && navigator.userAgent.toLowerCase().includes("android")

// Capability
export const isTouchable =
	typeof window !== "undefined" &&
	typeof navigator !== "undefined" &&
	("ontouchstart" in window || navigator.maxTouchPoints > 0)

// Browser
export const isChrome: boolean = typeof navigator !== "undefined" && /Chrome/.test(navigator.userAgent || "")
export const isSafari: boolean = typeof navigator !== "undefined" && /Safari/.test(navigator.userAgent || "")
export const isMobileWebSafari: boolean = isTouchable && isSafari
export const isMobileWebAndroid: boolean = isTouchable && isWebAndroid

// App
export const isExtension: boolean = process.env.IS_UNISWAP_EXTENSION === "true"
export const isMobileApp: boolean = false
export const isInterface: boolean = process.env.REACT_APP_IS_UNISWAP_INTERFACE === "true"

// iFrame
export const isIFramed = (redirect = false): boolean => {
	try {
		if (location.ancestorOrigins !== undefined) {
			// Does not exist in IE and firefox.
			// See https://developer.mozilla.org/en-US/docs/Web/API/Location/ancestorOrigins for how this works
			if (location.ancestorOrigins.length > 0) {
				if (redirect && top) {
					top.location = self.location
				}
				return true
			}
		}
		if (self !== top) {
			// For IE and Firefox
			// See https://cheatsheetseries.owasp.org/cheatsheets/Clickjacking_Defense_Cheat_Sheet.html for how this works
			if (redirect && top) {
				top.location = self.location
			}
			return true
		}
	} catch {
		// this should never be called, but just in we are being iframed in an old browser where an attacker can overwrite these variables
		if (redirect) {
			throw new Error("isIFramed: should redirect, but unable to determine if framed")
		}
		return true // fail closed and say we are being iframed if we can't determine if we are being iframed
	}

	return false
}
