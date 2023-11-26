import Viewport from '@/components/canvas/Viewport'
import { TabsProvider } from '@/contexts/TabsProvider'

import { GetServerSideProps, InferGetServerSidePropsType } from 'next'

import { NextPageWithLayout } from '@/lib/types'

const Page: NextPageWithLayout<
	InferGetServerSidePropsType<typeof getServerSideProps>
> = ({ id }) => {
	return <Viewport id={id} />
}

export const getServerSideProps = (async context => {
	const { id } = context.query

	if (!id || Array.isArray(id)) throw new Error('Single id required.')

	return {
		props: {
			id
		}
	}
}) satisfies GetServerSideProps<{
	id: string
}>

Page.getLayout = page => <TabsProvider>{page}</TabsProvider>

export default Page
