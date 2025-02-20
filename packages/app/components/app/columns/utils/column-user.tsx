import { formatAddress } from "@/lib"
import { useColumnStore } from "@/state/columns"
import { TransferRecipient } from "../../frames/assets/transfer-recipient"

export const ColumnUser = ({ index }: { index: number }) => {
	const { column } = useColumnStore(index)

	if (!column || !column.item) return null

	return (
		<div className="p-4">
			<TransferRecipient
				address={column?.transfer?.recipient ?? ""}
				handleSelect={() => {}}
			/>
		</div>
	)
}
