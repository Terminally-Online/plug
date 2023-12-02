import { useState } from 'react'

import type { GetServerSideProps, InferGetServerSidePropsType } from 'next'

import { getSession } from 'next-auth/react'

import { useMotionValueEvent, useScroll } from 'framer-motion'

import Block from '@/components/canvas/block'
import CanvasPreviewGrid from '@/components/canvas/preview-grid'
import { Search } from '@/components/canvas/search'
import { TabsProvider } from '@/contexts/TabsProvider'
import { api, RouterOutputs } from '@/lib/api'
import { type NextPageWithLayout } from '@/lib/types'

export const revalidate = 1

const Page: NextPageWithLayout<
	InferGetServerSidePropsType<typeof getServerSideProps>
> = ({ search }) => {
	const { scrollYProgress } = useScroll()

	const [page, setPage] = useState(0)
	const [count, setCount] = useState(0)
	const [loading, setLoading] = useState(false)

	const [canvases, setCanvases] = useState<RouterOutputs['canvas']['all']>([])

	const { fetchNextPage } = api.canvas.infinite.useInfiniteQuery(
		{ search },
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

	api.canvas.onCreate.useSubscription(undefined, {
		onData(canvas) {
			setCanvases(prevCanvases => [...prevCanvases, canvas])
		}
	})

	api.canvas.onUpdate.useSubscription(undefined, {
		onData(canvas) {
			setCanvases(prevCanvases => {
				const index = prevCanvases.findIndex(
					({ id }) => id === canvas.id
				)

				if (index === -1) {
					return prevCanvases
				}

				const updatedCanvases = [...prevCanvases]

				updatedCanvases[index] = canvas

				return updatedCanvases
			})
		}
	})

	useMotionValueEvent(scrollYProgress, 'change', latest => {
		if (loading) return
		if (latest < 0.8) return

		setLoading(true)
		fetchNextPage()
	})

	return (
		<>
			<Block />
			{canvases ? <Search results={count} /> : <></>}
			<CanvasPreviewGrid canvases={canvases} />
		</>
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
			search: context.query.search || ''
		}
	}
}) satisfies GetServerSideProps<{
	search: string | string[] | undefined
}>

Page.getLayout = page => <TabsProvider>{page}</TabsProvider>

export default Page
