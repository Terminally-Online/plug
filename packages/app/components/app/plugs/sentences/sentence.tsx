import { FC, HTMLAttributes } from "react"

import { Hash, X } from "lucide-react"

import { getInputPlaceholder } from "@terminallyonline/cord"

import { Frame } from "@/components/app/frames/base"
import { Checkbox } from "@/components/app/inputs/checkbox"
import { Search } from "@/components/app/inputs/search"
import { TokenImage } from "@/components/app/sockets/tokens/token-image"
import { Image } from "@/components/app/utils/image"
import { Button } from "@/components/shared/buttons/button"
import { Accordion } from "@/components/shared/utils/accordion"
import { Counter } from "@/components/shared/utils/counter"
import { Action, cn, formatTitle, Options, useCord } from "@/lib"
import { api } from "@/server/client"
import { useColumnStore } from "@/state/columns"
import { usePlugStore } from "@/state/plugs"

type SentenceProps = HTMLAttributes<HTMLButtonElement> & {
	index: number
	item: string
	preview?: boolean
	error?: boolean
	action: Action
	actionIndex: number
}

export const Sentence: FC<SentenceProps> = ({
	index,
	item,
	action,
	actionIndex,
	preview = false,
	error = false,
	className,
	...props
}) => {
	const {
		column,
		handle: { frame }
	} = useColumnStore(index)
	const {
		own,
		actions: plugActions,
		handle: {
			action: { edit }
		},
		solver: { actions: solverActions }
	} = usePlugStore(item, action)

	const actionSchema = solverActions ? solverActions[action.protocol] : undefined
	const sentence = actionSchema ? actionSchema.schema[action.action].sentence : ""

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
		helpers: { getInputName, getInputValue, getInputError, isValid, isComplete }
	} = useCord(sentence, values)

	const parts = parsed ? parsed.template.split(/(\{[^}]+\})/g) : []

	const handleValue = (index: number, value: string, isNumber?: boolean) => {
		const inputName = getInputName(index)

		if (!parsed || !inputName) return

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
									[index]: { value: isNumber ? parseInt(value) : value, name: inputName }
								}
							: action.values
				}))
			)
		})
	}

	if (!column || !solverActions || !actionSchema || !parsed) return null

	return (
		<>
			<Accordion
				className={cn(
					"cursor-default hover:bg-white",
					isValid && isComplete && !error
						? "border-plug-yellow hover:border-plug-yellow"
						: "border-plug-red hover:border-plug-red",
					className
				)}
				data-sentence
				data-chains={actionSchema.metadata.chains.join(",")}
				data-valid={isValid && isComplete}
				data-action-preview={item}
				{...props}
			>
				<div className={cn("flex flex-row items-center font-bold")}>
					<div className="flex w-full flex-wrap items-center gap-[4px]">
						<div className="relative h-6 w-10">
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

						<div className="flex flex-wrap items-center gap-1">
							{parts.map((part, partIndex) => {
								const match = part.match(/\{(\d+)(?:=>(\d+))?\}/)

								if (!match) return <span key={partIndex}>{part}</span>

								const inputIndex = parseInt(match[2] || match[1])
								const optionsIndex = match[2] ? parseInt(match[1]) : inputIndex
								const input = parsed.inputs.find(i => i.index === inputIndex)

								if (!input) return null

								const value = getInputValue(inputIndex)
								const inputError = getInputError(inputIndex)
								const dependentOnValue =
									(input.dependentOn !== undefined && getInputValue(input.dependentOn)?.value) ||
									undefined

								const sentenceOptions = solverActions[action.protocol].schema[action.action].options
								const options =
									sentenceOptions &&
									(Array.isArray(sentenceOptions[optionsIndex])
										? (sentenceOptions[optionsIndex] as Options)
										: sentenceOptions &&
											  typeof sentenceOptions?.[optionsIndex] === "object" &&
											  dependentOnValue
											? (sentenceOptions[optionsIndex] as Record<string, Options>)[
													dependentOnValue
												]
											: undefined)
								const isOptionBased = options !== undefined

								// NOTE: This is not the most performant way to do this, but for now it works.
								const option = Array.isArray(options)
									? options.find(option => option.value === value?.value)
									: undefined

								const isReady =
									(input.dependentOn !== undefined && getInputValue(input.dependentOn)?.value) ||
									input.dependentOn === undefined
								const isEmpty = !value?.value
								const isValid = !isEmpty && !inputError && !error

								return (
									<>
										<button
											className={cn(
												"rounded-sm bg-gradient-to-tr px-2 py-1 font-bold transition-all duration-200 ease-in-out",
												!isValid ? "text-plug-red" : "text-plug-green",
												own === true ? "cursor-pointer" : "cursor-default"
											)}
											style={{
												background: !isValid
													? "linear-gradient(to top right, rgba(255,0,0,0.1), rgba(255,0,0,0.1))"
													: `linear-gradient(to top right, rgba(56, 88, 66, 0.2), rgba(210, 243, 138, 0.2))`
											}}
											onClick={() => (own ? frame(`${actionIndex}-${inputIndex}`) : undefined)}
										>
											{(option && option.label) ||
												value?.value ||
												input.name
													?.replaceAll("_", " ")
													.replace(/([A-Z])/g, " $1")
													.toLowerCase() ||
												`Input #${input.index}`}
										</button>

										<Frame
											index={index}
											icon={
												<div className="relative h-10 min-w-10">
													<Image
														src={actionSchema.metadata.icon}
														alt={`Action ${actionIndex} icon`}
														width={64}
														height={64}
														className="absolute left-1/2 top-1/2 h-16 w-16 -translate-x-1/2 rounded-sm blur-2xl filter"
													/>
													<Image
														src={actionSchema.metadata.icon}
														alt={`Action ${actionIndex} icon`}
														width={64}
														height={64}
														className="relative left-1/2 top-1/2 h-8 w-8 -translate-x-1/2 -translate-y-1/2 rounded-sm"
													/>
												</div>
											}
											label={
												<span className="relative text-lg">
													<span className={"opacity-40"}>{formatTitle(action.action)}:</span>
													<span> {formatTitle(input.name ?? `Input #${inputIndex}`)}</span>
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
											<div className="flex flex-col gap-2 overflow-y-auto px-6">
												{!isReady && (
													<div className="mb-2 flex rounded-lg border-[1px] border-plug-green/10 p-4 py-4 text-center font-bold text-black/40">
														<p className="mx-auto max-w-[380px]">
															Please enter a value for{" "}
															{
																parsed.inputs.find(i => i.index === input.dependentOn)
																	?.name
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
															handleValue(
																input.index,
																data,
																input.type?.toString().includes("int")
															)
														}
														isNumber={input.type?.toString().includes("int")}
													/>
												)}

												{isReady && isOptionBased && (
													<>
														<div className="mb-4 flex w-full flex-col gap-2">
															{options.map((option, optionIndex) => (
																<div
																	key={`${index}-${actionIndex}-${optionIndex}`}
																	className="flex flex-row items-center gap-4"
																>
																	<Checkbox
																		checked={option.value === value?.value}
																		handleChange={() =>
																			handleValue(
																				input.index,
																				option.value === value?.value
																					? ""
																					: option.value
																			)
																		}
																	/>

																	<button
																		key={`${index}-${actionIndex}-${optionIndex}`}
																		className="group flex w-full flex-row items-center gap-4 truncate overflow-ellipsis whitespace-nowrap text-left font-bold"
																		onClick={() =>
																			handleValue(
																				input.index,
																				option.value === value?.value
																					? ""
																					: option.value
																			)
																		}
																	>
																		{option.icon && (
																			<div className="min-w-6">
																				<div className="flex items-center space-x-2">
																					{option.icon
																						.split("%7C")
																						.map(icon =>
																							decodeURIComponent(icon)
																						)
																						.map((icon, index) => (
																							<div
																								key={index}
																								className="flex items-center"
																							>
																								<TokenImage
																									logo={icon}
																									symbol={
																										option.label
																									}
																									size="xs"
																									blur={false}
																									className={cn(
																										index > 0
																											? "-ml-4"
																											: ""
																									)}
																								/>
																							</div>
																						))}
																				</div>
																			</div>
																		)}
																		<span className="truncate">{option.name}</span>
																		{option.info && (
																			<span className="ml-auto tabular-nums opacity-40">
																				<Counter count={option.info} />
																			</span>
																		)}
																	</button>
																</div>
															))}
														</div>
													</>
												)}
											</div>

											<div className="mt-auto bg-white">
												<div className="relative">
													{options && options.length > 0 && (
														<div className="pointer-events-none absolute -top-8 left-0 right-0 h-8 bg-gradient-to-b from-white/0 to-white" />
													)}
													<div className="mb-4 px-6">
														<Button
															variant={!isEmpty && !error ? "primary" : "primaryDisabled"}
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
										</Frame>
									</>
								)
							})}
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

			{preview === false && actionIndex < plugActions.length - 1 && (
				<div className="mx-auto h-2 w-[2px] bg-plug-green/5" />
			)}
		</>
	)
}
