import { useSession } from "next-auth/react"
import { FC, useMemo } from "react"

import { Hash } from "lucide-react"

import { Button, Checkbox, Counter, Frame, Image, Search, TokenImage } from "@/components"
import { Action, cn, formatInputName, formatTitle, getIndexes, Value } from "@/lib"
import { api } from "@/server/client"
import { useActions, useColumnStore, usePlugStore } from "@/state"

export const DynamicFragment: FC<{
	item: string
	index: number
	actionIndex: number
	fragmentIndex: number
	dynamicIndex: number
	action: Action
	fragment: string
	dynamic: string[]
	preview: boolean
}> = ({ index, item, action, actionIndex, dynamicIndex, fragment, dynamic, preview }) => {
	const { data: session } = useSession()
	const { isFrame, handle } = useColumnStore(index, `${actionIndex}-${dynamicIndex}`)
	const { plug, actions, handle: plugHandle } = usePlugStore(item)

	const { data: solverActions } = api.solver.actions.get.useQuery({
		protocol: action.protocol,
		action: action.action
	})

	const protocol = solverActions?.[action.protocol]
	const staticAction = protocol?.schema[action.action]

	const own = plug && session && session.address === plug.socketId

	const [childIndex, parentIndex] = useMemo(() => getIndexes(fragment), [fragment])

	const inputName = formatInputName(staticAction.fields[parentIndex]?.name)

	const label = useMemo(() => {
		const value = action.values[parentIndex]

		if (!value || value === "") return inputName

		return value instanceof Object ? value.label : value
	}, [action, parentIndex, inputName])

	const isReady = action && Boolean(action.values[parentIndex])

	const options = useMemo(() => {
		if (!action.values || !staticAction.fields[parentIndex].options) return

		if (childIndex === null) return staticAction.fields[parentIndex].options
	}, [staticAction, childIndex, parentIndex, action])

	const handleValue = (value: Value) => {
		const hasChanged =
			value instanceof Object && action.values[parentIndex] instanceof Object
				? value.value !== action.values[parentIndex].value
				: value !== action.values[parentIndex]

		if (hasChanged) {
			plugHandle.action.edit({
				id: plug?.id,
				actions: JSON.stringify(
					actions.map((action, nestedActionIndex) => ({
						...action,
						values: action.values.map((actionValue, valueIndex) => {
							if (actionIndex !== nestedActionIndex) return actionValue

							if (valueIndex === parentIndex) return value

							const fragment = dynamic[valueIndex]

							if (!fragment) return actionValue

							const [thisChildIndex] = getIndexes(fragment)

							const shouldReset =
								parentIndex === thisChildIndex ||
								action.values.some((_, idx) => {
									if (idx === valueIndex) return false
									const [depChildIndex] = getIndexes(dynamic[idx])
									if (!depChildIndex || !thisChildIndex) return false
									return (
										depChildIndex === thisChildIndex &&
										(idx === parentIndex || action.values[idx] === undefined)
									)
								})

							return shouldReset ? undefined : actionValue
						})
					}))
				)
			})
		}
	}

	return (
		<>
			<button
				className={cn(
					"rounded-sm bg-gradient-to-tr px-2 py-1 font-bold transition-all duration-200 ease-in-out",
					preview && isReady === false ? "text-plug-red" : "text-plug-green",
					own === true ? "cursor-pointer" : "cursor-default"
				)}
				style={{
					background:
						preview && isReady === false
							? "linear-gradient(to right, rgba(255,0,0,0.1), rgba(255,0,0,0.1))"
							: `linear-gradient(to right, rgba(0,239,54,0.1), rgba(147,223,0,0.1))`
				}}
				onClick={() => (own ? handle.frame() : undefined)}
			>
				{label}
			</button>

			<Frame
				index={index}
				icon={
					<div className="relative h-10 min-w-10">
						<Image
							src={protocol.metadata.icon}
							alt={action.protocol}
							width={64}
							height={64}
							className="absolute left-1/2 top-1/2 h-16 w-16 -translate-x-1/2 rounded-sm blur-2xl filter"
						/>
						<Image
							src={protocol.metadata.icon}
							alt={action.protocol}
							width={64}
							height={64}
							className="relative left-1/2 top-1/2 h-8 w-8 -translate-x-1/2 -translate-y-1/2 rounded-sm"
						/>
					</div>
				}
				label={
					<span className="relative">
						<span className="text-lg">
							<span className={cn(action.values.length > 1 && "opacity-40")}>
								{formatTitle(action.action)}
								{action.values.length > 1 && <span>:</span>}
							</span>
							{action.values.length > 1 && <span> {formatTitle(inputName)}</span>}
						</span>
					</span>
				}
				visible={isFrame}
				handleBack={dynamicIndex > 0 ? () => handle.frame(`${actionIndex}-${dynamicIndex - 1}`) : undefined}
				hasOverlay
				hasChildrenPadding={false}
				scrollBehavior="partial"
			>
				<div className="flex flex-col gap-2 overflow-y-auto px-6">
					{options === undefined && action.values[parentIndex] instanceof Object === false && (
						<Search
							className="mb-4"
							icon={<Hash size={14} />}
							placeholder={formatTitle(inputName)}
							// @ts-ignore
							search={action.values[parentIndex]}
							handleSearch={handleValue}
						/>
					)}

					{options !== undefined && (
						<div className="mb-4 flex w-full flex-col gap-2">
							{options.map((option, optionIndex) => (
								<div
									key={`${index}-${actionIndex}-${optionIndex}`}
									className="flex flex-row items-center gap-4"
								>
									<Checkbox
										checked={
											// @ts-ignore
											option.value === action.values[parentIndex]?.value
										}
										handleChange={() =>
											handleValue(
												// @ts-ignore
												option.value === action.values[parentIndex]?.value ? null : option
											)
										}
									/>

									<button
										key={`${index}-${actionIndex}-${optionIndex}`}
										className="group flex w-full flex-row items-center gap-4 truncate overflow-ellipsis whitespace-nowrap text-left font-bold"
										// @ts-ignore - I broke this while refactoring the way state is managed
										onClick={() => handleValue(option)}
									>
										{option.icon && (
											<div className="min-w-6">
												{option.icon.startsWith(
													"https://token-icons.llamao.fi/icons/tokens/"
												) ? (
													<TokenImage
														logo={option.icon}
														symbol={option.label}
														size="xs"
														blur={false}
													/>
												) : (
													<Image
														src={option.icon}
														alt={option.label}
														width={60}
														height={60}
														className="h-6 w-6 rounded-full"
														unoptimized
													/>
												)}
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
					)}
				</div>

				<div className="mt-auto bg-white">
					<div className="relative">
						{options && options.length > 0 && (
							<div className="pointer-events-none absolute -top-8 left-0 right-0 h-8 bg-gradient-to-b from-white/0 to-white" />
						)}
						<div className="mb-4 px-6">
							<Button
								variant={
									action.values[parentIndex] === "" ||
									(options !== undefined && action.values[parentIndex] === null)
										? "disabled"
										: "primary"
								}
								className="w-full py-4"
								onClick={() => {
									handle.frame(
										dynamicIndex + 1 < action.values.length
											? `${actionIndex}-${dynamicIndex + 1}`
											: undefined
									)
								}}
								disabled={
									action.values[parentIndex] === "" ||
									(options !== undefined && action.values[parentIndex] === null)
								}
							>
								{action.values[parentIndex] === "" ||
								(options !== undefined && action.values[parentIndex] === null)
									? "Choose an option"
									: action.values.length - 1 > dynamicIndex
										? "Next"
										: "Done"}
							</Button>
						</div>
					</div>
				</div>
			</Frame>
		</>
	)
}
