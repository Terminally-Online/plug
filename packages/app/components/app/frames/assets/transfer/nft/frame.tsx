import { FC, useCallback, useMemo } from "react"

import { useAtom, useAtomValue } from "jotai"

import { TransferRecipient } from "@/components/app/frames/assets/transfer/recipient/recipient"
import { Frame } from "@/components/app/frames/base"
import { CollectibleImage } from "@/components/app/sockets/collectibles/collectible-image"
import { Accordion } from "@/components/shared/utils/accordion"
import { cn, formatTitle, getChainId, useStateDebounce } from "@/lib"
import { api, RouterOutputs } from "@/server/client"
import { columnByIndexAtom, COLUMNS, isFrameAtom, useColumnActions } from "@/state/columns"
import { getAddress, isAddress } from "viem"
import { useSocket } from "@/state/authentication"
import { useAccount } from "@/lib/hooks/account/useAccount"
import { useSendTransaction } from "wagmi"
import { TransferSFTAmount } from "../sft/amount"
import { ScrollingError } from "../../scrolling-error"
import { ChainSpecificButton } from "@/components/shared/buttons/authenticate"
import { base } from "viem/chains"

type TransferNFTFrameProps = {
	index: number
	collectible?: NonNullable<RouterOutputs["service"]["zerion"]["nfts"]["detail"]["data"]>
	included?: NonNullable<RouterOutputs["service"]["zerion"]["nfts"]["detail"]["included"]>[number]
	color: string
	textColor: string
}
export const TransferNFTFrame: FC<TransferNFTFrameProps> = ({ index, collectible, included, color, textColor }) => {
	const [column] = useAtom(columnByIndexAtom(index))
	const frameKey = `collectible___${collectible?.id}___transfer-amount`
	const isFrame = useAtomValue(isFrameAtom)(column, frameKey)
	const { frame, navigate } = useColumnActions(index, frameKey)

	const { isAuthenticated } = useAccount()
	const { socket } = useSocket()
	const { error, sendTransaction, isPending } = useSendTransaction()

	const chainId = getChainId("base")
	const from = socket
		? index === COLUMNS.SIDEBAR_INDEX
			? getAddress(socket.id)
			: getAddress(socket.socketAddress)
		: ""

	const balance = 1
	const token = `${collectible?.attributes?.contract_address}:${collectible?.attributes?.token_id}:${collectible?.attributes?.interface === "ERC1155" ? 1155 : 721}`
	const amount = useStateDebounce<string>(`${parseInt(column?.transfer?.precise ?? "0")}`)
	const recipient =
		column && socket
			? index === COLUMNS.SIDEBAR_INDEX
				? getAddress(socket.socketAddress)
				: column.transfer?.recipient && isAddress(column.transfer?.recipient)
					? getAddress(column.transfer?.recipient)
					: ""
			: ""

	const isEOA = column && column.index === COLUMNS.SIDEBAR_INDEX
	const options = { isEOA, simulate: true }

	const isReady = useMemo(() => {
		const amount = parseInt(column?.transfer?.precise ?? "0")

		if (isNaN(amount)) return false

		return amount > 0 && amount <= balance
	}, [column?.transfer?.precise, balance])
	const enabled = !!column && !!socket && isAuthenticated && isReady && isFrame

	const { data: intent } = api.solver.actions.intent.useQuery({
		chainId,
		from,
		inputs: [{
			protocol: "plug",
			action: "transfer",
			token,
			amount,
			recipient
		}],
		options
	}, {
		enabled,
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

	return (
		<Frame
			index={index}
			icon={
				<div className="relative h-8 w-10">
					<div
						className="absolute left-1/2 top-1/2 -translate-x-1/2 -translate-y-1/2 animate-fade-in rounded-full bg-plug-green/10 blur-2xl filter"
						style={{
							backgroundImage: `url(${included?.attributes?.metadata?.icon?.url})`,
							backgroundSize: "cover",
							backgroundPosition: "center",
							backgroundRepeat: "no-repeat",
							width: "4rem",
							minWidth: "4rem",
							height: "4rem"
						}}
					/>
					<div
						className="absolute left-1/2 top-1/2 -translate-x-1/2 -translate-y-1/2 animate-fade-in rounded-full bg-plug-green/10"
						style={{
							backgroundImage: `url(${included?.attributes?.metadata?.icon?.url})`,
							backgroundSize: "cover",
							backgroundPosition: "center",
							backgroundRepeat: "no-repeat",
							width: "2rem",
							minWidth: "2rem",
							height: "2rem"
						}}
					/>
				</div>
			}
			label="Transfer"
			visible={isFrame}
			handleBack={() =>
				frame(index !== COLUMNS.SIDEBAR_INDEX
					? `collectible___${collectible?.id}___transfer-recipient`
					: `collectible___${collectible?.id}`
				)
			}
			hasChildrenPadding={false}
			hasOverlay
		>
			<div className="relative my-4 flex flex-col gap-2">
				<div className="flex flex-col gap-2">
					{index !== COLUMNS.SIDEBAR_INDEX && (
						<div className="px-6">
							<TransferRecipient
								address={column?.transfer?.recipient ?? ""}
								handleSelect={() => frame(`collectible___${collectible?.id}___transfer-recipient`)}
							/>
						</div>
					)}

					{collectible?.attributes?.interface !== "ERC1155" ? (
						<div className="px-6">
							<Accordion>
								<div className="flex w-full flex-row items-center gap-4">
									<div className="relative h-10 w-10 min-w-10">
										<CollectibleImage
											video={collectible?.attributes?.metadata?.content?.video?.url}
											image={collectible?.attributes?.metadata?.content?.detail?.url}
											name={collectible?.attributes?.metadata?.name ?? ""}
										/>
									</div>
									<div className="flex w-full flex-col truncate overflow-ellipsis">
										<div className="flex flex-row items-center justify-between">
											<p className="mr-auto font-bold">
												{formatTitle(collectible?.attributes?.metadata?.name || included?.attributes?.metadata?.name || "")}
											</p>
										</div>
										<div className="flex flex-row items-center justify-between">
											<p className="mr-auto truncate overflow-ellipsis text-sm font-bold tabular-nums opacity-40">
												#{collectible?.attributes?.token_id}
											</p>
										</div>
									</div>
								</div>
							</Accordion>
						</div>
					) : (
						<TransferSFTAmount
							index={index}
							collectible={collectible}
							included={included}
							color={color}
						/>
					)}
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
