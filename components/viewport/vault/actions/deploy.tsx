import type { FC, PropsWithChildren } from "react"
import { useMemo, useState } from "react"

import { Input } from "@/components/ui/input"
import { blockExplorerAddress, truncateAddress } from "@/lib/blockchain"

export const Deploy: FC<PropsWithChildren> = () => {
	const [amount, setAmount] = useState<number | undefined>(0.1)

	const chainId = 1
	const address = "0x62180042606624f02d8a130da8a3171e9b33894d"

	const displayAddress = useMemo(() => truncateAddress(address), [address])

	const explorerUrl = useMemo(() => {
		return blockExplorerAddress(chainId, address)
	}, [chainId, address])

	return (
		<div className="mt-[-46px] flex h-screen w-[360px] flex-col text-center">
			<div className="mt-[46px] space-y-4 border-b-[1px] border-stone-950 py-16">
				<h1 className="mx-auto text-2xl">Create a Vault</h1>
				<p className="px-12 text-sm opacity-60">
					Deploy and fund a new Vault on Base with a single click. You
					can choose to the same address on multiple chains in the
					future.
				</p>
			</div>

			<div className="mt-auto">
				<p className="mb-4 text-sm">
					Deploying To:{" "}
					<span className="opacity-60">
						<a target="_blank" href={explorerUrl}>
							{displayAddress}
						</a>
					</span>
				</p>

				<div className="flex flex-row items-center justify-center border-t-[1px] border-stone-950 hover:bg-stone-950">
					<Input
						name="amount"
						type="number"
						placeholder="AMOUNT"
						autoComplete="off"
						value={amount}
						step="0.01"
						min="0.1"
						onChange={e => setAmount(Number(e.target.value))}
						className="relative mb-auto w-full bg-transparent py-8 uppercase text-white"
					/>

					<p className="p-4 opacity-60">ETH</p>
				</div>

				<button
					onClick={() => {}}
					className="text-md group pointer-events-auto flex h-full h-min w-full items-center justify-center border-b-[1px] border-stone-950 bg-white p-4 text-stone-950 transition-all duration-200 ease-in-out hover:bg-stone-950 hover:text-white active:bg-white active:text-stone-950"
				>
					Submit
				</button>
			</div>
		</div>
	)
}

export default Deploy
