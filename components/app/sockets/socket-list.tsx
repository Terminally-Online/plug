import { Button } from "@/components/buttons"
import { useSockets } from "@/contexts"

import { SocketItem } from "./socket-item"

export const SocketList = () => {
	const { sockets, handleAdd: handleSocketAdd } = useSockets()

	const hasSockets = sockets && sockets.length > 0

	return (
		<>
			{hasSockets ? (
				<div className="flex flex-col gap-2">
					{sockets.map((socket, index) => (
						<SocketItem key={index} socket={socket} />
					))}
				</div>
			) : (
				<div className="my-32 flex flex-col gap-[30px]">
					<p className="mx-auto w-[80%] max-w-[360px] text-center text-lg opacity-60">
						A Socket is needed to execute your first Plug. Create
						and fund your Socket with a single transaction now.
					</p>

					<Button className="mx-auto w-max" onClick={handleSocketAdd}>
						Create Socket
					</Button>
				</div>
			)}
		</>
	)
}
