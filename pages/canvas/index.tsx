import { useState } from "react"

import type { GetServerSideProps, InferGetServerSidePropsType } from "next"

import { getSession } from "next-auth/react"

import { useMotionValueEvent, useScroll } from "framer-motion"

import Block from "@/components/canvas/block"
import CanvasPreviewGrid from "@/components/canvas/preview-grid"
import { Search } from "@/components/canvas/search"
import { TabsProvider } from "@/contexts/TabsProvider"
import { api, RouterOutputs } from "@/lib/api"
import { type NextPageWithLayout } from "@/lib/types"

const Page: NextPageWithLayout<
	InferGetServerSidePropsType<typeof getServerSideProps>
> = ({ search }) => {
	const { scrollYProgress } = useScroll()

	const [canvasState, setCanvasState] = useState<{
		loading?: boolean
		count?: number
		search?: string | string[]
		canvases: RouterOutputs["canvas"]["all"]
	}>({ search, canvases: [] })

	const { fetchNextPage } = api.canvas.infinite.useInfiniteQuery(
		{ search },
		{
			getNextPageParam(lastPage) {
				return lastPage.nextCursor
			},
			onSuccess(data) {
				setCanvasState(prevCanvasState => {
					const page = data.pages.length - 1

					// ? Clear the list and start over instead of appending.
					if (prevCanvasState.search !== search)
						return {
							loading: false,
							count: data.pages[page].count,
							search,
							canvases: data.pages[page].items
						}

					return {
						loading: false,
						count: data.pages[page].count,
						search,
						canvases: [
							...prevCanvasState.canvases,
							...data.pages[page].items
						]
					}
				})
			}
		}
	)

	api.canvas.onAdd.useSubscription(undefined, {
		onData(canvas) {
			setCanvasState(prevCanvasState => {
				return {
					...prevCanvasState,
					canvases: [canvas, ...prevCanvasState.canvases]
				}
			})
		}
	})

	api.canvas.onUpdate.useSubscription(undefined, {
		onData(canvas) {
			setCanvasState(prevState => {
				const index = prevState.canvases.findIndex(
					({ id }) => id === canvas.id
				)

				if (index === -1) {
					return prevState
				}

				const updatedCanvases = [...prevState.canvases]

				updatedCanvases[index] = canvas

				return { ...prevState, canvases: updatedCanvases }
			})
		}
	})

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
		<>
			<Block />

			{canvasState.canvases ? (
				<Search search={search} results={canvasState.count} />
			) : (
				<></>
			)}

			<CanvasPreviewGrid canvases={canvasState.canvases} />
		</>
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
			search: context.query.search || ""
		}
	}
}) satisfies GetServerSideProps<{
	search: string | string[] | undefined
}>

Page.getLayout = page => <TabsProvider>{page}</TabsProvider>

export default Page
