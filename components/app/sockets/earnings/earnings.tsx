import { FC, HTMLAttributes } from "react"

import { Trophy } from "lucide-react"

import { Button, Callout, Counter, Header, PlugGrid, SocketEarningsChart, StatCard } from "@/components"
import { useSockets } from "@/contexts"
import { cn } from "@/lib"
import { api } from "@/server/client"

export const SocketEarnings: FC<HTMLAttributes<HTMLDivElement> & { id: string }> = ({ id, className, ...props }) => {
	const { isAnonymous: anonymous } = useSockets()

	const { data: plugs } = api.plug.all.useQuery({
		target: "mine",
		limit: 8
	})

	return (
		<div className={cn("flex h-full flex-col gap-2", className)} {...props}>
			<Callout.Anonymous viewing="earnings" />

			{anonymous === false && (
				<>
					<SocketEarningsChart />

					<div className="mt-4 flex flex-col gap-2">
						<div className="flex flex-row gap-2">
							<StatCard>
								<Counter className="mr-auto w-max text-2xl font-bold" count={201} />
								<p className="font-bold opacity-40">Onboarded</p>
							</StatCard>
							<StatCard>
								<Counter className="mr-auto w-max text-2xl font-bold" count={9351} />
								<p className="font-bold opacity-40">Attributed Runs</p>
							</StatCard>
						</div>
						<div className="flex flex-row gap-2">
							<StatCard>
								<Counter className="mr-auto w-max text-2xl font-bold" count={201} />
								<p className="font-bold opacity-40">All Time Fees</p>
							</StatCard>
							<StatCard>
								<Counter className="mr-auto w-max text-2xl font-bold" count={9351} />
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

							<PlugGrid id={id} className="mb-4" from={"mine"} plugs={plugs} />
						</>
					)}
				</>
			)}
		</div>
	)
}
