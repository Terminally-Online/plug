import { getCsrfToken, signIn } from "next-auth/react"
import Image from "next/image"
import { FC, useCallback, useEffect, useMemo, useState } from "react"

import { EthereumProvider, EthereumProviderOptions } from "@walletconnect/ethereum-provider"
import { SiweMessage } from "siwe"
import { useChainId, useDisconnect, Connector as wagmiConnector } from "wagmi"

import { motion } from "framer-motion"
import { Loader2 } from "lucide-react"

import qrcode from "qrcode-generator"

import { Accordion, Animate, Button, Callout } from "@/components"
import { wagmiChains, WALLETCONNECT_PARAMS } from "@/contexts"
import {
	CONNECTOR_ICON_OVERRIDE_MAP,
	formatAddress,
	greenGradientStyle,
	useConnect,
	useOrderedConnections,
	useRecentConnectorId
} from "@/lib"

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

const ConnectorQrCode: FC<{ uri: string | undefined }> = ({ uri }) => {
	const size = 200

	const qrMatrix = useMemo(() => {
		if (uri === undefined) return

		const qr = qrcode(0, "L")
		qr.addData(uri)
		qr.make()
		return {
			moduleCount: qr.getModuleCount(),
			getModule: (row: number, col: number) => qr.isDark(row, col)
		}
	}, [uri])

	const isCorner = (row: number, col: number) => {
		if (qrMatrix === undefined) return false
		const lastIndex = qrMatrix.moduleCount - 1
		return (row < 7 && col < 7) || (row < 7 && col > lastIndex - 7) || (row > lastIndex - 7 && col < 7)
	}

	const pixelSpacing = 0.4
	const moduleSize = qrMatrix ? size / qrMatrix.moduleCount : 0
	const actualSize = moduleSize - moduleSize * pixelSpacing
	const offset = (moduleSize * pixelSpacing) / 2

	return (
		<div className="my-2 flex w-full flex-col items-center justify-center py-8">
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
							viewBox={`0 0 ${size} ${size}`}
							preserveAspectRatio="xMidYMid meet"
						>
							<rect width={size} height={size} fill={"#FFFFFF"} />
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
												rx={isCorner(row, col) ? 0 : "1px"}
												ry={isCorner(row, col) ? 0 : "1px"}
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

			<p className="pt-4 text-sm font-bold opacity-40">Scan to connect your wallet.</p>
		</div>
	)
}

const Connector: FC<{ connector: wagmiConnector }> = ({ connector }) => {
	const { connection } = useConnect()

	const isLoading = connection.isLoading && connection.variables?.connector === connector

	const isDetected = connector.isInjected as boolean
	const isRecent = connector.id === useRecentConnectorId()
	const icon = CONNECTOR_ICON_OVERRIDE_MAP[connector.id] ?? connector.icon
	// TODO(WEB-4173): Remove isIFrame check when we can update wagmi to version >= 2.9.4
	const isDisabled = Boolean(connection?.isLoading)
	// const isDisabled = Boolean(connection?.isLoading && !isIFramed())

	const Badge = () => {
		if (isLoading) return <Loader2 className="animate-spin opacity-60" size={14} />
		if (isRecent) return <span style={{ ...greenGradientStyle }}>Recent</span>
		if (isDetected) return <span className="opacity-40">Detected</span>
		return null
	}

	return (
		<Accordion onExpand={() => connection.connect({ connector })}>
			<div className="flex flex-row items-center gap-4">
				<ConnectorImage icon={icon} name={connector.name} />
				<p className="font-bold">{connector.name}</p>
				{isDetected && (
					<p className="ml-auto text-sm font-bold">
						<Badge />
					</p>
				)}
			</div>
		</Accordion>
	)
}

const Connectors = () => {
	const connectors = useOrderedConnections(true)

	const [initialized, setInitialized] = useState(false)
	const [provider, setProvider] = useState<Awaited<ReturnType<typeof EthereumProvider.init>>>()
	const [uri, setUri] = useState<string>()

	const init = async () => {
		const provider = await EthereumProvider.init({
			...WALLETCONNECT_PARAMS,
			showQrModal: false,
			optionalChains: wagmiChains.map(chain => chain.id)
		} as EthereumProviderOptions)

		setProvider(provider)
		setInitialized(true)
	}

	useEffect(() => {
		if (!initialized) init()
	}, [initialized])

	useEffect(() => {
		if (!provider) return

		const generateUri = async () => {
			if (!provider) return
			await provider.connect()
		}

		provider.on("display_uri", (uri: string) => setUri(uri))

		generateUri()
	}, [provider])

	return (
		<div className="mb-auto w-full pt-2">
			<ConnectorQrCode uri={uri} />
			<div className="h-[1px] w-full bg-grayscale-100" />
			<div className="px-4 pt-4">
				<Animate.List>
					{connectors.map(connector => (
						<Animate.ListItem key={connector.id}>
							<Connector connector={connector} />
						</Animate.ListItem>
					))}
				</Animate.List>
			</div>
		</div>
	)
}

export const ColumnAuthenticate = () => {
	const chainId = useChainId()
	const { account, sign } = useConnect()
	const { disconnect } = useDisconnect()

	const handleProof = useCallback(async () => {
		if (!account.address) throw new Error("Message cannot be signed without a connected wallet.")

		try {
			const message = new SiweMessage({
				domain: window.location.host,
				address: account.address,
				statement: `Access the Plug platform by proving your ownership of the address: ${account.address}.`,
				uri: window.location.origin,
				version: "1",
				chainId: chainId,
				nonce: await getCsrfToken()
			}).prepareMessage()

			sign.signMessage(
				{
					message
				},
				{
					onSuccess: signature =>
						signIn("credentials", {
							message: JSON.stringify(message),
							redirect: true,
							signature,
							callbackUrl: "/app/"
						}),
					onError: () => disconnect()
				}
			)
		} catch (e) {
			sign.reset()
			disconnect()
		}
	}, [chainId, account, sign, disconnect])

	// useEffect(() => {
	// 	if (isConnected === false || isLoading || isError) return

	// 	handleLogin()
	// }, [isConnected, isLoading, isError, handleLogin])

	const Column = useCallback(() => {
		if (account.address) {
			if (sign.isLoading)
				return (
					<Callout
						title="Proving ownership."
						description={`Completing the signing process to prove ownership of ${formatAddress(account.address)}`}
					>
						<Button className="mt-2" sizing="sm" disabled>
							Signing...
						</Button>
					</Callout>
				)

			const title = sign.failureReason ? "Signature error." : "Prove ownership."
			const description = sign.failureReason
				? "An internal error was received while signing the message. " +
					sign.failureReason.message.split("Details:")[1].split("Details:")[0].trim()
				: `Please sign the message to prove your ownership of ${formatAddress(account.address)}.`

			return (
				<Callout title={title} description={description}>
					<Button className="mt-2" sizing="sm" onClick={handleProof}>
						Sign Message
					</Button>
				</Callout>
			)
		}

		return <Connectors />
	}, [account.address, sign.failureReason, sign.isLoading, handleProof])

	return (
		<div className="flex h-full flex-col items-center justify-center text-center">
			<Column />
		</div>
	)
}
