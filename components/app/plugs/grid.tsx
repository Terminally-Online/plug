import type { FC } from "react"

import { PlugGridItem } from "@/components/app/plugs/grid-item"
import { Button } from "@/components/buttons"
import { usePlugs } from "@/contexts/PlugProvider"
import { routes } from "@/lib/constants"

type Props = { from: string; count?: number; all?: boolean }

export const PlugGrid: FC<Props> = ({ from, count, all = false }) => {
	const { plugs, search, handle } = usePlugs()

	if (plugs === undefined) return null

	return (
		<>
			{plugs && plugs.length > 0 ? (
				<div className="grid grid-cols-2 gap-1 lg:grid-cols-4">
					{plugs.slice(0, count || plugs.length).map(plug => (
						<PlugGridItem key={plug.id} from={from} plug={plug} />
					))}
				</div>
			) : search !== "" ? (
				<div className="my-64 flex flex-col gap-[30px]">
					<p className="mx-auto max-w-[65%] text-center text-lg opacity-60">
						No Plugs could be found that match that search.
					</p>
					<div className="mx-auto flex flex-row gap-1">
						<Button
							className="w-max"
							onClick={() => handle.search("")}
						>
							Clear Search
						</Button>
					</div>
				</div>
			) : (
				<div className="my-64 flex flex-col gap-[30px]">
					<p className="mx-auto max-w-[65%] text-center text-lg opacity-60">
						Create your first Plug from scratch by clicking the
						button below and assembling your strategy.
					</p>
					<div className="mx-auto flex flex-row gap-1">
						<Button
							variant="secondary"
							href={routes.app.plugs.create}
							className="w-max"
						>
							See Templates
						</Button>
						<Button
							href={`${routes.app.plugs.create}?from=${routes.app.plugs.mine}`}
							className="w-max"
						>
							Create
						</Button>
					</div>
				</div>
			)}
		</>
	)
}
