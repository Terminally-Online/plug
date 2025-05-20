import sharp from "sharp"

interface StatsWithPalette extends sharp.Stats {
	palette?: Record<string, number>
}

const calculateContrastRatio = (
	rgb1: { r: number; g: number; b: number },
	rgb2: { r: number; g: number; b: number }
): number => {
	const l1 = calculateRelativeLuminance(rgb1)
	const l2 = calculateRelativeLuminance(rgb2)

	return l1 > l2 ? (l1 + 0.05) / (l2 + 0.05) : (l2 + 0.05) / (l1 + 0.05)
}

const calculateRelativeLuminance = (rgb: { r: number; g: number; b: number }): number => {
	const r = rgb.r / 255
	const g = rgb.g / 255
	const b = rgb.b / 255


	const rsrgb = r <= 0.03928 ? r / 12.92 : Math.pow((r + 0.055) / 1.055, 2.4)
	const gsrgb = g <= 0.03928 ? g / 12.92 : Math.pow((g + 0.055) / 1.055, 2.4)
	const bsrgb = b <= 0.03928 ? b / 12.92 : Math.pow((b + 0.055) / 1.055, 2.4)


	return 0.2126 * rsrgb + 0.7152 * gsrgb + 0.0722 * bsrgb
}

const hexToRgb = (hex: string): { r: number; g: number; b: number } => {
	const result = /^#?([a-f\d]{2})([a-f\d]{2})([a-f\d]{2})$/i.exec(hex)
	return result
		? {
			r: parseInt(result[1], 16),
			g: parseInt(result[2], 16),
			b: parseInt(result[3], 16)
		}
		: { r: 0, g: 0, b: 0 }
}

export const getDominantColor = async (input: string) => {
	try {
		if (input.startsWith("http") === false) return undefined

		const response = await fetch(input)
		const arrayBuffer = await response.arrayBuffer()
		const buffer = Buffer.from(arrayBuffer)

		const image = sharp(buffer)

		const stats = (await image.resize(50, 50, { fit: "inside" }).stats()) as StatsWithPalette
		const bgColor = hexToRgb("#FDFFF7")
		const minContrastRatio = 4.5

		if (stats.palette) {
			const sortedColors = Object.entries(stats.palette).sort(([, countA], [, countB]) => countB - countA)

			for (const [colorHex] of sortedColors) {
				const rgb = {
					r: parseInt(colorHex.substring(0, 2), 16),
					g: parseInt(colorHex.substring(2, 4), 16),
					b: parseInt(colorHex.substring(4, 6), 16)
				}

				const contrastRatio = calculateContrastRatio(rgb, bgColor)

				if (contrastRatio >= minContrastRatio) {
					return `rgb(${rgb.r}, ${rgb.g}, ${rgb.b})`
				}
			}
		}


		const { dominant } = stats
		const contrastRatio = calculateContrastRatio(dominant, bgColor)

		if (contrastRatio < minContrastRatio) {
			const luminance = calculateRelativeLuminance(dominant)
			const bgLuminance = calculateRelativeLuminance(bgColor)

			const luminanceDiff = Math.abs(luminance - bgLuminance)
			let adjustmentFactor = 0.5

			if (luminanceDiff < 0.2) {
				adjustmentFactor = 0.7
			} else if (luminanceDiff < 0.4) {
				adjustmentFactor = 0.6
			}

			if (bgLuminance > 0.5) {
				const darkened = {
					r: Math.max(0, Math.floor(dominant.r * (1 - adjustmentFactor))),
					g: Math.max(0, Math.floor(dominant.g * (1 - adjustmentFactor))),
					b: Math.max(0, Math.floor(dominant.b * (1 - adjustmentFactor)))
				}
				return `rgb(${darkened.r}, ${darkened.g}, ${darkened.b})`
			} else {
				const lightened = {
					r: Math.min(255, Math.floor(dominant.r * (1 + adjustmentFactor))),
					g: Math.min(255, Math.floor(dominant.g * (1 + adjustmentFactor))),
					b: Math.min(255, Math.floor(dominant.b * (1 + adjustmentFactor)))
				}
				return `rgb(${lightened.r}, ${lightened.g}, ${lightened.b})`
			}
		}

		return `rgb(${dominant.r}, ${dominant.g}, ${dominant.b})`
	} catch (error) {
		console.error(error)
		return undefined
	}
}
