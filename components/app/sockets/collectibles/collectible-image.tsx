import { FC, useEffect, useRef, useState } from "react"

import Image from "next/image"

import { getAssetColor } from "@/lib"

export const CollectibleImage: FC<{
	image?: string
	name?: string
	size?: "xs" | "sm" | "md"
	handleColor?: (color: string) => void
}> = ({ image = "", name = "", size = "md", handleColor }) => {
	const canvasRef = useRef<HTMLCanvasElement>(null)
	const imgRef = useRef<HTMLImageElement>(null)

	const [imageError, setImageError] = useState(false)

	const dimensions = {
		blur: size === "xs" ? 2 : size === "sm" ? 3 : 4,
		imageBlur: size === "xs" ? 4 : size === "sm" ? 6 : 12,
		content: size === "xs" ? 1.5 : size === "sm" ? 2 : 2.5
	}

	const color = getAssetColor(name)

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

			const imageData = ctx.getImageData(
				0,
				0,
				canvas.width,
				canvas.height
			)
			const data = imageData.data

			const colorCounts: Record<string, number> = {}
			let maxCount = 0
			let dominantColor = ""

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

				if (colorCounts[rgb] > maxCount) {
					maxCount = colorCounts[rgb]
					dominantColor = `rgb(${rgb})`
				}
			}

			return dominantColor
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
	}, [image, name, color, handleColor])

	return (
		<div className="relative h-full w-full">
			<canvas ref={canvasRef} style={{ display: "none" }} />

			{image === "" || imageError ? (
				<>
					<div
						className="absolute left-1/2 top-1/2 -translate-x-1/2 -translate-y-1/2 animate-fade-in rounded-full blur-2xl filter transition-all duration-200 ease-in-out"
						style={{
							backgroundColor: getAssetColor(name),
							width: `${dimensions.blur}rem`,
							height: `${dimensions.blur}rem`
						}}
					/>
					<div
						className="absolute left-1/2 top-1/2 flex -translate-x-1/2 -translate-y-1/2 animate-fade-in items-center justify-center rounded-full"
						style={{
							backgroundColor: getAssetColor(name),
							width: `${dimensions.content}rem`,
							height: `${dimensions.content}rem`,
							minWidth: `${dimensions.content}rem`
						}}
					>
						<span className="font-bold text-white">
							{name.slice(0, 1).toUpperCase()}
						</span>
					</div>
				</>
			) : (
				<>
					<Image
						src={image}
						alt={name}
						className="mb-4 h-[200%] w-[200%] rounded-md object-cover blur-[200px]"
						width={1200}
						height={1200}
					/>
					<Image
						ref={imgRef}
						src={image}
						alt={name}
						className="absolute top-0 mb-4 h-full w-full rounded-md object-cover"
						width={1200}
						height={1200}
						onError={() => setImageError(true)}
					/>
				</>
			)}
		</div>
	)
}
