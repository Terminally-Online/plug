import { FC, HTMLAttributes, memo, useCallback, useMemo } from "react"

import { Hash, SearchIcon } from "lucide-react"

import { useAtom } from "jotai"

import { getInputPlaceholder, InputReference } from "@terminallyonline/cord"

import { Frame } from "@/components/app/frames/base"
import { Search } from "@/components/app/inputs/search"
import { TokenImage } from "@/components/app/sockets/tokens/token-image"
import { Image } from "@/components/app/utils/image"
import { Button } from "@/components/shared/buttons/button"
import { Accordion } from "@/components/shared/utils/accordion"
import { Counter } from "@/components/shared/utils/counter"
import { SchemasRequestAction, SchemasResponseCoils, cn, formatTitle, SchemasRequestValues, useCordStateless, useDebounce, SchemasResponseOptionsSet, SchemasRequestValuesSet } from "@/lib"
import { columnByIndexAtom, useColumnActions } from "@/state/columns"

type PartProps = HTMLAttributes<HTMLButtonElement> & {
	index: number
	frame: ReturnType<typeof useColumnActions>["frame"]
	own?: boolean
	preview?: boolean
	error?: boolean
	actionIcon: string
	action: SchemasRequestAction
	actionIndex: number
	parsed: NonNullable<ReturnType<typeof useCordStateless>["state"]["parsed"]>
	input: NonNullable<InputReference>
	inputIndex: number
	optionsIndex: number
	options?: SchemasResponseOptionsSet
	coils?: SchemasResponseCoils,
	search: Record<number, string | undefined>
	getInputValue: ReturnType<typeof useCordStateless>["helpers"]["getInputValue"]
	getInputError: ReturnType<typeof useCordStateless>["helpers"]["getInputError"]
	handleSearch: (search: string, index: number) => void
	handleValue: (args: HandleValueProps) => void
	validateType?: (coilName: string, expectedType: string) => boolean
	availableCoils?: Record<string, { type: string, actionIndex: number }>
}

export type HandleValueProps = {
	index: number
	key: string
	value?: string
	name: string
	isNumber?: boolean
} & Partial<SchemasRequestValues[number]>

