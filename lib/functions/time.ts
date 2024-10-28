export const getTimeInterval = (seconds: number) => {
	let interval = Math.floor(seconds / 31536000)
	if (interval >= 1) {
		return { interval, unit: "y" }
	}

	interval = Math.floor(seconds / 2592000)
	if (interval >= 1) {
		return { interval, unit: "mo" }
	}

	interval = Math.floor(seconds / 86400)
	if (interval >= 1) {
		return { interval, unit: "d" }
	}

	interval = Math.floor(seconds / 3600)
	if (interval >= 1) {
		return { interval, unit: "h" }
	}

	interval = Math.floor(seconds / 60)
	if (interval >= 1) {
		return { interval, unit: "m" }
	}

	return { interval: Math.floor(seconds), unit: "s" }
}
