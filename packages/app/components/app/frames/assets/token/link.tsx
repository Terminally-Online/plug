import { ArrowRight, ChartColumnStacked, ChartPie, Coins, Globe, MessageCircleIcon, Send, Twitter } from "lucide-react"

const TokenFrameExternalLinkIcon = ({ name }: { name: string }) => {
	switch (name.toLowerCase()) {
		case "website":
			return <Globe size={18} className="opacity-20" />
		case "twitter":
			return <Twitter size={18} className="opacity-20" />
		case "telegram":
			return <Send size={18} className="opacity-20" />
		case "discord":
			return <MessageCircleIcon size={18} className="opacity-20" />
		case "coingecko":
		case "coinmarketcap":
			return <Coins size={18} className="opacity-20" />
		case "dex.guru":
			return <ChartColumnStacked size={18} className="opacity-20" />
		default:
			return <ChartPie size={18} className="opacity-20" />
	}
}

export const TokenFrameExternalLink = ({ link }: { link: { name: string; url: string } }) => {
	return (
		<p className="flex w-full flex-row items-center gap-4 font-bold">
			<TokenFrameExternalLinkIcon name={link.name} />
			<a
				href={link.url}
				target="_blank"
				rel="noopener noreferrer"
				className="group flex w-full flex-row items-center justify-between"
			>
				<span className="mr-auto truncate whitespace-nowrap opacity-40 transition-opacity duration-200 ease-in-out group-hover:opacity-100">
					{link.name}
				</span>
				<span className="flex max-w-[120px] flex-row items-center truncate overflow-ellipsis whitespace-nowrap">
					<ArrowRight
						size={18}
						className="-rotate-45 opacity-20 transition-opacity duration-200 ease-in-out group-hover:opacity-100"
					/>
				</span>
			</a>
		</p>
	)
}
