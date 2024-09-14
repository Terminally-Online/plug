import Image from "next/image"
import { FC, useState } from "react"

import { TriangleAlert } from "lucide-react"

export const CollectibleImage: FC<{
	video?: string
	image?: string
	fallbackImage?: string
	name?: string
	size?: "xs" | "sm" | "md"
}> = ({ video, image = "", fallbackImage = "", name = "", size = "md" }) => {
	const [videoError, setVideoError] = useState(false)
	const [imageError, setImageError] = useState(false)

	return (
		<div className="relative h-full w-full">
			{video && videoError === false ? (
				<>
					<video
						src={video}
						className="absolute z-[-1] h-full w-full rounded-lg object-cover blur-2xl"
						autoPlay
						playsInline
						loop
						muted
					/>
					<video
						src={video}
						className="h-full w-full rounded-lg object-cover"
						autoPlay
						playsInline
						loop
						muted
						onError={() => setVideoError(true)}
					/>
				</>
			) : image && imageError === false ? (
				<>
					<Image
						src={image}
						alt={name}
						className="mb-4 h-full w-full rounded-lg object-cover blur-2xl"
						width={1200}
						height={1200}
					/>
					<Image
						src={image}
						alt={name}
						className="absolute top-0 mb-4 h-full w-full rounded-lg object-cover"
						width={1200}
						height={1200}
						onError={() => setImageError(true)}
					/>
				</>
			) : (
				<>
					<Image
						src={fallbackImage}
						alt={name}
						className="mb-4 h-full w-full rounded-lg object-cover blur-2xl"
						width={1200}
						height={1200}
					/>
					<Image
						src={fallbackImage}
						alt={name}
						className="absolute left-1/2 top-1/2 mb-4 h-[60%] w-[60%] -translate-x-1/2 -translate-y-1/2 animate-fade-in rounded-full"
						width={1200}
						height={1200}
						onError={() => setImageError(true)}
					/>
					<div className="absolute bottom-6 right-1/2 z-[2] mx-auto flex w-full translate-x-1/2 flex-row items-center justify-center gap-2 whitespace-nowrap text-center font-bold text-black text-opacity-60">
						<TriangleAlert size={18} className="opacity-40" />
						<p>Could not load image.</p>
					</div>
				</>
			)}
		</div>
	)
}
