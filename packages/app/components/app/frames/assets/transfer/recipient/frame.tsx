import { FC } from "react"

import { isAddress } from "viem"

import { SearchIcon } from "lucide-react"

import { useAtom, useAtomValue } from "jotai"

import { TransferRecipient } from "@/components/app/frames/assets/transfer/recipient/recipient"
import { Frame } from "@/components/app/frames/base"
import { Search } from "@/components/app/inputs/search"
import { TokenImage } from "@/components/app/sockets/tokens/token-image"
import { getZerionTokenIconUrl, ZerionPosition } from "@/lib"
import { useAccount } from "@/lib/hooks/account/useAccount"
import { RouterOutputs } from "@/server/client"
import { columnByIndexAtom, isFrameAtom, useColumnActions } from "@/state/columns"
import { useRecipients } from "@/state/recipients"

type TransferRecipientFrameProps = {
	index: number
	token?: ZerionPosition
	collectible?: NonNullable<RouterOutputs["service"]["zerion"]["nfts"]["detail"]["data"]>
	included?: NonNullable<RouterOutputs["service"]["zerion"]["nfts"]["detail"]["included"]>[number]
}

export const TransferRecipientFrame: FC<TransferRecipientFrameProps> = ({ index, token, collectible, included }) => {
	const account = useAccount()

	const [column] = useAtom(columnByIndexAtom(index))
	const frameKey = token
		? `${token.attributes.fungible_info.symbol}-transfer-recipient`
		: `collectible___${collectible?.id}___transfer-recipient`
	const isFrame = useAtomValue(isFrameAtom)(column, frameKey)
	const { frame, transfer } = useColumnActions(index, frameKey)

	const formatRecipientInput = (input: string): string => {
		if (!input) return ""
		if (isAddress(input)) return input
		if (input.startsWith("0x") === false && isAddress(`0x${input}`)) return `0x${input}`
		if (isAddress(input) === false && input.endsWith(".eth") === false) {
			if (input.endsWith(".et")) return `${input}h`
			if (input.endsWith(".e")) return `${input}th`
			if (input.endsWith(".")) return `${input}eth`
			return `${input}.eth`
		}
		return input
	}

	const formattedRecipient = formatRecipientInput(column?.transfer?.recipient ?? "")
	const { recipients, handleRecent } = useRecipients(formattedRecipient)

	const icon = token ? (
		<TokenImage logo={getZerionTokenIconUrl(token)} symbol={token.attributes.fungible_info.symbol} size="sm" />
	) : (
		<div
			className="relative h-8 w-8 rounded-full bg-cover bg-center bg-no-repeat"
			style={{
				backgroundImage: `url(${included?.attributes?.metadata?.icon?.url ?? ""})`
			}}
		/>
	)

	const handleSelect = (address: string) => {
		if (address !== account.address) handleRecent(address)

		transfer(prev => ({
			...prev,
			recipient: address
		}))

		if (address !== "") {
			if (token) {
				frame(`${token.attributes.fungible_info.symbol}-transfer-amount`)
			} else if (collectible) {
				frame(`collectible___${collectible?.id}___transfer-amount`)
			}
		}
	}

	const handleBack = () => {
		if (token) {
			frame(`${token.attributes.fungible_info.symbol}-token`)
		} else if (collectible) {
			frame(`collectible___${collectible?.id}`)
		}
	}

	if (!token && !collectible) return null

	return (
		<Frame
			index={index}
			className="min-h-[480px]"
			icon={<div className="relative h-8 w-10">{icon}</div>}
			label="Transfer Recipient"
			visible={isFrame}
			handleBack={handleBack}
			hasChildrenPadding={false}
			hasOverlay
		>
			<div className="relative flex h-full flex-col gap-2 px-6 pb-4">
				<Search
					icon={<SearchIcon size={14} className="opacity-60" />}
					placeholder="Search addresses or ENS"
					search={column?.transfer?.recipient ?? ""}
					handleSearch={recipient => transfer(prev => ({ ...prev, recipient }))}
					clear
				/>

				{/* Current recipient */}
				<TransferRecipient address={formattedRecipient} handleSelect={handleSelect} />

				{/* Connected wallet (if not current recipient) */}
				{account.address && column?.transfer?.recipient !== account.address && (
					<TransferRecipient address={account.address as string} handleSelect={handleSelect} />
				)}

				{/* Recent recipients */}
				{recipients.length > 0 ? (
					recipients
						.filter(recipient => recipient !== "")
						.slice(0, column?.transfer?.recipient === account.address ? 6 : 5)
						.map(recipient => (
							<TransferRecipient
								key={recipient}
								address={recipient}
								handleSelect={handleSelect}
								isRecent
							/>
						))
				) : (
					<p className="mx-auto my-24 max-w-[320px] text-sm font-bold opacity-40">
						Recently used recipients will appear here.
					</p>
				)}
			</div>
		</Frame>
	)
}
