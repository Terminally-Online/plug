import { FC, useCallback } from "react"

import { getAddress, isAddress } from "viem"
import { useSendTransaction } from "wagmi"

import { useAtom, useAtomValue } from "jotai"

import { Button } from "@/components/shared/buttons/button"
import { Frame } from "@/components/app/frames/base"
import { TokenImage } from "@/components/app/sockets/tokens/token-image"
import { cn, getChainId, NATIVE_TOKEN_ADDRESS, useConnect, useDebounceInline } from "@/lib"
import { api, RouterOutputs } from "@/server/client"
import { useSocket } from "@/state/authentication"
import { columnByIndexAtom, COLUMNS, isFrameAtom, useColumnActions } from "@/state/columns"

import { TransferRecipient } from "./transfer-recipient"
import { TransferTokenImplementation } from "./transfer-implementation"
import { ScrollingError } from "./scrolling-error"
import { base } from "viem/chains"


export const TransferAmountFrame: FC<{
	index: number
	token: NonNullable<RouterOutputs["socket"]["balances"]["positions"]>["tokens"][number]
	color: string
	textColor: string
}> = ({ index, token, color, textColor }) => {
	const [column] = useAtom(columnByIndexAtom(index))
	const frameKey = `${token?.symbol}-transfer-${index === COLUMNS.SIDEBAR_INDEX ? "deposit" : "amount"}`
	const isFrame = useAtomValue(isFrameAtom)(column, frameKey)
	const { frame, navigate } = useColumnActions(index, frameKey)

	const {
		account: { isAuthenticated }
	} = useConnect()
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

	const chain = "base"
	const chainId = getChainId(chain)
	const implementation = token?.implementations.find(implementation => implementation.chain === chain)
	const request = useDebounceInline({
		chainId,
		from,
		inputs: [
			{
				protocol: "plug",
				action: "transfer",
				amount: `${column?.transfer?.precise ?? 0}`,
				token: `${implementation?.contract ?? NATIVE_TOKEN_ADDRESS}:${implementation?.decimals ?? 18}:20`,
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

		toggleSavedMutation.mutate({ id: intent.intentId }, {
			onSuccess: () => {
				navigate({ index, key: COLUMNS.KEYS.ACTIVITY })
				frame(`${intent.intentId}-activity`)
			}
		})

	}, [intent])

	const handleTransactionOnchain = useCallback(() => {
		if (!column || !intent) return

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
		else
			handleTransactionOffchain()
	}, [column, intent])

	if (!token || !column) return null

	return (
		<>
			<Frame
				index={index}
				icon={
					<div className="relative h-8 w-10">
						<TokenImage
							logo={
								token?.icon ||
								`https://token-icons.llamao.fi/icons/tokens/${getChainId(token.implementations[0].chain)}/${token.implementations[0].contract}?h=240&w=240`
							}
							symbol={token.symbol}
							size="sm"
						/>
					</div>
				}
				label={`${index === COLUMNS.SIDEBAR_INDEX ? "Deposit" : "Transfer"}`}
				visible={isFrame}
				handleBack={() =>
					frame(
						index !== COLUMNS.SIDEBAR_INDEX ? `${token.symbol}-transfer-recipient` : `${token.symbol}-token`
					)
				}
				hasChildrenPadding={false}
				hasOverlay
			>
				<div className="mb-4 flex flex-col gap-2">
					{index !== COLUMNS.SIDEBAR_INDEX && (
						<div className="px-6">
							<TransferRecipient
								address={column?.transfer?.recipient ?? ""}
								handleSelect={() => frame(`${token.symbol}-transfer-recipient`)}
							/>
						</div>
					)}

					<div className="flex flex-col gap-2">
						{token.implementations.map((implementation, implementationIndex) => (
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

						<Button
							className={cn(
								"flex w-full items-center justify-center gap-2 rounded-lg border-[1px] py-4 font-bold transition-all duration-200 ease-in-out hover:opacity-90 hover:brightness-105",
								isReady === false && "transparent"
							)}
							style={{
								backgroundColor: isReady ? color : "transparent",
								color: isReady ? textColor : color,
								borderColor: isReady ? "#FFFFFF" : color
							}}
							disabled={(intent && isPending) || isReady === false}
							onClick={intent && !isPending && isReady ? handleTransactionOnchain : () => { }}
							chain={base}
						>
							{!isAuthenticated
								? "Connect Wallet"
								: isPending
									? "Transfering..."
									: isReady
										? index === COLUMNS.SIDEBAR_INDEX
											? "Deposit"
											: "Send"
										: "Enter Amount"}
						</Button>
					</div>
				</div>
			</Frame>
		</>
	)
}
