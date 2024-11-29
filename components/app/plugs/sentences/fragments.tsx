import Image from "next/image"
import { FC } from "react"

import { Hash } from "lucide-react"

import { getInputPlaceholder, InputState, ParsedCordSentence } from "@terminallyonline/cord"

import { Frame, Search} from "@/components"
import { cn, formatTitle } from "@/lib"
import { ACTION_REGEX, useColumnStore } from "@/state"

import { ValidationError } from "./useCord"

type FragmentProps = {
	index: number
	actionIndex: number
	item: string
	icon: string
	// action: Action
	preview: boolean
	own: boolean
	parsed: ParsedCordSentence | null
	setValue: (index: number, value: string) => void
	getInputValue: (index: number) => InputState | undefined
	getInputError: (index: number) => ValidationError | undefined
}

// export const Fragments: FC<FragmentProps> = ({
// 	index,
// 	actionIndex,
// 	item,
// 	icon,
// 	preview,
// 	own,
// 	parsed,
// 	setValue,
// 	getInputValue,
// 	getInputError
// }) => {
// 	const { column, handle } = useColumnStore(index)
//
// 	if (!parsed || !column) return null

	// const fragments = useMemo(() => {
	// 	return sentence.split(ACTION_REGEX) as string[]
	// }, [sentence])
	//
	// const dynamic = useMemo(() => {
	// 	return fragments.filter(fragment => fragment.match(ACTION_REGEX))
	// }, [fragments])
	//
	// let dynamicIndex = -1

	// return (
	// 	<>
	// 		{fragments.map((fragment, fragmentIndex) => {
	// 			if (fragment.match(ACTION_REGEX)) {
	// 				dynamicIndex++
	// 				return (
	// 					<DynamicFragment
	// 						key={`${actionIndex}-${fragmentIndex}`}
	// 						item={item}
	// 						index={index}
	// 						actionIndex={actionIndex}
	// 						fragmentIndex={fragmentIndex}
	// 						dynamicIndex={dynamicIndex}
	// 						fragment={fragment}
	// 						// protocol={protocol}
	// 						// action={action}
	// 						// dynamic={dynamic}
	// 						// preview={preview}
	// 					/>
	// 				)
	// 			}
	// 			return <StaticFragment key={`${actionIndex}-${fragmentIndex}`} fragment={fragment} />
	// 		})}
	// 	</>
	// )
// }
