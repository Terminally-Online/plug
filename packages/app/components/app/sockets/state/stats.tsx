import { Counter } from "@/components/shared/utils/counter"
import { api } from "@/server/client"
import { useSocket } from "@/state/authentication"
import { ArrowDownToLine, ArrowUpFromLine, LockIcon, Wallet, WavesLadder } from "lucide-react"
import { FC, HTMLAttributes, memo, useMemo } from "react"
import { ChainImage } from "../chains/chain.image"
import { formatTitle } from "@/lib"

export const SocketStateStats: FC<
	HTMLAttributes<HTMLDivElement> & {
		index: number
		address?: string
		isExpanded?: boolean
		count?: number
		isColumn?: boolean
	}
> = memo(({ index, address, isExpanded = false, count = 5, isColumn = true, className }) => {
	const { isAnonymous, socket } = useSocket()
	const { data } = api.service.zerion.wallet.portfolio.useQuery({
		path: { address: address || socket?.socketAddress },
		query: {
			filter: {
				positions: "no_filter"
			}
		}
	}, { enabled: !isAnonymous, placeholderData: prev => prev })
	const portfolio = useMemo(() => data?.data, [data])

	const distribution = useMemo(() => {
		if (!portfolio?.attributes.positions_distribution_by_chain) {
			return [];
		}

		const chains = Object.entries(portfolio.attributes.positions_distribution_by_chain);
		const total = chains.reduce((sum, [_, impl]) => sum + (impl || 0), 0);

		return chains.map(([chainId, impl]) => {
			const value = impl || 0;
			const percentage = total > 0 ? (value / total) * 100 : 0;

			return {
				chainId,
				value,
				percentage
			};
		}).sort((a, b) => b.percentage - a.percentage);
	}, [portfolio]);

	return <>
		<div className="flex flex-col px-6 pt-2">
			{portfolio?.attributes.positions_distribution_by_type.wallet !== undefined && <p className="flex w-full flex-row items-center gap-4 font-bold">
				<Wallet size={18} className="opacity-20" />
				<span className="mr-auto truncate whitespace-nowrap opacity-40">Idle</span>
				<span className="flex flex-row font-bold opacity-60">
					$<Counter count={((portfolio?.attributes.positions_distribution_by_type.wallet ?? 0) - (portfolio?.attributes.positions_distribution_by_type.deposited ?? 0)).toLocaleString("en-US", {
						minimumFractionDigits: 2,
						maximumFractionDigits: 2
					})} />
				</span>
				<span className="flex min-w-[72px] flex-row items-center text-right font-bold">
					<Counter count={
						portfolio?.attributes.positions_distribution_by_type.wallet > 0
							? (((portfolio?.attributes.positions_distribution_by_type.wallet ?? 0) - (portfolio?.attributes.positions_distribution_by_type.deposited ?? 0)) / portfolio?.attributes.positions_distribution_by_type.wallet) * 100
							: 0
					} decimals={2} />%
				</span>
			</p>}

			<p className="flex w-full flex-row items-center gap-4 font-bold">
				<ArrowDownToLine size={18} className="opacity-20" />
				<span className="mr-auto truncate whitespace-nowrap opacity-40">Deposited</span>
				<span className="flex flex-row font-bold opacity-60">
					$
					<Counter count={(portfolio?.attributes.positions_distribution_by_type.deposited ?? 0).toLocaleString("en-US", {
						minimumFractionDigits: 2,
						maximumFractionDigits: 2
					})} />
				</span>
			</p>

			<p className="flex w-full flex-row items-center gap-4 font-bold">
				<ArrowUpFromLine size={18} className="opacity-20" />
				<span className="mr-auto truncate whitespace-nowrap opacity-40">Borrowed</span>
				<div className="flex flex-col font-bold opacity-60">
					<p className="flex flex-row">
						$
						<Counter count={(portfolio?.attributes.positions_distribution_by_type.borrowed ?? 0).toLocaleString("en-US", {
							minimumFractionDigits: 2,
							maximumFractionDigits: 2
						})} />
					</p>
				</div>
			</p>

			<p className="flex w-full flex-row items-center gap-4 font-bold">
				<LockIcon size={18} className="opacity-20" />
				<span className="mr-auto truncate whitespace-nowrap opacity-40">Locked</span>
				<div className="flex flex-col font-bold opacity-60">
					<p className="flex flex-row">
						$
						<Counter count={(portfolio?.attributes.positions_distribution_by_type.locked ?? 0).toLocaleString("en-US", {
							minimumFractionDigits: 2,
							maximumFractionDigits: 2
						})} />
					</p>
				</div>
			</p>

			<p className="flex w-full flex-row items-center gap-4 font-bold">
				<WavesLadder size={18} className="opacity-20" />
				<span className="mr-auto truncate whitespace-nowrap opacity-40">Staked</span>
				<div className="flex flex-col font-bold opacity-60">
					<p className="flex flex-row">
						$
						<Counter count={(portfolio?.attributes.positions_distribution_by_type.staked ?? 0).toLocaleString("en-US", {
							minimumFractionDigits: 2,
							maximumFractionDigits: 2
						})} />
					</p>
				</div>
			</p>
		</div>

		<div className="h-[1px] w-full bg-plug-green/10 my-2" />

		<div className="px-6 flex flex-col">
			{distribution.map((item, index) => (
				<div key={index} className="flex flex-row items-center gap-4">
					<ChainImage chainId={item.chainId} size="sm" />
					<p className="mr-auto font-bold whitespace-nowrap truncate">{formatTitle(item.chainId ?? "Unknown")}</p>
					<p className="flex flex-row font-bold opacity-60">
						<Counter
							count={item.value < 0.01 ? "<$0.01" : `$${(item.value).toLocaleString("en-US", {
								minimumFractionDigits: 2,
								maximumFractionDigits: 2
							})}`}
							decimals={2}
						/>
					</p>
					<p className="flex min-w-[72px] flex-row items-center text-right font-bold">
						<Counter count={item.percentage < 0.01 ? "<0.01" : `${item.percentage.toLocaleString("en-US", {
							minimumFractionDigits: 2,
							maximumFractionDigits: 2
						})}`} decimals={2} />%
					</p>
				</div>
			))}
		</div>
	</>
})

SocketStateStats.displayName = "SocketStateStats"

