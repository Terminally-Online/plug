/* eslint-disable @next/next/no-img-element */
import { NextRequest } from "next/server"

import { ImageResponse } from "@vercel/og"

export const config = {
	runtime: "edge"
}

const nameLength = 60

export default async function handler(req: NextRequest) {
	const regular = await fetch(new URL("../../../assets/Satoshi-Regular.ttf", import.meta.url)).then(res =>
		res.arrayBuffer()
	)
	const bold = await fetch(new URL("../../../assets/Satoshi-Bold.ttf", import.meta.url)).then(res =>
		res.arrayBuffer()
	)

	try {
		const { searchParams, protocol: protocolParam, host } = req.nextUrl

		const name = searchParams.get("name") ?? "Untitled Plug"
		const protocols = searchParams.get("protocols")?.split(",") ?? ["aave", "gearbox", "yearn", "ens"]
		const sentences = searchParams.get("sentences")?.split(",") ?? [
			"Yoink [all] the money from Aave",
			"Deposit [1] [$USDC] into [USDC/WETH]",
			"Haha got all your money sir :)",
			"Renewed my name now"
		]
		const color = searchParams.get("color") ?? "#9F8AF3"

		const zipped = protocols.map((protocol, index) => ({
			protocol,
			sentence: sentences[index]
		}))

		const cleanedName = name.slice(0, nameLength).trim() + `${name.length > nameLength ? "..." : ""}`
		const Svg = () => {
			const svgs = [
				<svg
					key="1"
					width="1200"
					height="630"
					viewBox="0 0 1200 630"
					fill="none"
					xmlns="http://www.w3.org/2000/svg"
				>
					<g clipPath="url(#clip0_5066_135)">
						<path
							d="M568 551.064C739.225 551.064 752.919 479.987 874.369 493.775C968.706 504.484 1058.57 714.657 1251 591.337"
							stroke="url(#paint0_linear_5066_135)"
							strokeWidth="60"
							strokeLinecap="round"
						/>
					</g>
					<defs>
						<linearGradient
							id="paint0_linear_5066_135"
							x1="1214.82"
							y1="660.856"
							x2="772.146"
							y2="498.133"
							gradientUnits="userSpaceOnUse"
						>
							<stop stopColor="#D2F38A" />
							<stop offset="1" stopColor="#385842" />
						</linearGradient>
						<clipPath id="clip0_5066_135">
							<rect width="1200" height="630" fill="white" />
						</clipPath>
					</defs>
				</svg>,
				<svg
					key="3"
					width="1200"
					height="630"
					viewBox="0 0 1200 630"
					fill="none"
					xmlns="http://www.w3.org/2000/svg"
				>
					<g clipPath="url(#clip0_5066_137)">
						<path
							d="M786 97.3359C969.5 187.336 1064 -90.6641 1263 124.336"
							stroke="url(#paint0_linear_5066_137)"
							strokeWidth="60"
							strokeLinecap="round"
						/>
					</g>
					<defs>
						<linearGradient
							id="paint0_linear_5066_137"
							x1="1035.16"
							y1="886.946"
							x2="530.448"
							y2="321.221"
							gradientUnits="userSpaceOnUse"
						>
							<stop stopColor="#D2F38A" />
							<stop offset="1" stopColor="#385842" />
						</linearGradient>
						<clipPath id="clip0_5066_137">
							<rect width="1200" height="630" fill="white" />
						</clipPath>
					</defs>
				</svg>
			]

			return svgs[Math.floor(Math.random() * svgs.length)]
		}

		return new ImageResponse(
			(
				<div
					tw="flex flex-col text-[#385842] w-full h-full p-16 bg-[#FDFFF7] relative font-bold flex justify-center"
					style={{
						fontFamily: "Satoshi"
					}}
				>
					<div tw="w-[45%] absolute h-[100vh] top-0 right-0 bottom-0 flex items-center z-[20]">
						<div tw="w-[4px] h-full bg-[#385842]/10 absolute left-1/2 -translate-x-1/2 -z-[1]" />
						<div tw="flex flex-col w-full pl-8">
							{zipped.map((protocol, index) => (
								<div
									key={index}
									tw={`relative z-[20] w-full border-[4px] border-[#385842]/10 flex rounded-[16px] p-4 py-2 font-bold flex flex-row items-center bg-[#FDFFF7] ${
										index !== 0 && "mt-4"
									}`}
								>
									<img
										tw="w-8 h-8 rounded-[8px] mr-4"
										src={`${protocolParam}//${host}/protocols/${protocol.protocol}.png`}
										alt="fade"
									/>
									<p tw="flex items-center">
										{protocol.sentence.split(/(\[[^\]]*\])/).map((part, i) => {
											if (part.startsWith("[") && part.endsWith("]")) {
												return (
													<span
														key={i}
														tw="inline-flex items-center bg-[#385842]/20 mx-1 px-2 py-1 rounded-[6px]"
													>
														{part.slice(1, -1)}
													</span>
												)
											}
											return part
										})}
									</p>
								</div>
							))}
						</div>
					</div>

					<svg
						// @ts-ignore
						tw="absolute bottom-0 left-0 right-0 z-[4]"
						width="100%"
						height="50%"
						preserveAspectRatio="none"
					>
						<defs>
							<linearGradient id="fade" x1="0" y1="0" x2="0" y2="1">
								<stop offset="0%" stopColor="white" stopOpacity="0" />
								<stop offset="100%" stopColor="white" stopOpacity="1" />
							</linearGradient>
						</defs>
						<rect width="100%" height="100%" fill="url(#fade)" />
					</svg>

					<div tw="absolute top-0 left-0 right-0 bottom-0 flex z-[5]">
						<Svg />
					</div>

					<h1 tw="relative text-[120px] flex flex-col max-w-1/2" style={{ wordBreak: "break-word" }}>
						<img
							tw="w-16 h-16 mb-4 rounded-[8px] mr-4"
							src={`${protocolParam}//${host}/protocols/plug.png`}
							alt="fade"
						/>
						{cleanedName}
					</h1>
					<p tw="font-bold text-[24px] flex flex-row items-center w-[30%] justify-between opacity-60">
						{new Date().toLocaleDateString()}
					</p>
				</div>
			),
			{
				width: 1200,
				height: 630,
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
