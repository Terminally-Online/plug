import { FC, useEffect, useRef, useState } from "react"

import { motion } from "framer-motion"

import { useClient } from "@/lib"
import { cn } from "@/lib/utils"

const isDate = (input: string) => input.includes("/")

const formatForDisplay = (input: number | string, formatDecimals: boolean, decimals?: number) => {
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
	return []
}

const MinusColumn = () => <span>-</span>

const CommaColumn = () => <span>,</span>

const DecimalColumn = () => <span>.</span>

const SlashColumn = () => <span>/</span>

const NumberColumn: FC<{ digit: string }> = ({ digit }) => {
	const [y, setY] = useState(0)
	const columnContainer = useRef<HTMLSpanElement>(null)

	const setColumnToNumber = (number: string) => {
		if (columnContainer.current === null) return
		setY(columnContainer.current.clientHeight * parseInt(number))
	}

	useEffect(() => setColumnToNumber(digit), [digit])

	return (
		<span className="relative text-center" ref={columnContainer}>
			<motion.span className="absolute bottom-0 h-[1000%]" whileInView={{ y }}>
				{[9, 8, 7, 6, 5, 4, 3, 2, 1, 0].map(i => (
					<div key={i} className="h-[10%]">
						{i}
					</div>
				))}
			</motion.span>
			<span className="invisible">0</span>
		</span>
	)
}

export const Counter: FC<
	{
		count: number | string
		formatDecimals?: boolean
		decimals?: number
	} & React.HTMLAttributes<HTMLDivElement>
> = ({ count, formatDecimals = true, decimals, className = "", ...props }) => {
	const isClient = useClient()

	const numArray = formatForDisplay(count, formatDecimals, decimals)

	if (!isClient) return null

	return (
		<span className={cn("relative flex w-full flex-row-reverse overflow-hidden", className)} {...props}>
			{numArray.map((number, index) =>
				number === "." ? (
					<DecimalColumn key={index} />
				) : number === "-" ? (
					<MinusColumn key={index} />
				) : number === "," ? (
					<CommaColumn key={index} />
				) : number === "/" ? (
					<SlashColumn key={index} />
				) : (
					<NumberColumn key={index} digit={number} />
				)
			)}
		</span>
	)
}
