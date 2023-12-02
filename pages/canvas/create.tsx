import type { GetServerSideProps, InferGetServerSidePropsType } from 'next'

import { getSession } from 'next-auth/react'

import Block from '@/components/canvas/block'
import CanvasPreviewGrid from '@/components/canvas/preview-grid'
import { Search } from '@/components/canvas/search'
import { TabsProvider } from '@/contexts/TabsProvider'
import { api } from '@/lib/api'
import { type NextPageWithLayout } from '@/lib/types'

const Page: NextPageWithLayout<
	InferGetServerSidePropsType<typeof getServerSideProps>
> = ({ search }) => {
	const [canvases] = api.canvas.all.useSuspenseQuery(search)

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
