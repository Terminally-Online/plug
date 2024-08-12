export const getTextColor = (backgroundColor: string) => {
	let r: number, g: number, b: number

	if (backgroundColor.startsWith("rgb")) {
		const matches = backgroundColor.match(/\d+/g)
		if (matches && matches.length === 3) {
			;[r, g, b] = matches.map(Number)
		} else {
			throw new Error("Invalid RGB color format")
		}
	} else {
		const hex = backgroundColor.replace(/^#/, "")

		if (hex.length === 3) {
			// Convert shorthand hex to full form
			r = parseInt(hex[0] + hex[0], 16)
			g = parseInt(hex[1] + hex[1], 16)
			b = parseInt(hex[2] + hex[2], 16)
		} else if (hex.length === 6) {
			r = parseInt(hex.slice(0, 2), 16)
			g = parseInt(hex.slice(2, 4), 16)
			b = parseInt(hex.slice(4, 6), 16)
		} else {
			return "#FFFFFF"
		}
	}

	r = Math.min(255, Math.max(0, r))
	g = Math.min(255, Math.max(0, g))
	b = Math.min(255, Math.max(0, b))

	const luminance = (0.299 * r + 0.587 * g + 0.114 * b) / 255

	return luminance > 0.5 ? "#000000" : "#FFFFFF"
}
