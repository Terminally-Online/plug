import { useCallback } from "react"

import { atom, useAtom } from "jotai"

import { atomWithStorage } from "jotai/utils"

interface SidebarState {
	expanded: boolean
	activePane: "authenticating" | "searching" | "viewingAs" | null
	width: number
}

const DEFAULT_SIDEBAR_STATE: SidebarState = {
	expanded: false,
	activePane: null,
	width: 380
}

const sidebarAtom = atomWithStorage<SidebarState>("sidebar", DEFAULT_SIDEBAR_STATE)

const expandedAtom = atom(
	get => get(sidebarAtom).expanded,
	(get, set) => {
		const current = get(sidebarAtom)
		set(sidebarAtom, { ...current, expanded: !current.expanded })
	}
)

const activePaneAtom = atom(
	get => get(sidebarAtom).activePane,
	(get, set, newPane: SidebarState["activePane"]) => {
		const current = get(sidebarAtom)
		set(sidebarAtom, { ...current, activePane: current.activePane === newPane ? null : newPane })
	}
)

const widthAtom = atom(
	get => get(sidebarAtom).width,
	(get, set, newWidth: number) => {
		const current = get(sidebarAtom)
		set(sidebarAtom, { ...current, width: newWidth })
	}
)

export const useSidebar = () => {
	const [sidebarState] = useAtom(sidebarAtom)
	const [, toggleExpanded] = useAtom(expandedAtom)
	const [, setActivePane] = useAtom(activePaneAtom)
	const [width, setWidth] = useAtom(widthAtom)

	return {
		is: {
			...sidebarState,
			authenticating: sidebarState.activePane === "authenticating",
			searching: sidebarState.activePane === "searching",
			viewingAs: sidebarState.activePane === "viewingAs"
		},
		width,
		handleActivePane: useCallback((newPane: SidebarState["activePane"]) => setActivePane(newPane), [setActivePane]),
		toggleExpanded: useCallback(() => toggleExpanded(), [toggleExpanded]),
		toggleAuthenticating: useCallback(() => setActivePane("authenticating"), [setActivePane]),
		toggleSearching: useCallback(() => setActivePane("searching"), [setActivePane]),
		toggleViewingAs: useCallback(() => setActivePane("viewingAs"), [setActivePane]),
		resize: useCallback((newWidth: number) => setWidth(newWidth), [setWidth])
	}
}
