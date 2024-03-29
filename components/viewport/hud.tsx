import type { FC, PropsWithChildren } from "react"
import { useEffect, useState } from "react"

import { usePathname } from "next/navigation"
import { useRouter } from "next/router"

import { Cross1Icon, HomeIcon } from "@radix-ui/react-icons"

import { Toggler, Vault, Wallet } from "@/components/viewport/vault"
import { useTabs } from "@/contexts"
import { cn } from "@/lib/utils"

export const Hud: FC<PropsWithChildren> = ({ children }) => {
	const { tabs, handleAdd, handleRemove } = useTabs()

	const router = useRouter()
	const path = usePathname()

	const [client, setClient] = useState(false)

	const color = Math.floor(Math.random() * 16777215).toString(16)

	useEffect(() => {
		setClient(true)
	}, [])

	useEffect(() => {
		switch (path) {
			case "/canvas/create":
				handleAdd({
					label: `New Canvas`,
					color: `#${color}`,
					href: `/canvas/create`,
					active: true
				})
				break
			case "/canvas/templates":
				handleAdd({
					label: `Templates`,
					color: "#FF8C00",
					href: `/canvas/templates`,
					active: true
				})
				break
			default:
				break
		}
	}, [color, path, handleAdd])

	if (!client) return null

	return (
		<>
			<div className="fixed left-0 top-0 z-[99999] w-screen border-b-[1px] border-b-stone-950 bg-stone-900">
				<div className="flex h-12 flex-row items-center">
					<button
						onClick={() =>
							router.push({
								pathname: "/canvas/",
								query: {}
							})
						}
						className="text-md pointer-events-auto flex h-full items-center justify-center bg-stone-800 p-2 px-4 font-bold text-white/60 transition-all duration-200 ease-in-out hover:bg-white hover:text-stone-950"
					>
						<HomeIcon width={16} height={16} />
					</button>

					{tabs.map(({ label, color, href, active }, index) => (
						<button
							key={href}
							className={cn(
								"group flex h-full flex-row items-center gap-4 border-r-[1px] border-r-stone-950 px-4 text-sm text-white/60 transition-all duration-200 ease-in-out hover:bg-stone-950 hover:text-white active:bg-white active:text-stone-950/60 hover:active:bg-white active:hover:text-stone-950",
								active ? "active" : ""
							)}
							onClick={() => router.push(href)}
							suppressHydrationWarning
						>
							<div className="flex h-full flex-row items-center gap-4 ">
								<div
									className="h-2 w-2 rounded-full"
									style={{ backgroundColor: color }}
								/>
								<div className="max-w-[140px] overflow-hidden overflow-ellipsis whitespace-nowrap">
									{label}
								</div>
							</div>

							<button
								type="button"
								className={cn(
									"flex h-full items-center justify-center text-white/60 opacity-0 transition-all duration-200 ease-in-out hover:text-white group-hover:opacity-100 active:text-stone-950 active:text-stone-950/60",
									active ? "active" : ""
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

					<Toggler />
					<Wallet />
				</div>
			</div>

			<div className="inset-0 flex h-full min-h-screen flex-row overscroll-none bg-stone-900 pt-12">
				<div className="w-full">{children}</div>

				<Vault />
			</div>
		</>
	)
}

export default Hud
