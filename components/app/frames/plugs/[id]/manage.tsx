import { FC, useEffect } from "react"

import { PencilLine, Settings } from "lucide-react"

import { Button, Checkbox, Frame, Search } from "@/components"
import { cardColors, useDebounce } from "@/lib"
import { useColumnStore, usePlugStore } from "@/state"

export const ManagePlugFrame: FC<{ index: number; item: string; from?: string }> = ({ index, item, from }) => {
	const { isFrame } = useColumnStore(index, `${item}-manage`)
	const { plug, handle } = usePlugStore(item)

	const [name, debouncedName, handleName, nameRef] = useDebounce(plug?.name ?? "", 1000)

	useEffect(() => {
		if (!plug || nameRef.current === debouncedName) return

		nameRef.current = debouncedName

		handle.plug.edit({ ...plug, name: debouncedName })
	}, [nameRef, plug, debouncedName, handle])

	if (!plug) return null

	return (
		<Frame
			index={index}
			className="z-[2]"
			icon={<Settings size={18} />}
			label="Manage Plug"
			visible={isFrame}
			hasChildrenPadding={false}
		>
			<div className="mb-4 flex flex-col gap-4 px-6">
				<Search
					icon={<PencilLine size={14} />}
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
						{Object.keys(cardColors).map(color => (
							<div
								key={color}
								className="group flex h-6 w-6 cursor-pointer items-center justify-center rounded-full border-[2px]"
								style={{
									borderColor:
										plug.color === color
											? cardColors[color as keyof typeof cardColors]
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
										background: cardColors[color as keyof typeof cardColors]
									}}
								/>
							</div>
						))}
					</div>
				</div>

				<Button
					variant="destructive"
					className="w-full"
					onClick={() => handle.plug.delete({ plug: plug.id, index, from })}
				>
					Delete
				</Button>
			</div>
		</Frame>
	)
}
