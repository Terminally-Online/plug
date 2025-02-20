import { useEffect, useMemo, useRef, useState } from "react"

import { Activity, Cable, Cog, Coins, Globe, ImageIcon, PiggyBank, Plug, Plus, Star, LockIcon } from "lucide-react"

import { Accordion } from "@/components/shared/utils/accordion"
import { cn, formatTitle } from "@/lib"
import { useSocket } from "@/state/authentication"
import { COLUMNS, useColumnStore } from "@/state/columns"
import { Flag, useFlags } from "@/state/flags"
import { usePlugStore } from "@/state/plugs"

type Options = Array<{
	label: keyof (typeof COLUMNS)["KEYS"]
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

export const ColumnAdd = ({ index }: { index: number }) => {
	const resizeRef = useRef<HTMLDivElement>(null)

	const { getFlag } = useFlags()
	const { columns, handle } = useColumnStore()
	const { socket } = useSocket()
	const { handle: plugHandle } = usePlugStore()

	const [width, setWidth] = useState(COLUMNS.DEFAULT_WIDTH)
	const [isResizing, setIsResizing] = useState(false)

	useEffect(() => {
		const getBoundedWidth = (width: number) => Math.min(Math.max(width, 380), 620)

		const handleMouseMove = (e: MouseEvent) => {
			if (!resizeRef.current || !isResizing) return

			setWidth(getBoundedWidth(e.clientX - resizeRef.current.getBoundingClientRect().left))
		}

		const handleMouseUp = () => {
			setIsResizing(false)
		}

		if (isResizing) {
			window.addEventListener("mousemove", handleMouseMove)
			window.addEventListener("mouseup", handleMouseUp)
		}

		return () => {
			window.removeEventListener("mousemove", handleMouseMove)
			window.removeEventListener("mouseup", handleMouseUp)
		}
	}, [isResizing])

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

	const isApproved = Boolean(socket?.identity?.approvedAt)
	if (!isApproved) return null

	const options = [
		...flagOptions,
		...OPTIONS,
		{
			label: "SETTINGS",
			description: "View and manage your Plug settings.",
			icon: <Cog size={14} className="opacity-40" />
		},
		socket?.admin ? {
			label: "ADMIN",
			description: "Manage administrative settings.",
			icon: <LockIcon size={14} className="opacity-40" />
		} : {}
	]

	const isBody = index != columns.length - 2

	return (
		<>
			<div
				ref={resizeRef}
				className={cn(
					"relative flex select-none flex-col rounded-lg",
					!isBody && "my-2 border-[1px] border-plug-green/10 bg-white",
					columns.some(column => column.index >= 0) ? "" : "ml-2"
				)}
				style={{ width, minWidth: width }}
			>
				{!isBody && (
					<div className="relative flex cursor-pointer flex-row items-center overflow-hidden overflow-y-auto rounded-t-lg border-b-[1px] border-plug-green/10 bg-white transition-all duration-200 ease-in-out">
						<div className="flex w-full flex-row items-center gap-4 px-6 py-4">
							<Plus size={18} className="opacity-40" />
							<p className="overflow-hidden truncate overflow-ellipsis text-lg font-bold">Add Column</p>
						</div>
					</div>
				)}

				<div className="h-full overflow-y-scroll">
					<div className="flex h-full flex-col gap-2 p-4">
						<Accordion key={"add"} onExpand={() => plugHandle.plug.add(isBody ? { index } : undefined)}>
							<div className="flex flex-row items-center gap-2">
								<div className="flex h-10 w-10 min-w-10 items-center justify-center">
									<Plug size={14} className="opacity-40" />
								</div>

								<div className="flex flex-col items-start text-left font-bold">
									<p>Create Plug</p>
									<p className="text-sm opacity-40">Start creating a new Plug from scratch.</p>
								</div>
							</div>
						</Accordion>
						{options.map(option => (
							<Accordion
								key={option.label}
								onExpand={() =>
									isBody
										? handle.navigate({ index, key: option.label })
										: handle.add({ key: option.label })
								}
							>
								<div className="flex flex-row items-center gap-2">
									<div className="flex h-10 w-10 min-w-10 items-center justify-center">
										{option.icon}
									</div>

									<div className="flex flex-col items-start text-left font-bold">
										<p>{formatTitle(option?.label?.replace("_", " ").toLowerCase() ?? "")}</p>
										<p className="text-sm opacity-40">{option.description}</p>
									</div>
								</div>
							</Accordion>
						))}
					</div>
				</div>

			</div>

			<div
				className="h-full cursor-col-resize pl-2"
				onMouseDown={e => {
					e.preventDefault()
					setIsResizing(true)
				}}
			>
				<div className="h-full w-[1px] bg-plug-green/10" />
			</div>
		</>
	)
}
