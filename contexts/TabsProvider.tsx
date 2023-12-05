"use client"

import {
	createContext,
	FC,
	PropsWithChildren,
	useCallback,
	useContext,
	useEffect,
	useState
} from "react"

import { usePathname, useRouter } from "next/navigation"

import { Hud } from "@/components/viewport/hud"

type Tab = {
	label: string
	color: string
	href: string
	active?: boolean
}

const ephemeralTabs: string[] = ["/canvas/create", "/canvas/templates"]

export const TabsContext = createContext<{
	tabs: Tab[]
	ephemeralTabs: string[]
	handleAdd: (tab: Tab) => void
	handleRemove: (index: number) => void
	handleMove: (index: number, newIndex: number) => void
}>({
	tabs: [],
	ephemeralTabs: [],
	handleAdd: () => {},
	handleRemove: () => {},
	handleMove: () => {}
})

export const TabsProvider: FC<PropsWithChildren> = ({ children }) => {
	const router = useRouter()
	const path = usePathname()

	const [tabs, setTabs] = useState<Tab[]>(() => {
		if (typeof window === "undefined") return []

		const savedTabs = localStorage.getItem("tabs")

		return savedTabs ? JSON.parse(savedTabs) : []
	})

	const handleAdd = useCallback(
		(tab: Tab) => {
			setTabs(tabs => {
				// * If we already have a tab with the same href, we need
				//   to replace the element.
				if (tabs.find(t => t.href === tab.href)) {
					return tabs.map(t => (t.href === tab.href ? tab : t))
				}

				// * Go to the newly created page.
				router.push(tab.href)

				return [...tabs, tab]
			})
		},
		[router]
	)

	const handleRemove = useCallback(
		(index: number) => {
			setTabs(tabs => {
				const newTabs = [...tabs]

				newTabs.splice(index, 1)

				// * If we do not need a layout shift, we can just remove the tab.
				if (tabs[index].active) {
					// * When removing, we may remove the active tab. In that case, we need to
					//   redirect to the last tab in the list.
					if (newTabs.length > 0) {
						const lastTab = newTabs[newTabs.length - 1]

						router.push(lastTab.href)
					}
					// * If there are no more tabs, we need to redirect to the home page.
					else router.push("/canvas")
				}

				return newTabs
			})
		},
		[router]
	)

	const handleMove = useCallback((index: number, newIndex: number) => {
		setTabs(tabs => {
			const newTabs = [...tabs]

			const [removed] = newTabs.splice(index, 1)

			newTabs.splice(newIndex, 0, removed)

			return newTabs
		})
	}, [])

	useEffect(() => {
		localStorage.setItem("tabs", JSON.stringify(tabs))
	}, [tabs])

	useEffect(() => {
		// * Mark the selected tab as active when fit.
		setTabs(tabs => {
			const activatedTabs = tabs.map(tab => ({
				...tab,
				active: tab.href === path
			}))

			// * If we have any ephemeral tabs that we are not on, remove them.
			return activatedTabs.filter(
				tab => !ephemeralTabs.includes(tab.href) || tab.active
			)
		})
	}, [path])

	return (
		<TabsContext.Provider
			value={{
				tabs,
				ephemeralTabs,
				handleAdd,
				handleRemove,
				handleMove
			}}
		>
			<Hud>{children}</Hud>
		</TabsContext.Provider>
	)
}

export const useTabs = () => useContext(TabsContext)
