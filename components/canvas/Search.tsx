import type { FC } from 'react'
import { useEffect } from 'react'

import { useRouter, useSearchParams } from 'next/navigation'

import { ChevronRightIcon, Cross2Icon } from '@radix-ui/react-icons'

import { Input } from '@/components/ui/input'
import { useDebounce } from '@/lib/hooks/useDebounce'

export type SearchProps = {
	baseUrl?: string
	results?: number
}

export const Search: FC<SearchProps> = ({ baseUrl, results = 0 }) => {
	const router = useRouter()
	const searchParams = useSearchParams()

	const search = searchParams.get('search') ?? ''

	const { debounce, value, debounced } = useDebounce({ initial: search })

	useEffect(() => {
		if (search !== debounced) {
			const searchParams = new URLSearchParams()
			searchParams.set('search', debounced)
			// * Add the search param to the URL.
			const newUrl = `${
				baseUrl ?? '/canvas/create'
			}?${searchParams.toString()}`
			// * Push the new URL to the router.
			router.push(newUrl)
		}
	}, [baseUrl, router, search, debounced])

	return (
		<div className="group flex flex-row items-center bg-stone-900 px-4">
			<div className="flex w-full flex-row items-center">
				<ChevronRightIcon
					width={16}
					height={16}
					className="group:hover:opacity-100 flex h-full text-white opacity-60"
				/>
				<Input
					placeholder="SEARCH ALL CANVASES"
					className="w-full py-8 text-white"
					value={value}
					onChange={e => {
						debounce(e.target.value)
					}}
				/>
				{search && search !== '' && (
					<button onClick={() => debounce('')} className="ml-auto">
						<Cross2Icon
							width={16}
							height={16}
							className="text-white opacity-60"
						/>
					</button>
				)}
				<p className="ml-auto block w-max min-w-[100px] text-right text-sm tabular-nums text-white/60">
					{results} results.
				</p>
			</div>
		</div>
	)
}
