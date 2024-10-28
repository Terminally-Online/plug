import { FC, HTMLAttributes, useMemo } from "react"

import { PlugZap, Puzzle } from "lucide-react"

import { Callout, Header, PlugGrid } from "@/components"
import { api } from "@/server/client"
import { COLUMN_KEYS, useColumns } from "@/state"

const Discover: FC<{ index: number }> = ({ index }) => {
	const { navigate } = useColumns(index)

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
					navigate({
						index,
						key: COLUMN_KEYS.DISCOVER,
						from: COLUMN_KEYS.HOME
					})
				}
				nextLabel="See All"
			/>

			<Callout.EmptyPlugs className="my-24" index={index} isEmpty={plugs !== undefined && plugs.length === 0} />
			<PlugGrid index={index} from={COLUMN_KEYS.HOME} plugs={visiblePlugs} />
		</div>
	)
}

const Mine: FC<{ index: number }> = ({ index }) => {
	const { navigate } = useColumns(index)

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
					navigate({
						index,
						key: COLUMN_KEYS.MY_PLUGS,
						from: COLUMN_KEYS.HOME
					})
				}
				nextLabel="See All"
			/>

			<Callout.EmptyPlugs className="my-24" index={index} isEmpty={plugs !== undefined && plugs.length === 0} />
			<PlugGrid index={index} from={COLUMN_KEYS.HOME} plugs={visiblePlugs} />
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
