import { FC, HTMLAttributes, memo } from "react"

import { Hash, SearchIcon } from "lucide-react"

import { getInputPlaceholder, InputReference } from "@terminallyonline/cord"

import { Frame } from "@/components/app/frames/base"
import { Search } from "@/components/app/inputs/search"
import { TokenImage } from "@/components/app/sockets/tokens/token-image"
import { Image } from "@/components/app/utils/image"
import { Button } from "@/components/shared/buttons/button"
import { Accordion } from "@/components/shared/utils/accordion"
import { Counter } from "@/components/shared/utils/counter"
import { Action, cn, formatTitle, Options, useCord, useDebounce } from "@/lib"
import { useColumnStore } from "@/state/columns"

type PartProps = HTMLAttributes<HTMLButtonElement> & {
	index: number
	column: NonNullable<ReturnType<typeof useColumnStore>['column']>,
	frame: ReturnType<typeof useColumnStore>['handle']['frame'],
	own?: boolean
	preview?: boolean
	error?: boolean
	actionIcon: string
	action: Action
	actionIndex: number,
	parsed: NonNullable<ReturnType<typeof useCord>['state']['parsed']>,
	input: NonNullable<InputReference>,
	inputIndex: number,
	optionsIndex: number,
	options: Record<string, Options | Record<string, Options>> | undefined,
	search: Record<number, string | undefined>,
	getInputValue: ReturnType<typeof useCord>['helpers']['getInputValue'],
	getInputError: ReturnType<typeof useCord>['helpers']['getInputError'],
	handleSearch: (search: string, index: number) => void,
	handleValue: (args: HandleValueProps) => void
}

export type HandleValueProps = {
	index: number,
	value: string,
	name: string,
	isNumber?: boolean
} & Partial<Options[number]>

