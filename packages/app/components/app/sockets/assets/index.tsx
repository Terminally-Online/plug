import { signOut } from "next-auth/react"
import { FC, HTMLAttributes, useState } from "react"

import { useDisconnect } from "wagmi"

import { Check, CircleDollarSign, ImageIcon, LogOut, Share } from "lucide-react"

import {
	Button,
	Callout,
	Container,
	Header,
	SocketCollectionList,
	SocketPositionList,
	SocketTokenList
} from "@/components"
import { cn } from "@/lib"
import { useHoldings, useSocket } from "@/state"
import { COLUMNS, useColumnStore } from "@/state"

export const SocketAssets: FC<
	HTMLAttributes<HTMLDivElement> & {
		index?: number
		address?: string
		hasTokens?: boolean
		hasPositions?: boolean
		hasCollectibles?: boolean
	}
> = ({
	index = -1,
	address,
	hasTokens = false,
	hasPositions = false,
	hasCollectibles = false,
	className,
	...props
}) => {
	const { isAnonymous, socket } = useSocket()
	const { collectibles, tokens, protocols } = useHoldings(address)
	const { handle } = useColumnStore(COLUMNS.MOBILE_INDEX)
	const [copied, setCopied] = useState(false)
	const { disconnect } = useDisconnect({
		mutation: {
			onSuccess: () => signOut({ callbackUrl: "/" })
		}
	})

	return (
		<>
			<Container className="border-grayscale-100 fixed left-0 right-0 top-0 z-[10] border-b-[1px] bg-white md:hidden">
				<div className="flex flex-row items-center gap-4 py-4">
					<span className="font-bold">Account</span>

					<div className="ml-auto flex flex-row gap-2">
						<Button
							variant="secondary"
							className="group rounded-sm p-1"
							onClick={async () => {
								try {
									const shareUrl = `${window.location.origin}/app?rfid=${socket?.identity?.referralCode}`
									await navigator.clipboard.writeText(shareUrl)
									setCopied(true)
									setTimeout(() => setCopied(false), 2000)
								} catch (err) {
									console.error("Failed to copy link:", err)
								}
							}}
						>
							{copied ? (
								<Check size={14} className="opacity-60 transition-all" />
							) : (
								<Share size={14} className="opacity-60 transition-opacity group-hover:opacity-100" />
							)}
						</Button>

						<Button variant="secondary" className="group rounded-sm p-1" onClick={() => disconnect()}>
							<LogOut size={14} className="opacity-60 transition-opacity group-hover:opacity-100" />
						</Button>
					</div>
				</div>
			</Container>

			<div className={cn("flex flex-col gap-2", "mt-16 md:mt-0", "pb-24 md:pb-0", className)} {...props}>
				<Callout.Anonymous index={index} viewing="assets" />
				<Callout.EmptyAssets
					index={index}
					isEmpty={[collectibles, tokens, protocols].every(basket => basket.length === 0)}
				/>
				{isAnonymous === false && (
					<>
						{hasTokens && tokens && tokens.length > 0 && (
							<SocketTokenList className="h-max" index={index} expanded={true} isColumn={false} />
						)}

						{hasPositions && protocols && protocols.length > 0 && (
							<>
								<Header
									size="sm"
									icon={<CircleDollarSign size={14} className="opacity-40" />}
									label="Positions"
								/>
								<SocketPositionList index={index} expanded={true} isColumn={false} />
							</>
						)}

						{hasCollectibles && collectibles && collectibles.length > 0 && (
							<>
								<Header
									size="sm"
									icon={<ImageIcon size={14} className="opacity-40" />}
									label="Collectibles"
								/>
								<SocketCollectionList index={index} expanded={true} isColumn={false} />
							</>
						)}
					</>
				)}
			</div>
		</>
	)
}