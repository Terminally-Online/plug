'use client'

import { FC, PropsWithChildren, useEffect, useState } from 'react'

import Link from 'next/link'
import { usePathname, useRouter } from 'next/navigation'

import { useTabs } from '@/contexts/TabsProvider'
import { cn } from '@/lib/utils'

import { Cross1Icon, HomeIcon, PlusIcon } from '@radix-ui/react-icons'

export const Hud: FC<PropsWithChildren> = ({ children }) => {
	const { tabs, createTab, handleAdd, handleRemove } = useTabs()

	const router = useRouter()
	const path = usePathname()

	const [client, setClient] = useState(false)

	useEffect(() => {
		setClient(true)
	}, [])

	useEffect(() => {
		switch (path) {
			case '/canvas/create':
				handleAdd({
					label: `New Canvas`,
					color: `#${Math.floor(Math.random() * 16777215).toString(
						16
					)}`,
					href: `/canvas/create`,
					active: true
				})
				break
			default:
				break
		}
	}, [path, handleAdd])

	if (!client) return null

	return (
		<>
			<div className="bg-stone-900 border-b-[1px] border-b-stone-950 fixed top-0 left-0 w-screen z-[99999]">
				<div className="flex flex-row items-center h-8">
					<Link
						href="/canvas"
						className="bg-stone-800 text-white/60 p-2 h-full flex items-center justify-center text-md font-bold pointer-events-auto hover:bg-white hover:text-stone-950 transition-all duration-200 ease-in-out"
					>
						<HomeIcon width={16} height={16} />
					</Link>

					{tabs.map(({ label, color, href, active }, index) => (
						<button
							key={href}
							className={cn(
								'group border-l-[1px] border-l-stone-950 h-full px-4 text-white/60 hover:bg-stone-950 hover:text-white active:bg-white hover:active:bg-white active:text-stone-950/60 active:hover:text-stone-950 text-sm transition-all duration-200 ease-in-out flex flex-row items-center gap-4',
								active ? 'active' : ''
							)}
							onClick={() => router.push(href)}
							suppressHydrationWarning
						>
							<div className="h-full flex flex-row items-center gap-4">
								<div
									className="w-2 h-2 rounded-full"
									style={{ backgroundColor: color }}
								/>
								{label}
							</div>

							<button
								type="button"
								className={cn(
									'h-full flex items-center justify-center opacity-0 group-hover:opacity-100 text-white/60 hover:text-white active:text-stone-950 transition-all duration-200 ease-in-out active:text-stone-950/60',
									active ? 'active' : ''
								)}
								onClick={e => {
									// ? Without this here, the close button would also fire navigation to the canvas route.
									e.stopPropagation()
									handleRemove(index)
								}}
							>
								<Cross1Icon width={12} height={12} />
							</button>
						</button>
					))}

					{createTab === undefined ? (
						<button
							type="button"
							className="px-2 h-full flex items-center justify-center border-x-[1px] border-x-stone-950 bg-stone-800 text-white/60 hover:bg-white hover:text-stone-950 transition-all duration-200 ease-in-out"
							onClick={() => {
								router.push('/canvas/create')
							}}
						>
							<PlusIcon width={16} height={16} />
						</button>
					) : null}
				</div>
			</div>

			<div className="pt-8 h-screen">{children}</div>
		</>
	)
}

export default Hud
