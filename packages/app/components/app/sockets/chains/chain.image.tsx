import { useMemo } from "react"

import { ChainId, chains, cn, getZerionChainIconUrl } from "@/lib"

import { Image } from "../../utils/image"

type ChainImageProps = {
	chainId: string | ChainId
	size?: "xs" | "sm" | "md" | "lg"
}

const sizeClasses = {
	xs: "h-4 w-4 rounded-xs",
	sm: "h-5 w-5 rounded-sm",
	md: "h-8 w-8 rounded-sm",
	lg: "h-10 w-10 rounded-sm"
}

export const ChainImage = ({ chainId, size = "sm" }: ChainImageProps) => {
	const src = useMemo(
		() =>
			(typeof chainId === "string" ? getZerionChainIconUrl(chainId) : chains[chainId]?.logo) ||
			"/protocols/plug.png",
		[chainId]
	)

	if (!chainId) return null

	return <Image src={src} alt={chainId.toString()} width={128} height={128} className={cn(sizeClasses[size])} />
}
