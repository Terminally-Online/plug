import { FC, useEffect, useRef, useState } from "react"

import Image from "next/image"

import { getAssetColor } from "@/lib"

export const TokenImage: FC<{
	logo?: string
	symbol?: string
	size?: "xs" | "sm" | "md"
	blur?: boolean
	handleColor?: (color: string) => void
}> = ({ logo = "", symbol = "", size = "md", blur = true, handleColor }) => {
	const canvasRef = useRef<HTMLCanvasElement>(null)
	const imgRef = useRef<HTMLImageElement>(null)

	const [imageError, setImageError] = useState(false)

	const dimensions = {
		blur: size === "xs" ? 2 : size === "sm" ? 3 : 4,
		imageBlur: size === "xs" ? 4 : size === "sm" ? 6 : 12,
		content: size === "xs" ? 1 : size === "sm" ? 2 : 2.5
	}

	const color = getAssetColor(symbol)

	useEffect(() => {
		if (!handleColor) return

		const getAverageColor = () => {
			const canvas = canvasRef.current
			const img = imgRef.current
			const ctx = canvas?.getContext("2d")

			if (!ctx || !canvas || !img) return color

			canvas.width = img.naturalWidth
			canvas.height = img.naturalHeight
			ctx.drawImage(img, 0, 0, img.naturalWidth, img.naturalHeight)

			const imageData = ctx.getImageData(0, 0, canvas.width, canvas.height)
			const data = imageData.data

			const colorCounts: Record<string, number> = {}
			const dominantColors: Array<{ color: string; count: number }> = []

			for (let i = 0; i < data.length; i += 4) {
				const r = data[i]
				const g = data[i + 1]
				const b = data[i + 2]
				const rgb = `${Math.floor(r / 10) * 10},${Math.floor(g / 10) * 10},${Math.floor(b / 10) * 10}`

				if (colorCounts[rgb]) {
					colorCounts[rgb]++
				} else {
					colorCounts[rgb] = 1
				}
			}

			for (const [rgb, count] of Object.entries(colorCounts)) {
				dominantColors.push({ color: `rgb(${rgb})`, count })
			}

			dominantColors.sort((a, b) => b.count - a.count)

			const isReadable = (color: string) => {
				const rgb = color.match(/\d+/g)
				if (!rgb) return false
				const [r, g, b] = rgb.map(Number)
				const luminance = (0.299 * r + 0.587 * g + 0.114 * b) / 255
				return luminance < 0.7
			}

			const readableColors = dominantColors.filter(({ color }) => isReadable(color))

			if (readableColors.length > 0) {
				const nonBlackColor = readableColors.find(({ color }) => {
					const [r, g, b] = color.match(/\d+/g)!.map(Number)
					return r > 20 || g > 20 || b > 20
				})
				return nonBlackColor ? nonBlackColor.color : readableColors[0].color
			}

			return color
		}

		const img = imgRef.current

		if (!img) {
			handleColor(color)
			return
		}

		if (img.complete) {
			handleColor(getAverageColor())
		} else {
			img.onload = () => handleColor?.(getAverageColor())
		}
	}, [logo, symbol, color, handleColor])

	return (
		<div
			className="relative h-10"
			style={{
				width: `${dimensions.content}rem`,
				height: `${dimensions.content}rem`
			}}
		>
			<canvas ref={canvasRef} style={{ display: "none" }} />

			{logo === "" || imageError ? (
				<>
					{blur && (
						<div
							className="absolute left-1/2 top-1/2 -translate-x-1/2 -translate-y-1/2 animate-fade-in rounded-full blur-2xl filter transition-all duration-200 ease-in-out"
							style={{
								backgroundColor: getAssetColor(symbol),
								width: `${dimensions.blur}rem`,
								height: `${dimensions.blur}rem`
							}}
						/>
					)}
					<div
						className="absolute left-1/2 top-1/2 flex -translate-x-1/2 -translate-y-1/2 animate-fade-in items-center justify-center rounded-full"
						style={{
							backgroundColor: getAssetColor(symbol),
							width: `${dimensions.content}rem`,
							height: `${dimensions.content}rem`,
							minWidth: `${dimensions.content}rem`
						}}
					>
						<span className="font-bold text-white">{symbol.slice(0, 1).toUpperCase()}</span>
					</div>
				</>
			) : (
				<>
					{blur && (
						<Image
							src={logo}
							alt={symbol}
							className="absolute left-1/2 top-1/2 -translate-x-1/2 -translate-y-1/2 animate-fade-in rounded-full blur-2xl filter transition-all duration-200 ease-in-out"
							style={{
								width: `${dimensions.imageBlur}rem`,
								height: `${dimensions.imageBlur}rem`
							}}
							width={240}
							height={240}
						/>
					)}
					<Image
						ref={imgRef}
						src={logo}
						alt={symbol}
						className="absolute left-1/2 top-1/2 -translate-x-1/2 -translate-y-1/2 animate-fade-in rounded-full bg-grayscale-100"
						style={{
							width: `${dimensions.content}rem`,
							minWidth: `${dimensions.content}rem`
						}}
						width={240}
						height={240}
						onError={() => setImageError(true)}
					/>
				</>
			)}
		</div>
	)
}