export const Part: FC<PartProps> = memo(
	({
		index,
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
		coils,
		search,
		getInputValue,
		getInputError,
		handleSearch,
		handleValue,
		validateType,
		availableCoils
	}) => {
		const [searching, , handleDebounce] = useDebounce(search[optionsIndex] ?? "", 250, debounced =>
			handleSearch(debounced, input.index)
		)

		const [column] = useAtom(columnByIndexAtom(index))

		const value = getInputValue(inputIndex)
		const inputError = getInputError(inputIndex)
		const dependentOnValue =
			(input.dependentOn !== undefined && getInputValue(input.dependentOn)?.value) || undefined

		const validateLinkedInput = useCallback((linkedValue: string | undefined, expectedType: string | undefined): boolean => {
			if (!linkedValue || !validateType || !availableCoils || !expectedType) return false

			const match = linkedValue.match(/^<-\{(.+)\}$/)
			if (!match) return false

			const coilName = match[1]
			return validateType(coilName, expectedType)
		}, [availableCoils, validateType])

		const indexedOptions =
			options &&
			(Array.isArray(options[optionsIndex])
				? (options[optionsIndex] as SchemasRequestValues)
				: options && typeof options?.[optionsIndex] === "object" && dependentOnValue
					? (options[optionsIndex] as SchemasRequestValuesSet)[parseInt(dependentOnValue)]
					: undefined)
		const isOptionBased = indexedOptions !== undefined
		const option = Array.isArray(indexedOptions)
			? indexedOptions.find(option => option.value === value?.value)
			: undefined

		const isReady =
			(input.dependentOn !== undefined && getInputValue(input.dependentOn)?.value) ||
			input.dependentOn === undefined

		const validCoils = useMemo(() => {
			if (!coils) return {}

			return Object.fromEntries(Object.keys(coils).filter(name =>
				validateLinkedInput(`<-{${name}}`, coils[name])).map(key => [key, coils[key]])
			)
		}, [coils, input, validateLinkedInput])

		const isLinked = typeof value?.value === "string" && value?.value?.startsWith("<-{") && value?.value?.endsWith("}")
		const isCompatibleCoil = useMemo(() =>
			typeof value?.value === "string" && isLinked && Object.keys(validCoils).includes(value?.value.replace("<-{", "").replace("}", "")),
			[isLinked, value, validCoils]
		)
		const isEmpty = !value?.value || (isLinked && !isCompatibleCoil)
		const isValid = !isEmpty && !inputError && !error

		// Use the data directly from our centralized store
		// This removes the need to separate value.value vs action.values[index]
		const icon = (option && option?.icon?.default) || action.values?.[input.index]?.icon?.default
		const label =
			(option && option.label) ||
			((isValid || (isLinked && isCompatibleCoil)) && value?.value) ||
			(!isLinked && (isLinked && isCompatibleCoil) && action.values?.[input.index]?.label) ||
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
						"mx-1 flex flex-row items-center gap-2 rounded-sm px-2 py-1 font-bold text-black/60 transition-all duration-200 ease-in-out",
						isCompatibleCoil
							? "bg-orange-300/60"
							: isValid
								? "bg-plug-yellow/60"
								: "bg-plug-red/60",
						own && !preview ? "cursor-pointer" : "cursor-default"
					)}
					style={{
						background: isLinked && isCompatibleCoil
							? "bg-orange-300/60"
							: !isValid ? "bg-plug-red" : "bg-plug-yellow"
					}}
					onClick={() => (own && !preview ? frame(`${actionIndex}-${inputIndex}`) : undefined)}
				>
					{icon && (
						<div className="flex items-center space-x-2">
							{icon
								.split("%7C")
								.map(icon => decodeURIComponent(icon))
								.map((icon, tokenIndex) => (
									<TokenImage
										key={tokenIndex}
										logo={icon}
										symbol={icon}
										className={cn(tokenIndex > 0 ? "-ml-24" : "")}
										size="xs"
									/>
								))}
						</div>
					)}
					<span className="max-w-[150px] overflow-hidden truncate text-ellipsis">
						{isLinked && isCompatibleCoil && label.startsWith("<-{")
							? label.replace("<-{", "").replace("}", "")
							: label}
					</span>
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
							<span className={"opacity-40"}>{formatTitle(action.action)}:</span>
							<span> {formatTitle(input.name ?? `Input #${inputIndex}`)}</span>
						</span>
					}
					visible={column.frame === `${actionIndex}-${inputIndex}`}
					handleBack={inputIndex > 0 ? () => frame(`${actionIndex}-${inputIndex - 1}`) : undefined}
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
											{parsed.inputs.find(i => i.index === input.dependentOn)?.name} before
											continuing.
										</p>
									</div>
								)}

								{isReady && !isOptionBased && Object.keys(validCoils).length > 0 && (
									<>
										<div className="relative flex flex-row flex-wrap gap-2 overflow-hidden">
											{Object.keys(validCoils).map((coil, index) => {
												if (value?.value === `<-{${coil}}`) return null

												return (
													<Button
														key={index}
														variant="secondary"
														sizing="sm"
														onClick={() =>
															handleValue({
																index: input?.index ?? "",
																key: input?.name ?? "",
																name: input?.name ?? "",
																label: label,
																value: `<-{${coil}}`,
																isNumber: input.type?.toString().includes("int")
															})
														}
														className="group/coil flex flex-row gap-2 px-2 pr-3"
													>
														<div className="h-4 w-4 rounded-[4px] bg-orange-300 group-hover/coil:bg-orange-400 transition-all duration-200 ease-in-out flex items-center justify-center">
															<p className="text-xs font-bold text-plug-white">#</p>
														</div>

														{coil}
													</Button>
												)
											})}
										</div>

										{isLinked ? (
											<>
												<button
													className={cn(
														`mb-4 flex w-full cursor-pointer items-center gap-4 rounded-[16px] border-[1px] p-4 px-6 transition-colors duration-200 ease-in-out`,
														isCompatibleCoil ? "border-plug-green/10" : "border-plug-red/10"
													)}
													onClick={() =>
														handleValue({
															index: input?.index ?? "",
															key: input?.name ?? "",
															name: input?.name ?? "",
															label: label,
															value: undefined,
															isNumber: input.type?.toString().includes("int")
														})
													}
												>
													<div className={`flex h-4 w-4 items-center justify-center rounded-[4px] ${isCompatibleCoil ? "bg-orange-300" : "bg-plug-red"
														}`}>
														<p className="text-xs font-bold text-plug-white">
															{isCompatibleCoil ? "#" : "!"}
														</p>
													</div>
													<span className={isCompatibleCoil ? "" : "text-red-600 font-semibold"}>
														{typeof value?.value === "string" && value?.value.startsWith("<-{")
															? isCompatibleCoil
																? formatTitle(value?.value.replace("<-{", "").replace("}", ""))
																: "Invalid link: Coil not available in this position"
															: getInputPlaceholder(input.type)}
													</span>
												</button>
											</>
										) : (
											<Search
												className="mb-4"
												icon={<Hash size={14} />}
												placeholder={
													typeof value?.value === "string" && value?.value.startsWith("<-")
														? "Amount: number"
														: getInputPlaceholder(input.type)
												}
												search={value?.value ? String(value?.value) : ""}
												handleSearch={data =>
													handleValue({
														index: input?.index ?? "",
														key: input?.name ?? "",
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
												focus
											/>
										)}
									</>
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
															key: input?.name ?? "",
															// NOTE: Support toggling of the option by clicking it again.
															value: option.value === value?.value ? "" : option.value
														})
													}
													className="relative"
												>
													{option.value === value?.value && (
														<div className="absolute bottom-0 right-0 h-24 w-24 bg-plug-yellow blur-[80px] filter" />
													)}

													<div className="flex flex-row items-center gap-4">
														{option?.icon?.default && (
															<div className="flex items-center space-x-2">
																{option.icon.default
																	.split("%7C")
																	.map(icon => decodeURIComponent(icon))
																	.map((icon, tokenIndex) => (
																		<TokenImage
																			key={tokenIndex}
																			logo={icon}
																			symbol={option.label}
																			className={cn(
																				tokenIndex > 0 ? "-ml-24" : ""
																			)}
																		/>
																	))}
															</div>
														)}
														<div className="flex w-full flex-col">
															<p className="flex w-full flex-row justify-between gap-2 truncate">
																{option.name}
																{option.info && (
																	<span className="ml-auto tabular-nums">
																		<Counter count={option.info.value} />
																	</span>
																)}
															</p>
															<p className="flex flex-row items-center justify-between gap-2 whitespace-nowrap text-sm tabular-nums text-black/40">
																{option?.icon?.secondary && (
																	<Image
																		className="h-4 w-4 rounded-[4px]"
																		src={option.icon.secondary}
																		alt="secondary option icon"
																		width={32}
																		height={32}
																	/>
																)}
																{option.label}
																{option.info && (
																	<Counter
																		className="ml-auto tabular-nums"
																		count={option.info.label}
																	/>
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
						</>
					) : (
						<></>
					)}
				</Frame>
			</>
		)
	}
)

Part.displayName = "SentencePart"
