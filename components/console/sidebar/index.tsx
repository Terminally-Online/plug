import Image from "next/image"

import BlockiesSvg from "blockies-react-svg"
import { AnimatePresence, motion } from "framer-motion"
import { Book, ClipboardCheck, Github, LogOut, Plus } from "lucide-react"

import { Button } from "@/components/shared"
import { useSockets } from "@/contexts"
import { useClipboard } from "@/lib"

export const ConsoleSidebar = () => {
	const { address, ensAvatar } = useSockets()
	const { copied, handleCopied } = useClipboard(address ?? "")

	return (
		<div className="flex h-screen min-w-20 flex-col items-center border-r-[1px] border-grayscale-100 bg-white py-4">
			<div className="mb-auto flex flex-col items-center justify-center gap-4">
				{address && (
					<button
						className="relative h-12 w-12 rounded-sm bg-grayscale-0"
						onClick={() => handleCopied()}
					>
						<motion.div
							initial={{ opacity: 0 }}
							animate={{ opacity: 1 }}
							exit={{ opacity: 0 }}
							transition={{ duration: 0.2 }}
						>
							{ensAvatar ? (
								<Image
									src={ensAvatar}
									alt="ENS Avatar"
									width={72}
									height={72}
									className="h-full w-full rounded-sm"
								/>
							) : (
								<BlockiesSvg
									className="h-full w-full rounded-sm"
									address={address}
								/>
							)}
						</motion.div>

						<AnimatePresence>
							{copied && (
								<motion.div
									className="absolute -bottom-2 -right-2 rounded-full border-[1px] border-grayscale-0 bg-white p-1"
									initial={{ opacity: 0 }}
									animate={{ opacity: 1 }}
									exit={{ opacity: 0 }}
									transition={{ duration: 0.2 }}
								>
									<ClipboardCheck
										size={14}
										className="opacity-40"
									/>
								</motion.div>
							)}
						</AnimatePresence>
					</button>
				)}

				<div className="flex w-full flex-col gap-1 text-center">
					<p className="text-xs opacity-60">Lvl. 12</p>
					<div className="h-[2px] w-full bg-grayscale-100">
						<div className="h-[2px] w-[50%] bg-gradient-to-r from-plug-green to-plug-yellow" />
					</div>
				</div>

				<Button
					variant="primary"
					href="https://docs.onplug.io"
					sizing="sm"
					className="rounded-sm p-1 outline-none"
				>
					<Plus size={18} className="opacity-60" />
				</Button>
			</div>

			<div className="mt-auto flex flex-col items-center justify-center gap-4">
				<Button
					variant="secondary"
					href="https://docs.onplug.io"
					sizing="sm"
					className="rounded-sm p-1 outline-none"
				>
					<Book size={18} className="opacity-60" />
				</Button>

				<Button
					variant="secondary"
					href="https://github.com/nftchance/plug"
					sizing="sm"
					className="rounded-sm p-1 outline-none"
				>
					<Github size={18} className="opacity-60" />
				</Button>

				<Button
					variant="secondary"
					onClick={() => {}}
					sizing="sm"
					className="rounded-sm p-1 outline-none"
				>
					<LogOut size={18} className="rotate-180 opacity-60" />
				</Button>
			</div>
		</div>
	)
}
