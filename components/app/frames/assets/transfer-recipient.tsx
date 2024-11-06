import Image from "next/image"
import { FC, HTMLAttributes, useEffect } from "react"

import { isAddress, zeroAddress } from "viem"
import { useEnsAddress, useEnsAvatar, useEnsName } from "wagmi"

import { SearchIcon } from "lucide-react"

import { Accordion, Avatar, Frame, Search, TokenImage } from "@/components"
import { formatAddress, getChainId, greenGradientStyle, useConnect } from "@/lib"
import { RouterOutputs } from "@/server/client"
import { useColumns, useRecipients } from "@/state"

type TokenType = NonNullable<RouterOutputs["socket"]["balances"]["positions"]>["tokens"][number]
type CollectibleType = NonNullable<RouterOutputs["socket"]["balances"]["collectibles"]>[number]["collectibles"][number]
type CollectionType = NonNullable<RouterOutputs["socket"]["balances"]["collectibles"]>[number]

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

export const TransferRecipient: FC<
	HTMLAttributes<HTMLDivElement> & {
		isRecent?: boolean
		address: string
		handleSelect: (address: string) => void
	}
> = ({ isRecent = false, address = "", handleSelect, className }) => {
	const { account } = useConnect()

	const { data: ensAddress } = useEnsAddress({
		name: address,
		query: {
			enabled: address?.endsWith("eth") || false
		}
	})
	const {
		data: ensName,
		isFetching,
		isFetched
	} = useEnsName({
		address: ensAddress ?? (address as `0x${string}`),
		query: {
			enabled: (ensAddress ?? address ?? "").startsWith("0x") === true
		}
	})
	const { data: ensAvatar } = useEnsAvatar({
		name: ensName ?? "",
		query: {
			enabled: ensName !== undefined || address?.endsWith(".eth") || false
		}
	})

	const Badge = () => {
		if (isRecent)
			return (
				<p className="rounded-md bg-plug-green/10 px-2 py-1 text-sm">
					<span style={{ ...greenGradientStyle }}>Recent</span>
				</p>
			)
		if (account.address === address)
			return (
				<p className="rounded-md bg-grayscale-0 px-2 py-1 text-sm">
					<span className="opacity-40">Connected</span>
				</p>
			)
		return null
	}

	return (
		<Accordion className={className} onExpand={() => handleSelect(address)}>
			<div className="flex flex-row items-center gap-4">
				<div className="relative h-10 w-10 min-w-10 rounded-sm">
					{ensAvatar ? (
						<>
							<Image
								className="absolute left-0 left-1/2 top-1/2 h-12 w-48 -translate-x-1/2 blur-2xl filter"
								src={ensAvatar}
								alt="ENS Avatar"
								width={240}
								height={240}
							/>
							<Image
								className="relative rounded-sm"
								src={ensAvatar}
								alt="ENS Avatar"
								width={240}
								height={240}
							/>
						</>
					) : (
						<>
							<div className="absolute left-0 top-1/2 blur-lg filter">
								<Avatar name={address} className="rounded-sm" />
							</div>
							<Avatar name={address} className="rounded-sm" />
						</>
					)}
				</div>
				<div className="flex w-max flex-col truncate overflow-ellipsis text-left">
					<p className="font-bold">
						{isFetching && isFetched === false
							? "Loading..."
							: isFetched
								? ((address.includes("eth") ? address : ensName) ?? formatAddress(address))
								: address === ""
									? "Enter your search"
									: "No match found"}
					</p>
					<p className="text-sm font-bold opacity-40">
						{isFetching && isFetched === false
							? formatAddress(zeroAddress)
							: address.endsWith(".eth") && !ensName && isFetched === false
								? "No ENS record found for this address."
								: address.endsWith(".eth") || ensName
									? formatAddress(ensAddress ?? address)
									: "External account"}
					</p>
				</div>
				<div className="my-auto ml-auto w-max font-bold">
					<Badge />
				</div>
			</div>
		</Accordion>
	)
}

interface TransferRecipientFrameProps {
	index: number
	token?: TokenType
	collectible?: CollectibleType
	collection?: CollectionType
}

export const TransferRecipientFrame: FC<TransferRecipientFrameProps> = ({ index, token, collectible, collection }) => {
	const { account } = useConnect()
	const { column, isFrame, frame, transfer } = useColumns(
		index,
		token
			? `${token.symbol}-transfer-recipient`
			: `${collection?.address}-${collection?.chain}-${collectible?.tokenId}-transfer-recipient`
	)
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
				{column?.transfer?.recipient !== account.address && (
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
