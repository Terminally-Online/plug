import { FC, useEffect, useState } from "react"

import Image from "next/image"

import axios from "axios"
import { ExternalLink, Send } from "lucide-react"

import { Button } from "@/components/shared"
import { useFrame } from "@/contexts"
import { getAPIKey } from "@/lib"
import { RouterOutputs } from "@/server/client"

import { CollectibleImage } from "../../sockets/collectibles/collectible-image"
import { Frame } from "../base"

export const CollectibleFrame: FC<{
	collection: NonNullable<
		RouterOutputs["socket"]["balances"]["collectibles"]
	>[number]
	collectible?: NonNullable<
		RouterOutputs["socket"]["balances"]["collectibles"]
	>[number]["collectibles"][number]
}> = ({ collection, collectible }) => {
	const { frameVisible } = useFrame()

	const [color, setColor] = useState<string>("")
	const [traits, setTraits] = useState<
		| Array<{
				trait_type: string
				display_type: number
				max_value: number
				value: number
		  }>
		| undefined
	>(undefined)

	const isFrame =
		frameVisible ===
		`${collection.slug}-${collectible?.contract}-${collectible?.identifier}`

	useEffect(() => {
		if (!collection || !collectible || !isFrame || traits) return

		const getMetadata = async () => {
			const url = `https://api.opensea.io/api/v2/chain/${collection.chain}/contract/${collectible?.contract}/nfts/${collectible?.identifier}`
			const response = await axios.get(url, {
				headers: {
					Accept: "application/json",
					"x-api-key": "47f4db9595fa4eb3bdd0be743354f94b"
				}
			})
			setTraits(response.data.nft.traits)
			console.log(response.data.nft.traits)
		}

		getMetadata()
	}, [collection, collectible, isFrame, traits])

	return (
		<Frame
			icon={
				<div className="relative h-10 w-10">
					{/* <Image
						src={collection.imageUrl}
						alt={collection.name}
						className="absolute left-1/2 top-1/2 h-24 w-24 -translate-x-1/2 -translate-y-1/2 animate-fade-in rounded-full blur-lg filter transition-all duration-200 ease-in-out"
						width={140}
						height={140}
						// onError={() => setError(true)}
					/> */}
					<div
						className="absolute left-1/2 top-1/2 h-10 w-10 min-w-10 -translate-x-1/2 -translate-y-1/2 animate-fade-in rounded-full bg-grayscale-100"
						style={{
							backgroundImage: `url(${collection.imageUrl})`,
							backgroundSize: "cover",
							backgroundPosition: "center",
							backgroundRepeat: "no-repeat"
						}}
					/>
				</div>
			}
			label={collection.name}
			visible={isFrame}
			hasOverlay={true}
			hasChildrenPadding={false}
		>
			<div className="flex flex-col gap-2 px-6 pb-2">
				<CollectibleImage
					image={collectible?.displayImageUrl || collection.imageUrl}
					name={collectible?.name || collection.name}
					handleColor={setColor}
				/>

				<div className="pb-2 pt-4">
					<p className="flex flex-row items-center gap-2 font-bold text-black text-opacity-40">
						#{collectible?.identifier}
					</p>
					<p className="text-lg font-bold">{collectible?.name}</p>
				</div>
			</div>

			<div className="flex flex-row gap-2 px-6">
				<button
					className="flex w-full items-center justify-center gap-2 rounded-lg py-4 font-bold text-white"
					style={{ backgroundColor: color }}
				>
					<ExternalLink size={14} className="opacity-60" />
					Opensea
				</button>
				<button
					className="flex w-full items-center justify-center gap-2 rounded-lg py-4 font-bold text-white"
					style={{ backgroundColor: color }}
				>
					<Send size={14} className="opacity-60" />
					Transfer
				</button>
			</div>

			{traits && (
				<div className="grid grid-cols-2 gap-2 px-6 pt-4">
					{traits.map((trait, index) => (
						<div
							key={index}
							className="flex flex-col rounded-md bg-gradient-to-tr from-grayscale-100 to-grayscale-0 px-4 py-2"
						>
							<p className="truncate overflow-ellipsis whitespace-nowrap text-sm font-bold opacity-40">
								{trait.trait_type}
							</p>
							<p className="flex flex-row items-center gap-2 truncate overflow-ellipsis whitespace-nowrap font-bold">
								{trait.value}
							</p>
						</div>
					))}
				</div>
			)}

			<div className="pb-4 pt-4">
				<p className="border-t-[1px] border-grayscale-100 px-6 pt-4 opacity-60">
					{collectible?.description || collection.description}
				</p>
			</div>
		</Frame>
	)
}
