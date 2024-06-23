import { useEffect, useRef } from "react"

import { PencilLine, Settings } from "lucide-react"

import { Frame } from "@/components/app/frames/base"
import { Button } from "@/components/buttons"
import { Checkbox, Search } from "@/components/inputs"
import { useFrame, usePlugs } from "@/contexts"
import { colors, useDebounce } from "@/lib"
import { useNavigation } from "@/lib/hooks/useNavigation"

export const ManageFrame = () => {
	const { id, from } = useNavigation()
	const { frameVisible, handleFrameVisible } = useFrame()
	const { plug, handle } = usePlugs()

	const [name, debouncedName, handleName, nameRef] = useDebounce(
		plug?.name ?? "",
		1000
	)

	useEffect(() => {
		if (!plug || nameRef.current === debouncedName) return

		nameRef.current = debouncedName

		handle.plug.edit({ ...plug, name: debouncedName })
	}, [nameRef, plug, debouncedName, handle])

	if (!plug) return null

	return (
		<Frame
			className="z-[2]"
			icon={<Settings size={18} className="opacity-60" />}
			label="Manage Plug"
			visible={frameVisible === "manage"}
			handleVisibleToggle={() => handleFrameVisible(undefined)}
		>
			<div className="flex flex-col gap-4">
				<Search
					icon={<PencilLine size={14} className="opacity-60" />}
					placeholder="Plug name"
					search={name}
					handleSearch={handleName}
				/>

				<div className="flex flex-row items-center gap-2">
					<p className="mr-auto font-bold">Private</p>

					<Checkbox
						checked={plug.isPrivate}
						handleChange={(checked: boolean) =>
							handle.plug.edit({
								...plug,
								isPrivate: checked
							})
						}
					/>
				</div>

				<div className="flex flex-row items-center gap-2">
					<p className="font-bold">Color</p>

					<div className="ml-auto flex flex-wrap items-center gap-1">
						{Object.keys(colors).map(color => (
							<div
								key={color}
								className="group flex h-6 w-6 cursor-pointer items-center justify-center rounded-full border-[2px]"
								style={{
									borderColor:
										plug.color === color
											? colors[
													color as keyof typeof colors
												]
											: "transparent"
								}}
								onClick={() =>
									handle.plug.edit({
										...plug,
										color
									})
								}
							>
								<div
									className="h-full w-full rounded-full border-[2px] border-white transition-all duration-200 ease-in-out"
									style={{
										backgroundColor:
											colors[color as keyof typeof colors]
									}}
								/>
							</div>
						))}
					</div>
				</div>
			</div>

			<Button
				variant="destructive"
				className="mt-[20px] w-full"
				onClick={() => handle.plug.delete({ id, from })}
			>
				Delete
			</Button>
		</Frame>
	)
}
