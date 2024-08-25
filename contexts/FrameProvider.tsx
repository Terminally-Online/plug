import { createContext, Dispatch, FC, PropsWithChildren, SetStateAction, useContext, useEffect, useState } from "react"

import { useRouter } from "next/router"

export const FrameContext = createContext<{
	frames: Record<string, string | undefined>
	handleFrames: Dispatch<SetStateAction<Record<string, string | undefined>>>
}>({
	frames: {},
	handleFrames: () => {}
})

export const FrameProvider: FC<PropsWithChildren> = ({ children }) => {
	const router = useRouter()

	const [frames, setFrames] = useState<Record<string, string | undefined>>({})

	useEffect(() => {
		setFrames({})
	}, [router.pathname])

	return (
		<FrameContext.Provider
			value={{
				frames,
				handleFrames: setFrames
			}}
		>
			{children}
		</FrameContext.Provider>
	)
}

export const useFrame = ({ id, key, seperator }: { id?: string; key?: string; seperator?: string }) => {
	const { frames, handleFrames } = useContext(FrameContext)

	const frame = seperator && key ? key.split(seperator)[0] : key
	const prevFrame = seperator && key ? key.split(seperator)[1] : undefined

	const isFrame = id ? frames[id] === frame : false

	const handleFrame = (key?: string) => {
		const frameKey = key ?? frame

		if (!id) return

		handleFrames(prevFrames => ({
			...prevFrames,
			[id]: frames[id] === frameKey ? undefined : frameKey
		}))
	}

	return { id, key, isFrame, prevFrame, frames, handleFrame }
}
