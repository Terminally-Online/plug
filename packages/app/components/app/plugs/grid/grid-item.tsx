import { useSession } from "next-auth/react"
import Image from "next/image"
import { FC } from "react"

import { Eye, GitFork, Play } from "lucide-react"

import { Accordion } from "@/components/shared/utils/accordion"
import { Counter } from "@/components/shared/utils/counter"
import { cn, colors, formatAddress, formatTitle } from "@/lib"
import { RouterOutputs } from "@/server/client"
import { COLUMNS, useColumnStore } from "@/state/columns"

import { Avatar } from "../../sockets/profile"

type Props = { index: number; from: string; plug: RouterOutputs["plugs"]["all"][number] | undefined }

export const PlugGridItem: FC<Props> = ({ index, from, plug }) => {
	const { handle } = useColumnStore()
	const { data: session } = useSession()

	return (
		<Accordion
			onExpand={
				plug
					? () =>
							handle.navigate({
								index,
								key: COLUMNS.KEYS.PLUG,
								item: plug.id,
								from
							})
					: undefined
			}
			loading={!plug}
			className="relative flex h-[160px] w-full flex-col overflow-hidden bg-plug-white text-left"
			noPadding
		>
			<div className="relative z-[2] flex h-full w-full flex-row items-start justify-between bg-plug-white p-4 transition-all duration-200 ease-in-out group-hover:bg-transparent">
				<div className="ml-auto flex flex-col">
					<p className="ml-auto flex flex-row gap-2 truncate text-sm font-bold tabular-nums">
						<Eye
							size={14}
							className="h-4 w-4 opacity-40 transition-all duration-200 ease-in-out group-hover:opacity-60"
						/>
						<span className="opacity-40 transition-all duration-200 ease-in-out group-hover:opacity-100">
							<Counter count={plug?.views?.[0]?.views ?? 0} />
						</span>
					</p>
					<p className="ml-auto flex flex-row items-end gap-2 truncate text-sm font-bold tabular-nums">
						<Play
							size={14}
							className="h-4 w-4 opacity-40 transition-all duration-200 ease-in-out group-hover:opacity-60"
						/>
						<span className="opacity-40 transition-all duration-200 ease-in-out group-hover:opacity-100">
							<Counter count={plug?._count?.executions ?? 0} />
						</span>
					</p>
					<p className="ml-auto flex flex-row items-end gap-2 truncate text-sm font-bold tabular-nums">
						<GitFork
							size={14}
							className="h-4 w-4 opacity-40 transition-all duration-200 ease-in-out group-hover:opacity-60"
						/>
						<span className="opacity-40 transition-all duration-200 ease-in-out group-hover:opacity-100">
							<Counter count={plug?.forkCount ?? 0} />
						</span>
					</p>
				</div>
			</div>

			<div className="flex flex-row items-start justify-between gap-8 overflow-visible border-t-[1px] border-t-plug-green/10 bg-plug-white p-4">
				<div className="absolute inset-0 overflow-hidden">
					<div
						className="absolute -bottom-full -left-1/3 h-full w-3/4 rounded-full blur-[80px] filter"
						style={{
							backgroundColor: plug ? colors[plug.color as keyof typeof colors] : undefined
						}}
					></div>
				</div>

				<span className={cn("break-words pb-1 font-bold", !plug && "invisible")}>
					{plug ? (plug.name === "" ? "Untitled Plug" : formatTitle(plug.name)) : "."}
				</span>

				{plug?.socketId !== session?.user.id && (
					<p className="ml-auto flex flex-row items-center gap-2 truncate text-sm font-bold tabular-nums">
						<span className="relative h-6 w-6">
							{plug?.socket?.identity?.ens?.avatar ? (
								<Image
									src={plug?.socket.identity.ens.avatar ?? ""}
									alt="ENS Avatar"
									width={64}
									height={64}
									className="absolute left-0 top-0 h-12 w-12 rounded-sm blur-xl filter"
								/>
							) : (
								<Avatar name={plug?.socketId ?? ""} />
							)}

							{plug?.socket?.identity?.ens?.avatar ? (
								<Image
									src={plug?.socket.identity.ens.avatar ?? ""}
									alt="ENS Avatar"
									width={64}
									height={64}
									className="relative h-full w-full rounded-sm"
								/>
							) : (
								<Avatar name={plug?.socketId ?? ""} />
							)}
						</span>
					</p>
				)}
			</div>
		</Accordion>
	)
}
