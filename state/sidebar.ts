import { useCallback } from "react"

import { atom, useAtom } from "jotai"

import { atomWithStorage } from "jotai/utils"

interface SidebarState {
	expanded: boolean
	searching: boolean
	viewingAs: boolean
	width: number
}

const DEFAULT_SIDEBAR_STATE: SidebarState = {
	expanded: false,
	searching: false,
	viewingAs: false,
	width: 380
}

const sidebarAtom = atomWithStorage<SidebarState>("sidebar", DEFAULT_SIDEBAR_STATE)

const createToggleAction = (key: keyof SidebarState) =>
	atom(
		get => get(sidebarAtom)[key],
		(get, set) => {
			const current = get(sidebarAtom)
			set(sidebarAtom, {
				...current,
				[key]: !current[key],
				...(key === "searching" ? { viewingAs: false } : {}),
				...(key === "viewingAs" ? { searching: false } : {})
			})
		}
	)

const expandedAtom = createToggleAction("expanded")
const searchingAtom = createToggleAction("searching")
const viewingAsAtom = createToggleAction("viewingAs")

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
	const [, toggleSearching] = useAtom(searchingAtom)
	const [, toggleViewingAs] = useAtom(viewingAsAtom)
	const [width, setWidth] = useAtom(widthAtom)

	return {
		is: sidebarState,
		width,
		toggleExpanded: useCallback(() => toggleExpanded(), [toggleExpanded]),
		toggleSearching: useCallback(() => toggleSearching(), [toggleSearching]),
		toggleViewingAs: useCallback(() => toggleViewingAs(), [toggleViewingAs]),
		resize: useCallback((newWidth: number) => setWidth(newWidth), [setWidth])
	}
}
