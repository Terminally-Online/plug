import { FC, useCallback } from "react"

import { Image } from "@/components/app/utils/image"
import { Accordion } from "@/components/shared/utils/accordion"
import { SchemasResponseSchema, formatTitle, getValues } from "@/lib"
import { useColumnActions } from "@/state/columns"
import { editPlugAtom, plugByIdAtom } from "@/state/plugs"
import { useAtom, useSetAtom } from "jotai"
import { api } from "@/server/client"

export const ActionItem: FC<{
	index: number
	item: string
	actionName: string
	protocol: string
	action: SchemasResponseSchema
	image?: boolean
}> = ({ index, item, protocol, actionName, action }) => {
	const { frame } = useColumnActions(index)

	const [plug] = useAtom(plugByIdAtom(item))
	const editPlug = useSetAtom(editPlugAtom)
	const actionMutation = api.plugs.action.edit.useMutation({
		onSuccess: result => editPlug(result)
	})
	const edit = useCallback(
		(...params: Parameters<typeof actionMutation.mutate>) => actionMutation.mutate(...params),
		[actionMutation]
	)

	if (!plug) return null

	return (
		<Accordion
			onExpand={() => {
				edit({
					id: plug.id,
					actions: JSON.stringify([
						...plug.actions,
						{
							protocol,
							action: actionName,
							id: Math.floor(Math.random() * 100_000_000_000),
							...getValues(action.schema[actionName].sentence)
						}
					])
				})
				frame()
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
