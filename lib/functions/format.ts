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
