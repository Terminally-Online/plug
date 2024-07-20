import { useMemo, useState } from "react"

import Image from "next/image"

import { motion } from "framer-motion"
import {
	ArrowLeftRight,
	ArrowRight,
	ChevronDown,
	ChevronRight,
	Globe,
	ReceiptText,
	SearchIcon,
	SquareArrowDownRight,
	User,
	Wallet
} from "lucide-react"
import { createPublicClient, http, isAddress, zeroAddress } from "viem"
import { useAccount, useEnsAddress, useEnsAvatar } from "wagmi"

import { Button, Footer, Frame, Search, SocketTokenList } from "@/components"
import { useBalances, useFrame, useSockets } from "@/contexts"
import { getChainImage } from "@/lib"

import { mainnet } from "viem/chains"

// TODO: Implement the functionality to actually withdraw and deposit funds from and to the Socket.
// NOTE: This will be pending until a Socket is deployed onto testnet to enable the ability of writing Socket and Wallet interactions at the same time.

const DEFAULT_TRANSFER = {
	token: undefined,
	chain: undefined,
	amount: undefined,
	to: undefined
}

export const TransferFrame = () => {
	const { frameVisible, handleFrameVisible } = useFrame()

	// TODO: Probably use the wallet client unless it would require them to change their network
	const client = createPublicClient({
		chain: mainnet,
		transport: http()
	})

	// const [transferVisible, setTransferVisible] = useState(false)
	const [transfer, setTransfer] = useState<{
		action?: "receive" | "send"
		token?: NonNullable<typeof balances>[number]
		chain?: NonNullable<typeof balances>[number]["chains"][0]
		amount?: string | undefined
		to?: string
	}>(DEFAULT_TRANSFER)
	const [advanced, setAdvanced] = useState(false)

	const { address } = useAccount()
	const { socket } = useSockets()

	const fromAddress =
		(transfer.action !== undefined
			? // Depositing tokens from wallet into Socket.
				transfer.action === "receive"
				? address
				: // Withdrawing tokens from Socket into wallet.
					socket?.socketAddress
			: "") || ""

	const { search, balances, handleSearch } = useBalances({
		address: fromAddress
	})

	// NOTE: This is not used currently, but it will be for the transaction functionality so leaving it here until then.
	const toAddress =
		(transfer.to !== undefined
			? transfer.to
			: fromAddress === socket?.socketAddress
				? address
				: socket?.socketAddress) || ""

	const { data: ensAddress } = useEnsAddress({
		name: transfer.to ?? zeroAddress
	})
	const { data: ensAvatar } = useEnsAvatar({
		name: transfer.to ?? zeroAddress
	})

	const transferValid = useMemo(() => {
		if (transfer.amount === undefined || transfer.amount === "0")
			return [false, "Enter Amount"]
		if (isNaN(Number(transfer.amount))) return [false, "Invalid Amount"]
		if (transfer.to !== undefined && transfer.to !== "") {
			if (transfer.to.startsWith("0x")) {
				if (!isAddress(transfer.to)) return [false, "Invalid Address"]
			} else {
				if (!ensAddress === undefined || ensAddress === null)
					return [false, "Invalid ENS"]
			}
		}
		if (
			Number(transfer.amount ?? 0) >
			Number(transfer.token?.balanceFormatted ?? 0)
		)
			return [false, "Insufficient Balance"]

		return [true, transfer.action === "send" ? "Withdraw" : "Deposit"]
	}, [ensAddress, transfer])

	// useEffect(() => {
	// 	if (transferVisible === false) setTransfer(DEFAULT_TRANSFER)
	// }, [transferVisible])

	return (
		<>
			<Footer>
				<Button
					className="w-full"
					onClick={() => handleFrameVisible("transfer")}
				>
					Transfer
				</Button>
			</Footer>

			<Frame
				className="z-[2]"
				icon={<ArrowLeftRight size={18} />}
				label="Choose Transfer Direction"
				visible={
					frameVisible === "transfer" && transfer.token === undefined
				}
			>
				<div className="flex flex-col gap-4">
					<button
						className="group flex flex-col gap-2 text-left"
						onClick={() =>
							setTransfer({ ...transfer, action: "receive" })
						}
					>
						<div className="flex w-full flex-row gap-2">
							<div className="flex w-full flex-row items-center gap-2">
								<SquareArrowDownRight size={14} />
								<p className="font-bold">To Socket</p>
							</div>

							<Button
								variant="secondary"
								className="p-1 group-hover:bg-grayscale-100"
								onClick={() =>
									setTransfer({
										...transfer,
										action: "receive"
									})
								}
							>
								<ArrowRight size={14} className="ml-auto " />
							</Button>
						</div>
						<p className="max-w-[85%] opacity-60">
							Deposit assets into your Socket from your connected
							wallet.
						</p>
					</button>

					<div className="h-0 border-[1px] border-b-grayscale-100" />

					<button
						className="group flex flex-col gap-2 text-left"
						onClick={() =>
							setTransfer({ ...transfer, action: "send" })
						}
					>
						<div className="flex w-full flex-row gap-2">
							<div className="flex w-full flex-row items-center gap-2">
								<SquareArrowDownRight
									size={14}
									className="rotate-[270deg]"
								/>
								<p className="font-bold">From Socket</p>
							</div>

							<Button
								variant="secondary"
								className="p-1 group-hover:bg-grayscale-100"
								onClick={() =>
									setTransfer({ ...transfer, action: "send" })
								}
							>
								<ArrowRight size={14} className="ml-auto" />
							</Button>
						</div>
						<p className="max-w-[85%] opacity-60">
							Withdraw assets from your Socket to your connected
							wallet.
						</p>
					</button>
				</div>
			</Frame>

			<Frame
				className="scrollbar-hide z-[3] max-h-[calc(100vh-80px)] overflow-y-auto"
				icon={
					transfer.action === "send" ? (
						<SquareArrowDownRight
							size={14}
							className="rotate-[270deg]"
						/>
					) : (
						<SquareArrowDownRight size={14} />
					)
				}
				label={`${transfer.action === "send" ? "Withdraw" : "Deposit"} Token`}
				visible={
					transfer.action !== undefined &&
					transfer.token === undefined
				}
				handleBack={() =>
					setTransfer({ ...transfer, action: undefined })
				}
			>
				<div className="flex h-full min-h-[280px] flex-col gap-4">
					<Search
						icon={<SearchIcon size={14} />}
						placeholder="Search activity and assets"
						search={search}
						handleSearch={(search: string) => handleSearch(search)}
					/>

					<SocketTokenList
						// balances={balances}
						handleSelect={(
							token: NonNullable<typeof balances>[number]
						) => setTransfer({ ...transfer, token })}
					/>
				</div>
			</Frame>

			<Frame
				className="scrollbar-hide z-[4] max-h-[calc(100vh-80px)] overflow-y-auto"
				icon={<Globe size={14} />}
				label={`${transfer.action === "send" ? "Withdraw" : "Deposit"} On`}
				visible={transfer.token !== undefined}
				handleBack={() =>
					setTransfer({ ...transfer, token: undefined })
				}
				hasOverlay={true}
			>
				{transfer.token && (
					<div className="flex flex-col gap-4">
						{transfer.token.chains.map((chain, index) => {
							return (
								<button
									key={index}
									className="group flex w-full cursor-pointer flex-row gap-2"
									onClick={() =>
										setTransfer({ ...transfer, chain })
									}
								>
									<div className="flex w-full flex-row items-center gap-2">
										<Image
											src={getChainImage(chain.chainId)}
											alt="Ethereum"
											className="h-6 w-6 rounded-full"
											width={48}
											height={48}
										/>
										<p className="mr-auto font-bold">
											{chain.chainName}
										</p>

										<p className="ml-auto flex flex-row items-center gap-2 tabular-nums">
											<span className="opacity-60">
												{chain.balanceFormatted.toLocaleString()}
											</span>
											<Image
												className="h-4 w-4 rounded-full"
												src={
													transfer.token?.logoURI ??
													""
												}
												alt={
													transfer.token?.symbol ?? ""
												}
												width={48}
												height={48}
											/>
											<span className="opacity-60">
												{transfer.token?.symbol ?? ""}
											</span>
										</p>
									</div>

									<Button
										variant="secondary"
										className="p-1 group-hover:bg-grayscale-100"
										onClick={() =>
											setTransfer({ ...transfer, chain })
										}
									>
										<ChevronRight size={14} />
									</Button>
								</button>
							)
						})}
					</div>
				)}
			</Frame>

			<Frame
				className="scrollbar-hide z-[5] max-h-[calc(100vh-80px)] overflow-y-auto"
				icon={<ReceiptText size={18} />}
				label={`${transfer.action === "send" ? "Withdraw" : "Deposit"} Details`}
				visible={transfer.chain !== undefined}
				handleBack={() =>
					setTransfer({
						...transfer,
						chain: undefined,
						amount: undefined,
						to: undefined
					})
				}
			>
				<div className="flex flex-col gap-4">
					<div className="flex flex-col gap-2">
						<Search
							icon={<Wallet size={14} />}
							placeholder="Amount"
							search={transfer.amount?.toString() ?? ""}
							handleSearch={(amount: string) =>
								setTransfer({ ...transfer, amount })
							}
						>
							<span className="flex w-max flex-row items-center gap-2">
								<Image
									className="h-4 w-4 rounded-full"
									src={transfer.token?.logoURI ?? ""}
									alt={transfer.token?.symbol ?? ""}
									width={48}
									height={48}
								/>
								<span className="opacity-40">
									{transfer.token?.symbol ?? ""} |{" "}
									<button
										onClick={() =>
											setTransfer({
												...transfer,
												amount: transfer.token?.balanceFormatted.toString()
											})
										}
									>
										MAX
									</button>
								</span>
							</span>
						</Search>

						<button
							className="group mt-4 flex w-full cursor-pointer flex-row gap-2 font-bold opacity-40"
							onClick={() => setAdvanced(!advanced)}
						>
							<span className="mr-auto">Advanced</span>
							<Button
								variant="secondary"
								className="p-1 group-hover:bg-grayscale-100"
								onClick={() => setAdvanced(!advanced)}
							>
								<motion.div
									animate={{ rotate: advanced ? 180 : 0 }}
									transition={{ duration: 0.2 }}
								>
									<ChevronDown size={14} />
								</motion.div>
							</Button>
						</button>

						{advanced && transfer.action === "send" && (
							<Search
								icon={
									ensAvatar ? (
										// NOTE: Ignoring the eslint rule here because the ENS avatar may resolve to where they set it and is not
										//       something that we control soooooo... yeah.
										// eslint-disable-next-line @next/next/no-img-element
										<img
											className="user-select-none rounded-full"
											src={ensAvatar}
											alt={`Avatar for ${transfer.to}`}
											width={24}
											height={24}
										/>
									) : (
										<User size={14} />
									)
								}
								placeholder="Recipient"
								search={transfer.to ?? ""}
								handleSearch={(address: string) =>
									setTransfer({
										...transfer,
										to: address
									})
								}
							/>
						)}
					</div>

					<Button
						variant={
							Number(transfer.amount ?? 0) >
								Number(transfer.token?.balanceFormatted ?? 0) ||
							transferValid[0] === false
								? "disabled"
								: "primary"
						}
						className="w-full"
						onClick={() => {
							// TODO: Chain interaction functionality here when Sockets are deployed on testnet.
						}}
					>
						{transferValid[1]}
					</Button>
				</div>
			</Frame>
		</>
	)
}
