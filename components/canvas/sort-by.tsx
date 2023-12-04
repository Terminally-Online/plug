import { useEffect, useState } from "react"

import { useRouter } from "next/router"

import { AnimatePresence, motion } from "framer-motion"
import { ChevronDownIcon, FilterIcon } from "lucide-react"

import { cn } from "@/lib/utils"

const keys = {
	default: "",
	newest: "Newest",
	oldest: "Oldest",
	active: "Active"
}

export const SortBy = () => {
	const router = useRouter()

	const [collapsed, setCollapsed] = useState(true)
	const [selected, setSelected] =
		useState<(typeof keys)[keyof typeof keys]>("")

	const handleSelect = (key: keyof typeof keys) => {
		setCollapsed(true)
		setSelected(key)
	}

	useEffect(() => {
		if (selected === "" || selected === "default") {
			const query = { ...router.query }
			delete query.sort

			router.push({
				query
			})

			return
		}

		router.push({
			query: {
				...router.query,
				sort: selected
			}
		})
	}, [selected])

	return (
		<>
			<div
				className={cn(
					"group flex cursor-pointer select-none flex-row items-center gap-4 border-b-[1px] border-stone-950 p-4 text-sm text-muted-foreground transition-all duration-200 ease-in-out hover:bg-stone-950 hover:text-white",
					collapsed === false ? "bg-stone-950 text-white" : ""
				)}
				onClick={() => setCollapsed(prev => !prev)}
			>
				<FilterIcon
					width={10}
					height={10}
					className="text-white opacity-40 transition-all duration-200 ease-in-out group-hover:opacity-100"
				/>
				<span>
					SORT BY
					{selected !== "" && (
						<>
							{": "}
							<span className="ml-auto text-white">
								{keys[
									selected as keyof typeof keys
								].toUpperCase()}
							</span>
						</>
					)}
				</span>
				<motion.span
					className="ml-auto"
					animate={{
						rotate: collapsed ? 0 : 180
					}}
				>
					<ChevronDownIcon
						width={10}
						height={10}
						className="text-white opacity-60"
					/>
				</motion.span>
			</div>

			{collapsed === false && (
				<AnimatePresence key="accordion">
					<motion.div
						className="flex flex-col text-sm transition-all duration-200 ease-in-out"
						initial={{ opacity: 0 }}
						animate={{ opacity: 1 }}
						exit={{ opacity: 0 }}
						transition={{ duration: 0.2 }}
					>
						{Object.entries(keys)
							.filter(([key]) => key !== "default")
							.map(([key, value]) => (
								<motion.button
									key={key}
									className="group flex cursor-pointer flex-row items-center gap-2 border-b-[1px] border-stone-950 bg-stone-900 p-4 text-muted-foreground transition-all duration-200 ease-in-out hover:bg-stone-950 hover:text-white"
									initial={{ opacity: 0 }}
									animate={{ opacity: 1 }}
									exit={{ opacity: 0 }}
									transition={{ duration: 0.2 }}
									onClick={() =>
										handleSelect(key as keyof typeof keys)
									}
								>
									<label
										htmlFor={key}
										className="border-l-[1px] border-muted-foreground pl-8 group-hover:border-white"
									>
										{value.toUpperCase()}
									</label>
								</motion.button>
							))}
					</motion.div>
				</AnimatePresence>
			)}
		</>
	)
}

export default SortBy
