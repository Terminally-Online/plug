import type { FC, PropsWithChildren } from "react"
import { useEffect, useRef } from "react"

import { SewingPinIcon } from "@radix-ui/react-icons"

import { api } from "@/lib/api"
import { useDebounce } from "@/lib/hooks/useDebounce"
import CanvasStore from "@/lib/store"

import { Input } from "../ui/input"

export type ToolbarProps = { id: string; name: string }

export const Toolbar: FC<PropsWithChildren<ToolbarProps>> = ({ id, name }) => {
	const rightSideRef = useRef<HTMLDivElement>(null)

	const { debounce, value, debounced } = useDebounce({ initial: name })

	const updateCanvas = api.canvas.update.useMutation()

	const zoom = Math.round(CanvasStore.scale.x * 100)

	const handleCenter = () => {
		CanvasStore.centerCamera()
	}

	useEffect(() => {
		const handleZoom = (e: KeyboardEvent) => {
			const keys = [
				["+", -10],
				["-", 10],
				["=", -10]
			]

			for (const [key, value] of keys) {
				if (e.key === key) {
					e.preventDefault()

					const zoom =
						typeof value === "number" ? value : parseInt(value)

					CanvasStore.zoomCamera(zoom, zoom)
				}
			}
		}

		window.addEventListener("keydown", handleZoom)

		return () => window.removeEventListener("keydown", handleZoom)
	}, [])

	useEffect(() => {
		if (!debounced || !updateCanvas) return
		if (name === debounced) return

		updateCanvas.mutate({
			id,
			name: debounced
		})
	}, [debounced, id, name, updateCanvas])

	return (
		<>
			<div className="fixed left-0 right-0 top-12 border-b-[1px] border-stone-950">
				<div className="relative left-0 flex flex-row items-center bg-stone-900 text-white">
					<div className="flex w-full items-center justify-center">
						<Input
							value={value ?? name}
							placeholder={name ?? "Untitled Canvas"}
							className="max-w-[380px] flex-grow overflow-hidden text-ellipsis bg-stone-900 p-4 text-center text-xs tabular-nums text-white/60 placeholder:text-muted-foreground"
							style={{
								marginLeft: rightSideRef.current?.clientWidth
							}}
							onChange={e => debounce(e.target.value)}
							onBlur={() => {
								// ? If they deleted the name, but left it empty, return to the previous value.
								if (value === "") debounce(name)
							}}
						/>
					</div>

					<div ref={rightSideRef} className="flex flex-row">
						<button
							className="group ml-auto w-min border-l-[1px] border-stone-950 p-4 text-center text-xs tabular-nums text-white/60 transition-all duration-200 ease-in-out hover:bg-white hover:text-stone-950"
							onClick={handleCenter}
						>
							<SewingPinIcon
								className="group:hover:opacity-100 opacity-60"
								width={16}
								height={16}
							/>
						</button>

						<p className="border-l-[1px] border-stone-950 p-4 text-center text-xs tabular-nums text-white/60">
							{zoom}%
						</p>
					</div>
				</div>
			</div>

			<div className="fixed bottom-0 left-0 right-0 flex flex-row items-center gap-2 border-t-[1px] border-stone-950 bg-stone-900">
				<div className="mx-auto p-4 text-white">
					<p className="text-xs opacity-60">
						Tip: Double click anywhere to start a new plug
					</p>
				</div>
			</div>

			{
				<div className="fixed right-0 top-24 z-10 m-2 rounded-sm bg-red-400 p-2 font-bold text-red-700">
					<p>
						Camera: {Math.round(CanvasStore.camera.x)},{" "}
						{Math.round(CanvasStore.camera.y)},{" "}
						{Math.round(CanvasStore.camera.z)}
					</p>
					<p>
						Container: {CanvasStore.container.width} x{" "}
						{CanvasStore.container.height}
					</p>
					<p>Locked: {CanvasStore.camera.locked.toString()}</p>
					<p>
						Scale: {Math.round(CanvasStore.scale.x)},{" "}
						{Math.round(CanvasStore.scale.y)}
					</p>
					<p>
						Screen: {Math.round(CanvasStore.screen.x)},{" "}
						{Math.round(CanvasStore.screen.y)}
					</p>
					<p>
						Pointer: {Math.round(CanvasStore.pointer.x)},{" "}
						{Math.round(CanvasStore.pointer.y)}
					</p>
				</div>
			}
		</>
	)
}

export default Toolbar
