import { useCallback } from "react"

import { useConnect} from "wagmi"

import { motion } from "framer-motion"
import { Loader2 } from "lucide-react"

import { useAtomValue } from "jotai"

import { Image } from "@/components/app/utils/image"
import {
	cn,
} from "@/lib"
import { walletConnectURIMatrixAtom } from "@/state/authentication"

const QR_CODE_SIZE = 200
const QR_CODE_PIXEL_SPACING = 0.3

export const ConnectorQrCode = () => {
	const connection = useConnect()

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
					: connection.isPending
						? "Open the wallet you selected to confirm the connection with Plug."
						: "Scan the QR code to connect your wallet from your camera or the in-wallet scanner."}
			</p>
		</div>
	)
}

