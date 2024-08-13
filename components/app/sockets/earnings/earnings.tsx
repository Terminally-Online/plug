import { FC, HTMLAttributes } from "react"

import { SocketEarningsChart } from "."
import { Trophy } from "lucide-react"

import { Button, Counter } from "@/components/shared"
import { api } from "@/server/client"

import { StatCard } from "../../cards"
import { Header } from "../../layout"
import { PlugGrid } from "../../plugs"

export const SocketEarnings: FC<
	HTMLAttributes<HTMLDivElement> & { id: string }
> = ({ id, ...props }) => {
	const { data: plugs } = api.plug.all.useQuery({
		target: "mine",
		limit: 8
	})

	return (
		<div {...props}>
			{/* <div className="flex flex-col font-bold">
				<div className="flex flex-row items-center gap-2">
					<div className="h-3 w-3 rounded-full bg-gradient-to-tr from-plug-green to-plug-yellow" />
					<p className="opacity-40">Forks</p>
				</div>
				<div className="flex flex-row items-center gap-2">
					<div className="to-sun-yellow from-sun-orange h-3 w-3 rounded-full bg-gradient-to-tr" />
					<p className="opacity-40">Runs</p>
				</div>
			</div> */}

			<SocketEarningsChart />

			<div className="mt-4 flex flex-col gap-2">
				<div className="flex flex-row gap-2">
					<StatCard>
						<Counter
							className="mr-auto w-max text-2xl font-bold"
							count={201}
						/>
						<p className="font-bold opacity-40">Onboarded</p>
					</StatCard>
					<StatCard>
						<Counter
							className="mr-auto w-max text-2xl font-bold"
							count={9351}
						/>
						<p className="font-bold opacity-40">Attributed Runs</p>
					</StatCard>
				</div>
				<div className="flex flex-row gap-2">
					<StatCard>
						<Counter
							className="mr-auto w-max text-2xl font-bold"
							count={201}
						/>
						<p className="font-bold opacity-40">All Time Fees</p>
					</StatCard>
					<StatCard>
						<Counter
							className="mr-auto w-max text-2xl font-bold"
							count={9351}
						/>
						<p className="font-bold opacity-40">Unclaimed Fees</p>
					</StatCard>
				</div>
				<Button className="w-full" onClick={() => {}}>
					Claim Fees
				</Button>
			</div>

			{plugs && (
				<>
					<Header
						size="md"
						icon={<Trophy size={14} className="opacity-40" />}
						label="Top Performers"
					/>

					<PlugGrid
						id={id}
						className="mb-4"
						from={"mine"}
						plugs={plugs}
					/>
				</>
			)}
		</div>
	)
}
