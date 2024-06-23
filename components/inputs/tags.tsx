import { usePlugs } from "@/contexts/PlugProvider"
import { tags } from "@/lib/constants"
import { cn } from "@/lib/utils"

export const Tags = () => {
	const { tag, handle } = usePlugs()

	return (
		<div className="relative mb-[20px] mt-2">
			<div className="scrollbar-hide flex flex-row gap-2 overflow-x-auto whitespace-nowrap">
				{tags.map((tagItem, index) => (
					<button
						key={tag}
						className={cn(
							"cursor-pointer rounded-full p-2 px-4 text-xs font-bold transition-all duration-200 ease-in-out",
							tagItem === tag
								? "bg-grayscale-100"
								: "bg-grayscale-0",
							index === 0 && "ml-4",
							index === tags.length - 1 && "mr-24"
						)}
						onClick={() => handle.tag(tag)}
					>
						{tagItem}
					</button>
				))}
			</div>

			<div className="absolute bottom-0 right-0 top-0 w-24 bg-gradient-to-r from-white/0 to-white" />
		</div>
	)
}
