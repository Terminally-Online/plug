import type { FC, PropsWithChildren } from "react"

import BlockiesSvg from "blockies-react-svg"
import { CheckIcon, ClipboardIcon } from "lucide-react"

import { CaretDownIcon } from "@radix-ui/react-icons"

import { useTabs, useVaults } from "@/contexts"
import { INITIAL_PANE } from "@/contexts/TabsProvider"
import { truncateAddress } from "@/lib/blockchain"
import { useClipboard } from "@/lib/hooks"

export const Selector: FC<PropsWithChildren> = () => {
	const { pane, handlePane } = useTabs()
	const { copy, isCopied } = useClipboard()

	const { expanded } = useTabs()
	const { vault } = useVaults()

	const focused = pane === "vaults"

	return expanded ? (
		<div className="flex flex-row border-b-[1px] border-stone-950">
			{vault ? (
				<>
					<button
						onClick={() =>
							handlePane(focused ? INITIAL_PANE : "vaults")
						}
						className="z-[99] flex h-min w-full cursor-pointer flex-row items-center justify-center p-4 text-white hover:bg-stone-950 active:bg-white active:text-stone-950"
					>
						<BlockiesSvg
							size={8}
							scale={8}
							address={vault.address}
							caseSensitive={true}
							className="mr-4 h-4 w-4 rounded-full"
						/>
						{truncateAddress(vault.address)}
						<CaretDownIcon
							width={12}
							height={12}
							className="ml-auto opacity-60"
						/>
					</button>

					<button
						onClick={() => copy(vault.address)}
						className="flex h-full w-min cursor-pointer flex-row items-center justify-center border-l-[1px] border-stone-950 p-4 text-sm text-white/60 hover:bg-stone-950 active:bg-white active:text-stone-950"
					>
						{isCopied ? (
							<CheckIcon
								width={16}
								height={16}
								className="opacity-60"
							/>
						) : (
							<ClipboardIcon
								width={16}
								height={16}
								className="opacity-60"
							/>
						)}
					</button>
				</>
			) : null}
		</div>
	) : null
}

export default Selector
