import { FC, HTMLAttributes, memo, useCallback, useEffect, useState } from "react"

import { motion } from "framer-motion"
import { Hash, X } from "lucide-react"

import { Image } from "@/components/app/utils/image"
import { Button } from "@/components/shared/buttons/button"
import { Accordion } from "@/components/shared/utils/accordion"
import { SchemasRequestAction, SchemasResponseCoils, cn, SchemasRequestValue, useConnect, useCordStateless } from "@/lib"
import { columnByIndexAtom, useColumnActions } from "@/state/columns"
import { editPlugAtom, plugByIdAtom, plugsAtom } from "@/state/plugs"
import { sentenceValidStateAtom } from "@/state/sentences"

import { HandleValueProps, Part } from "./part"
import { useAtom, useSetAtom } from "jotai"
import { api } from "@/server/client"

type SentenceProps = HTMLAttributes<HTMLDivElement> & {
	index: number
	item: string
	action: SchemasRequestAction
	actionIndex: number
	preview?: boolean
	error?: boolean
	linked?: SchemasRequestValue[]
	prevCoils?: SchemasResponseCoils
	dragging?: boolean
	handleValueChange?: (inputIndex: string, value: string, additionalData?: any) => void
	handleRemoveAction?: () => void
	validateType?: (coilName: string, expectedType: string) => boolean
	availableCoils?: Record<string, { type: string, actionIndex: number }>
}

