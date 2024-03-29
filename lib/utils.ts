import { type ClassValue, clsx } from "clsx"
import { twMerge } from "tailwind-merge"

export const cn = (...inputs: ClassValue[]) => twMerge(clsx(inputs))

export const formatNumber = (value: number) => {
	const fixed = 4

	if (value < 1e3)
		return parseFloat(value.toFixed(fixed).toString()).toString()
	if (value >= 1e3 && value < 1e6)
		return parseFloat((+(value / 1e3).toFixed(fixed)).toString()) + "K"
	if (value >= 1e6 && value < 1e9)
		return parseFloat((+(value / 1e6).toFixed(fixed)).toString()) + "M"
	if (value >= 1e9 && value < 1e12)
		return parseFloat((+(value / 1e9).toFixed(fixed)).toString()) + "B"
	return parseFloat((+(value / 1e12).toFixed(fixed)).toString()) + "T"
}

export const formatFloat = (value: number) =>
	parseFloat(formatNumber(value).toString())

export const levenshteinDistance = (a: string, b: string): number => {
	const matrix: number[][] = []

	for (let i = 0; i <= b.length; i++) {
		matrix[i] = [i]
	}

	for (let j = 0; j <= a.length; j++) {
		matrix[0][j] = j
	}

	for (let i = 1; i <= b.length; i++) {
		for (let j = 1; j <= a.length; j++) {
			if (b.charAt(i - 1) == a.charAt(j - 1)) {
				matrix[i][j] = matrix[i - 1][j - 1]
			} else {
				matrix[i][j] = Math.min(
					matrix[i - 1][j - 1] + 1,
					Math.min(matrix[i][j - 1] + 1, matrix[i - 1][j] + 1)
				)
			}
		}
	}

	return matrix[b.length][a.length]
}
