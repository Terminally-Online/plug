import { useState } from "react"

import { CheckCheck, SearchIcon } from "lucide-react"

import { MinimalUserSocketModel, UserSocketModel } from "@/prisma/types"
import { api } from "@/server/client"

import { Accordion, Avatar, Image, Search } from "@/components"
import { Column, formatAddress, formatTitle, useDebounce, VIEW_KEYS } from "@/lib"
import { useColumns, useSocket } from "@/state"

const EXCLUDED_KEYS = [VIEW_KEYS.HOME, VIEW_KEYS.VIEW_AS, VIEW_KEYS.ADD, VIEW_KEYS.PLUG]

export const ColumnViewAs = () => {
	const { socket } = useSocket()
	const { columns, as } = useColumns()

	const [sockets, setSockets] = useState<UserSocketModel[]>([])
	const [search, debouncedSearch, setSearch] = useDebounce("")
	const [view, setView] = useState<MinimalUserSocketModel | undefined>(socket)

	const options = socket && sockets ? [socket, ...sockets] : sockets

	api.socket.search.useQuery(
		{
			search: debouncedSearch
		},
		{
			onSettled: data => data && setSockets(data)
		}
	)

	return (
		<div className="flex h-full flex-col py-4">
			<div className="px-4">
				<Search
					className="mb-2"
					icon={<SearchIcon size={14} className="opacity-40" />}
					placeholder="Search account or socket"
					search={search}
					handleSearch={setSearch}
					clear
				/>

				{socket && options && options.length > 0 && (
					<div className="flex flex-col gap-2">
						{options.map(option => (
							<Accordion key={option.id} onExpand={() => setView(option)}>
								<div className="flex flex-row items-center gap-4 whitespace-nowrap">
									<div className="relative h-8 w-8 min-w-8 rounded-sm">
										{option.identity?.ens?.avatar ? (
											<>
												<Image
													className="absolute left-0 top-0 blur-xl filter"
													src={option.identity.ens.avatar}
													alt="ENS Avatar"
													width={240}
													height={240}
												/>
												<Image
													className="relative rounded-md"
													src={option.identity.ens.avatar}
													alt="ENS Avatar"
													width={240}
													height={240}
												/>
											</>
										) : (
											<>
												<div className="absolute top-1/2 blur-xl filter">
													<Avatar name={option.id} />
												</div>
												<Avatar name={option.id} />
											</>
										)}
									</div>

									<div className="flex flex-col items-start">
										<p className="font-bold">
											{option.id === socket.id
												? "Yourself"
												: (option.identity?.ens?.name ?? formatAddress(option.id))}
										</p>
										<p className="text-sm font-bold opacity-40">
											{option.socketAddress ? formatAddress(option.socketAddress) : "Anonymous"}
										</p>
									</div>

									{view && option.id === view.id && (
										<CheckCheck size={14} className="ml-auto text-plug-green" />
									)}
								</div>
							</Accordion>
						))}
					</div>
				)}
			</div>

			<div className="my-4 h-[1px] w-full bg-grayscale-100" />

			<div className="px-4">
				<div className="flex flex-col gap-2">
					{socket &&
						view &&
						columns
							.filter(column => EXCLUDED_KEYS.includes(column.key) === false)
							.map((column: Column) => (
								<Accordion key={column.index} onExpand={() => as({ index: column.index, as: view })}>
									<div className="flex flex-row items-center justify-between gap-4">
										<p className="font-bold">{formatTitle(column.key.toLowerCase())}</p>
										<>
											<div className="relative h-6 w-6 min-w-6 overflow-hidden rounded-sm">
												{column.viewAs && column.viewAs.identity?.ens?.avatar ? (
													<>
														<Image
															className="absolute left-0 top-0 blur-xl filter"
															src={column.viewAs.identity.ens.avatar}
															alt="ENS Avatar"
															width={240}
															height={240}
														/>
														<Image
															className="relative rounded-sm"
															src={column.viewAs.identity.ens.avatar}
															alt="ENS Avatar"
															width={240}
															height={240}
														/>
													</>
												) : (
													<>
														<div className="absolute left-0 top-0 blur-xl filter">
															<Avatar
																name={column.viewAs?.id ?? socket.id}
																className="rounded-sm"
															/>
														</div>
														<Avatar
															name={column.viewAs?.id ?? socket.id}
															className="rounded-sm"
														/>
													</>
												)}
											</div>
										</>
									</div>
								</Accordion>
							))}
				</div>
			</div>
		</div>
	)
}
