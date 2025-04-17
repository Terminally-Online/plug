import { useSession } from "next-auth/react"
import Image from "next/image"
import { FC } from "react"

import { Eye, GitFork } from "lucide-react"

import { Accordion } from "@/components/shared/utils/accordion"
import { Counter } from "@/components/shared/utils/counter"
import { cn, colors, formatTitle } from "@/lib"
import { RouterOutputs } from "@/server/client"
import { useSocket } from "@/state/authentication"
import { COLUMNS, useColumnActions } from "@/state/columns"

import { Avatar } from "../../sockets/profile"

type Props = { index: number; from: string; plug: RouterOutputs["plugs"]["all"][number] | undefined }

export const PlugGridItem: FC<Props> = ({ index, from, plug }) => {
	const { navigate } = useColumnActions(index)
	const { socket } = useSocket()

	return (
		<>
			<Accordion
				onExpand={
					plug
						? () =>
								navigate({
									index,
									key: COLUMNS.KEYS.PLUG,
									item: plug.id,
									from
								})
						: undefined
				}
				loading={!plug}
				className={cn(
					"relative flex h-[160px] w-full flex-col overflow-hidden text-left",
					!!plug && "bg-plug-white"
				)}
				noPadding
			>
				{plug === undefined ? (
					<div className="invisible">
						<p>.</p>
						<p>.</p>
					</div>
				) : (
					<>
						<div className="relative z-[999] flex h-full w-full flex-row justify-between bg-plug-white p-4 transition-all duration-200 ease-in-out group-hover:bg-transparent">
							<div className="mt-auto flex flex-col">
								<p className="flex flex-row gap-2 truncate text-sm font-bold tabular-nums">
									<Eye
										size={14}
										className="h-4 w-4 opacity-40 transition-all duration-200 ease-in-out group-hover:opacity-60"
									/>
									<span className="opacity-40 transition-all duration-200 ease-in-out group-hover:opacity-100">
										<Counter count={plug?.views?.[0]?.views ?? 0} />
									</span>
								</p>
								{/*<p className="flex flex-row items-end gap-2 truncate text-sm font-bold tabular-nums">
							<Play
								size={14}
								className="h-4 w-4 opacity-40 transition-all duration-200 ease-in-out group-hover:opacity-60"
							/>
							<span className="opacity-40 transition-all duration-200 ease-in-out group-hover:opacity-100">
								<Counter count={plug?._count?.executions ?? 0} />
							</span>
						</p>*/}
							</div>

							<p className="flex flex-row items-end gap-2 truncate text-sm font-bold tabular-nums">
								<GitFork
									size={14}
									className="h-4 w-4 opacity-40 transition-all duration-200 ease-in-out group-hover:opacity-60"
								/>
								<span className="opacity-40 transition-all duration-200 ease-in-out group-hover:opacity-100">
									<Counter count={plug?.forkCount ?? 0} />
								</span>
							</p>
						</div>

						<div className="relative flex flex-row items-start justify-between gap-8 rounded-b-lg border-t-[1px] border-t-plug-green/10 bg-plug-white p-4">
							<div className="absolute inset-0">
								<div
									className="absolute -bottom-full -left-1/3 h-full w-3/4 rounded-full blur-[80px] filter"
									style={{
										backgroundColor: plug ? colors[plug.color as keyof typeof colors] : undefined
									}}
								></div>
							</div>

							<div className="relative z-10 flex-1">
								<p
									className={cn(
										"line-clamp-2 break-words font-bold leading-snug",
										!plug && "invisible"
									)}
								>
									{plug ? (plug.name === "" ? "Untitled Plug" : formatTitle(plug.name)) : "."}
								</p>
							</div>

							{plug?.socketId !== socket?.id && (
								<div className="relative ml-auto h-6 w-6 min-w-6 shrink-0">
									<div className="relative z-[21] h-6 w-6">
										{plug?.socket?.identity?.ens?.avatar ? (
											<Image
												src={plug?.socket.identity.ens.avatar ?? ""}
												alt="ENS Avatar"
												width={64}
												height={64}
												className="h-full w-full rounded-sm"
											/>
										) : (
											<div className="h-full w-full rounded-sm">
												<Avatar name={plug?.socketId ?? ""} />
											</div>
										)}
									</div>
									<div className="absolute -right-4 -top-4 z-20 h-16 w-16">
										{plug?.socket?.identity?.ens?.avatar ? (
											<Image
												src={plug?.socket.identity.ens.avatar ?? ""}
												alt="ENS Avatar"
												width={64}
												height={64}
												className="h-full w-full rounded-sm blur-[60px] filter"
											/>
										) : (
											<div className="h-full w-full rounded-sm blur-[60px] filter">
												<Avatar name={plug?.socketId ?? ""} />
											</div>
										)}
									</div>
								</div>
							)}
						</div>
					</>
				)}
			</Accordion>
		</>
	)
}
