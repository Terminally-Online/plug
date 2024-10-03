import { FC, useMemo } from "react"

import { Blocks, SearchIcon } from "lucide-react"

import { ActionItem, Button, Frame, Image, Search } from "@/components"
import { abis, categories, formatAddress, formatTitle, actions as staticActions, useDebounce } from "@/lib"
import { useColumns } from "@/state"

export const ActionsFrame: FC<{ index: number; item: string }> = ({ index, item }) => {
	const { column, isFrame, frame } = useColumns(index, `${index}-${item}-actions`)
	const [search, debouncedSearch, handleDebounce] = useDebounce("")

	const allFilteredActions = useMemo(
		() =>
			Object.keys(staticActions).flatMap(categoryName => {
				if (categoryName.toLowerCase().includes(debouncedSearch.toLowerCase())) {
					return Object.keys(staticActions[categoryName]).map(actionName => ({
						categoryName,
						actionName
					}))
				}

				return Object.keys(staticActions[categoryName])
					.filter(actionName => actionName.toLowerCase().includes(debouncedSearch.toLowerCase()))
					.map(actionName => ({
						categoryName,
						actionName
					}))
			}),
		[debouncedSearch]
	)

	if (!column) return null

	return (
		<>
			<Frame
				index={index}
				className="scrollbar-hide z-[1] max-h-[85%] overflow-y-auto"
				icon={<Blocks size={18} className="opacity-60" />}
				label="Add Action"
				visible={isFrame}
				hasChildrenPadding={false}
			>
				<div className="flex flex-col px-6">
					<Search
						className="mb-4"
						icon={<SearchIcon size={14} />}
						placeholder="Search protocols and actions"
						search={search}
						handleSearch={handleDebounce}
						clear={true}
					/>

					<div className="flex flex-col gap-2">
						{allFilteredActions.length > 0 ? (
							allFilteredActions.map(({ categoryName, actionName }) => (
								<ActionItem
									key={`${categoryName}-${actionName}`}
									index={index}
									item={item}
									categoryName={categoryName}
									actionName={actionName}
									image={true}
								/>
							))
						) : (
							<>
								<div className="mx-auto my-64 flex h-full max-w-[80%] flex-col gap-2 text-center">
									<p className="text-lg font-bold">No matching actions.</p>
									<p className="opacity-60">
										Search for another action or protocol or request a new integration.
									</p>

									<Button
										className="mx-auto mt-4 w-max"
										onClick={() => frame("featureRequest-actions")}
									>
										Request Integration
									</Button>
								</div>
							</>
						)}
					</div>
				</div>
			</Frame>

			{Object.keys(categories).map(categoryName => {
				return Object.keys(staticActions[categoryName]).map(actionName => {
					const action = staticActions[categoryName][actionName]

					return (
						<Frame
							index={index}
							key={`${categoryName}-${actionName}-info`}
							className="scrollbar-hide z-[3] max-h-[calc(100vh-80px)] overflow-y-auto"
							icon={
								<Image
									src={categories[categoryName].image}
									alt={categoryName}
									width={24}
									height={24}
									className="h-6 w-6 rounded-md"
								/>
							}
							label={formatTitle(actionName)}
							visible={column.frame === `${index}-${categoryName}-${actionName}`}
							handleBack={() => frame(`${index}-${item}-actions`)}
							hasOverlay={true}
							hasChildrenPadding={false}
						>
							<div className="flex flex-col gap-8 px-6">
								<p className="opacity-60">{action.info}</p>
								<div className="flex flex-col gap-2">
									<p className="font-bold">Input Data</p>
									{action.inputs.length > 0 &&
										action.inputs.map((input, index) => (
											<p key={index} className="flex w-full flex-row gap-2">
												<span className="font-bold opacity-40">
													{formatTitle(input.name?.replace("$", "") ?? "")}
												</span>
												<span className="ml-auto opacity-60">{input.type}</span>
											</p>
										))}
								</div>
								<div className="flex flex-col gap-2">
									<p className="font-bold">Constraint</p>
									<p className="flex flex-row gap-2">
										<span className="font-bold opacity-40">Address</span>
										<span className="ml-auto opacity-60">{formatAddress(action.address)}</span>
									</p>
									<p className="flex flex-row gap-2">
										<span className="font-bold opacity-40">Function</span>
										<span className="ml-auto opacity-60">
											{abis[categoryName][actionName]
												? abis[categoryName][actionName].split("(")[0].replace("function ", "")
												: "Unknown"}
										</span>
									</p>
									<p className="flex flex-row gap-2">
										<span className="mr-auto font-bold opacity-40">Supported Chains</span>
										{categories[categoryName].chains.map(chain => (
											<Image
												key={chain}
												className="ml-[-20px] h-6 w-6"
												src={`/blockchain/${chain}.png`}
												alt={chain}
												width={24}
												height={24}
											/>
										))}
									</p>
								</div>
							</div>
						</Frame>
					)
				})
			})}
		</>
	)
}
