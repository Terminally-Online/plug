/* eslint-disable @next/next/no-img-element */
import { NextRequest } from "next/server"
import type { CSSProperties } from "react"

import { ImageResponse } from "@vercel/og"

const colors = ["#F3EF8A", "#8AF3E6", "#EB8AF3", "#9F8AF3", "#F3908A", "#F3B08A", "#8AAEF3", "#92F38A"]

export const config = {
	runtime: "edge"
}

// export const getStaticPaths = () => {
// 	const paths = colors.map(color => ({ params: { color } }))
// 	return { paths, fallback: "blocking" }
// }

// export const getStaticProps = ({ params }: { params: { color?: string; number?: string } }) => ({
// 	props: { color: params?.color, number: params?.number },
// 	revalidate: 60 * 60 * 24 * 30
// })

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

	const { searchParams, host, protocol } = req.nextUrl
	const color = searchParams.get("color")
	const numberParam = searchParams.get("number")

	if (!color && !numberParam) {
		return new Response("Missing required parameters: color or number", {
			status: 400
		})
	}

	const finalColor = color ?? numberParam ?? "FDFFF7"
	const number = numberParam ? parseInt(numberParam) : null

	return new ImageResponse(
		(
			<div
				tw="flex flex-col text-[#385842] w-full h-full p-16 pb-4 relative font-bold flex flex-col"
				style={{
					fontFamily: "Satoshi",
					background: `linear-gradient(to top, #FDFFF7, #FDFFF7, #FDFFF7, #${finalColor}, #${finalColor})`,
					overflow: "hidden"
				}}
			>
				{/* <div
					tw="absolute inset-0"
					style={{
						background: `linear-gradient(135deg, 
							transparent 0%,
							rgba(254,255,247,0.4) 35%,
							rgba(254,255,247,0.7) 45%,
							#${finalColor}33 50%,
							transparent 65%
						)`,
						mixBlendMode: "soft-light",
						filter: "blur(1px)"
					}}
				/> */}

				<div
					tw="flex w-full h-[70vh] rounded-[40px] relative p-90"
					style={{
						background: `linear-gradient(45deg, #38584319, #${finalColor}19, #79BE9119)`
					}}
				>
					<img
						tw="mx-auto h-full opacity-60"
						src={`${protocol}//${host}/dna.png`}
						alt="Dna image"
						width={200}
						height={1000}
					/>
				</div>

				<div tw="mt-auto flex relative h-max w-full mb-4">
					<h1
						tw="text-[120px] font-black w-[60%]"
						style={{
							background: `linear-gradient(45deg, #385842, #${finalColor}, #385842)`,
							backgroundClip: "text",
							WebkitBackgroundClip: "text",
							WebkitTextFillColor: "transparent",
							textShadow: `0 0 20px #${finalColor}33`
						}}
					>
						FOUNDING USER
					</h1>
					{number && (
						<div tw="flex flex-col bottom-6 right-0 absolute justify-end items-end font-bold">
							<h3 tw="text-[80px] opacity-60">#{number.toLocaleString()}</h3>
						</div>
					)}
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
}
