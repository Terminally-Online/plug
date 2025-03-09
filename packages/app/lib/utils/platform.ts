export const isMobileWeb: boolean =
	// https://stackoverflow.com/a/29509267
	typeof navigator !== "undefined" && /iPhone|iPad|iPod|Android|Mobi/i.test(navigator.userAgent)

export const isTouchable =
	typeof window !== "undefined" &&
	typeof navigator !== "undefined" &&
	("ontouchstart" in window || navigator.maxTouchPoints > 0)
