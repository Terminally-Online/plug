import { FC, useState } from "react"

import Image from "next/image"

import { keccak256 } from "viem"

import { cn } from "@/lib"

const colors = [
	"#f87171",
	"#fb923c",
	"#fbbf24",
	"#facc15",
	"#a3e635",
	"#4ade80",
	"#34d399",
	"#2dd4bf",
	"#22d3ee",
	"#38bdf8",
	"#60a5fa",
	"#818cf8",
	"#a78bfa",
	"#c084fc",
	"#e879f9",
	"#f472b6",
	"#fb7185"
]

// Take a symbol and calculate a deterministic color for it.
const getColor = (symbol: string) => {
	// Get the ASCII code of the first character of the symbol
	const asciiValue = symbol.charCodeAt(0)

	// Determine the index in the colors array using modulo operation
	const index = asciiValue % colors.length

	// Return the color at that index
	return colors[index]
}

export const TokenImage: FC<{ logo?: string; symbol?: string }> = ({
	logo = "",
	symbol = ""
}) => {
	const [imageError, setImageError] = useState(false)

	return (
		<div className="relative h-10 w-10 bg-grayscale-0">
			{logo === "" || imageError ? (
				<>
					<div
						className="absolute left-1/2 top-1/2 h-16 w-16 -translate-x-1/2 -translate-y-1/2 animate-fade-in rounded-full blur-2xl filter transition-all duration-200 ease-in-out"
						style={{ backgroundColor: getColor(symbol) }}
					/>
					<div
						className="absolute left-1/2 top-1/2 flex h-10 w-10 min-w-10 -translate-x-1/2 -translate-y-1/2 animate-fade-in items-center justify-center rounded-full"
						style={{ backgroundColor: getColor(symbol) }}
					>
						<p className="font-bold text-white">
							{symbol.slice(0, 1).toUpperCase()}
						</p>
					</div>
				</>
			) : (
				<>
					<Image
						src={logo}
						alt={symbol}
						className="absolute left-1/2 top-1/2 h-48 w-48 -translate-x-1/2 -translate-y-1/2 animate-fade-in rounded-full blur-2xl filter transition-all duration-200 ease-in-out"
						width={140}
						height={140}
						onError={() => setImageError(true)}
					/>
					<Image
						src={logo}
						alt={symbol}
						className="absolute left-1/2 top-1/2 h-10 w-10 min-w-10 -translate-x-1/2 -translate-y-1/2 animate-fade-in rounded-full bg-grayscale-100"
						width={140}
						height={140}
						onError={() => setImageError(true)}
					/>
				</>
			)}
		</div>
	)
}
