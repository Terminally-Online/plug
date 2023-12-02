import { useState } from 'react'

import type { GetServerSideProps, InferGetServerSidePropsType } from 'next'

import { getSession } from 'next-auth/react'

import Block from '@/components/canvas/block'
import CanvasPreviewGrid from '@/components/canvas/preview-grid'
import { Search } from '@/components/canvas/search'
import { TabsProvider } from '@/contexts/TabsProvider'
import { api } from '@/lib/api'
import { type NextPageWithLayout } from '@/lib/types'

export const revalidate = 0

const Page: NextPageWithLayout<
	InferGetServerSidePropsType<typeof getServerSideProps>
> = ({ search }) => {
	const [initialCanvases] = api.canvas.all.useSuspenseQuery(search)
	const [canvases, setCanvases] = useState(initialCanvases)

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

	return (
		<>
			<Block />
			{canvases ? <Search results={canvases.length} /> : <></>}
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
