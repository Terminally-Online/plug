import { useEffect, useState } from "react"

import { useRouter } from "next/router"

import BlockiesSvg from "blockies-react-svg"
import { Ellipsis } from "lucide-react"

import {
	Container,
	Header,
	ManageSocketFrame,
	SocketActivity,
	SocketPositionList,
	SocketTabs,
	SocketTokens
} from "@/components"
import { useFrame, useSockets } from "@/contexts"
import { NextPageWithLayout, routes } from "@/lib"

const Page: NextPageWithLayout = () => {
	const router = useRouter()
	const address = router.query.address as string

	const { socket, handleSelect } = useSockets()
	const { handleFrameVisible } = useFrame()

	const [selected, setSelected] = useState(0)

	useEffect(() => handleSelect(address), [handleSelect, address])

	if (!socket) return null

	return (
		<>
			<Container>
				<Header
					size="lg"
					back={routes.app.index}
					icon={
						<BlockiesSvg
							address={socket.socketAddress}
							className="h-6 w-6 rounded-sm"
						/>
					}
					label={socket.name}
					nextOnClick={() => handleFrameVisible("manage")}
					nextLabel={<Ellipsis size={14} />}
				/>
			</Container>

			<SocketTabs selected={selected} onSelect={setSelected} />

			<Container>
				{selected === 0 ? (
					<SocketActivity />
				) : selected === 1 ? (
					<SocketTokens />
				) : (
					<SocketPositionList />
				)}

				<ManageSocketFrame />
			</Container>
		</>
	)
}

export default Page
