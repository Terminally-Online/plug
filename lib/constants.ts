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
