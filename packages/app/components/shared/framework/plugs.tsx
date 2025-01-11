import { FC, HTMLAttributes, useMemo } from "react"

import { PlugZap, Puzzle } from "lucide-react"

import { Header } from "@/components/app/layout/header"
import { PlugGrid } from "@/components/app/plugs/grid/grid"
import { Callout } from "@/components/app/utils/callout"
import { api } from "@/server/client"
import { COLUMNS, useColumnStore } from "@/state/columns"

const Discover: FC<{ index: number }> = ({ index }) => {
	const { handle } = useColumnStore(index)

	const { data: plugs, isLoading } = api.plugs.all.useQuery({
		target: "others",
		limit: 4
	})

	const visiblePlugs = useMemo(() => {
		if (isLoading || plugs === undefined || plugs.length === 0) return Array(6).fill(undefined)
		return plugs
	}, [isLoading, plugs])

	return (
		<div className="relative">
			<Header
				size="md"
				icon={<Puzzle size={14} className="opacity-40" />}
				label="Discover"
				nextOnClick={() =>
					handle.navigate({
						index,
						key: COLUMNS.KEYS.DISCOVER,
						from: COLUMNS.KEYS.HOME
					})
				}
				nextLabel="See All"
			/>

			<Callout.EmptyPlugs className="my-24" index={index} isEmpty={plugs !== undefined && plugs.length === 0} />
			<PlugGrid index={index} from={COLUMNS.KEYS.HOME} plugs={visiblePlugs} />
		</div>
	)
}

const Mine: FC<{ index: number }> = ({ index }) => {
	const { handle } = useColumnStore(index)

	const { data: plugs, isLoading } = api.plugs.all.useQuery({
		target: "mine",
		limit: 12
	})

	const visiblePlugs = useMemo(() => {
		if (isLoading || plugs === undefined || plugs.length === 0) return Array(6).fill(undefined)
		return plugs
	}, [isLoading, plugs])

	return (
		<div className="relative">
			<Header
				size="md"
				icon={<PlugZap size={14} className="opacity-40" />}
				label="My Plugs"
				nextOnClick={() =>
					handle.navigate({
						index,
						key: COLUMNS.KEYS.MY_PLUGS,
						from: COLUMNS.KEYS.HOME
					})
				}
				nextLabel="See All"
			/>

			<Callout.EmptyPlugs className="my-24" index={index} isEmpty={plugs !== undefined && plugs.length === 0} />
			<PlugGrid index={index} from={COLUMNS.KEYS.HOME} plugs={visiblePlugs} />
		</div>
	)
}

export const Plugs: FC<HTMLAttributes<HTMLDivElement> & { index?: number; hideEmpty?: boolean }> = ({
	index = -1,
	hideEmpty = false,
	...props
}) => (
	<div {...props}>
		<Discover index={index} />
		<Mine index={index} />
	</div>
)
