import { FC, HTMLAttributes, PropsWithChildren, useMemo } from "react"

import { Activity, Cable, Coins, Globe, ImageIcon, PiggyBank, Plus, ShieldAlert, Star, User } from "lucide-react"

import { Header } from "@/components"
import { Accordion } from "@/components/shared"
import { cn, formatTitle } from "@/lib"
import { COLUMN_KEYS, Flag, useColumns, useFlags, useSocket } from "@/state"

type Options = Array<{
	label: keyof typeof COLUMN_KEYS
	description: string
	icon: JSX.Element
}>

export const ANONYMOUS_OPTIONS: Options = [
	{
		label: "DISCOVER",
		description: "Discover curated and community Plugs.",
		icon: <Globe size={14} className="opacity-40" />
	},
	{
		label: "MY_PLUGS",
		description: "Create, edit, and run your Plugs.",
		icon: <Cable size={14} className="opacity-40" />
	}
]

export const OPTIONS: Options = [
	...ANONYMOUS_OPTIONS,
	{
		label: "ACTIVITY",
		description: "View the simulations and runs of your Plugs.",
		icon: <Activity size={14} className="opacity-40" />
	},
	{
		label: "TOKENS",
		description: "View your tokens and manage them.",
		icon: <Coins size={14} className="opacity-40" />
	},
	{
		label: "COLLECTIBLES",
		description: "View your collectibles and manage them.",
		icon: <ImageIcon size={14} className="opacity-40" />
	},
	{
		label: "POSITIONS",
		description: "View your positions and manage them.",
		icon: <PiggyBank size={14} className="opacity-40" />
	}
] as const

export const ADMIN_OPTIONS: Options = [
	...OPTIONS,
	{
		label: "ADMIN",
		description: "View and manage the admin panel.",
		icon: <ShieldAlert size={14} className="opacity-40" />
	},
	{
		label: "PROFILE",
		description: "View your profile.",
		icon: <User size={14} className="opacity-40" />
	}
] as const

export const ColumnAdd = () => {
	const { getFlag } = useFlags()
	const { socket } = useSocket()
	const { columns, add } = useColumns()

	const flagOptions = useMemo(() => {
		const options: Options = []

		if (getFlag(Flag.SHOW_PWA))
			options.push({
				label: "APPLICATION",
				description: "Install Plug as an app on your device.",
				icon: <Star size={14} className="opacity-40" />
			})

		return options
	}, [getFlag])

	const isAdmin = socket?.admin ?? false
	const options = isAdmin ? ADMIN_OPTIONS : [...flagOptions, ...OPTIONS]

	return (
		<div
			className={cn(
				"relative my-2 mr-2 mr-48 flex h-full select-none flex-col rounded-lg border-[1px] border-grayscale-100 bg-white",
				columns.length === 2 && "ml-2"
			)}
			style={{ minWidth: `${480}px` }}
		>
			<div className="relative flex cursor-pointer flex-row items-center overflow-hidden overflow-y-auto rounded-t-lg border-b-[1px] border-grayscale-100 bg-white transition-all duration-200 ease-in-out">
				<div className="flex w-full flex-row items-center gap-4 px-6 py-4">
					<Plus size={18} className="opacity-40" />
					<p className="overflow-hidden truncate overflow-ellipsis text-lg font-bold">Add Column</p>
				</div>
			</div>

			<div className="h-full overflow-y-scroll">
				<div className="flex h-full flex-col gap-2 p-4">
					{options.map(option => (
						<Accordion
							key={option.label}
							onExpand={() => {
								add({ key: option.label })
							}}
						>
							<div className="flex flex-row items-center gap-2">
								<div className="flex h-10 w-10 min-w-10 items-center justify-center">{option.icon}</div>

								<div className="flex flex-col items-start text-left font-bold">
									<p>{formatTitle(option.label.replace("_", " ").toLowerCase())}</p>
									<p className="text-sm opacity-40">{option.description}</p>
								</div>
							</div>
						</Accordion>
					))}
				</div>
			</div>
		</div>
	)
}
