import {
	createContext,
	FC,
	PropsWithChildren,
	useContext,
	useState
} from "react"

import { Page } from "@/lib"

const DEFAULT_PAGE: Page = {
	key: "home"
} as const

export const PageContext = createContext<{
	page: Page
	handlePage: (page: Page) => void
}>({
	page: DEFAULT_PAGE,
	handlePage: () => {}
})

export const PageProvider: FC<PropsWithChildren> = ({ children }) => {
	const [page, handlePage] = useState<Page>(DEFAULT_PAGE)

	return (
		<PageContext.Provider
			value={{
				page,
				handlePage
			}}
		>
			{children}
		</PageContext.Provider>
	)
}

export const usePage = () => useContext(PageContext)
