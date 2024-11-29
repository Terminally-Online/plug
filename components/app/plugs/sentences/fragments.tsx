import { FC } from "react"

import { DynamicFragment, Frame, Search, StaticFragment } from "@/components"
import { ACTION_REGEX, useColumnStore } from "@/state"
import { ValidationError } from "./useCord"
import { getInputPlaceholder, InputState, ParsedCordSentence } from "@terminallyonline/cord"
import { cn, formatTitle } from "@/lib"
import Image from "next/image"
import { Hash } from "lucide-react"

type FragmentProps = {
	index: number
	actionIndex: number
	item: string
	icon: string
	// action: Action
	preview: boolean
	own: boolean
	parsed: ParsedCordSentence | null;
	setValue: (index: number, value: string) => void;
	getInputValue: (index: number) => InputState | undefined;
	getInputError: (index: number) => ValidationError | undefined;
}

export const Fragments: FC<FragmentProps> = ({
	index,
	actionIndex,
	item,
	icon,
	preview,
	own,
	parsed,
	setValue,
	getInputValue,
	getInputError
}) => {
	const { column, handle } = useColumnStore(index)

	if (!parsed || !column) return null

	return (
		<div className="flex flex-row">
			{parsed.inputs.map((input, inputIndex) => {
				const value = getInputValue(input.index)
				const error = getInputError(input.index)
				const isEmpty = !value?.value.trim()
				const isValid = !isEmpty && !error

				return (
					<>
						<button
							className={cn(
								"rounded-sm bg-gradient-to-tr px-2 py-1 font-bold transition-all duration-200 ease-in-out",
								preview && !isValid ? "text-plug-red" : "text-plug-green",
								own === true ? "cursor-pointer" : "cursor-default"
							)}
							style={{
								background:
									preview && !isValid
										? "linear-gradient(to right, rgba(255,0,0,0.1), rgba(255,0,0,0.1))"
										: `linear-gradient(to right, rgba(0,239,54,0.1), rgba(147,223,0,0.1))`
							}}
							onClick={() => (own ? handle.frame(`${actionIndex}-${inputIndex}`) : undefined)}
						>
							{value?.value ?? input.name ?? `Input #${input.index}`}
						</button>

						<Frame
							index={index}
							icon={
								<div className="relative h-10 min-w-10">
									<Image
										src={icon}
										alt={`Action ${actionIndex} icon`}
										width={64}
										height={64}
										className="absolute left-1/2 top-1/2 h-16 w-16 -translate-x-1/2 rounded-sm blur-2xl filter"
									/>
									<Image
										src={icon}
										alt={`Action ${actionIndex} icon`}
										width={64}
										height={64}
										className="relative left-1/2 top-1/2 h-8 w-8 -translate-x-1/2 -translate-y-1/2 rounded-sm"
									/>
								</div>
							}
							label={
								<span className="relative">
									<span className="text-lg">
										{/*
										<span className={cn(parsed.inputs.length > 1 && "opacity-40")}>
											{formatTitle(action.action)}
											{parsed.inputs.length > 1 && <span>:</span>}
										</span>
										*/}
										{parsed.inputs.length > 1 && <span> {formatTitle(input.name ?? `Input #${inputIndex}`)}</span>}
									</span>
								</span>
							}
							visible={column.frame === `${actionIndex}-${inputIndex}`}
							handleBack={inputIndex > 0 ? () => handle.frame(`${actionIndex}-${inputIndex - 1}`) : undefined}
							hasOverlay
							hasChildrenPadding={false}
							scrollBehavior="partial"
						>
							<div className="flex flex-col gap-2 overflow-y-auto px-6">
								<Search
									className="mb-4"
									icon={<Hash size={14} />}
									placeholder={getInputPlaceholder(input.type)}
									search={value?.value}
									handleSearch={data => setValue(input.index, data)}
								/>
							</div>
						</Frame>
					</>
				)
			})}
		</div>
	)

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
}
