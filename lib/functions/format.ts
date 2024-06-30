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

export const formatChainName = (name: string) =>
	name.replace("Mainnet", "").replace("Testnet", "").replace("OP", "Optimism")

export const formatAddress = (address: string) => {
	return `${address.slice(0, 6)}...${address.slice(-4)}`
}

export const formatBalance = (
	value: string | bigint | bigint | undefined,
	decimals: number | undefined
) => {
	if (!value || !decimals) return 0

	return Number.parseFloat(
		(Number(value) / 10 ** Number(decimals)).toFixed(4)
	)
}

export const formatTitle = (title: string) =>
	title
		.replace(/([a-z])([A-Z])|([A-Z])([A-Z][a-z])/g, "$1$3 $2$4")
		.split(" ")
		.map(word => {
			// If the word is all uppercase and more than one letter, keep it as an acronym
			if (word.toUpperCase() === word && word.length > 1) {
				return word
			} else {
				// Capitalize the first letter of other words
				return (
					word.charAt(0).toUpperCase() + word.slice(1).toLowerCase()
				)
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

// Format a date in the format of "name of the week, month, date"
// Example: "Wed Apr 24"
export const formatDate = (date: Date): string => {
	const dayOfWeek = date.toLocaleString("en-US", { weekday: "short" }) // "Wed"
	const month = date.toLocaleString("en-US", { month: "short" }) // "Apr"
	const dayOfMonth = date.getDate() // 24

	return `${dayOfWeek} ${month} ${dayOfMonth}`
}
