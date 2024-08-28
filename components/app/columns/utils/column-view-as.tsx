import { useEffect, useState } from "react"

import Image from "next/image"

import BoringAvatar from "boring-avatars"
import { CheckCheck, SearchIcon } from "lucide-react"

import { Accordion } from "@/components/shared"
import { useSockets } from "@/contexts"
import { formatAddress, formatTitle, useDebounce, VIEW_KEYS } from "@/lib"
import { UserSocketModel } from "@/prisma/types"
import { api } from "@/server/client"

import { Search } from "../../inputs"

const EXCLUDED_KEYS = [VIEW_KEYS.HOME, VIEW_KEYS.VIEW_AS, VIEW_KEYS.ADD, VIEW_KEYS.PLUG]

export const ColumnViewAs = () => {
	const { socket, handle } = useSockets()

	const [sockets, setSockets] = useState<UserSocketModel[]>([])
	const [search, debouncedSearch, setSearch] = useDebounce("")
	const [as, setAs] = useState<UserSocketModel | undefined>(socket)
	const [columns, setColumns] = useState<
		Array<{ column: UserSocketModel["columns"][number]; as: UserSocketModel }> | undefined
	>(undefined)

	const options = socket && sockets ? [socket, ...sockets] : sockets

	// NOTE: This is a kind of hacky way to prevent flashing when searching for sockets. This way,
	//       we update the local state when the search hits instead of showing a loading state.
	api.socket.search.useQuery(
		{
			search: debouncedSearch
		},
		{
			onSettled: data => data && setSockets(data)
		}
	)

	useEffect(() => {
		if (!socket || columns !== undefined) return
		setColumns(socket.columns.map(column => ({ column, as: socket })))
	}, [socket, columns])

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
							<Accordion
								key={option.id}
								onExpand={() => setAs(prev => (prev && prev.id === option.id ? undefined : option))}
							>
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
												<div className="absolute left-0 top-0 blur-xl filter">
													<BoringAvatar
														variant="beam"
														name={option.id}
														size={"100%"}
														colors={["#00E100", "#A3F700"]}
														square
													/>
												</div>
												<div className="relative overflow-hidden rounded-md">
													<BoringAvatar
														variant="beam"
														name={option.id}
														size={"100%"}
														colors={["#00E100", "#A3F700"]}
														square
													/>
												</div>
											</>
										)}
									</div>

									<div className="flex flex-col items-start">
										<p className="font-bold">
											{option.id === socket.id
												? "Yourself"
												: (option.identity?.ens?.name ?? option.id)}
										</p>
										<p className="text-sm font-bold opacity-40">
											{option.socketAddress ? formatAddress(option.socketAddress) : "Anonymous"}
										</p>
									</div>

									{as && option.id === as.id && (
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
						columns &&
						as &&
						columns
							.filter(column => EXCLUDED_KEYS.includes(column.column.key) === false)
							.map(column => (
								<Accordion
									key={column.column.id}
									onExpand={() => {
										handle.columns.as({ id: column.column.id, as: as.id })
										setColumns(
											prev =>
												prev &&
												prev.map(prevColumn => ({
													...prevColumn,
													as: prevColumn.column.id === column.column.id ? as : prevColumn.as
												}))
										)
									}}
								>
									<div className="flex flex-row items-center justify-between gap-4">
										<p className="font-bold">{formatTitle(column.column.key.toLowerCase())}</p>
										<>
											<div className="relative h-6 w-6 min-w-6 rounded-sm">
												{column.as && column.as.identity?.ens?.avatar ? (
													<>
														<Image
															className="absolute left-0 top-0 blur-xl filter"
															src={column.as.identity.ens.avatar}
															alt="ENS Avatar"
															width={240}
															height={240}
														/>
														<Image
															className="relative rounded-sm"
															src={column.as.identity.ens.avatar}
															alt="ENS Avatar"
															width={240}
															height={240}
														/>
													</>
												) : (
													<>
														<div className="absolute left-0 top-0 blur-xl filter">
															<BoringAvatar
																variant="beam"
																name={column.as?.id ?? socket.id}
																size={"100%"}
																colors={["#00E100", "#A3F700"]}
																square
															/>
														</div>
														<div className="relative overflow-hidden rounded-sm">
															<BoringAvatar
																variant="beam"
																name={column.as?.id ?? socket.id}
																size={"100%"}
																colors={["#00E100", "#A3F700"]}
																square
															/>
														</div>
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
