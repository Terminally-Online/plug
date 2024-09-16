import { useRouter } from "next/router"
import { useEffect } from "react"

import { atom, useAtom } from "jotai"

const framesAtom = atom<Record<string, string | undefined>>({})

export const useFrames = () => {
	const router = useRouter()
	const [frames, setFrames] = useAtom(framesAtom)

	useEffect(() => {
		setFrames({})
	}, [router.pathname, setFrames])

	return { frames, setFrames }
}

export const useFrame = ({ index, key, separator }: { index?: number; key?: string; separator?: string }) => {
	const { frames, setFrames } = useFrames()

	const frame = separator && key ? key.split(separator)[0] : key
	const prevFrame = separator && key ? key.split(separator)[1] : undefined

	const isFrame = index ? frames[index] === frame : false

	const handleFrame = (key?: string) => {
		const frameKey = key ?? frame

		if (!index) return

		setFrames(prevFrames => ({
			...prevFrames,
			[index]: prevFrames[index] === frameKey ? undefined : frameKey
		}))
	}

	return { index, key, isFrame, prevFrame, frames, handleFrame }
}
