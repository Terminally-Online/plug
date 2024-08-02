import {
	createContext,
	FC,
	PropsWithChildren,
	useContext,
	useState
} from "react"

type Page = "home" | "discover" | "mine" | "activity" | "create" | "plug"

export const PageContext = createContext<{
	page: Page
	handlePage: (page: Page) => void
}>({
	page: "home",
	handlePage: () => {}
})

export const PageProvider: FC<PropsWithChildren> = ({ children }) => {
	const [page, handlePage] = useState<Page>("home")

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