export const Part: FC<PartProps> = memo(({
	index,
	column,
	frame,
	own,
	preview,
	error,
	actionIcon,
	action,
	actionIndex,
	parsed,
	input,
	inputIndex,
	optionsIndex,
	options,
	search,
	getInputValue,
	getInputError,
	handleSearch,
	handleValue
}) => {
	const [searching, , handleDebounce] = useDebounce(
		search[optionsIndex] ?? "",
		250,
		debounced => handleSearch(debounced, input.index)
	)

	const value = getInputValue(inputIndex)
	const inputError = getInputError(inputIndex)
	const dependentOnValue =
		(input.dependentOn !== undefined && getInputValue(input.dependentOn)?.value) ||
		undefined

	const indexedOptions =
		options &&
		(Array.isArray(options[optionsIndex])
			? (options[optionsIndex] as Options)
			: options &&
				typeof options?.[optionsIndex] === "object" &&
				dependentOnValue
				? (options[optionsIndex] as Record<string, Options>)[
				dependentOnValue
				]
				: undefined)
	const isOptionBased = indexedOptions !== undefined
	const option = Array.isArray(indexedOptions)
		? indexedOptions.find(option => option.value === value?.value)
		: undefined

	const isReady =
		(input.dependentOn !== undefined && getInputValue(input.dependentOn)?.value) ||
		input.dependentOn === undefined
	const isEmpty = !value?.value
	const isValid = !isEmpty && !inputError && !error

	// NOTE: These are using saved option data from the database when it exists. For example,
	//       this means that if the user enters an ENS and they choose one, then when they refresh
	//       we will still have the 'nftchance.eth' as the label even before the refresh and the
	//       option values are retrieved from the Solver. This also means that if an option
	//       is no longer supported or shown in the list existing Plugs will still function as 
	//       expected and the user will have the ability to choose an up to date option in the
	//       future if they see fit. 
	// TODO: In some rare cases, we will have to pause plugs that are using a version of an 
	//       action that is not supported.
	const icon = (action.values?.[input.index]?.icon?.default) || (option && option.icon.default)
	const label = (option && option.label) ||
		value?.value ||
		(action.values?.[input.index]?.label) ||
		input.name
			?.replaceAll("_", " ")
			.replace(/([A-Z])/g, " $1")
			.toLowerCase() ||
		`Input #${input.index}`

	if (!column) return null

	return (
		<>
			<button
				className={cn(
					"rounded-sm mx-1 px-2 py-1 font-bold transition-all duration-200 ease-in-out flex flex-row items-center gap-2 text-black/60",
					isValid ? "bg-plug-yellow/60" : "bg-plug-red/60",
					own && !preview ? "cursor-pointer" : "cursor-default"
				)}
				style={{
					background: !isValid
						? "bg-plug-red"
						: "bg-plug-yellow"
				}}
				onClick={() =>
					own && !preview ? frame(`${actionIndex}-${inputIndex}`) : undefined
				}
			>
				{icon && (
					<div className="flex items-center space-x-2">
						{icon
							.split("%7C")
							.map(icon =>
								decodeURIComponent(
									icon
								)
							)
							.map(
								(
									icon,
									tokenIndex
								) => (
									<TokenImage
										key={
											tokenIndex
										}
										logo={icon}
										symbol={icon}
										className={cn(
											tokenIndex >
												0
												? "-ml-24"
												: ""
										)}
										size="xs"
									/>
								)
							)}
					</div>
				)}
				{label}
			</button>

			<Frame
				index={index}
				icon={
					<div className="relative h-10 min-w-10">
						<Image
							src={actionIcon}
							alt={`Action ${actionIndex} icon`}
							width={64}
							height={64}
							className="absolute left-1/2 top-1/2 h-16 w-16 -translate-x-1/2 rounded-sm blur-2xl filter"
						/>
						<Image
							src={actionIcon}
							alt={`Action ${actionIndex} icon`}
							width={64}
							height={64}
							className="relative left-1/2 top-1/2 h-8 w-8 -translate-x-1/2 -translate-y-1/2 rounded-sm"
						/>
					</div>
				}
				label={
					<span className="relative text-lg">
						<span className={"opacity-40"}>
							{formatTitle(action.action)}:
						</span>
						<span>
							{" "}
							{formatTitle(input.name ?? `Input #${inputIndex}`)}
						</span>
					</span>
				}
				visible={column.frame === `${actionIndex}-${inputIndex}`}
				handleBack={
					inputIndex > 0
						? () => frame(`${actionIndex}-${inputIndex - 1}`)
						: undefined
				}
				hasOverlay
				hasChildrenPadding={false}
				scrollBehavior="partial"
			>
				{column.frame === `${actionIndex}-${inputIndex}` ? (
					<>
						<div className="flex flex-col gap-2 overflow-y-auto px-6">
							{!isReady && (
								<div className="mb-2 flex rounded-lg border-[1px] border-plug-green/10 p-4 py-4 text-center font-bold text-black/40">
									<p className="mx-auto max-w-[380px]">
										Please enter a value for{" "}
										{
											parsed.inputs.find(
												i => i.index === input.dependentOn
											)?.name
										}{" "}
										before continuing.
									</p>
								</div>
							)}

							{isReady && !isOptionBased && (
								<Search
									className="mb-4"
									icon={<Hash size={14} />}
									placeholder={getInputPlaceholder(input.type)}
									search={value?.value}
									handleSearch={data =>
										handleValue({
											index: input?.index ?? "",
											name: input?.name ?? "",
											label,
											value: data,
											isNumber: input.type?.toString().includes("int")
										})
									}
									isNumber={
										input.type?.toString().includes("int") ||
										input.type?.toString().includes("float")
									}
									focus={true}
								/>
							)}

							{isReady && isOptionBased && (
								<>
									<Search
										icon={<SearchIcon size={14} />}
										placeholder="Search options"
										search={searching ?? ""}
										handleSearch={handleDebounce}
										focus
										clear
									/>

									<div className="mb-4 flex w-full flex-col gap-2">
										{indexedOptions.map((option, optionIndex) => (
											<Accordion
												key={`${index}-${actionIndex}-${optionIndex}`}
												onExpand={() =>
													handleValue({
														...option,
														index: input.index,
														// NOTE: Support toggling of the option by clicking it again.
														value: option.value === value?.value
															? ""
															: option.value
													})
												}
												className="relative"
											>
												{option.value === value?.value && (
													<div className="absolute bottom-0 right-0 h-24 w-24 bg-plug-yellow blur-[80px] filter" />
												)}

												<div className="flex flex-row items-center gap-4">
													{option.icon.default && (
														<div className="flex items-center space-x-2">
															{option.icon.default
																.split("%7C")
																.map(icon => decodeURIComponent(icon))
																.map(
																	(
																		icon,
																		tokenIndex
																	) => (
																		<TokenImage
																			key={tokenIndex}
																			logo={icon}
																			symbol={option.label}
																			className={cn(
																				tokenIndex >
																					0
																					? "-ml-24"
																					: ""
																			)}
																		/>
																	)
																)}
														</div>
													)}
													<div className="flex w-full flex-col">
														<p className="flex w-full flex-row justify-between gap-2 truncate">
															{option.name}
															{option.info && (
																<span className="ml-auto tabular-nums">
																	<Counter
																		count={
																			option.info
																				.value
																		}
																	/>
																</span>
															)}
														</p>
														<p className="whitespace-nowrap flex flex-row items-center justify-between gap-2 text-sm tabular-nums text-black/40">
															{option.icon.secondary && (
																<Image className="rounded-[4px] w-4 h-4" src={option.icon.secondary} alt="secondary option icon" width={32} height={32} />
															)}
															{option.label}
															{option.info && (
																<Counter className="ml-auto tabular-nums" count={option.info.label} />
															)}
														</p>
													</div>
												</div>
											</Accordion>
										))}
									</div>
								</>
							)}
						</div>
						<div className="mt-auto bg-white">
							<div className="relative">
								{indexedOptions && indexedOptions.length > 0 && (
									<div className="pointer-events-none absolute -top-8 left-0 right-0 h-8 bg-gradient-to-b from-white/0 to-white" />
								)}
								<div className="mb-4 px-6">
									<Button
										variant={
											!isEmpty && !error
												? "primary"
												: "primaryDisabled"
										}
										className="w-full py-4"
										onClick={() =>
											frame(
												inputIndex + 1 < parsed.inputs.length
													? `${actionIndex}-${inputIndex + 1}`
													: undefined
											)
										}
										disabled={isEmpty || error}
									>
										{isOptionBased && isEmpty
											? "Choose option"
											: isEmpty || error
												? inputError?.message || "Enter value"
												: parsed.inputs.length - 1 > inputIndex
													? "Next"
													: "Done"}
									</Button>
								</div>
							</div>
						</div>
					</>
				) : (
					<></>
				)}
			</Frame>
		</>
	)
})

Part.displayName = "SentencePart"
