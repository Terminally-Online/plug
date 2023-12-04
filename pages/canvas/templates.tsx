import { useState } from "react"

import { GetServerSideProps, InferGetServerSidePropsType } from "next"

import { getSession } from "next-auth/react"

import { useMotionValueEvent, useScroll } from "framer-motion"
import { HandIcon } from "lucide-react"

import CanvasPreviewGrid from "@/components/canvas/preview-grid"
import { Search } from "@/components/canvas/search"
import SortBy from "@/components/canvas/sort-by"
import { TabsProvider } from "@/contexts/TabsProvider"
import { api, RouterOutputs } from "@/lib/api"
import { NextPageWithLayout } from "@/lib/types"

// TODO: Update the search bar to just work with any url.
// TODO: Implement a select that allows you to choose from a dropdown.
//	- This select should be used for:
//		- Sorting
// TODO: Sort by createdAt, updatedAt, and name.

// TODO: After alpha we will want to add selection of tags.

const Templates: NextPageWithLayout<
	InferGetServerSidePropsType<typeof getServerSideProps>
> = ({ search }) => {
	const { scrollYProgress } = useScroll()

	const [page, setPage] = useState(0)
	const [count, setCount] = useState(0)
	const [loading, setLoading] = useState(false)

	const [canvases, setCanvases] = useState<RouterOutputs["canvas"]["all"]>([])

	const { fetchNextPage } = api.canvas.infinite.useInfiniteQuery(
		{ search, limit: 20 },
		{
			getNextPageParam(lastPage) {
				return lastPage.nextCursor
			},
			onSuccess(data) {
				if (!data.pages[page]) return

				setCanvases(prevCanvases => [
					...prevCanvases,
					...data.pages[page].items
				])

				setCount(data.pages[page].count)
				setPage(prevPage => prevPage + 1)

				setLoading(false)
			}
		}
	)

	useMotionValueEvent(scrollYProgress, "change", latest => {
		if (loading) return
		if (latest < 0.8) return

		setLoading(true)
		fetchNextPage()
	})

	return (
		<div className="grid grid-cols-12 grid-rows-1">
			<div className="col-span-3 flex flex-col">
				<div className="group col-span-3 flex flex-col items-center justify-center gap-2 border-[1px] border-l-[0px] border-t-[0px] border-stone-950 bg-stone-900 p-8 py-24 text-center text-white transition-all duration-200 ease-in-out hover:bg-stone-950 hover:text-white">
					<div className="w-min rounded-full border-[1px] border-stone-950 bg-transparent p-2 transition-all duration-200 ease-in-out group-hover:border-white/20">
						<HandIcon
							width={18}
							height={18}
							className="opacity-60"
						/>
					</div>

					<h1 className="max-w-[280px] text-2xl">
						Welcome to the Plug Discovery Hub.
					</h1>
					<p className="max-w-[240px] text-sm opacity-60">
						Here you can find templates to get you started that have
						been made by the Plug community and team.
					</p>
				</div>

				<div className="top-18 sticky col-span-3 h-full border-r-[1px] border-stone-950">
					<Search baseUrl={"/canvas/templates"} results={count} />
					<SortBy />
				</div>
			</div>

			<div className="col-span-9">
				<CanvasPreviewGrid canvases={canvases} />
			</div>
		</div>
	)
}

export const getServerSideProps = (async context => {
	if (!(await getSession(context))) {
		return {
			redirect: {
				destination: `/connect`,
				permanent: false
			}
		}
	}

	return {
		props: {
			search: context.query.search || ""
		}
	}
}) satisfies GetServerSideProps<{
	search: string | string[] | undefined
}>

Templates.getLayout = page => <TabsProvider>{page}</TabsProvider>

export default Templates
