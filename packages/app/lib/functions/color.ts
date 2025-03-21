import { SchemasRequestActions } from "@/lib/types"

import { colors } from "../constants/colors"
import { ASSET_COLORS } from "./blockchain"

export const getTextColor = (backgroundColor: string) => {
	let r: number, g: number, b: number

	if (ASSET_COLORS.includes(backgroundColor)) return "#FFFFFF"

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

export const getDominantProtocolColor = async (
	actions: SchemasRequestActions,
	plugId: string,
	ctx: any
): Promise<string> => {
	// Default to Plug color if no actions
	if (!actions?.length) {
		const color = colors.plug
		await updateWorkflowColor(plugId, color, ctx)
		return color
	}

	// Count protocol frequency
	const protocolFrequency: Record<string, number> = {}

	for (const action of actions) {
		if (!action?.protocol) continue

		// Normalize protocol name
		const normalizedProtocol = action.protocol
			.split("_")[0] // Remove version numbers
			.toLowerCase()

		protocolFrequency[normalizedProtocol] = (protocolFrequency[normalizedProtocol] || 0) + 1
	}

	// Find protocol with highest frequency
	const entries = Object.entries(protocolFrequency)
	if (!entries.length) {
		const color = colors.plug
		await updateWorkflowColor(plugId, color, ctx)
		return color
	}

	const dominantProtocol = entries.reduce((a, b) => (a[1] > b[1] ? a : b))[0]

	// Get color from colors constant
	const color = colors[dominantProtocol as keyof typeof colors] || colors.plug
	await updateWorkflowColor(plugId, color, ctx)
	return color
}

const updateWorkflowColor = async (plugId: string, color: string, ctx: any) => {
	await ctx.db.plug.update({
		where: { id: plugId },
		data: { color }
	})
}
