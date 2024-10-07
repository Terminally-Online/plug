import NextImage from "next/image"
import { FC, HTMLAttributes, Ref } from "react"

const loader = ({ src, width }: { src: string; width: number }) => {
	if (src.includes("data:image")) {
		const byteString = atob(src.split(",")[1])
		const mimeString = src.split(",")[0].split(":")[1].split(";")[0]
		const ab = new ArrayBuffer(byteString.length)
		const ia = new Uint8Array(ab)
		for (let i = 0; i < byteString.length; i++) {
			ia[i] = byteString.charCodeAt(i)
		}
		const blob = new Blob([ab], { type: mimeString })
		return URL.createObjectURL(blob)
	}
	return `${src}${src.includes("?") ? "&" : "?"}w=${width}`
}

export const Image: FC<
	React.PropsWithRef<HTMLAttributes<HTMLImageElement>> & {
		src: string
		width: number
		height: number
		alt: string
		ref?: Ref<HTMLImageElement>
		priority?: boolean
		quality?: number
		blurSrc?: string
	}
> = ({ src, width, height, alt, ...props }) => (
	<NextImage
		loader={loader}
		src={src}
		width={width}
		height={height}
		alt={alt}
		placeholder={props.blurSrc ? "blur" : "empty"}
		blurDataURL={props.blurSrc}
		{...props}
	/>
)
