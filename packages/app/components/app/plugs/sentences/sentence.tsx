import { FC, HTMLAttributes, useCallback, useEffect, useState } from "react"

import { motion } from "framer-motion"
import { Hash, X } from "lucide-react"

import { useAtom, useSetAtom } from "jotai"

import { Image } from "@/components/app/utils/image"
import { Button } from "@/components/shared/buttons/button"
import { Accordion } from "@/components/shared/utils/accordion"
import {
	cn,
	SchemasRequestAction,
	SchemasRequestValue,
	SchemasResponseCoils,
	useCordStateless
} from "@/lib"
import { api } from "@/server/client"
import { useActions } from "@/state/actions"
import { columnByIndexAtom, useColumnActions } from "@/state/columns"
import { editPlugAtom, plugByIdAtom, plugsAtom } from "@/state/plugs"
import { sentenceValidStateAtom } from "@/state/sentences"

import { HandleValueProps, Part } from "./part"
import { useAccount } from "@/lib/hooks/account/useAccount"

type SentenceProps = HTMLAttributes<HTMLDivElement> & {
	index: number
	item?: string
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
	availableCoils?: Record<string, { type: string; actionIndex: number }>
}

export const Sentence: FC<SentenceProps> = ({
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
	const { id } = useAccount()

	const [column] = useAtom(columnByIndexAtom(index))
	const { frame } = useColumnActions(index)

	const [plug] = useAtom(plugByIdAtom(item ?? ""))
	const own = (plug && id === plug.socketId) || false

	const editPlug = useSetAtom(editPlugAtom)
	const actionMutation = api.plugs.action.edit.useMutation({
		onSuccess: result => editPlug(result)
	})
	const edit = useCallback(
		(...params: Parameters<typeof actionMutation.mutate>) => actionMutation.mutate(...params),
		[actionMutation]
	)

	const setSentenceValidState = useSetAtom(sentenceValidStateAtom)
	const setPlugs = useSetAtom(plugsAtom)
	const handleCordValueUpdate = useCallback(
		(index: number, rawValue: string | undefined, _?: string) => {
			if (!plug) return

			const isNumber = typeof rawValue === "string" && /^[0-9]+(\.[0-9]+)?$/.test(rawValue)
			const value = isNumber && rawValue ? parseFloat(rawValue) : rawValue
			const updatedValue = {
				// Preserve any existing metadata
				...action.values?.[index],
				// Always include these basics
				index: String(index),
				key: String(index),
				name: String(index),
				// Store the actual value
				value: value
			}
			const updatedActions = plug.actions.map((action, nestedActionIndex) => {
				if (nestedActionIndex !== actionIndex) return action

				return {
					...action,
					values: {
						...action.values,
						[index]: updatedValue
					}
				}
			})

			setPlugs(prev =>
				prev.map(p =>
					p.id === plug.id ? { ...p, actions: JSON.stringify(updatedActions), updatedAt: new Date() } : p
				)
			)

			edit({
				id: item,
				actions: JSON.stringify(updatedActions)
			})
		},
		[plug, action, actionIndex, item, setPlugs, edit]
	)

	const [minimalActions] = useActions()
	const sentence = minimalActions[action.protocol]?.schema[action.action]?.sentence
	const coils = minimalActions[action.protocol]?.schema[action.action]?.coils ?? {}
	const values = Object.entries(action.values ?? []).reduce(
		(acc, [key, value]) => {
			if (value && value.value !== undefined) {
				acc[key] = String(value.value)
			}
			return acc
		},
		{} as Record<string, string>
	)

	const {
		state: { parsed, parts },
		helpers: { getInputValue, getInputError, isValid, isComplete }
	} = useCordStateless(sentence, values, handleCordValueUpdate)

	const [search, setSearch] = useState<Record<number, string | undefined>>({})
	const [debouncedSearch, setDebouncedSearch] = useState<typeof search>({})
	const handleSearch = useCallback((s: string | null, index: number) => {
		const parsedIndex = parseInt(String(index))
		const newValue = s ?? undefined

		setSearch(prev =>
			prev[parsedIndex] === newValue
				? prev
				: {
					...prev,
					[parsedIndex]: newValue
				}
		)

		const timeoutId = setTimeout(() => {
			setDebouncedSearch(prev =>
				prev[parsedIndex] === newValue
					? prev
					: {
						...prev,
						[parsedIndex]: newValue
					}
			)
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
	const options = actionSchema ? actionSchema.schema[action.action].options : undefined

	useEffect(() => {
		if(!item) return

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
	}, [item, isValid, isComplete, actionSchema, actionIndex, plug, setSentenceValidState])

	const handleValue = ({ index, value }: HandleValueProps) => {
		handleCordValueUpdate(index, value)
	}

	if (!column) return null

	if (!minimalActions && !solverActions)
		return (
			<motion.div
				className="mb-2 h-16 w-full animate-loading rounded-lg border-[1px] border-plug-green/10 bg-gradient-animated bg-[length:200%_200%] p-4"
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
				{...props}
			>
				<div className={cn("flex flex-row items-center p-4 font-bold")}>
					<div className="flex w-full flex-wrap items-center gap-[4px]">
						<div className="flex flex-row items-start gap-[4px]">
							<div className="relative mt-1 h-6 w-10 flex-shrink-0">
								<Image
									className="absolute mr-2 h-6 w-6 rounded-sm blur-xl filter"
									src={minimalActions[action.protocol].metadata.icon}
									alt={`Icon for ${action.protocol}`}
									width={64}
									height={64}
								/>
								<Image
									className="absolute mr-2 h-6 w-6 rounded-sm"
									src={minimalActions[action.protocol].metadata.icon}
									alt={`Icon for ${action.protocol}`}
									width={64}
									height={64}
								/>
							</div>

							<div className="flex flex-wrap items-center gap-y-1">
								{parts.map((part, partIndex) => {
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
											actionIcon={minimalActions[action.protocol].metadata.icon}
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

				{coils && Object.keys(coils).length > 0 && (
					<div className="border-t-[1px] border-plug-green/10 px-4 pb-2 pt-2 text-sm">
						{Object.keys(coils).map((name, coilIndex) => (
							<p key={coilIndex} className="flex w-full flex-row items-center gap-2 font-bold">
								<Hash size={14} className="opacity-20" />
								{name}
								<span className="ml-auto opacity-40">{coils[name]}</span>
							</p>
						))}
					</div>
				)}
			</Accordion>

			<div
				className={cn(
					"mx-auto h-2 w-[1px] transition-all duration-200 ease-in-out",
					isValid && isComplete && !error
						? "bg-plug-yellow hover:border-plug-yellow"
						: "bg-plug-red hover:border-plug-red",
					!(!dragging && plug?.actions && actionIndex < plug?.actions.length - 1) && "bg-plug-white"
				)}
			/>
		</div>
	)
}
