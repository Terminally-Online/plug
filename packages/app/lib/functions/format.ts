export const formatNumber = (value: number) => {
	const fixed = 4

	if (value < 1e3) return parseFloat(value.toFixed(fixed).toString()).toString()
	if (value >= 1e3 && value < 1e6) return parseFloat((+(value / 1e3).toFixed(fixed)).toString()) + "K"
	if (value >= 1e6 && value < 1e9) return parseFloat((+(value / 1e6).toFixed(fixed)).toString()) + "M"
	if (value >= 1e9 && value < 1e12) return parseFloat((+(value / 1e9).toFixed(fixed)).toString()) + "B"
	return parseFloat((+(value / 1e12).toFixed(fixed)).toString()) + "T"
}

export const formatFloat = (value: number) => parseFloat(formatNumber(value).toString())

export const formatChainName = (name: string) =>
	name.replace("Mainnet", "").replace("Testnet", "").replace("OP", "Optimism")

export const formatAddress = (address: string | `0x${string}`, characters = 4) => {
	return `${address.slice(0, characters)}...${address.slice(-characters)}`
}

export const formatBalance = (value: string | bigint | bigint | undefined, decimals: number | undefined) => {
	if (!value || !decimals) return undefined

	return Number.parseFloat((Number(value) / 10 ** Number(decimals)).toFixed(4))
}

export const formatTitle = (title: string) =>
	title
		.replace("_", " ")
		.replace(/([a-z])([A-Z])|([A-Z])([A-Z][a-z])/g, "$1$3 $2$4")
		.split(" ")
		.map(word => {
			// If the word is all uppercase and more than one letter, keep it as an acronym
			if (word.toUpperCase() === word && word.length > 1) {
				return word
			} else {
				// Capitalize the first letter of other words
				return word.charAt(0).toUpperCase() + word.slice(1).toLowerCase()
			}
		})
		.join(" ")

export const formatInputName = (input: string | undefined) =>
	input
		?.replace("$", "")
		.replace(/([a-z])([A-Z])|([A-Z])([A-Z][a-z])/g, "$1$3 $2$4")
		.split(" ")
		.join(" ")
		.toLowerCase() ?? ""

export const formatTimeSince = (date: Date): string => {
	const now = new Date()
	const seconds = Math.floor((now.getTime() - date.getTime()) / 1000)

	let interval = seconds / 31536000

	if (seconds < 30) return "just now"

	if (interval > 1) {
		return `${Math.floor(interval)} year${Math.floor(interval) > 1 ? "s" : ""} ago`
	}
	interval = seconds / 2592000 // Number of seconds in one month
	if (interval > 1) {
		return `${Math.floor(interval)} month${Math.floor(interval) > 1 ? "s" : ""} ago`
	}
	interval = seconds / 86400 // Number of seconds in one day
	if (interval > 1) {
		return `${Math.floor(interval)} day${Math.floor(interval) > 1 ? "s" : ""} ago`
	}
	interval = seconds / 3600 // Number of seconds in one hour
	if (interval > 1) {
		return `${Math.floor(interval)} hour${Math.floor(interval) > 1 ? "s" : ""} ago`
	}
	interval = seconds / 60 // Number of seconds in one minute
	if (interval > 1) {
		return `${Math.floor(interval)} minute${Math.floor(interval) > 1 ? "s" : ""} ago`
	}
	return `${Math.floor(seconds)} second${seconds > 1 ? "s" : ""} ago`
}

export const formatDate = (date: Date): string => {
	const dayOfWeek = date.toLocaleString("en-US", { weekday: "short" })
	const month = date.toLocaleString("en-US", { month: "short" })
	const dayOfMonth = date.getDate()

	return `${dayOfWeek} ${month} ${dayOfMonth}`
}

export const formatLongString = (str: string, maxLength?: number) => {
	if (maxLength === undefined) return { data: str, truncated: false }

	if (str.length > maxLength) {
		return {
			data: str.slice(0, maxLength) + "...",
			truncated: true
		}
	}

	return { data: str, truncated: false }
}

export const formatTokenStandard = (tokenStandard: string) => {
	switch (tokenStandard) {
		case "erc721":
			return "ERC-721"
		case "erc1155":
			return "ERC-1155"
		default:
			return tokenStandard
	}
}

const isDate = (input: string) => input.includes("/")

export const formatForDisplay = (input: number | string, formatDecimals: boolean, decimals?: number) => {
	if (typeof input === "number") {
		const absCount = Math.abs(input)
		let formattedNumber: string[]

		if (formatDecimals) {
			if (Number.isInteger(absCount) && decimals === undefined) {
				// For whole numbers without specified decimals, don't add decimal places
				formattedNumber = absCount.toLocaleString("en-US").split("")
			} else {
				// Convert to string to preserve all decimal places
				let stringNumber = absCount.toString()

				let decimalPlacesToKeep: number

				if (decimals !== undefined) {
					// Use the manually specified number of decimal places
					decimalPlacesToKeep = decimals
				} else {
					// Find the index of the first non-zero digit after decimal point
					const decimalIndex = stringNumber.indexOf(".")
					let significantDigitIndex = decimalIndex + 1
					while (stringNumber[significantDigitIndex] === "0") {
						significantDigitIndex++
					}

					// Calculate how many decimal places to keep
					decimalPlacesToKeep = Math.max(3, significantDigitIndex - decimalIndex + 2)
				}

				// Round to the calculated number of decimal places
				const rounded = Number(absCount.toFixed(decimalPlacesToKeep))

				// Format the number, ensuring at least 2 decimal places or the targetDecimals
				let formatted = rounded.toFixed(Math.max(2, decimalPlacesToKeep))

				// Trim trailing zeros, but keep at least 2 decimal places or the targetDecimals
				if (decimals === undefined) {
					formatted = formatted.replace(/\.?0+$/, "")
					if (formatted.includes(".") && formatted.split(".")[1].length < 2) {
						formatted += "0"
					}
				}

				formattedNumber = Number(formatted)
					.toLocaleString("en-US", {
						maximumFractionDigits: 20
					})
					.split("")
			}
		} else {
			formattedNumber = absCount
				.toLocaleString("en-US", {
					maximumFractionDigits: 20
				})
				.split("")
		}

		if (input < 0) formattedNumber.unshift("-")
		return formattedNumber.reverse()
	} else if (typeof input === "string" && isDate(input)) {
		return input.split("").reverse()
	}

	return input.split("").reverse()
}

export const frequencies = [
	{ label: "Once", value: "0" },
	{ label: "Daily", value: "1" },
	{ label: "Weekly", value: "7" },
	{ label: "Monthly", value: "30" },
	{ label: "Quarterly", value: "90" },
	{ label: "Yearly", value: "365" }
]

export const formatFrequency = (frequencyValue: number) => {
	const frequencyIndex = frequencies.findIndex(frequency => parseInt(frequency.value) === frequencyValue)
	return frequencies[frequencyIndex].label
}