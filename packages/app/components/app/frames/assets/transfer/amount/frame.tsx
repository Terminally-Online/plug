import { FC, useCallback } from "react"

import { getAddress, isAddress } from "viem"
import { base } from "viem/chains"
import { useSendTransaction } from "wagmi"

import { useAtom, useAtomValue } from "jotai"

import { TransferTokenImplementation } from "@/components/app/frames/assets/transfer/amount/implementation"
import { TransferRecipient } from "@/components/app/frames/assets/transfer/recipient/recipient"
import { Frame } from "@/components/app/frames/base"
import { TokenImage } from "@/components/app/sockets/tokens/token-image"
import { ChainSpecificButton } from "@/components/shared/buttons/authenticate"
import { cn, getZerionTokenIconUrl, NATIVE_TOKEN_ADDRESS, useDebounceInline, ZerionPosition } from "@/lib"
import { useAccount } from "@/lib/hooks/account/useAccount"
import { api } from "@/server/client"
import { useSocket } from "@/state/authentication"
import { columnByIndexAtom, COLUMNS, isFrameAtom, useColumnActions } from "@/state/columns"

import { ScrollingError } from "../../scrolling-error"

type TransferAmountFrameProps = {
	index: number
	token: ZerionPosition
	color: string
	textColor: string
}

export const TransferAmountFrame: FC<TransferAmountFrameProps> = ({ index, token, color, textColor }) => {
	const [column] = useAtom(columnByIndexAtom(index))
	const frameKey = `${token.attributes.fungible_info.symbol}-transfer-${index === COLUMNS.SIDEBAR_INDEX ? "deposit" : "amount"}`
	const isFrame = useAtomValue(isFrameAtom)(column, frameKey)
	const { frame, navigate } = useColumnActions(index, frameKey)

	const { isAuthenticated } = useAccount()
	const { socket } = useSocket()
	const { error, sendTransaction, isPending } = useSendTransaction()

	const isReady = token && column && parseFloat(column?.transfer?.precise ?? "0") > 0 && !isPending
	const from = socket
		? index === COLUMNS.SIDEBAR_INDEX
			? getAddress(socket.id)
			: getAddress(socket.socketAddress)
		: ""
	const recipient =
		column && socket
			? index === COLUMNS.SIDEBAR_INDEX
				? getAddress(socket.socketAddress)
				: column.transfer?.recipient && isAddress(column.transfer?.recipient)
					? getAddress(column.transfer?.recipient)
					: ""
			: ""

	const implementation = token.attributes.fungible_info.implementations.find(
		implementation => implementation.chain_id === "base"
	)
	const request = useDebounceInline({
		chainId: base.id,
		from,
		inputs: [
			{
				protocol: "plug",
				action: "transfer",
				amount: `${column?.transfer?.precise ?? 0}`,
				token: `${implementation?.address ?? NATIVE_TOKEN_ADDRESS}:${implementation?.decimals ?? 18}:20`,
				recipient
			}
		],
		options: {
			isEOA: column && column.index === COLUMNS.SIDEBAR_INDEX,
			simulate: true
		}
	})
	const { data: intent } = api.solver.actions.intent.useQuery(request, {
		enabled: isFrame && isReady && !!column && !!socket && !!implementation
	})

	const toggleSavedMutation = api.plugs.activity.toggleSaved.useMutation()
	const handleTransactionOffchain = useCallback(() => {
		if (!intent) return

		toggleSavedMutation.mutate(
			{ id: intent.intentId },
			{
				onSuccess: () => {
					navigate({ index, key: COLUMNS.KEYS.ACTIVITY })
					frame(`${intent.intentId}-activity`)
				}
			}
		)
	}, [intent, frame, index, navigate, toggleSavedMutation])

	const handleTransactionOnchain = useCallback(() => {
		if (!column || !intent || !isReady || isPending) return

		if (column.index === COLUMNS.SIDEBAR_INDEX)
			sendTransaction(
				{
					to: intent.transactions[0].to,
					data: intent.transactions[0].data,
					value: intent.transactions[0].value
				},
				{
					onSuccess: handleTransactionOffchain
				}
			)
		else handleTransactionOffchain()
	}, [column, intent, handleTransactionOffchain, isPending, isReady, sendTransaction])

	const isDisabled = (intent && isPending) || isReady === false

	if (!token || !column) return null

	return (
		<Frame
			index={index}
			icon={
				<div className="relative h-8 w-10">
					<TokenImage
						logo={getZerionTokenIconUrl(token)}
						symbol={token.attributes.fungible_info.symbol}
						size="sm"
					/>
				</div>
			}
			label={`${index === COLUMNS.SIDEBAR_INDEX ? "Deposit" : "Transfer"}`}
			visible={isFrame}
			handleBack={() =>
				frame(
					index !== COLUMNS.SIDEBAR_INDEX
						? `${token.attributes.fungible_info.symbol}-transfer-recipient`
						: `${token.attributes.fungible_info.symbol}-token`
				)
			}
			hasChildrenPadding={false}
			hasOverlay
		>
			<div className="py-4 flex flex-col gap-2">
				{index !== COLUMNS.SIDEBAR_INDEX && (
					<div className="px-6">
						<TransferRecipient
							address={column?.transfer?.recipient ?? ""}
							handleSelect={() => frame(`${token.attributes.fungible_info.symbol}-transfer-recipient`)}
						/>
					</div>
				)}

				<div className="flex flex-col gap-2">
					{token.attributes.fungible_info.implementations
						// TODO: (#797) Implement support based on `connectedChains`
						.filter(implementation => implementation.chain_id === "base")
						.filter(implementation => implementation.balance && implementation.balance > 0)
						.map((implementation, implementationIndex) => (
							<TransferTokenImplementation
								key={implementationIndex}
								index={index}
								implementation={implementation}
								token={token}
								color={color}
							/>
						))}
				</div>

				<div className="mx-6 mt-2 flex flex-col gap-4">
					<ScrollingError error={error?.message ?? ""} />

					<ChainSpecificButton
						className={cn(
							"flex w-full items-center justify-center gap-2 rounded-lg border-[1px] py-4 font-bold transition-all duration-200 ease-in-out hover:opacity-90 hover:brightness-105",
							isReady === false && "transparent"
						)}
						style={{
							backgroundColor: isReady ? color : "transparent",
							color: isReady ? textColor : color,
							borderColor: isReady ? "#FFFFFF" : color
						}}
						chainId={base.id}
						onClick={handleTransactionOnchain}
						disabled={isDisabled}
					>
						{!isAuthenticated
							? "Connect Wallet"
							: isPending
								? "Transferring..."
								: isReady
									? index === COLUMNS.SIDEBAR_INDEX
										? "Deposit"
										: "Send"
									: "Enter Amount"}
					</ChainSpecificButton>
				</div>
			</div>
		</Frame>
	)
}
