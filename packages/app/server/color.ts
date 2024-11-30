import sharp from "sharp"

export const getDominantColor = async (input: string) => {
	try {
		if (input.startsWith("http") === false) return undefined

		const response = await fetch(input)
		const arrayBuffer = await response.arrayBuffer()
		const buffer = Buffer.from(arrayBuffer)

		const image = sharp(buffer)
		const { dominant } = await image.resize(50, 50, { fit: "inside" }).stats()

		return `rgb(${dominant.r}, ${dominant.g}, ${dominant.b})`
	} catch (error) {
		console.error(error)
		return undefined
	}
}
