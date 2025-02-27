import { FC } from "react"

import { Image } from "@/components/app/utils/image"
import { Accordion } from "@/components/shared/utils/accordion"
import { connectedChains } from "@/contexts"
import { ActionSchema, formatTitle, getValues } from "@/lib"
import { useColumnStore } from "@/state/columns"
import { usePlugStore } from "@/state/plugs"

import { ChainImage } from "../../sockets/chains/chain.image"

export const ActionItem: FC<{
	index: number
	item: string
	actionName: string
	protocol: string
	action: ActionSchema
	image?: boolean
}> = ({ index, item, protocol, actionName, action }) => {
	const { handle } = useColumnStore(index)
	const { plug, actions, handle: plugHandle } = usePlugStore(item)

	if (!plug) return null

	return (
		<Accordion
			onExpand={() => {
				plugHandle.action.edit({
					id: plug.id,
					actions: JSON.stringify([
						...actions,
						{
							protocol,
							action: actionName,
							id: Math.floor(Math.random() * 100_000_000_000),
							...getValues(action.schema[actionName].sentence)
						}
					])
				})
				handle.frame()
			}}
		>
			<div className="flex flex-row items-center gap-2">
				<div className="relative h-6 w-10 min-w-10">
					<Image
						src={action.metadata.icon}
						alt={"icon"}
						width={128}
						height={128}
						className="absolute left-1/2 top-1/2 mr-2 h-12 w-12 -translate-x-1/2 -translate-y-1/2 rounded-full blur-2xl filter"
					/>
					<Image
						src={action.metadata.icon}
						alt={"icon"}
						width={128}
						height={128}
						className="absolute left-1/2 top-1/2 mr-2 h-6 w-6 -translate-x-1/2 -translate-y-1/2 rounded-sm"
					/>
				</div>

				<p className="font-bold">
					<span className="opacity-40">{formatTitle(protocol)}: </span>
					{formatTitle(actionName)}
				</p>

				{/* TODO: Re-implement this once we have more than one chain id active otherwise it is
					      just a redundant display because they all support the same chain.
				<div className="ml-auto flex flex-row items-center">
					<pre>{JSON.stringify(
						action.metadata.chains
						, null, 2)}</pre>

					{action.metadata.chains
						.filter(chain => connectedChains.some(c => chain.chainIds.includes(c.id)))
						.map(chain => (
							<div key={chain} className="-ml-1">
								<ChainImage chainId={chain} />
							</div>
						))}
				</div>
				*/}
			</div>
		</Accordion>
	)
}
