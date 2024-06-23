import { FC } from "react"

import Image from "next/image"

import { AlertCircle, CheckCircle, ChevronRight, XCircle } from "lucide-react"

import { Button } from "@/components/buttons"
import { useFrame } from "@/contexts"
import { colors } from "@/lib"
import { formatTitle } from "@/lib/functions"

import { Frame } from "../frames/base"

const getStatusIcon = (status: string) => {
	switch (status) {
		case "success":
			return (
				<div className="rounded-full bg-plug-green/10 p-1">
					<CheckCircle className="text-plug-green" size={16} />
				</div>
			)
		case "error":
			return (
				<div className="rounded-full bg-red-400/10 p-1">
					<XCircle className="text-red-400" size={16} />
				</div>
			)
		case "warning":
			return (
				<div className="rounded-full bg-yellow-600/10 p-1">
					<AlertCircle className="text-yellow-600" size={16} />
				</div>
			)
		default:
			return <></>
	}
}

export const ActivityItem: FC<{
	text: string
	color: keyof typeof colors
	status: string
	time: string
}> = ({ text, status, time }) => {
	const icon = getStatusIcon(status)
	const { handleFrameVisible, frameVisible } = useFrame()

	const handleClick = () => {
		handleFrameVisible(`activityDetails-${time}`)
	}

	return (
		<div className="border-b border-grayscale-100 py-2">
			<button
				className="group flex w-full flex-row items-center gap-4"
				onClick={handleClick}
			>
				{icon}
				<span className="flex-1 truncate text-left font-bold">
					{text}
				</span>
				<span className="opacity-60">{time}</span>
				<Button
					variant="secondary"
					className="p-1 group-hover:bg-grayscale-100"
					onClick={handleClick}
				>
					<ChevronRight className="opacity-60" size={14} />
				</Button>
			</button>

			<Frame
				className="scrollbar-hide z-[1] flex flex-col gap-4 overflow-y-auto"
				icon={
					<div
						className="h-6 w-6 rounded-md"
						style={{ backgroundColor: colors["blue"] }}
					/>
				}
				label={text}
				visible={frameVisible === `activityDetails-${time}`}
				handleVisibleToggle={() => handleFrameVisible(undefined)}
			>
				<p className="font-bold opacity-60">Actions</p>

				<div className="flex items-center gap-4">
					<Image
						src="/protocols/aave.png"
						alt="Aave"
						width={32}
						height={32}
						className="h-6 w-6"
					/>
					<p className="font-bold">
						<span
							style={{
								background: `linear-gradient(to right, rgba(0,239,54,0.1), rgba(147,223,0,0.1))`,
								color: `#00EF35`
							}}
							className="rounded px-2 py-1 font-bold"
						>
							lend{" "}
						</span>
						rate for{" "}
						<span
							style={{
								background: `linear-gradient(to right, rgba(0,239,54,0.1), rgba(147,223,0,0.1))`,
								color: `#00EF35`
							}}
							className="rounded px-2 py-1 font-bold"
						>
							ETH
						</span>{" "}
						is{" "}
						<span
							style={{
								background: `linear-gradient(to right, rgba(0,239,54,0.1), rgba(147,223,0,0.1))`,
								color: `#00EF35`
							}}
							className="rounded px-2 py-1 font-bold"
						>
							greater than
						</span>{" "}
						<span
							style={{
								background: `linear-gradient(to right, rgba(0,239,54,0.1), rgba(147,223,0,0.1))`,
								color: `#00EF35`
							}}
							className="rounded px-2 py-1 font-bold"
						>
							12%
						</span>
					</p>
				</div>
				<div className="flex items-center gap-4">
					<Image
						src="/protocols/uniswap.png"
						alt="Uniswap"
						width={32}
						height={32}
						className="h-6 w-6"
					/>
					<p className="font-bold">
						Swap{" "}
						<span
							style={{
								background: `linear-gradient(to right, rgba(0,239,54,0.1), rgba(147,223,0,0.1))`,
								color: `#00EF35`
							}}
							className="rounded px-2 py-1 font-bold"
						>
							10,000
						</span>{" "}
						<span
							style={{
								background: `linear-gradient(to right, rgba(0,239,54,0.1), rgba(147,223,0,0.1))`,
								color: `#00EF35`
							}}
							className="rounded px-2 py-1 font-bold"
						>
							USDC
						</span>{" "}
						for{" "}
						<span
							style={{
								background: `linear-gradient(to right, rgba(0,239,54,0.1), rgba(147,223,0,0.1))`,
								color: `#00EF35`
							}}
							className="rounded px-2 py-1 font-bold"
						>
							2.5
						</span>{" "}
						<span
							style={{
								background: `linear-gradient(to right, rgba(0,239,54,0.1), rgba(147,223,0,0.1))`,
								color: `#00EF35`
							}}
							className="rounded px-2 py-1 font-bold"
						>
							ETH
						</span>
					</p>
				</div>
				<div className="flex items-center gap-4">
					<Image
						src="/protocols/aave.png"
						alt="Aave"
						width={32}
						height={32}
						className="h-6 w-6"
					/>
					<p className="font-bold">
						Increase collateral in{" "}
						<span
							style={{
								background: `linear-gradient(to right, rgba(0,239,54,0.1), rgba(147,223,0,0.1))`,
								color: `#00EF35`
							}}
							className="rounded px-2 py-1 font-bold"
						>
							ETH
						</span>{" "}
						with{" "}
						<span
							style={{
								background: `linear-gradient(to right, rgba(0,239,54,0.1), rgba(147,223,0,0.1))`,
								color: `#00EF35`
							}}
							className="rounded px-2 py-1 font-bold"
						>
							2.5
						</span>{" "}
						<span
							style={{
								background: `linear-gradient(to right, rgba(0,239,54,0.1), rgba(147,223,0,0.1))`,
								color: `#00EF35`
							}}
							className="rounded px-2 py-1 font-bold"
						>
							ETH
						</span>
					</p>
				</div>
				<div className="flex items-center gap-4">
					<Image
						src="/protocols/uniswap.png"
						alt="Uniswap"
						width={32}
						height={32}
						className="h-6 w-6"
					/>
					<p className="font-bold">
						Swap{" "}
						<span
							style={{
								background: `linear-gradient(to right, rgba(0,239,54,0.1), rgba(147,223,0,0.1))`,
								color: `#00EF35`
							}}
							className="rounded px-2 py-1 font-bold"
						>
							max
						</span>{" "}
						<span
							style={{
								background: `linear-gradient(to right, rgba(0,239,54,0.1), rgba(147,223,0,0.1))`,
								color: `#00EF35`
							}}
							className="rounded px-2 py-1 font-bold"
						>
							ETH
						</span>{" "}
						for{" "}
						<span
							style={{
								background: `linear-gradient(to right, rgba(0,239,54,0.1), rgba(147,223,0,0.1))`,
								color: `#00EF35`
							}}
							className="rounded px-2 py-1 font-bold"
						>
							USDC
						</span>
					</p>
				</div>

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
							<span>{getStatusIcon(status)}</span>
							{formatTitle(status)}
						</p>
					</div>
				</p>

				{status === "success" && (
					<div className="mt-4 flex flex-row gap-2">
						<Button
							variant="secondary"
							className="w-max"
							onClick={() => {}}
						>
							Explorer
						</Button>
						<Button className="w-full" onClick={() => {}}>
							Share
						</Button>
					</div>
				)}
			</Frame>
		</div>
	)
}
