import { FC } from "react"

import {
	Activity,
	Cable,
	CircleFadingPlus,
	Coins,
	ImageIcon,
	Landmark,
	PiggyBank,
	Settings,
	ShieldAlert,
	User,
	Wallet
} from "lucide-react"

import { useSockets } from "@/contexts"
import { formatTitle, VIEW_KEYS } from "@/lib"

type Props = {
	id: string
}

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
	...ANONYMOUS_OPTIONS,
	{
		label: "ACTIVITY",
		description: "View the simulations and runs of your Plugs.",
		icon: <Activity size={14} className="opacity-40" />
	},
	{
		label: "ASSETS",
		description: "View your tokens, collectibles and positions.",
		icon: <Wallet size={14} className="opacity-40" />
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
	},
	{
		label: "EARNINGS",
		description: "View your earnings and manage them.",
		icon: <Landmark size={14} className="opacity-40" />
	}
] as const

const ADMIN_OPTIONS: Options = [
	...ANONYMOUS_OPTIONS,
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

export const ColumnAddOptions: FC<Props> = ({ id }) => {
	const { anonymous, socket, handle } = useSockets()

	const isAdmin = socket?.admin ?? false

	const options = isAdmin ? ADMIN_OPTIONS : anonymous ? ANONYMOUS_OPTIONS : OPTIONS

	return (
		<div className="flex h-full flex-col">
			{options.map(option => (
				<button
					key={option.label}
					className="cursor-pointer border-b-[1px] border-grayscale-100 px-4 py-2 text-left transition-all duration-200 ease-in-out hover:bg-grayscale-0"
					onClick={() => handle.columns.add({ key: option.label, id })}
				>
					<div className="flex flex-row items-center gap-4">
						{option.icon}

						<div className="flex flex-col">
							<p className="font-bold">{formatTitle(option.label.replace("_", " ").toLowerCase())}</p>
							<p className="text-sm font-bold opacity-40">{option.description}</p>
						</div>
					</div>
				</button>
			))}

			{anonymous && (
				<p className="max-w-[380px] p-4 text-sm font-bold opacity-40">
					Several options are unavailble because you are using an anonymous account.
				</p>
			)}
		</div>
	)
}
