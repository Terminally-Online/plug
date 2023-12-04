import { useEffect, useRef, useState } from "react"

import { getRenderLoop, RenderLoop } from "../render"
import CanvasStore from "../store"

export const useRenderLoop = (fps: number = 15) => {
	const [frame, setFrame] = useState("0")
	const loop = useRef<RenderLoop>(
		getRenderLoop(fps, () => {
			if (CanvasStore.shouldRender) {
				setFrame(`${performance.now()}`)
				CanvasStore.shouldRender = false
			}
		})
	)

	useEffect(() => {
		CanvasStore.shouldRender = true
		loop.current.start()

		return () => loop.current.stop()
	}, [])
	return frame
}

export default useRenderLoop
