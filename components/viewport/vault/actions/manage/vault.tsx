import { type FC, type PropsWithChildren } from "react"

import {
	ArchiveIcon,
	ArrowBottomLeftIcon,
	ArrowTopRightIcon,
	ListBulletIcon
} from "@radix-ui/react-icons"

import { Input } from "@/components/ui/input"
import { useTabs, useVaults } from "@/contexts"
import { api } from "@/lib/api"

export const Vault: FC<PropsWithChildren> = () => {
	const { pane, Panel, handlePane } = useTabs()
	const { vault } = useVaults()

	const handleVaultName = api.account.vaults.rename.useMutation()

	return (
		<div className="w-[360px]">
			<Input
				name="amount"
				type="text"
				placeholder="VAULT NAME"
				autoComplete="off"
				value={vault?.name ?? undefined}
				onChange={e =>
					handleVaultName.mutate({
						address: vault?.address || "",
						name: e.target.value
					})
				}
				className="relative mb-auto w-full border-b-[1px] border-stone-950 bg-transparent py-8 uppercase text-white outline-none transition-all duration-200 ease-in-out hover:bg-stone-950"
			/>

			<div className="w-full border-b-[1px] border-stone-950 text-center">
				<div className="flex flex-row border-b-[1px] border-stone-950">
					<button
						onClick={() => handlePane("tokens")}
						className={`text-md group pointer-events-auto flex h-full w-1/2 items-center justify-center border-r-[1px] border-stone-950 p-4 outline-none transition-all duration-200 ease-in-out hover:bg-stone-950 hover:text-white active:bg-white active:text-stone-950 ${
							pane === "tokens" ? "active" : ""
						}`}
					>
						<ArchiveIcon
							className="mr-2 opacity-60"
							width={16}
							height={16}
						/>
						Tokens
					</button>

					<button
						onClick={() => handlePane("activity")}
						className={`text-md group pointer-events-auto flex h-full w-1/2 items-center justify-center border-r-[1px] border-stone-950 p-4 outline-none transition-all duration-200 ease-in-out hover:bg-stone-950 hover:text-white active:bg-white active:text-stone-950 ${
							pane === "activity" ? "active" : ""
						}`}
					>
						<ListBulletIcon
							className="mr-2 opacity-60"
							width={16}
							height={16}
						/>
						Activity
					</button>
				</div>
				<div className="flex flex-row">
					<button
						onClick={() => handlePane("deposit")}
						className={`text-md group pointer-events-auto flex h-full w-1/2 items-center justify-center border-r-[1px] border-stone-950 p-4 outline-none transition-all duration-200 ease-in-out hover:bg-stone-950 hover:text-white active:bg-white active:text-stone-950 ${
							pane === "deposit" ? "active" : ""
						}`}
					>
						<ArrowBottomLeftIcon
							className="mr-2 opacity-60"
							width={16}
							height={16}
						/>
						Deposit
					</button>

					<button
						onClick={() => handlePane("withdraw")}
						className={`text-md group pointer-events-auto flex h-full w-1/2 items-center justify-center border-r-[1px] border-stone-950 p-4 outline-none transition-all duration-200 ease-in-out hover:bg-stone-950 hover:text-white active:bg-white active:text-stone-950 ${
							pane === "withdraw" ? "active" : ""
						}`}
					>
						<ArrowTopRightIcon
							className="mr-2 opacity-60"
							width={16}
							height={16}
						/>
						Withdraw
					</button>
				</div>
			</div>

			<div className="flex h-full w-full flex-col">{Panel}</div>
		</div>
	)
}

export default Vault
