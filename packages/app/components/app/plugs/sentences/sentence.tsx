import { FC, HTMLAttributes, memo, useState } from "react"
import { motion } from "framer-motion"

import { Hash, SearchIcon, X } from "lucide-react"

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
import { usePlugStore } from "@/state/plugs"

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

type HandleValueProps = {
	index: number,
	value: string,
	name: string,
	isNumber?: boolean
} & Partial<Options[number]>

type SentenceProps = HTMLAttributes<HTMLButtonElement> & {
	index: number
	item: string
	preview?: boolean
	error?: boolean
	action: Action
	actionIndex: number,
}

const Part: FC<PartProps> = memo(({
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

export const Sentence: FC<SentenceProps> = memo(({
	index,
	item,
	action,
	actionIndex,
	preview = false,
	error = false,
	className,
	...props
}) => {
	const [search, setSearch] = useState<Record<number, string | undefined>>({})

	const { column, handle: { frame } } = useColumnStore(index)
	const {
		own,
		actions: plugActions,
		handle: {
			action: { edit }
		},
		solver: { actions: solverActions }
	} = usePlugStore(item, { protocol: action.protocol, action: action.action, search })

	const actionSchema = solverActions ? solverActions[action.protocol] : undefined
	const sentence = actionSchema ? actionSchema.schema[action.action].sentence : ""
	const options = actionSchema ? actionSchema.schema[action.action].options : undefined

	const values = Object.entries(action.values ?? []).reduce(
		(acc, [key, value]) => {
			if (value) {
				acc[key] = value.value
			}
			return acc
		},
		{} as Record<string, string>
	)

	const {
		state: { parsed },
		actions: { setValue },
		helpers: { getInputValue, getInputError, isValid, isComplete }
	} = useCord(sentence, values)

	const parts = parsed
		? parsed.template
			.split(/(\{[^}]+\})/g)
			.map(part => {
				if (part.match(/\{[^}]+\}/)) return [part]
				return part.split(/(\s+)/g)
			})
			.flat()
		: []

	const handleValue = ({ index, value, isNumber, ...rest }: HandleValueProps) => {
		setValue(index, value)

		edit({
			id: item,
			actions: JSON.stringify(
				plugActions.map((action, nestedActionIndex) => ({
					...action,
					values:
						nestedActionIndex === actionIndex
							? {
								...action.values,
								[index]: {
									...rest,
									value: isNumber ? parseFloat(value) : value,
								}
							}
							: action.values
				}))
			)
		})
	}

	if (!column) return null

	if (!solverActions || !actionSchema) return <motion.div
		className="h-16 mb-2 border-[1px] border-plug-green/10 rounded-lg p-4 animate-loading bg-gradient-animated bg-[length:200%_200%"
		initial={{ y: 20 }}
		animate={{ y: 0 }}
	>
		<p className="font-bold hidden py-4">.</p>
	</motion.div>

	if (!parsed) return <motion.div
		className="mb-2 border-[1px] border-plug-red rounded-lg p-4"
		initial={{ y: 20 }}
		animate={{ y: 0 }}
	>
		<p className="font-bold text-plug-red">Failed to parse: <span className="opacity-60">{sentence}</span></p>
	</motion.div>

	return (
		<motion.div initial={{ y: 20 }} animate={{ y: 0 }}>
			<Accordion
				className={cn(
					"cursor-default hover:bg-white",
					isValid && isComplete && !error
						? "border-plug-yellow hover:border-plug-yellow"
						: "border-plug-red hover:border-plug-red",
					className
				)}
				data-sentence
				data-chains={actionSchema?.metadata.chains.join(",") ?? ""}
				data-valid={isValid && isComplete}
				data-action-preview={item}
				{...props}
			>
				<div className={cn("flex flex-row items-center font-bold")}>
					<div className="flex w-full flex-wrap items-center gap-[4px]">
						<div className="flex flex-row items-start gap-[4px]">
							<div className="relative mt-1 h-6 w-10 flex-shrink-0">
								<Image
									className="absolute mr-2 h-6 w-6 rounded-sm blur-xl filter"
									src={actionSchema.metadata.icon}
									alt={`Icon for ${action.protocol}`}
									width={64}
									height={64}
								/>
								<Image
									className="absolute mr-2 h-6 w-6 rounded-sm"
									src={actionSchema.metadata.icon}
									alt={`Icon for ${action.protocol}`}
									width={64}
									height={64}
								/>
							</div>

							<div className="flex flex-wrap items-center gap-y-1">
								{!solverActions && <p>Failed to retrieve action schema: {action.protocol}</p>}

								{solverActions && parts.map((part, partIndex) => {
									const match = part.match(/\{(\d+)(?:=>(\d+))?\}/)

									if (!match)
										return (
											<span key={partIndex} className="whitespace-pre">
												{part}
											</span>
										)

									const inputIndex = parseInt(match[2] || match[1])
									const optionsIndex = match[2] ? parseInt(match[1]) : inputIndex
									const input = parsed.inputs.find(i => i.index === inputIndex)

									if (!input) return null

									return (
										<Part
											key={`${index}-${actionIndex}-${actionIndex}-${partIndex}`}
											index={index}
											column={column}
											frame={frame}
											own={own}
											preview={preview}
											error={error}
											actionIcon={actionSchema.metadata.icon}
											action={action}
											actionIndex={actionIndex}
											parsed={parsed}
											input={input}
											inputIndex={inputIndex}
											optionsIndex={optionsIndex}
											options={options}
											getInputValue={getInputValue}
											getInputError={getInputError}
											search={search}
											handleSearch={(s, index) => setSearch(prev => ({ ...prev, [parseInt(String(index))]: s ?? undefined }))}
											handleValue={handleValue}
										/>

									)
								})}
							</div>
						</div>
					</div>

					{preview === false && own && (
						<Button
							variant="secondary"
							className="mb-auto ml-4 mt-[4px] rounded-sm p-1"
							onClick={() =>
								edit({
									id: item,
									actions: JSON.stringify(plugActions.filter((_, i) => i !== actionIndex))
								})
							}
						>
							<X size={14} className="opacity-60" />
						</Button>
					)}
				</div>
			</Accordion>

			{actionIndex < plugActions.length - 1 && (
				<div
					className={cn(
						"mx-auto h-2 w-[2px]",
						isValid && isComplete && !error
							? "bg-plug-yellow hover:border-plug-yellow"
							: "bg-plug-red hover:border-plug-red"
					)}
				/>
			)}
		</motion.div>
	)
})

Sentence.displayName = "Sentence"
