import type { FC, PropsWithChildren } from "react"
import { createContext, useContext, useEffect, useMemo, useState } from "react"

import { useRouter } from "next/router"

export const FrameContext = createContext<{
	frameVisible: string | undefined
	handleFrameVisible: (key: string | undefined) => void
}>({
	frameVisible: undefined,
	handleFrameVisible: () => {}
})

export const FrameProvider: FC<PropsWithChildren> = ({ children }) => {
	const router = useRouter()

	const [frameVisible, handleFrameVisible] = useState<string | undefined>()

	useEffect(() => {
		handleFrameVisible(undefined)
	}, [router.pathname])

	return (
		<FrameContext.Provider
			value={{
				frameVisible,
				handleFrameVisible
			}}
		>
			{children}
		</FrameContext.Provider>
	)
}

export const useFrame = () => useContext(FrameContext)
