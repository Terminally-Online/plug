/* eslint-disable @next/next/no-img-element */
import { NextRequest } from "next/server"
import type { CSSProperties } from "react"

import { ImageResponse } from "@vercel/og"

export const config = {
	runtime: "edge"
}

export default async function handler(req: NextRequest) {
	const regular = await fetch(new URL("../../../assets/Satoshi-Regular.ttf", import.meta.url)).then(res =>
		res.arrayBuffer()
	)
	const bold = await fetch(new URL("../../../assets/Satoshi-Bold.ttf", import.meta.url)).then(res =>
		res.arrayBuffer()
	)
	const black = await fetch(new URL("../../../assets/Satoshi-Black.ttf", import.meta.url)).then(res =>
		res.arrayBuffer()
	)

	try {
		const { searchParams } = req.nextUrl
		const color = searchParams.get("color")
		const numberParam = searchParams.get("number")

		if (!color && !numberParam) {
			return new Response("Missing required parameters: color or number", {
				status: 400
			})
		}

		const finalColor = color ?? numberParam ?? "FDFFF7"
		const number = numberParam
			? parseInt(numberParam)
			: Math.floor(Math.random() * 100000)

		const isRare = finalColor !== "FDFFF7"
		const background = isRare
			? `linear-gradient(to top, #FDFFF7, #FDFFF7, #FDFFF7, #${finalColor}, #${finalColor})`
			: "#FDFFF7"

		const foilOverlay: CSSProperties = isRare
			? {
					background: `linear-gradient(135deg, 
				transparent 0%,
				rgba(254,255,247,0.4) 35%,
				rgba(254,255,247,0.7) 45%,
				#${finalColor}33 50%,
				transparent 65%
			)`,
					position: "absolute",
					top: 0,
					left: 0,
					right: 0,
					bottom: 0,
					mixBlendMode: "soft-light" as const,
					filter: "blur(1px)"
				}
			: {}

		return new ImageResponse(
			(
				<div
					tw="flex flex-col text-[#385842] w-full h-full p-16 pb-4 relative font-bold flex flex-col"
					style={{
						fontFamily: "Satoshi",
						background: background,
						overflow: "hidden"
					}}
				>
					{isRare && <div style={foilOverlay} />}
					<div
						tw="flex w-full h-[70vh] rounded-[40px] relative p-90"
						style={{
							background: isRare
								? `linear-gradient(45deg, #38584319, #${finalColor}19, #79BE9119)`
								: "linear-gradient(to bottom, #38584319, #79BE9119)"
						}}
					>
						<img
							tw="mx-auto h-full opacity-60"
							src={`${process.env.NEXT_PUBLIC_APP_URL}/dna.png`}
							alt="Dna image"
							width={200}
							height={1000}
						/>
					</div>

					<div tw="mt-auto flex relative h-max w-full mb-4">
						<h1
							tw="text-[120px] font-black whitespace-wrap w-[60%]"
							style={
								isRare
									? {
											background: `linear-gradient(45deg, #385842, #${finalColor}, #385842)`,
											backgroundClip: "text",
											WebkitBackgroundClip: "text",
											WebkitTextFillColor: "transparent",
											textShadow: `0 0 20px #${finalColor}33`
										}
									: {}
							}
						>
							FOUNDING USER
						</h1>
						<div tw="flex flex-col bottom-6 right-0 absolute justify-end items-end font-bold">
							<h3 tw="text-[80px] opacity-60">#{number.toLocaleString()}</h3>
						</div>
					</div>
				</div>
			),
			{
				width: 1000,
				height: 1600,
				fonts: [
					{
						name: "Satoshi",
						data: regular,
						style: "normal",
						weight: 400
					},
					{
						name: "Satoshi",
						data: bold,
						style: "normal",
						weight: 700
					},
					{
						name: "Satoshi",
						data: black,
						style: "normal",
						weight: 900
					}
				]
			}
		)
	} catch (e: any) {
		return new Response("Failed to generate image, " + e.message, {
			status: 500
		})
	}
}
