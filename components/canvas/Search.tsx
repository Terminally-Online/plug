import type { FC } from 'react'
import { useEffect } from 'react'

import { useRouter, useSearchParams } from 'next/navigation'

import { ChevronRightIcon } from '@radix-ui/react-icons'

import { Input } from '@/components/ui/input'
import { useDebounce } from '@/lib/hooks/useDebounce'

export type SearchProps = {
	baseUrl?: string
}

export const Search: FC<SearchProps> = ({ baseUrl }) => {
	const router = useRouter()
	const searchParams = useSearchParams()

	const search = searchParams.get('search') ?? ''

	const { debounce, value, debounced } = useDebounce({ initial: search })

	useEffect(() => {
		// * This is designed to fire once debounced is ahead of search.
		if (search !== debounced)
			router.push(`${baseUrl ?? '/canvas/create'}?search=${debounced}`)
	}, [baseUrl, router, search, debounced])

	return (
		<div className="group flex flex-row items-center bg-stone-900 px-4">
			<ChevronRightIcon
				width={18}
				height={18}
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
		</div>
	)
}
