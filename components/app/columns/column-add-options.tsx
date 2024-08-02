import { FC } from "react"

import {
	Activity,
	Cable,
	Coins,
	ImageIcon,
	Landmark,
	LayoutPanelTop,
	PiggyBank,
	Settings,
	Wallet
} from "lucide-react"

import { Header } from "@/components"
import { useSockets } from "@/contexts"
import { formatTitle } from "@/lib"
import { COLUMN_KEYS } from "@/server/api/routers/socket/columns"

type Props = {
	id: string
}

const options: Array<{
	label: keyof typeof COLUMN_KEYS
	description: string
	icon: JSX.Element
}> = [
	{
		label: "DISCOVER",
		description: "Discover curated and community Plugs.",
		icon: <Cable size={14} className="opacity-40" />
	},
	{
		label: "MY_PLUGS",
		description: "Create, edit, and run your Plugs.",
		icon: <Cable size={14} className="opacity-40" />
	},
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
	},
	{
		label: "SETTINGS",
		description: "View your settings and manage them.",
		icon: <Settings size={14} className="opacity-40" />
	}
] as const

export const ConsoleColumnAddOptions: FC<Props> = ({ id }) => {
	const { handle } = useSockets()

	return (
		<>
			<Header
				size="md"
				className="px-4"
				icon={<LayoutPanelTop size={14} className="opacity-40" />}
				label="Choose One"
			/>

			<div className="flex flex-col border-t-[1px] border-grayscale-100">
				{options.map(option => (
					<button
						key={option.label}
						className="cursor-pointer border-b-[1px] border-grayscale-100 px-4 py-2 text-left transition-all duration-200 ease-in-out hover:bg-grayscale-0"
						onClick={() =>
							handle.columns.add({ key: option.label, id })
						}
					>
						<div className="flex flex-row items-center gap-4">
							{option.icon}

							<div className="flex flex-col">
								<p className="font-bold opacity-40">
									{formatTitle(
										option.label
											.replace("_", " ")
											.toLowerCase()
									)}
								</p>
								<p className="text-sm opacity-60">
									{option.description}
								</p>
							</div>
						</div>
					</button>
				))}
			</div>
		</>
	)
}
