import { FC } from "react"

import { isAddress } from "viem"

import { SearchIcon } from "lucide-react"

import { useAtom, useAtomValue } from "jotai"

import { TransferRecipient } from "@/components/app/frames/assets/transfer/recipient/recipient"
import { Frame } from "@/components/app/frames/base"
import { Search } from "@/components/app/inputs/search"
import { TokenImage } from "@/components/app/sockets/tokens/token-image"
import { getChainId } from "@/lib"
import { useAccount } from "@/lib/hooks/account/useAccount"
import { RouterOutputs } from "@/server/client"
import { columnByIndexAtom, isFrameAtom, useColumnActions } from "@/state/columns"
import { useRecipients } from "@/state/recipients"

type TokenType = NonNullable<RouterOutputs["socket"]["balances"]["positions"]>["tokens"][number]
type CollectibleType = NonNullable<RouterOutputs["socket"]["balances"]["collectibles"]>[number]["collectibles"][number]
type CollectionType = NonNullable<RouterOutputs["socket"]["balances"]["collectibles"]>[number]

type TransferRecipientFrameProps = {
	index: number
	token?: TokenType
	collectible?: CollectibleType
	collection?: CollectionType
}

export const TransferRecipientFrame: FC<TransferRecipientFrameProps> = ({ index, token, collectible, collection }) => {
	const account = useAccount()

	const [column] = useAtom(columnByIndexAtom(index))
	const frameKey = token
		? `${token.symbol}-transfer-recipient`
		: `${collection?.address}-${collection?.chain}-${collectible?.tokenId}-transfer-recipient`
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
		<TokenImage
			logo={
				token.icon ||
				`https://token-icons.llamao.fi/icons/tokens/${getChainId(token.implementations[0].chain)}/${token.implementations[0].contract}?h=240&w=240`
			}
			symbol={token.symbol}
			size="sm"
		/>
	) : (
		<div
			className="relative h-8 w-8 rounded-full bg-cover bg-center bg-no-repeat"
			style={{
				backgroundImage: `url(${collection?.iconUrl})`
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
				frame(`${token.symbol}-transfer-amount`)
			} else if (collectible && collection) {
				frame(`${collection.address}-${collection.chain}-${collectible.tokenId}-transfer-amount`)
			}
		}
	}

	const handleBack = () => {
		if (token) {
			frame(`${token.symbol}-token`)
		} else if (collectible && collection) {
			frame(`${collection.address}-${collection.chain}-${collectible.tokenId}`)
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
