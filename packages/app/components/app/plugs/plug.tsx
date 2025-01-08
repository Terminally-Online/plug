import { useSession } from "next-auth/react"
import { FC, HTMLAttributes, useEffect, useState } from "react"

import { Check, ChevronLeft, Settings, Share } from "lucide-react"

import {
	ActionsFrame,
	ActionView,
	AuthRequiredFrame,
	Button,
	Container,
	DeletedFrame,
	ExecuteFrame,
	ManagePlugFrame,
	Search,
	ShareFrame
} from "@/components"
import { cardColors, cn } from "@/lib"
import { COLUMNS, useColumnStore, usePlugData } from "@/state"

export const Plug: FC<HTMLAttributes<HTMLDivElement> & { index?: number; item?: string; from?: string }> = ({
	index = COLUMNS.MOBILE_INDEX,
	item,
	from,
	...props
}) => {
	const { data: session } = useSession()
	const { handle } = useColumnStore(index)
	const { plug } = usePlugData(item)
	const [hasOpenedActions, setHasOpenedActions] = useState(false)
	const [copied, setCopied] = useState(false)

	const own = plug !== undefined && session && session.address === plug.socketId

	useEffect(() => {
		if (!plug || plug.actions !== "[]" || hasOpenedActions) return
		handle.frame(`${item}-actions`)
		setHasOpenedActions(true)
	}, [item, handle, plug, hasOpenedActions])

	useEffect(() => {
		const handleResize = () => {
			if (window.innerWidth >= 768 && index === COLUMNS.MOBILE_INDEX) {
				handle.add({
					key: COLUMNS.KEYS.PLUG,
					item: item,
					from: from
				})
			}
		}

		window.addEventListener("resize", handleResize)
		return () => window.removeEventListener("resize", handleResize)
	}, [handle, index, item, from])

	if (!session) return null

	return (
		<div {...props}>
			{plug ? (
				<>
					<Container className="border-grayscale-100 fixed left-0 right-0 top-0 z-[10] border-b-[1px] bg-white md:hidden">
						<div className="flex flex-row items-center gap-4 py-4">
							<Button
								variant="secondary"
								className="rounded-sm p-1"
								onClick={() =>
									handle.navigate({
										index: -1,
										key: from ?? COLUMNS.KEYS.HOME
									})
								}
							>
								<ChevronLeft size={14} />
							</Button>

							<div className="flex flex-row items-center gap-2">
								<div
									className="h-6 w-6 min-w-6 rounded-md"
									style={{
										backgroundImage: cardColors[plug.color]
									}}
								/>
								<span className="font-bold">{plug.name || "Untitled Plug"}</span>
							</div>

							<div className="ml-auto flex flex-row gap-2">
								{own && (
									<Button
										variant="secondary"
										className="rounded-sm p-1"
										onClick={() => handle.frame("manage")}
									>
										<Settings size={14} />
									</Button>
								)}
								<Button
									variant="secondary"
									className="rounded-sm p-1"
									onClick={async () => {
										try {
											const shareUrl = `${window.location.origin}/app/?id=${plug.id}`
											await navigator.clipboard.writeText(shareUrl)
											setCopied(true)
											setTimeout(() => setCopied(false), 2000)
										} catch (err) {
											console.error("Failed to copy link:", err)
										}
									}}
								>
									{copied ? (
										<Check size={14} className="opacity-60 transition-all" />
									) : (
										<Share size={14} />
									)}
								</Button>
							</div>
						</div>
					</Container>

					<div className="mt-16 md:mt-0">
						<ActionView index={index} />

						{/* FOOTER SECTION */}
						<div className="absolute bottom-0 left-0 z-[2] mb-4 flex w-full flex-col gap-2 overflow-y-visible">
							{/* Gradient overlay for visual transition */}
							<div className="pointer-events-none absolute bottom-[120px] left-0 right-0 top-0 z-[-1] bg-gradient-to-t from-white to-white/0" />

							{/* White background container */}
							<div
								className={cn(
									"absolute -bottom-4 left-0 right-0 z-[-1] h-[140px] bg-white",
									index !== COLUMNS.MOBILE_INDEX && "rounded-b-lg"
								)}
							/>

							{/* Search bar - only visible to plug owner */}
							{own && (
								<Search
									className="px-4 pt-16"
									icon={<Settings size={14} className="opacity-60" />}
									placeholder="Search protocols and actions"
									handleOnClick={() => handle.frame(`${item}-actions`)}
								/>
							)}

							{/* Action buttons - ALWAYS visible */}
							<div className="relative flex flex-row gap-2 px-4">
								<Button
									variant="secondary"
									className="w-max bg-white py-4"
									onClick={() => {
										handle.schedule()
										handle.frame("run")
									}}
								>
									Run
								</Button>

								<Button className="w-full py-4" onClick={() => handle.frame("schedule")}>
									Schedule
								</Button>
							</div>
						</div>
					</div>

					{item && (
						<>
							<AuthRequiredFrame index={index} />
							<ExecuteFrame index={index} item={item} />
							<ManagePlugFrame index={index} item={item} from={from} />
							<ActionsFrame index={index} item={item} />
							<ShareFrame index={index} item={item} />
						</>
					)}
				</>
			) : item ? (
				<DeletedFrame index={index} />
			) : null}
		</div>
	)
}