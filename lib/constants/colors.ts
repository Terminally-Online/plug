export const colors = {
	red: "#FF5154",
	orange: "#FF9900",
	yellow: "#EEB902",
	lightgreen: "#43D262",
	green: "#28BC7E",
	cyan: "#3FCAB1",
	blue: "#2D7DD2",
	magenta: "#D345F6",
	purple: "#9665FF"
} as const

export const tagColors = {
	red: "rgba(197,0,3,0.5)",
	orange: "rgba(162,97,0,0.5)",
	yellow: "rgba(178,128,0,0.5)",
	lightgreen: "rgba(0,153,34,0.5)",
	green: "rgba(0,110,64,0.5)",
	cyan: "rgba(0,112,92,0.5)",
	blue: "rgba(0,74,152,0.5)",
	magenta: "rgba(134,0,167,0.5)",
	purple: "rgba(76,9,218,0.5)"
} as const

export const lightenHexColor = (hex: string, amount = 0.4) => {
	amount = Math.max(0, Math.min(1, amount))

	let r = parseInt(hex.slice(1, 3), 16)
	let g = parseInt(hex.slice(3, 5), 16)
	let b = parseInt(hex.slice(5, 7), 16)

	r += Math.floor((255 - r) * amount)
	g += Math.floor((255 - g) * amount)
	b += Math.floor((255 - b) * amount)

	return "#" + [r, g, b].map(x => x.toString(16).padStart(2, "0")).join("")
}

export const cardColors = Object.fromEntries(
	Object.entries(colors).map(([key, value]) => [
		key,
		`radial-gradient(circle at 0% 100%, ${value}, ${value}, ${lightenHexColor(value, 0.4)})`
	])
)
