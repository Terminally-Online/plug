import { useSession } from "next-auth/react"
import { FC, useCallback } from "react"

import { Connector as wagmiConnector } from "wagmi"

import { motion } from "framer-motion"
import { Loader2 } from "lucide-react"

import { useAtomValue, useSetAtom } from "jotai"

import { Accordion, Animate, Button, Callout, Image } from "@/components"
import {
	cn,
	CONNECTOR_ICON_OVERRIDE_MAP,
	formatAddress,
	greenGradientStyle,
	recentConnectorIdAtom,
	useConnect,
	useMediaQuery,
	useOrderedConnections,
	useRecentConnectorId
} from "@/lib"
import { authenticationAtom, useColumnData, walletConnectURIMatrixAtom } from "@/state"

const QR_CODE_SIZE = 200
const QR_CODE_PIXEL_SPACING = 0.3

const ConnectorQrCode = () => {
	const { connection } = useConnect()
	const { md } = useMediaQuery()
	const qrMatrix = useAtomValue(walletConnectURIMatrixAtom)

	const isCorner = useCallback(
		(row: number, col: number) => {
			if (qrMatrix === undefined) return false

			return (
				(row < 7 && col < 7) ||
				(row < 7 && col > qrMatrix.moduleCount - 1 - 7) ||
				(row > qrMatrix.moduleCount - 1 - 7 && col < 7)
			)
		},
		[qrMatrix]
	)

	const moduleSize = qrMatrix ? QR_CODE_SIZE / qrMatrix.moduleCount : 0
	const actualSize = moduleSize - moduleSize * QR_CODE_PIXEL_SPACING
	const offset = (moduleSize * QR_CODE_PIXEL_SPACING) / 2

	return (
		<div className="my-2 flex w-full flex-col items-center justify-center py-8">
			{!md && (
				<>
					<h1 className="mb-8 text-2xl font-bold">Welcome to Plug</h1>
					<p className="mb-8 text-center font-bold text-black/40">Connect your wallet to get started</p>
				</>
			)}
			<div className="relative w-full max-w-[300px]">
				{qrMatrix ? (
					<motion.div
						className="relative w-full max-w-[300px]"
						initial={{ opacity: 0 }}
						animate={{ opacity: 1 }}
						transition={{ duration: 0.2 }}
					>
						<svg
							width="100%"
							height="100%"
							viewBox={`0 0 ${QR_CODE_SIZE} ${QR_CODE_SIZE}`}
							preserveAspectRatio="xMidYMid meet"
						>
							<rect width={QR_CODE_SIZE} height={QR_CODE_SIZE} fill={"#FFFFFF"} />
							{Array.from({ length: qrMatrix.moduleCount }, (_, row) =>
								Array.from(
									{ length: qrMatrix.moduleCount },
									(_, col) =>
										qrMatrix.getModule(row, col) && (
											<rect
												key={`${row}-${col}`}
												x={col * moduleSize + (isCorner(row, col) ? 0 : offset)}
												y={row * moduleSize + (isCorner(row, col) ? 0 : offset)}
												width={isCorner(row, col) ? moduleSize : actualSize}
												height={isCorner(row, col) ? moduleSize : actualSize}
												rx={isCorner(row, col) ? "2px" : "1px"}
												ry={isCorner(row, col) ? "2px" : "1px"}
												fill={"#000000"}
											/>
										)
								)
							)}
						</svg>

						<div className="absolute left-1/2 top-1/2 -translate-x-1/2 -translate-y-1/2 overflow-hidden rounded-lg border-[6px] border-white bg-white p-2">
							<Image
								className="absolute h-10 w-10 rounded-md blur-lg filter"
								src={"/wallets/walletconnect-icon.svg"}
								alt={"wallet connect"}
								width={48}
								height={48}
							/>
							<Image
								className="relative h-10 w-10 rounded-md"
								src={"/wallets/walletconnect-icon.svg"}
								alt={"wallet connect"}
								width={48}
								height={48}
							/>
						</div>
					</motion.div>
				) : (
					<div className="flex h-[300px] w-[300px] items-center justify-center">
						<Loader2 size={24} className="animate-spin opacity-60" />
					</div>
				)}
			</div>

			<p
				className={cn(
					"max-w-[320px] pt-4 text-sm font-bold",
					connection.isError ? "text-red-500" : "opacity-40"
				)}
			>
				{connection.isError
					? "Error connecting. Please click try and again and follow the steps to connect in your wallet."
					: connection.isLoading
						? "Open the wallet you selected to confirm the connection with Plug."
						: "Scan the QR code to connect your wallet from your camera or the in-wallet scanner."}
			</p>
		</div>
	)
}

