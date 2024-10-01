import Image from "next/image"
import { FC } from "react"

import { Badge, Link, Send, Twitter } from "lucide-react"

import { Button, Frame } from "@/components"
import { usePlugs } from "@/contexts"
import { routes, useClipboard } from "@/lib"
import { useColumns } from "@/state"

export const ShareFrame: FC<{ index: number; item: string }> = ({ index, item }) => {
	const { isFrame } = useColumns(index, "share")
	const { plug } = usePlugs(item)

	const { copied, handleCopied } = useClipboard(`${window?.location.origin}/app/${plug ? `?id=${plug.id}` : ""}`)

	if (!plug) return null

	return (
		<Frame index={index} className="z-[2]" icon={<Badge size={18} />} label="Share Plug" visible={isFrame}>
			<div className="flex flex-col gap-2">
				<div className="flex flex-row items-center gap-2">
					<Link size={14} />
					<p className="font-bold">Direct Link</p>
					<Button variant="secondary" sizing="sm" className="ml-auto" onClick={handleCopied}>
						{copied ? "Copied" : "Copy"}
					</Button>
				</div>

				<div className="flex flex-row items-center gap-2">
					<Twitter size={14} />
					<p className="font-bold">Twitter</p>
					<a
						className="ml-auto"
						href={`https://twitter.com/intent/tweet?text=${plug.name} using @onplug_io:%0A%0A${window.location.origin}${routes.app}/?id=${plug.id}`}
						target="_blank"
						rel="noopener noreferrer"
					>
						<button className="rounded-full bg-gradient-to-tr from-[#0085CE] to-[#00A2FB] px-[24px] py-[8px] text-xs font-bold text-white hover:from-[#0085CE]/90 hover:to-[#00A2FB]/90">
							Tweet
						</button>
					</a>
				</div>

				<div className="flex flex-row items-center gap-2">
					<Send size={14} />
					<p className="font-bold">Telegram</p>
					<a
						className="ml-auto"
						href={`https://t.me/share/url?url=${window.location.origin}${routes.app}/?id=${plug.id}&text=${plug.name} using @onplug_io`}
						target="_blank"
						rel="noopener noreferrer"
					>
						<button className="rounded-full bg-gradient-to-tr from-[#00A2E3] to-[#67D4FF] px-[24px] py-[8px] text-xs font-bold text-white hover:from-[#00A2E3]/90 hover:to-[#67D4FF]/90">
							Share
						</button>
					</a>
				</div>

				<div className="flex flex-row items-center gap-2">
					<Image src="/icons/farcaster.svg" alt="Farcaster" width={14} height={14} />
					<p className="font-bold">Warpcast</p>
					<a
						className="ml-auto"
						href={`https://warpcast.com/~/compose?text=https://twitter.com/intent/tweet?text=${plug.name}%20using%20@onplug_io&embeds[]=${window.location.origin}${routes.app}/?id=${plug.id}`}
						target="_blank"
						rel="noopener noreferrer"
					>
						<button className="rounded-full bg-[#472A91] px-[24px] py-[8px] text-xs font-bold text-white hover:bg-[#472A91]/90">
							Cast
						</button>
					</a>
				</div>
			</div>
		</Frame>
	)
}
