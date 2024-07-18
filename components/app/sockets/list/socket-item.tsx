import { FC } from "react"

import Link from "next/link"

import BlockiesSvg from "blockies-react-svg"

import { formatAddress } from "@/lib/functions"
import { cn } from "@/lib/utils"
import { UserSocket } from "@/server/api/routers/socket"

type Props = { socket: UserSocket }

export const SocketItem: FC<Props> = ({ socket }) => {
	// TODO: Implement the backend functionality for this.
	const activity = Array.from({ length: 10 }, () => Math.random())
	// TODO: Implement display of the deployed chains.

	return (
		<Link
			className="flex flex-row items-center gap-4"
			href={"/app/sockets/" + socket.socketAddress}
		>
			<BlockiesSvg
				address={socket.socketAddress}
				className="h-10 w-10 rounded-md"
			/>

			<div className="flex flex-col">
				<p className="font-bold">{socket.name}</p>
				<p className="font-light text-black/65">
					{formatAddress(socket.socketAddress)}
				</p>
			</div>

			<div className="ml-auto mt-auto flex flex-row gap-1">
				{activity.map((activity, index) => (
					<div
						key={index}
						className={cn(
							"mt-auto w-4 rounded-sm",
							index === 9
								? "bg-gradient-to-tr from-[#00E100] to-[#A3F700]"
								: "bg-grayscale-100"
						)}
						style={{
							height: `${Math.max(10, activity * 32)}px`
						}}
					/>
				))}
			</div>
		</Link>
	)
}
