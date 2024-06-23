import { useEffect, useMemo, useState } from "react"

import Image from "next/image"

import BlockiesSvg from "blockies-react-svg"
import {
	CheckCircle,
	ChevronRight,
	Eye,
	Globe,
	LoaderCircle,
	Users
} from "lucide-react"

import { Frame } from "@/components/app/frames/base"
import { ActionPreview } from "@/components/app/plugs/actions"
import { Button } from "@/components/buttons"
import { Checkbox } from "@/components/inputs"
import { useFrame, usePlugs, useSockets } from "@/contexts"
import { actionCategories } from "@/lib"
import { formatTitle } from "@/lib/functions"

export const ExecuteFrame = () => {
	const { socket, sockets, handleSelect } = useSockets()
	const { actions, plug } = usePlugs()
	const { frameVisible, handleFrameVisible } = useFrame()

	const [chainsSelected, setChainsSelected] = useState<string[]>([])

	const chainsAvailable = useMemo(() => {
		return Array.from(
			actions
				.map(
					action =>
						new Set(actionCategories[action.categoryName].chains)
				)
				// @ts-ignore -- Don't feel like properly typing this right now.
				.reduce((acc, curr) => {
					if (acc === null) return curr

					return new Set([...acc].filter(chain => curr.has(chain)))
				}, null)
		)
	}, [actions])

	const handleSocketSelect = (socketAddress: string) => {
		handleSelect(socketAddress)
		handleFrameVisible("chain")
	}

	const handleChainSelect = (chain: string) => {
		setChainsSelected(prev =>
			prev.includes(chain)
				? prev.filter(c => c !== chain)
				: [...prev, chain]
		)
	}

	useEffect(() => {
		if (frameVisible === "chain" && chainsAvailable.length === 1)
			setChainsSelected([chainsAvailable[0]])

		if (socket === undefined && sockets && sockets.length === 1) {
			handleSelect(sockets[0].socketAddress)
		}

		if (frameVisible === "socket") handleFrameVisible("chain")
	}, [
		chainsAvailable,
		socket,
		sockets,
		frameVisible,
		handleSelect,
		handleFrameVisible
	])

	useEffect(() => {
		if (frameVisible === "executing") {
			const timeout = setTimeout(
				() => handleFrameVisible("executed"),
				3000
			)

			return () => clearTimeout(timeout)
		}
	}, [frameVisible, handleFrameVisible])

	if (!plug) return null

	return (
		<>
			<Frame
				className="z-[2]"
				icon={<Users size={18} className="opacity-60" />}
				label="Choose Socket"
				visible={frameVisible === "socket"}
				handleVisibleToggle={() => handleFrameVisible(undefined)}
			>
				<div className="flex flex-col gap-4">
					{sockets && sockets.length > 0 ? (
						sockets.map((socketInMap, index) => (
							<div
								key={index}
								className="group flex cursor-pointer flex-row items-center gap-4"
								onClick={() =>
									handleSocketSelect(
										socketInMap.socketAddress
									)
								}
							>
								<BlockiesSvg
									address={socketInMap.socketAddress}
									className="h-6 w-6 rounded-md"
								/>
								<p className="mr-auto font-bold">
									{socketInMap.name}
								</p>

								<Button
									variant="secondary"
									className="ml-auto p-1 group-hover:bg-grayscale-100"
									onClick={() =>
										handleSocketSelect(
											socketInMap.socketAddress
										)
									}
								>
									<ChevronRight
										size={14}
										className="opacity-60"
									/>
								</Button>
							</div>
						))
					) : (
						<p>TODO</p>
					)}
				</div>
			</Frame>

			<Frame
				className="z-[2]"
				handleBack={
					sockets && sockets.length > 1
						? () => handleFrameVisible("socket")
						: undefined
				}
				icon={<Globe size={18} className="opacity-60" />}
				label={"Choose Chain" + (chainsAvailable.length > 1 ? "s" : "")}
				visible={frameVisible === "chain"}
				handleVisibleToggle={() => handleFrameVisible(undefined)}
				hasOverlay={true}
			>
				<div className="flex flex-col gap-4">
					{chainsAvailable.map((chain, index) => (
						<div
							key={`chain-${index}`}
							className="flex flex-row items-center gap-4"
						>
							<Checkbox
								checked={chainsSelected.includes(chain)}
								handleChange={() => handleChainSelect(chain)}
							/>

							<div className="mr-auto flex flex-row gap-2">
								<Image
									src={`/blockchain/${chain}.png`}
									alt={formatTitle(chain)}
									width={64}
									height={64}
									className="h-6 w-6"
								/>
								<p className="font-bold">
									{formatTitle(chain)}
								</p>
							</div>

							<p className="tabular-nums opacity-60">
								23.005 ETH
							</p>
						</div>
					))}

					<div className="mt-4 flex flex-row gap-4">
						<Button
							variant="secondary"
							className="w-max"
							onClick={() => handleFrameVisible("run")}
						>
							Execute
						</Button>

						<Button
							className="w-full"
							onClick={() => handleFrameVisible("schedule")}
						>
							Schedule
						</Button>
					</div>
				</div>
			</Frame>

			<Frame
				className="z-[2]"
				handleBack={() => handleFrameVisible("chain")}
				icon={<Eye size={18} className="opacity-60" />}
				label="Transaction Preview"
				visible={frameVisible === "run"}
				handleVisibleToggle={() => handleFrameVisible(undefined)}
				hasOverlay={true}
			>
				<div className="flex flex-col gap-4">
					<p className="font-bold opacity-60">Actions</p>
					<ActionPreview />

					{socket && (
						<p className="mt-4 flex font-bold">
							<span className="mr-auto opacity-60">
								Use Socket
							</span>
							<div className="flex flex-row items-center gap-2">
								<BlockiesSvg
									address={socket.socketAddress}
									className="h-5 w-5 rounded-md"
								/>
								<p className="mr-auto font-bold">
									{socket.name}
								</p>
							</div>
						</p>
					)}

					<p className="flex font-bold">
						<span className="mr-auto opacity-60">Run On</span>
						{chainsSelected.map(chain => (
							<Image
								key={chain}
								className="ml-[-10px] h-6 w-6"
								src={`/blockchain/${chain}.png`}
								alt={chain}
								width={24}
								height={24}
							/>
						))}
					</p>

					<p className="flex font-bold">
						<span className="mr-auto opacity-60">Fee</span>
						<span className="flex flex-row gap-2">
							<span className="opacity-40">0.0011 ETH</span>
							<span>$4.19</span>
						</span>
					</p>

					<Button
						className="mt-4 w-full"
						onClick={() => handleFrameVisible("executing")}
					>
						Submit Transaction
					</Button>
				</div>
			</Frame>

			<Frame
				className="z-[2]"
				handleBack={() => handleFrameVisible("chain")}
				icon={<Eye size={18} className="opacity-60" />}
				label="Schedule Transaction"
				visible={frameVisible === "schedule"}
				handleVisibleToggle={() => handleFrameVisible(undefined)}
				hasOverlay={true}
			>
				<div className="flex flex-col gap-4">
					{/* <p className="font-bold opacity-60">Actions</p> */}
					{/* <ActionPreview /> */}
					<p className="font-bold opacity-60">Actions</p>

					<div className="flex items-center gap-4">
						<Image
							src="/protocols/plug.png"
							alt="Plug"
							width={32}
							height={32}
							className="h-6 w-6"
						/>
						<p className="font-bold">
							Time stamp is
							<span
								style={{
									background: `linear-gradient(to right, rgba(0,239,54,0.1), rgba(147,223,0,0.1))`,
									color: `#00EF35`
								}}
								className="rounded px-2 py-1 font-bold"
							>
								greater than{" "}
							</span>
							<span
								style={{
									background: `linear-gradient(to right, rgba(0,239,54,0.1), rgba(147,223,0,0.1))`,
									color: `#00EF35`
								}}
								className="rounded px-2 py-1 font-bold"
							>
								06/16/2024
							</span>
						</p>
					</div>
					<div className="flex items-center gap-4">
						<Image
							src="/protocols/plug.png"
							alt="Plug"
							width={32}
							height={32}
							className="h-6 w-6"
						/>
						<p className="font-bold">
							Can only be called{" "}
							<span
								style={{
									background: `linear-gradient(to right, rgba(0,239,54,0.1), rgba(147,223,0,0.1))`,
									color: `#00EF35`
								}}
								className="rounded px-2 py-1 font-bold"
							>
								1
							</span>{" "}
							times every{" "}
							<span
								style={{
									background: `linear-gradient(to right, rgba(0,239,54,0.1), rgba(147,223,0,0.1))`,
									color: `#00EF35`
								}}
								className="rounded px-2 py-1 font-bold"
							>
								day
							</span>
						</p>
					</div>
					<div className="flex items-center gap-4">
						<Image
							src="/protocols/nouns.png"
							alt="Nouns"
							width={32}
							height={32}
							className="h-6 w-6"
						/>
						<p className="font-bold">
							Noun has{" "}
							<span
								style={{
									background: `linear-gradient(to right, rgba(0,239,54,0.1), rgba(147,223,0,0.1))`,
									color: `#00EF35`
								}}
								className="rounded px-2 py-1 font-bold"
							>
								glasses
							</span>{" "}
							of{" "}
							<span
								style={{
									background: `linear-gradient(to right, rgba(0,239,54,0.1), rgba(147,223,0,0.1))`,
									color: `#00EF35`
								}}
								className="rounded px-2 py-1 font-bold"
							>
								yellow
							</span>{" "}
						</p>
					</div>
					<div className="flex items-center gap-4">
						<Image
							src="/protocols/nouns.png"
							alt="Nouns"
							width={32}
							height={32}
							className="h-6 w-6"
						/>
						<p className="font-bold">
							Bid on Noun with{" "}
							<span
								style={{
									background: `linear-gradient(to right, rgba(0,239,54,0.1), rgba(147,223,0,0.1))`,
									color: `#00EF35`
								}}
								className="rounded px-2 py-1 font-bold"
							>
								8
							</span>{" "}
							$ETH
						</p>
					</div>

					{socket && (
						<p className="flex font-bold">
							<span className="mr-auto opacity-60">
								Use Socket
							</span>
							<div className="flex flex-row items-center gap-2">
								<BlockiesSvg
									address={socket.socketAddress}
									className="h-5 w-5 rounded-md"
								/>
								<p className="mr-auto font-bold">
									{socket.name}
								</p>
							</div>
						</p>
					)}

					<p className="flex font-bold">
						<span className="mr-auto opacity-60">Run On</span>
						{chainsSelected.map(chain => (
							<>
								<Image
									className="ml-[-10px] h-6 w-6"
									src={`/blockchain/${chain}.png`}
									alt={chain}
									width={24}
									height={24}
								/>
							</>
						))}
					</p>
					<p className="flex font-bold">
						<span className="mr-auto opacity-60">Total</span>
						<div className="flex flex-row gap-2">
							<span className="opacity-40">0.0011 ETH</span>
							<span>$4.19</span>
						</div>
					</p>

					<Button
						className="mt-4 w-full"
						onClick={() => handleFrameVisible("signing")}
					>
						Sign Intent
					</Button>
				</div>
			</Frame>

			<Frame
				className="z-[2]"
				icon={
					<LoaderCircle
						size={18}
						className="animate-spin opacity-60"
					/>
				}
				label="Execution Processing"
				visible={frameVisible === "executing"}
				handleVisibleToggle={() => handleFrameVisible(undefined)}
			>
				<div className="flex flex-col gap-8">
					<p className="leading-6">
						<span className="opacity-60">Your</span>
						<span
							className="rounded-lg bg-gradient-to-tr px-2 py-1 font-bold"
							style={{
								background: `linear-gradient(to right, rgba(0,239,54,0.1), rgba(147,223,0,0.1))`,
								color: `#00E100`
							}}
						>
							{plug.name}
						</span>{" "}
						<span className="opacity-60">
							execution is currently processing.
						</span>
					</p>
				</div>
			</Frame>

			<Frame
				className="z-[2]"
				icon={<CheckCircle size={18} className="opacity-60" />}
				label="Transaction Executed"
				visible={frameVisible === "executed"}
				handleVisibleToggle={() => handleFrameVisible(undefined)}
			>
				<div className="flex flex-col gap-2">
					<p>
						<span className="opacity-60">Your</span>
						<span
							className="rounded-lg bg-gradient-to-tr px-2 py-1 font-bold"
							style={{
								background: `linear-gradient(to right, rgba(0,239,54,0.1), rgba(147,223,0,0.1))`,
								color: `#00E100`
							}}
						>
							{plug.name}
						</span>{" "}
						<span className="opacity-60">
							execution was successfully executed. Your balances
							have updated and you can now resume editing your
							Plug.
						</span>
					</p>

					<p className="mt-4 flex font-bold">
						<span className="mr-auto opacity-60">Run On</span>

						<Image
							className="ml-[-10px] h-6 w-6"
							src={`/blockchain/ethereum.png`}
							alt={"Ethereum"}
							width={24}
							height={24}
						/>
					</p>
					<p className="flex font-bold">
						<span className="mr-auto opacity-60">Total</span>
						<div className="flex flex-row gap-2">
							<span className="opacity-40">0.00135 ETH</span>
							<span>$6.11</span>
						</div>
					</p>
					<p className="flex font-bold">
						<span className="mr-auto opacity-60">Status</span>
						<div className="flex flex-row gap-2">
							<p className="flex flex-row gap-2">
								<span>
									<div className="rounded-full bg-plug-green/10 p-1">
										<CheckCircle
											className="text-plug-green"
											size={16}
										/>
									</div>
								</span>
								Success
							</p>
						</div>
					</p>

					<div className="mt-4">
						<Button
							variant="secondary"
							className="w-full"
							onClick={() => handleFrameVisible("chain")}
						>
							View on Explorer
						</Button>
					</div>
				</div>
			</Frame>
		</>
	)
}
