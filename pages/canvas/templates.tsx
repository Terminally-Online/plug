import { useState } from "react"

import { GetServerSideProps, InferGetServerSidePropsType } from "next"

import { getSession } from "next-auth/react"

import { useMotionValueEvent, useScroll } from "framer-motion"

import CanvasPreviewGrid from "@/components/canvas/preview-grid"
import { Search } from "@/components/canvas/search"
import SortBy from "@/components/canvas/sort-by"
import Welcome from "@/components/canvas/welcome"
import { TabsProvider } from "@/contexts/TabsProvider"
import { api, RouterOutputs } from "@/lib/api"
import { NextPageWithLayout } from "@/lib/types"

// TODO: After alpha we will want to add selection of tags.

const Templates: NextPageWithLayout<
	InferGetServerSidePropsType<typeof getServerSideProps>
> = ({ search, sort }) => {
	const { scrollYProgress } = useScroll()

	const [canvasState, setCanvasState] = useState<{
		loading?: boolean
		count?: number
		search?: string | string[]
		sort?: string | string[]
		canvases: RouterOutputs["canvas"]["all"]
	}>({ search, canvases: [] })

	const { fetchNextPage } = api.canvas.infinite.useInfiniteQuery(
		{ search, sort, limit: 20 },
		{
			getNextPageParam(lastPage) {
				return lastPage.nextCursor
			},
			onSuccess(data) {
				setCanvasState(prevCanvasState => {
					const page = data.pages.length - 1

					// ? Clear the list and start over instead of appending.
					if (
						prevCanvasState.search !== search ||
						prevCanvasState.sort !== sort
					)
						return {
							loading: false,
							count: data.pages[page].count,
							search,
							sort,
							canvases: data.pages[page].items
						}

					return {
						loading: false,
						count: data.pages[page].count,
						search,
						sort,
						canvases: [
							...prevCanvasState.canvases,
							...data.pages[page].items
						]
					}
				})
			}
		}
	)

	useMotionValueEvent(scrollYProgress, "change", latest => {
		if (canvasState.loading) return
		if (latest < 0.8) return

		const hasNext = canvasState.canvases.length < (canvasState.count ?? 0)

		// ? Do not keep fetching if we've reached the end of the list.
		if (hasNext) {
			setCanvasState(prevState => {
				fetchNextPage()

				return { ...prevState, loading: true }
			})
		}
	})

	return (
		<div className="grid grid-cols-12 grid-rows-1">
			<div className="col-span-3 flex flex-col">
				<Welcome />

				<div className="top-18 sticky col-span-3 h-full border-r-[1px] border-stone-950">
					<Search search={search} results={canvasState.count} />
					<SortBy />
				</div>
			</div>

			<div className="col-span-9">
				<CanvasPreviewGrid canvases={canvasState.canvases} />
			</div>
		</div>
	)
}

export const getServerSideProps = (async context => {
	if (!(await getSession(context))) {
		return {
			redirect: {
				destination: `/?connect=true`,
				permanent: false
			}
		}
	}

	return {
		props: {
			search: context.query.search || "",
			sort: context.query.sort || ""
		}
	}
}) satisfies GetServerSideProps<{
	search: string | string[] | undefined
	sort: string | string[] | undefined
}>

Templates.getLayout = page => <TabsProvider>{page}</TabsProvider>

export default Templates
