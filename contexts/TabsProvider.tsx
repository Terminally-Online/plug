import type { FC, PropsWithChildren } from "react"
import {
	createContext,
	useCallback,
	useContext,
	useEffect,
	useMemo,
	useState
} from "react"

import { usePathname, useRouter } from "next/navigation"

import { Activity } from "lucide-react"

import { Hud } from "@/components/viewport/hud"
import { Deploy, Deposit, Tokens, Withdraw } from "@/components/viewport/vault"
import { Wallet } from "@/components/viewport/vault/actions/manage"

type Tab = {
	label: string
	color: string
	href: string
	active?: boolean
}

const ephemeralTabs: string[] = ["/canvas/create", "/canvas/templates"]

export const INITIAL_PANE = "tokens"

export const TabsContext = createContext<{
	pane: string
	Panel: JSX.Element
	tabs: Tab[]
	ephemeralTabs: string[]
	expanded: boolean
	handlePane: (pane: string | undefined) => void
	handleAdd: (tab: Tab) => void
	handleRemove: (index: number) => void
	handleMove: (index: number, newIndex: number) => void
	handleExpanded: () => void
}>({
	pane: INITIAL_PANE,
	Panel: <Tokens />,
	tabs: [],
	ephemeralTabs: [],
	expanded: true,
	handlePane: () => {},
	handleAdd: () => {},
	handleRemove: () => {},
	handleMove: () => {},
	handleExpanded: () => {}
})

export const TabsProvider: FC<PropsWithChildren> = ({ children }) => {
	const router = useRouter()
	const path = usePathname()

	const [expanded, setExpanded] = useState(true)
	const [tabs, setTabs] = useState<Tab[]>(() => {
		if (typeof window === "undefined") return []

		const savedTabs = localStorage.getItem("tabs")

		return savedTabs ? JSON.parse(savedTabs) : []
	})
	const [pane, setPane] = useState(INITIAL_PANE)
	const [nextPane, setNextPane] = useState(INITIAL_PANE)

	const Panel = useMemo(() => {
		switch (pane) {
			case "tokens":
				return <Tokens />
			case "activity":
				return <Activity />
			case "withdraw":
				return <Withdraw />
			case "wallet":
				return <Wallet />
			case "deploy":
				return <Deploy />
			default:
				return <Deposit />
		}
	}, [pane])

	const handlePane = useCallback((pane: string | undefined) => {
		setPane(previousPane => {
			const newPane = pane || nextPane

			setNextPane(previousPane)

			return newPane
		})
	}, [])

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

	const handleExpanded = useCallback(() => {
		setExpanded(expanded => !expanded)
	}, [expanded])

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
				pane,
				Panel,
				tabs,
				ephemeralTabs,
				expanded,
				handlePane,
				handleAdd,
				handleRemove,
				handleMove,
				handleExpanded
			}}
		>
			<Hud>{children}</Hud>
		</TabsContext.Provider>
	)
}

export const useTabs = () => useContext(TabsContext)
