import { FC } from "react"

import { Image } from "@/components/app/utils/image"

export const ConnectorImage: FC<{ icon: string | undefined; name: string }> = ({ icon, name }) => {
	const dimensions = {
		blur: 4,
		content: 2.5
	}

	if (!icon) return null

	return (
		<div
			className="relative h-10"
			style={{
				width: `${dimensions.content}rem`,
				height: `${dimensions.content}rem`
			}}
		>
			<Image
				className="absolute left-1/2 top-1/2 h-12 w-12 -translate-x-1/2 -translate-y-1/2 rounded-md blur-xl filter"
				src={icon.trimStart().trim()}
				alt={name}
				style={{
					height: `${dimensions.blur}rem`,
					width: `${dimensions.blur}rem`
				}}
				width={48}
				height={48}
			/>
			<Image
				className="absolute left-1/2 top-1/2 -translate-x-1/2 -translate-y-1/2 rounded-md"
				src={icon.trimStart().trim()}
				alt={name}
				style={{
					width: `${dimensions.content}rem`,
					height: `${dimensions.content}rem`
				}}
				width={48}
				height={48}
			/>
		</div>
	)
}


