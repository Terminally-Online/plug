import { FC, HTMLAttributes, useMemo } from "react"

import { Hash, X } from "lucide-react"

import { getInputPlaceholder, shouldRenderInput } from "@terminallyonline/cord"

import { Accordion, Button, Frame, Image, Search } from "@/components"
import { Action, cn, formatTitle } from "@/lib"
import { api } from "@/server/client"
import { useActions, useColumnStore, usePlugStore } from "@/state"

import { useCord } from "./useCord"

export const Sentence: FC<
	HTMLAttributes<HTMLButtonElement> & {
		index: number
		item: string
		action: Action
		actionIndex: number
		preview?: boolean
	}
> = ({ index, item, action, actionIndex, preview = false, className, ...props }) => {
	const {
		column,
		handle: { frame }
	} = useColumnStore(index)
	const { own, actions: plugActions, handle } = usePlugStore(item)
	// const [solverActions] = useActions()

	const { data: solverActions } = api.solver.actions.get.useQuery({
		protocol: action.protocol,
		action: action.action
	})
	const actionSchema = solverActions ? solverActions[action.protocol] : undefined

	// const sentence = useMemo(() => actionSchema.schema[action.action].sentence, [actionSchema, action.action])
	const sentence =
		"Transfer {0<amount:[(1.1)==721?1:uint256]>} {1<token:address=0x62180042606624f02d8a130da8a3171e9b33894d:uint256=721>} {2<id:[(1.1)>20?uint256:null]>}"

	const {
		state: { parsed },
		actions: { setValue },
		helpers: { getInputValue, getInputError }
	} = useCord(sentence)

	const parts = parsed ? parsed.template.split(/(\{[^}]+\})/g) : []

	if (!column || !actionSchema || !parsed) return null

	return (
		<>
			<Accordion
				className={cn("hover:cursor-auto hover:border-grayscale-100 hover:bg-white", className)}
				{...props}
			>
				<div className={cn("flex flex-row items-center font-bold")}>
					<p className="flex w-full flex-wrap items-center gap-[4px]">
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
								const match = part.match(/\{(\d+)\}/)

								if (!match) return <span key={partIndex}>{part}</span>

								const inputIndex = parseInt(match[1])
								const input = parsed.inputs.find(i => i.index === inputIndex)

								if (!input) return null

								const value = getInputValue(inputIndex)
								const error = getInputError(inputIndex)

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
											onClick={() => (own ? frame(`${actionIndex}-${inputIndex}`) : undefined)}
										>
											{value?.value || input.name || `Input #${input.index}`}
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
												<span className="relative">
													<span className="text-lg">
														<span className={cn(parsed.inputs.length > 1 && "opacity-40")}>
															{formatTitle(action.action)}
															{parsed.inputs.length > 1 && <span>:</span>}
														</span>
														{parsed.inputs.length > 1 && (
															<span>
																{" "}
																{formatTitle(input.name ?? `Input #${inputIndex}`)}
															</span>
														)}
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
											<div className="flex flex-col gap-2 overflow-y-auto px-6">
												<Search
													className="mb-4"
													icon={<Hash size={14} />}
													placeholder={getInputPlaceholder(input.type)}
													search={value?.value}
													handleSearch={data => setValue(input.index, data)}
												/>
											</div>

											<div className="mt-auto bg-white">
												<div className="relative">
													{/*
													{options && options.length > 0 && (
														<div className="pointer-events-none absolute -top-8 left-0 right-0 h-8 bg-gradient-to-b from-white/0 to-white" />
													)}
													*/}
													<div className="mb-4 px-6">
														<Button
															variant={!isEmpty && !error ? "primary" : "disabled"}
															className="w-full py-4"
															onClick={() =>
																frame(
																	inputIndex + 1 < parsed.inputs.length
																		? `${actionIndex}-${inputIndex + 1}`
																		: undefined
																)
															}
															disabled={isEmpty || error !== undefined}
														>
															{isEmpty || error
																? error?.message || "Enter value"
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
					</p>

					{preview === false && own && (
						<Button
							variant="secondary"
							className="mb-auto ml-4 mt-[4px] rounded-sm p-1"
							onClick={() =>
								handle.action.edit({
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
				<div className="mx-auto h-2 w-[2px] bg-grayscale-0" />
			)}
		</>
	)
}
