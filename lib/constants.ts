import { z } from "zod"

import { radians } from "./functions/math-utils"
import { Pins } from "./types"

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
}

export const greenGradientStyle = {
	background: "linear-gradient(45deg, #00EF35, #93DF00)",
	WebkitBackgroundClip: "text",
	WebkitTextFillColor: "transparent"
}

export const DEBUG = true

export const CAMERA_ANGLE = radians(30)
export const RECT_W = 1000
export const RECT_H = 1000

export const ItemTypes = {
	Markdown: "MARKDOWN",
	Box: "BOX",
	Plug: "PLUG"
} as const

const nounsSchema = z.object({
	price: z.number()
})

const thresholdSchema = z.object({
	threshold: z.number().default(Date.now())
})

export const pins: Pins = [
	{
		label: "Nouns",
		pins: [
			{
				label: "Can Bid on Noun",
				value: "can-bid",
				type: "if",
				schema: nounsSchema
			},
			{
				label: "Place Bid on Noun",
				value: "place-bid",
				type: "then",
				schema: nounsSchema
			}
		]
	},
	{
		label: "Schedule",
		pins: [
			{
				label: "Within Window",
				value: "within-window",
				type: "if",
				schema: thresholdSchema
			},
			{
				label: "Before Block Number",
				value: "before-block-number",
				type: "if",
				schema: thresholdSchema
			},
			{
				label: "After Block Number",
				value: "after-block-number",
				type: "if",
				schema: thresholdSchema
			},
			{
				label: "Before Timestamp",
				value: "before-timestamp",
				type: "if",
				schema: thresholdSchema
			},
			{
				label: "After Timestamp",
				value: "after-timestamp",
				type: "if",
				schema: thresholdSchema
			}
		]
	}
]
