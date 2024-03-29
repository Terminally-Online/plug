import type { FC } from "react"
import { useEffect } from "react"

import { useRouter } from "next/router"

import { ChevronRightIcon, Cross2Icon } from "@radix-ui/react-icons"

import { Input } from "@/components/ui/input"
import { useDebounce } from "@/lib/hooks/useDebounce"
import { cn } from "@/lib/utils"

export type SearchProps = {
	search?: string | string[]
	results?: number
}

export const Search: FC<SearchProps> = ({ search, results = 0 }) => {
	const router = useRouter()

	const { debounce, value, debounced } = useDebounce({ initial: search })

	useEffect(() => {
		if (search !== debounced) {
			// * If all the text was removed, remove the url param from the query.
			if (debounced === "") {
				const query = { ...router.query }
				delete query.search

				router.push({
					query
				})

				return
			}

			router.push({
				query: {
					...router.query,
					search: debounced
				}
			})
		}
	}, [router, search, debounced])

	return (
		<div
			className={cn(
				"transition-bg group sticky top-12 z-[99] flex w-full flex-row items-center border-b-[1px] border-r-[1px] border-stone-950 bg-stone-900 px-4 duration-200 ease-in-out hover:bg-stone-950",
				search && search.length > 0 ? "bg-stone-950" : ""
			)}
		>
			<ChevronRightIcon
				width={16}
				height={16}
				className="flex h-full text-white opacity-60 group-hover:opacity-100"
			/>

			<Input
				placeholder="SEARCH ALL CANVASES"
				className="relative w-full bg-transparent py-8 uppercase text-white"
				value={value ?? search}
				onChange={e => {
					debounce(e.target.value)
				}}
			/>
			{search && search !== "" && (
				<button onClick={() => debounce("")} className="ml-auto">
					<Cross2Icon
						width={16}
						height={16}
						className="text-white opacity-60"
					/>
				</button>
			)}
			<p className="ml-auto block w-max min-w-[100px] select-none text-right text-sm tabular-nums text-white opacity-60 group-hover:opacity-100">
				{results} results
			</p>
		</div>
	)
}
