import { FC, PropsWithChildren } from "react"

import { PinLeftIcon, PinRightIcon } from "@radix-ui/react-icons"

import { useTabs } from "@/contexts/TabsProvider"

export const Toggler: FC<PropsWithChildren> = () => {
	const { expanded, handleExpanded } = useTabs()

	return (
		<button
			onClick={handleExpanded}
			className="text-md pointer-events-auto ml-auto flex h-full items-center justify-center bg-stone-800 p-2 px-4 font-bold text-white/60 outline-none transition-all duration-200 ease-in-out hover:bg-white hover:text-stone-950"
		>
			{expanded ? (
				<PinRightIcon width={16} height={16} />
			) : (
				<PinLeftIcon width={16} height={16} />
			)}
		</button>
	)
}

export default Toggler
