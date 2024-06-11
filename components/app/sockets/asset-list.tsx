import type { FC } from "react"

import { LoaderCircle } from "lucide-react"

import { useBalances } from "@/contexts"

import { TransferFrame } from "../frames/transfer"
import { SocketAssetItem } from "./asset-item"

type Props = {
	balances: ReturnType<typeof useBalances>["balances"]
	hasFrame?: boolean
	handleSelect?: (
		token: NonNullable<ReturnType<typeof useBalances>["balances"]>[number]
	) => void
}

export const SocketAssetList: FC<Props> = ({
	balances,
	hasFrame = true,
	handleSelect
}) => {
	if (balances === undefined)
		return (
			<div className="my-8 flex items-center justify-center">
				<p className="flex flex-row items-center gap-2">
					<span className="origin-center animate-spin opacity-40">
						<LoaderCircle size={14} />
					</span>
					<span className="opacity-60">Loading...</span>
				</p>
			</div>
		)

	return (
		<>
			<div className="flex flex-col gap-4">
				{balances.map(
					(token, index) =>
						token && (
							<SocketAssetItem
								key={index}
								token={token}
								handleSelect={handleSelect}
							/>
						)
				)}
			</div>

			{hasFrame && <TransferFrame />}
		</>
	)
}
