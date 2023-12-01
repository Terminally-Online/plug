import { FC, PropsWithChildren, useEffect, useState } from 'react'

import Link from 'next/link'
import { usePathname, useRouter } from 'next/navigation'

import { Cross1Icon, HomeIcon, PlusIcon } from '@radix-ui/react-icons'

import { useTabs } from '@/contexts/TabsProvider'
import { cn } from '@/lib/utils'

export const Hud: FC<PropsWithChildren> = ({ children }) => {
	const { tabs, createTab, handleAdd, handleRemove } = useTabs()

	const router = useRouter()
	const path = usePathname()

	const [client, setClient] = useState(false)

	useEffect(() => {
		setClient(true)
	}, [])

	useEffect(() => {
		// * We create a custom tab for the canvas creation route.
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
			<div className="fixed left-0 top-0 z-[99999] w-screen border-b-[1px] border-b-stone-950 bg-stone-900">
				<div className="flex h-8 flex-row items-center">
					<Link
						href="/canvas"
						className="text-md pointer-events-auto flex h-full items-center justify-center bg-stone-800 p-2 font-bold text-white/60 transition-all duration-200 ease-in-out hover:bg-white hover:text-stone-950"
					>
						<HomeIcon width={16} height={16} />
					</Link>

					{tabs.map(({ label, color, href, active }, index) => (
						<button
							key={href}
							className={cn(
								'group flex h-full flex-row items-center gap-4 border-l-[1px] border-l-stone-950 px-4 text-sm text-white/60 transition-all duration-200 ease-in-out hover:bg-stone-950 hover:text-white active:bg-white active:text-stone-950/60 hover:active:bg-white active:hover:text-stone-950',
								active ? 'active' : ''
							)}
							onClick={() => router.push(href)}
							suppressHydrationWarning
						>
							<div className="flex h-full flex-row items-center gap-4">
								<div
									className="h-2 w-2 rounded-full"
									style={{ backgroundColor: color }}
								/>
								{label}
							</div>

							<button
								type="button"
								className={cn(
									'flex h-full items-center justify-center text-white/60 opacity-0 transition-all duration-200 ease-in-out hover:text-white group-hover:opacity-100 active:text-stone-950 active:text-stone-950/60',
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
							className="flex h-full items-center justify-center border-x-[1px] border-x-stone-950 bg-stone-800 px-2 text-white/60 transition-all duration-200 ease-in-out hover:bg-white hover:text-stone-950"
							onClick={() => {
								router.push('/canvas/create')
							}}
						>
							<PlusIcon width={16} height={16} />
						</button>
					) : null}
				</div>
			</div>

			<div className="flex h-screen flex-col overscroll-none pt-8">
				{children}
			</div>
		</>
	)
}

export default Hud
