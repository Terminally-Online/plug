import { useCallback } from "react"

import { atom, useAtom } from "jotai"

import { atomWithStorage } from "jotai/utils"

interface SidebarState {
	expanded: boolean
	searching: boolean
	viewingAs: boolean
}

const DEFAULT_SIDEBAR_STATE: SidebarState = {
	expanded: false,
	searching: false,
	viewingAs: false
}

const sidebarAtom = atomWithStorage<SidebarState>("sidebar", DEFAULT_SIDEBAR_STATE)

const createToggleAction = (key: keyof SidebarState) =>
	atom(
		get => get(sidebarAtom)[key],
		(get, set) => {
			const current = get(sidebarAtom)
			set(sidebarAtom, { ...current, [key]: !current[key] })
		}
	)

const expandedAtom = createToggleAction("expanded")
const searchingAtom = createToggleAction("searching")
const viewingAsAtom = createToggleAction("viewingAs")

export const useSidebar = () => {
	const [sidebarState] = useAtom(sidebarAtom)
	const [, toggleExpanded] = useAtom(expandedAtom)
	const [, toggleSearching] = useAtom(searchingAtom)
	const [, toggleViewingAs] = useAtom(viewingAsAtom)

	return {
		is: sidebarState,
		toggleExpanded: useCallback(() => toggleExpanded(), [toggleExpanded]),
		toggleSearching: useCallback(() => toggleSearching(), [toggleSearching]),
		toggleViewingAs: useCallback(() => toggleViewingAs(), [toggleViewingAs])
	}
}
