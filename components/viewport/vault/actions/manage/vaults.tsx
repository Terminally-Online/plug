import { type FC, type PropsWithChildren } from "react"

import BlockiesSvg from "blockies-react-svg"
import { PlusIcon } from "lucide-react"

import { CheckIcon } from "@radix-ui/react-icons"

import { useTabs, useVaults } from "@/contexts"
import { INITIAL_PANE } from "@/contexts/TabsProvider"
import { truncateAddress } from "@/lib/blockchain"

export const Vaults: FC<PropsWithChildren> = () => {
	const { handlePane } = useTabs()
	const { vault, vaults, handleAdd, handleSelect } = useVaults()

	const handleVaultSelect = (address: string) => {
		handlePane(INITIAL_PANE)
		handleSelect(address)
	}

	return (
		<div className="mt-[-46px] h-screen w-[360px]">
			{vault && vaults ? (
				<div className="mt-[46px] flex w-full flex-col text-center">
					{vaults
						.filter(
							filteredVault =>
								filteredVault.address != vault.address
						)
						.map((vault, index) => (
							<button
								key={index}
								onClick={() => handleVaultSelect(vault.address)}
								className="flex w-full items-center justify-center border-b-[1px] border-stone-950 p-4 transition-all duration-200 ease-in-out hover:bg-stone-950"
							>
								<BlockiesSvg
									size={8}
									scale={8}
									address={vault.address}
									caseSensitive={true}
									className="mr-4 h-4 w-4 rounded-full"
								/>
								{vault.name || truncateAddress(vault.address)}
								<span className="ml-auto">
									<CheckIcon
										className="ml-2 opacity-60"
										width={16}
										height={16}
									/>
								</span>
							</button>
						))}
				</div>
			) : null}

			<button
				onClick={() => handleAdd([1])}
				className="group pointer-events-auto flex h-min w-full items-center justify-center border-b-[1px] border-stone-950 p-4 transition-all duration-200 ease-in-out hover:bg-stone-950"
			>
				<PlusIcon className="mr-2 opacity-60" width={16} height={16} />
				Create New
			</button>
		</div>
	)
}
