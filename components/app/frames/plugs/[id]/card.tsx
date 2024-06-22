// // i think this file can go?
// import { FC } from "react"

// import Image from "next/image"

// import { Badge, Link, Send, Twitter } from "lucide-react"

// import { Button } from "@/components/buttons"
// import { useFrame, usePlugs } from "@/contexts"
// import { routes, useClipboard } from "@/lib"

// import { Frame } from "../../base"

// export const CardFrame: FC = () => {
// 	const { frameVisible, handleFrameVisible } = useFrame()
// 	// const [search, setSearch] = useState("")

// 	return (
// 		<Frame
// 			className="z-[2]"
// 			icon={<Badge size={18} className="opacity-60" />}
// 			label="Share Plug"
// 			visible={frameVisible === "share"}
// 			handleVisibleToggle={() => handleFrameVisible(undefined)}
// 		>
// 			<div className="plug-card w-full rounded-lg bg-white p-6 shadow-lg">
// 				<div className="mb-4 flex items-center">
// 					<img
// 						src="@/public/wallets/danner.png"
// 						className="mr-4 rounded-full"
// 					/>
// 					<div className="flex flex-col">
// 						<span className="text-lg font-bold">nftchance.eth</span>
// 						<span className="text-sm text-gray-500">
// 							0xe7af6...5ea429
// 						</span>
// 					</div>
// 				</div>
// 				<div className="space-y-2">
// 					<div className="flex items-center text-green-500">
// 						<span className="mr-2">🟢</span>
// 						<span>
// 							Can be run after <strong>01/01/2024</strong>
// 						</span>
// 					</div>
// 					<div className="flex items-center text-green-500">
// 						<span className="mr-2">🟢</span>
// 						<span>
// 							Can be run <strong>1</strong> time a{" "}
// 							<strong>day</strong>
// 						</span>
// 					</div>
// 					<div className="flex items-center text-yellow-500">
// 						<span className="mr-2">🟠</span>
// 						<span>
// 							Can bid <strong>4 ETH</strong> to win Noun
// 						</span>
// 					</div>
// 					<div className="flex items-center text-yellow-500">
// 						<span className="mr-2">🟠</span>
// 						<span>
// 							Can bid on Noun with a{" "}
// 							<strong>Pineapple Hat</strong>
// 						</span>
// 					</div>
// 				</div>
// 				<button className="mt-4 w-full rounded-lg bg-green-500 py-2 text-white">
// 					Share
// 				</button>
// 			</div>
// 		</Frame>
// 	)
// }

// export default MyComponent
