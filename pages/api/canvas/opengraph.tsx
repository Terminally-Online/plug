import { NextRequest } from "next/server"

import { ImageResponse } from "@vercel/og"

export const config = {
	runtime: "edge"
}

export default async function handler(req: NextRequest) {
	const fontData = await fetch(new URL("../../../assets/Satoshi-Regular.ttf", import.meta.url)).then(res => res.arrayBuffer())

	try {
		const { searchParams } = req.nextUrl

		const name = searchParams.get("name") ?? "Untitled Canvas"
		const color = searchParams.get("color") ?? "#FF8C00"
		const updatedAt = searchParams.get("updatedAt") ?? Date.now()
		const updatedAtDate = new Date(updatedAt)
		const duration = Date.now() - updatedAtDate.getTime()

		const seconds = Math.floor(duration / 1000)
		const minutes = Math.floor(seconds / 60)
		const hours = Math.floor(minutes / 60)

		let time = "just now"
		if (hours > 0) {
			time = `${hours} hour${hours > 1 ? "s" : ""} ago`
		} else if (minutes > 0) {
			time = `${minutes} minute${minutes > 1 ? "s" : ""} ago`
		}

		const nameLength = 60
		const cleanedName = name.slice(0, nameLength).trim() + `${name.length > nameLength ? "..." : ""}`

		return new ImageResponse(
			(
				<div
					tw="flex flex-col justify-end text-white w-full h-full p-16 bg-stone-900"
					style={{
						fontFamily: "Satoshi"
					}}
				>
					<div tw="flex flex-row items-center">
						<div
							tw="rounded-full border-[1px] border-[#0c0a09] w-[35px] h-[35px] mr-12"
							style={{
								backgroundColor: color
							}}
						/>
						<h1 tw="text-8xl" style={{ wordBreak: "break-word" }}>
							{cleanedName}
						</h1>
					</div>

					<p tw="text-4xl opacity-60">Edited {time}</p>
				</div>
			),
			{
				width: 1200,
				height: 630,
				fonts: [
					{
						name: "Satoshi",
						data: fontData,
						style: "normal"
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
