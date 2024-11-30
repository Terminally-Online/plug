import { FC } from "react"

export const ChainFrame: FC<{ index: number; item: string }> = ({ index, item }) => {
	index
	item
	// const { isFrame, frame } = useColumns(index, "chain")
	// const { chains, actions, handle } = usePlugs(item)

	// const isDisabled = chains.length === 0

	// useEffect(() => {
	// 	if (isFrame === false) return
	//
	// 	if (chains.length < 2 && chains.length === 1) handleFrame(nextFrame)
	// }, [chains, isFrame])

	return <></>

	// return (
	// // <Frame
	// // 	id={id}
	// // 	className="z-[2]"
	// // 	icon={<Globe size={18} />}
	// // 	label={"Choose Chain" + (chainsAvailable.length > 1 ? "s" : "")}
	// // 	visible={isFrame}
	// // 	hasOverlay={true}
	// // >
	// 	{/* <div className="flex flex-col gap-4">
	// 		{chainsAvailable.map((chain, index) => (
	// 			<div
	// 				key={`chain-${index}`}
	// 				className="flex flex-row items-center gap-4"
	// 			>
	// 				<Checkbox
	// 					checked={chains.includes(chain)}
	// 					handleChange={() => handle.chain.select(chain)}
	// 				/>

	// 				<div className="mr-auto flex flex-row gap-2">
	// 					<Image
	// 						src={`/blockchain/${chain}.png`}
	// 						alt={formatTitle(chain)}
	// 						width={64}
	// 						height={64}
	// 						className="h-6 w-6"
	// 					/>
	// 					<p className="font-bold">{formatTitle(chain)}</p>
	// 				</div>

	// 				<p className="tabular-nums opacity-60">23.005 ETH</p>
	// 			</div>
	// 		))}

	// 		<Button
	// 			variant={isDisabled ? "pointer disabled" : "primary"}
	// 			className="mt-4"
	// 			onClick={() => handleFrame(nextFrame)}
	// 			disabled={isDisabled}
	// 		>
	// 			{isDisabled ? "Choose a Chain" : "Next"}
	// 		</Button>
	// 	</div> */}
	// // </Frame>
	// )
}
