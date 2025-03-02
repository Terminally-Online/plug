import { FC, useMemo } from "react"

import { Sentence } from "@/components/app/plugs/sentences/sentence"
import { Actions } from "@/lib"
import { plugByIdAtom } from "@/state/plugs"
import { useAtom } from "jotai"

export const ActionPreview: FC<{ index: number; item: string; actions?: Actions; errors?: Array<string | null> }> = ({
	index,
	item,
	actions,
	errors = []
}) => {
	const [plug] = useAtom(plugByIdAtom(item))
	const plugActions: Actions = useMemo(() => {
		if (!plug) return []
		try {
			return JSON.parse(plug.actions)
		} catch {
			return []
		}
	}, [plug])

	actions = actions ?? plugActions

	return (
		<div className="flex flex-col gap-2">
			{actions.map((action, actionIndex) => (
				<div key={`${index}-${actionIndex}`}>
					<Sentence
						index={index}
						item={item}
						actionIndex={actionIndex}
						action={action}
						preview
						error={errors && errors[actionIndex] ? true : false}
					/>

					{errors && errors[actionIndex] && (
						<p className="text-sm font-bold text-plug-red">Error: {errors[actionIndex]}</p>
					)}
				</div>
			))}
		</div>
	)
}
