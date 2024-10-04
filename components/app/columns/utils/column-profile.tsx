import { FC } from "react"

import { Avatar, Image } from "@/components"
import { formatAddress } from "@/lib"
import { useSocket } from "@/state"

export const ColumnProfile: FC<{ index: number }> = () => {
	const { name, avatar, socket } = useSocket()

	const columns = 4

	if (!socket) return null

	return (
		<div className="flex h-full flex-col gap-4 text-center py-4 overflow-hidden">
			<div className="flex flex-row gap-8 pb-4 px-4 items-center">
				<div className="w-24 h-24 relative">
					{avatar ? (
						<Image
							src={avatar}
							alt="ENS Avatar"
							width={64}
							height={64}
							className="h-full w-full rounded-sm"
						/>
					) : (
						<>
							<div className="absolute filter blur-[80px]">
								<Avatar name={socket?.id ?? ""} />
							</div>
							<Avatar name={socket?.id ?? ""} />
						</>
					)}
				</div>


				<div className="flex flex-col">
					<p className="font-bold text-lg mr-auto">{name !== "" ? name : formatAddress(socket.id, 6)}</p>
					<p className="opacity-40 font-bold text-lg mr-auto">Joined: {new Date().toLocaleDateString()}</p>
				</div>
			</div>

			<div className="relative flex flex-row gap-2 px-4">
				{Array.from({ length: columns }).map((_, i) => (
					<div key={i} className="h-32 w-full bg-grayscale-0 rounded-md" />
				))}
			</div>

			<div className="px-4 flex flex-col gap-2">
				<div className="flex flex-row gap-2">
					<div className="bg-grayscale-0 rounded-md w-full h-32" />
					<div className="bg-grayscale-0 rounded-md w-full h-32" />
				</div>
				<div className="flex flex-row gap-2">
					<div className="bg-grayscale-0 rounded-md w-full h-32" />
					<div className="bg-grayscale-0 rounded-md w-full h-32" />
				</div>
			</div>
		</div>
	)
}
