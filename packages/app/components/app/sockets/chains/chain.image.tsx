import { ChainId, chains, cn } from "@/lib"

import { Image } from "../../utils/image"

type ChainImageProps = {
	chainId: ChainId
	size?: "xs" | "sm" | "md" | "lg"
}

export const ChainImage = ({ chainId, size = "sm" }: ChainImageProps) => {
	const src = chains[chainId]?.logo || "/protocols/plug.png"

	const sizeClasses = {
		xs: "h-4 w-4 rounded-xs",
		sm: "h-6 w-6 rounded-sm",
		md: "h-8 w-8 rounded-sm",
		lg: "h-10 w-10 rounded-sm"
	}

	if (!chainId) return null

	return <Image src={src} alt={chainId.toString()} width={128} height={128} className={cn(sizeClasses[size])} />
}