const ConnectorImage: FC<{ icon: string | undefined; name: string }> = ({ icon, name }) => {
	const dimensions = {
		blur: 4,
		content: 2.5
	}

	if (!icon) return null

	return (
		<div
			className="relative h-10"
			style={{
				width: `${dimensions.content}rem`,
				height: `${dimensions.content}rem`
			}}
		>
			<Image
				className="absolute left-1/2 top-1/2 h-12 w-12 -translate-x-1/2 -translate-y-1/2 rounded-md blur-xl filter"
				src={icon}
				alt={name}
				style={{
					height: `${dimensions.blur}rem`,
					width: `${dimensions.blur}rem`
				}}
				width={48}
				height={48}
			/>
			<Image
				className="absolute left-1/2 top-1/2 -translate-x-1/2 -translate-y-1/2 rounded-md"
				src={icon}
				alt={name}
				style={{
					width: `${dimensions.content}rem`,
					height: `${dimensions.content}rem`
				}}
				width={48}
				height={48}
			/>
		</div>
	)
}

const Connector: FC<{ connector: wagmiConnector; index: number; from?: string }> = ({ connector, index, from }) => {
	const { connection, prove } = useConnect()

	const updateRecentConnectorId = useSetAtom(recentConnectorIdAtom)

	const isLoading = connection.isLoading && connection.variables?.connector === connector
	const isRecent = connector.id === useRecentConnectorId()
	const isDetected = connector.isInjected as boolean
	const isDisabled = Boolean(connection?.isLoading)
	const icon = CONNECTOR_ICON_OVERRIDE_MAP[connector.id] ?? connector.icon

	const Badge = () => {
		if (isLoading) return <Loader2 className="animate-spin opacity-40" size={14} />
		if (isRecent) return <span style={{ ...greenGradientStyle }}>Recent</span>
		if (isDetected) return <span className="opacity-40">Detected</span>
		return null
	}

	return (
		<Accordion
			className={cn(isDisabled && "cursor-not-allowed bg-plug-green/5")}
			onExpand={
				isDisabled
					? undefined
					: () =>
							connection.connect(
								{ connector },
								{
									onSuccess: data => {
										updateRecentConnectorId(connector.id)
										prove(index, from, data.accounts[0])
									}
								}
							)
			}
		>
			<div className="flex flex-row items-center gap-4">
				<ConnectorImage icon={icon} name={connector.name} />
				<p className="font-bold">{connector.name}</p>
				<p className="ml-auto text-sm font-bold">
					<Badge />
				</p>
			</div>
		</Accordion>
	)
}

const Connectors: FC<{ index: number; from?: string }> = ({ index, from }) => {
	const connectors = useOrderedConnections(true)

	return (
		<div className="mb-auto w-full pt-2">
			<ConnectorQrCode />
			<div className="h-[1px] w-full bg-plug-green/10" />
			<div className="px-4 pt-4">
				<Animate.List>
					{connectors.map(connector => (
						<Animate.ListItem key={connector.id}>
							<Connector connector={connector} index={index} from={from} />
						</Animate.ListItem>
					))}
				</Animate.List>
			</div>
		</div>
	)
}

export const ColumnAuthenticate: FC<{ index: number }> = ({ index }) => {
	const { data: session } = useSession()
	const { account, sign, prove } = useConnect()
	const { column } = useColumnData(index)

	const authentication = useAtomValue(authenticationAtom)

	return (
		<div className="flex h-full flex-col items-center justify-center text-center">
			{session?.user?.id === account.address && (
				<Callout
					title="You are authenticated."
					description="You should not be seeing this message. Please refresh the page."
				/>
			)}

			{authentication.isLoading && (
				<Callout
					title="Authentication loading."
					description="We are loading all the state of your account. One moment please."
				/>
			)}

			{session?.user?.id !== account.address &&
				account.address &&
				sign.isLoading === false &&
				authentication.isLoading === false && (
					<Callout
						title={sign.failureReason ? "Signature error." : "Prove ownership."}
						description={
							sign.failureReason
								? "An internal error was received while signing the message. " +
									sign.failureReason.message.split("Details:")[1].split("Details:")[0].trim()
								: `Please sign the message to prove your ownership of ${formatAddress(account.address)}.`
						}
					>
						<Button className="mt-2" sizing="sm" onClick={() => prove(index, column?.from)}>
							Sign Message
						</Button>
					</Callout>
				)}

			{account.address && sign.isLoading && (
				<Callout
					title="Proving ownership."
					description={`Completing the signing process to prove ownership of ${formatAddress(account.address)}`}
				>
					<Button className="mt-2" sizing="sm" disabled>
						Signing...
					</Button>
				</Callout>
			)}

			{account.address === undefined && <Connectors index={index} from={column?.from} />}
		</div>
	)
}