export const Sentence: FC<SentenceProps> = memo(
	({
		index,
		item,
		action,
		actionIndex,
		preview = false,
		error = false,
		linked = [],
		prevCoils = {},
		dragging = false,
		handleValueChange,
		handleRemoveAction,
		validateType,
		availableCoils = {},
		className,
		...props
	}) => {
		const { account: { session } } = useConnect()

		const [column] = useAtom(columnByIndexAtom(index))
		const { frame } = useColumnActions(index)

		const [plug] = useAtom(plugByIdAtom(item))
		const own = plug && session && session.address === plug.socketId || false

		const editPlug = useSetAtom(editPlugAtom)
		const actionMutation = api.plugs.action.edit.useMutation({
			onSuccess: result => editPlug(result)
		})
		const edit = useCallback(
			(...params: Parameters<typeof actionMutation.mutate>) => actionMutation.mutate(...params),
			[actionMutation]
		)
		
		// Use atoms for sentence validation state and global plugs state
		const setSentenceValidState = useSetAtom(sentenceValidStateAtom)
		const setPlugs = useSetAtom(plugsAtom)

		const [search, setSearch] = useState<Record<number, string | undefined>>({})
		const [debouncedSearch, setDebouncedSearch] = useState<typeof search>({})
		const handleSearch = useCallback((s: string | null, index: number) => {
			const parsedIndex = parseInt(String(index))
			const newValue = s ?? undefined

			setSearch(prev => prev[parsedIndex] === newValue ? prev : {
				...prev,
				[parsedIndex]: newValue
			})

			const timeoutId = setTimeout(() => {
				setDebouncedSearch(prev => prev[parsedIndex] === newValue ? prev : {
					...prev,
					[parsedIndex]: newValue
				})
			}, 300)

			return () => clearTimeout(timeoutId)
		}, [])
		const { data: solverActions } = api.solver.actions.schemas.useQuery(
			{
				chainId: 8453,
				protocol: action?.protocol,
				action: action?.action,
				search: Object.entries(debouncedSearch ?? {}).map(([key, value]) => `search[${key}]=${value}`)
			},
			{ enabled: Boolean(action?.protocol ?? action?.action) ?? false, placeholderData: prev => prev }
		)
		const actionSchema = solverActions ? solverActions[action.protocol] : undefined
		const sentence = actionSchema ? actionSchema.schema[action.action].sentence : ""
		const options = actionSchema ? actionSchema.schema[action.action].options : undefined
		const coils = actionSchema ? actionSchema.schema[action.action].coils ?? {} : {}

		// Create a simpler values representation for cord
		const values = Object.entries(action.values ?? []).reduce(
			(acc, [key, value]) => {
				if (value && value.value !== undefined) {
					// Always store as string for cord
					acc[key] = String(value.value)
				}
				return acc
			},
			{} as Record<string, string>
		)

        // Single unified handler for all value updates
        const handleCordValueUpdate = useCallback((index: number, rawValue: string | undefined, error?: string) => {
            if (!plug) return;
            
            // Determine whether we need to convert the value (for numbers)
            const isNumber = typeof rawValue === 'string' && /^[0-9]+(\.[0-9]+)?$/.test(rawValue);
            const value = isNumber && rawValue ? parseFloat(rawValue) : rawValue;
            
            // Create a standardized value object that maintains consistency
            const updatedValue = {
                // Preserve any existing metadata
                ...action.values?.[index],
                // Always include these basics
                index: String(index),
                key: String(index),
                name: String(index),
                // Store the actual value
                value: value
            };
            
            // Create updated actions with the new value
            const updatedActions = plug.actions.map((action, nestedActionIndex) => {
                if (nestedActionIndex !== actionIndex) return action;
                
                return {
                    ...action,
                    values: {
                        ...action.values,
                        [index]: updatedValue
                    }
                };
            });
            
            // Update local state immediately for synchronization across columns
            setPlugs(prev => prev.map(p => 
                p.id === item 
                    ? { ...p, actions: JSON.stringify(updatedActions), updatedAt: new Date() } 
                    : p
            ));
            
            // Send to API
            edit({
                id: item,
                actions: JSON.stringify(updatedActions)
            });
        }, [plug, action, actionIndex, item, setPlugs, edit]);
        
        // Use the stateless cord hook that receives updates through our callback
        const {
            state: { parsed, parts },
            helpers: { getInputValue, getInputError, isValid, isComplete }
        } = useCordStateless(sentence, values, handleCordValueUpdate);
			
		// Update the global validation state whenever isValid or isComplete change
		useEffect(() => {
			if (actionSchema?.metadata?.chains) {
				setSentenceValidState(prev => ({
					...prev,
					[`${item}-${actionIndex}`]: { 
						isValid, 
						isComplete, 
						chains: actionSchema.metadata.chains.join(",") ?? "",
						actionPreview: item 
					}
				}))
			}
		}, [isValid, isComplete, actionSchema, actionIndex, item, setSentenceValidState])

		// Simplified handleValue that doesn't duplicate transformation logic
		const handleValue = ({ index, value, ...rest }: HandleValueProps) => {
			// Let handleCordValueUpdate deal with the typing and formatting
			handleCordValueUpdate(index, value);
		}

		// TODO: Need to make sure that we can properly evaluate a sentence when it is using a coil.
		//       Right now when we select a coil value the sentence is marked as invalid even though
		//       on refersh it is somehow now marked as valid and complete. Doesn't really make
		//       any sense to me how they work differently, but we need to find out.

		if (!column) return null

		if (!solverActions || !actionSchema)
			return (
				<motion.div
					className="bg-[length:200%_200%] w-full mb-2 h-16 animate-loading rounded-lg border-[1px] border-plug-green/10 bg-gradient-animated p-4"
					initial={{ y: 20 }}
					animate={{ y: 0 }}
				>
					<p className="hidden py-4 font-bold">.</p>
				</motion.div>
			)

		if (!parsed)
			return (
				<motion.div
					className="mb-2 rounded-lg border-[1px] border-plug-red p-4"
					initial={{ y: 20 }}
					animate={{ y: 0 }}
				>
					<p className="font-bold text-plug-red">
						Failed to parse: <span className="opacity-60">{sentence}</span>
					</p>
				</motion.div>
			)

		return (
			<div className="w-full">
				<Accordion
					className={cn(
						"cursor-default hover:bg-white",
						isValid && isComplete && !error
							? "border-plug-yellow hover:border-plug-yellow"
							: "border-plug-red hover:border-plug-red",
						className
					)}
					noPadding
					data-sentence
					data-chains={actionSchema?.metadata.chains.join(",") ?? ""}
					data-action-preview={item}
					{...props}
				>
					<div className={cn("flex flex-row items-center font-bold p-4")}>
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

									{solverActions &&
										parts.map((part, partIndex) => {
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
													coils={prevCoils}
													search={search}
													getInputValue={getInputValue}
													getInputError={getInputError}
													handleSearch={handleSearch}
													handleValue={handleValue}
													validateType={validateType}
													availableCoils={availableCoils}
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
								onClick={handleRemoveAction}
							>
								<X size={14} className="opacity-60" />
							</Button>
						)}
					</div>

					{coils && Object.keys(coils).length > 0 && <div className="border-t-[1px] border-plug-green/10 pt-2 text-sm px-4 pb-2">
						{Object.keys(coils).map((name, coilIndex) => (
							<p key={coilIndex} className="font-bold w-full flex flex-row gap-2 items-center">
								<Hash size={14} className="opacity-20" />
								{name}
								<span className="ml-auto opacity-40">{coils[name]}</span>

							</p>
						))}
					</div>}
				</Accordion>

				<div
					className={cn(
						"mx-auto h-2 w-[1px] transition-all duration-200 ease-in-out",
						isValid && isComplete && !error
							? "bg-plug-yellow hover:border-plug-yellow"
							: "bg-plug-red hover:border-plug-red",
						linked && linked.length > 0 && "bg-orange-300 hover:border-orange:300",
						!(!dragging && plug?.actions && actionIndex < plug?.actions.length - 1) && "bg-plug-white"
					)}
				/>
			</div>
		)
	}
)

Sentence.displayName = "Sentence"
