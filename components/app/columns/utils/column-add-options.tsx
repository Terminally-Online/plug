import { FC, HTMLAttributes, PropsWithChildren, useMemo } from "react"

import { Activity, Cable, Coins, ImageIcon, Landmark, PiggyBank, ShieldAlert, Star, User, Wallet } from "lucide-react"

import { Accordion } from "@/components/shared"
import { cn, formatTitle, VIEW_KEYS } from "@/lib"
import { Flag, useColumns, useFlags, useSocket } from "@/state"

type Options = Array<{
	label: keyof typeof VIEW_KEYS
	description: string
	icon: JSX.Element
}>

const ANONYMOUS_OPTIONS: Options = [
	{
		label: "DISCOVER",
		description: "Discover curated and community Plugs.",
		icon: <Cable size={14} className="opacity-40" />
	},
	{
		label: "MY_PLUGS",
		description: "Create, edit, and run your Plugs.",
		icon: <Cable size={14} className="opacity-40" />
	}
]

const OPTIONS: Options = [
	{
		label: "APPLICATION",
		description: "Save Plug as an app to your device.",
		icon: <Star size={14} className="opacity-40" />
	},
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

const ADMIN_OPTIONS: Options = [
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

export const ColumnAddOptions: FC<
	HTMLAttributes<HTMLDivElement> &
		PropsWithChildren<{
			index: number
		}>
> = ({ index, className, ...props }) => {
	const { flags } = useFlags()
	const { socket } = useSocket()
	const { navigate } = useColumns()

	const flagOptions = useMemo(() => {
		const options = []

		if (flags[Flag.SHOW_PWA])
			options.push({
				label: "Application",
				description: "Install Plug as an app on your device.",
				icon: <Star className="h-4 w-4" />
			})

		return options
	}, [flags])

	const isAdmin = socket?.admin ?? false
	const options = isAdmin ? ADMIN_OPTIONS : [...ANONYMOUS_OPTIONS, ...flagOptions, ...OPTIONS]

	return (
		<div className={cn("flex h-full flex-col gap-2", className)} {...props}>
			{options.map(option => (
				<Accordion key={option.label} onExpand={() => navigate({ key: option.label, index })}>
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
	)
}
