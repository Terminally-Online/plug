import { useState } from 'react'

import type { GetServerSideProps } from 'next'

import { Session } from 'next-auth'
import { getSession } from 'next-auth/react'

import CanvasPreviewGrid from '@/components/canvas/preview-grid'
import { TabsProvider } from '@/contexts/TabsProvider'
import { api } from '@/lib/api'
import { NextPageWithLayout } from '@/lib/types'

const Page: NextPageWithLayout = () => {
	const [initialCanvases] = api.canvas.all.useSuspenseQuery()
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

	return <CanvasPreviewGrid canvases={canvases} />
}

export const getServerSideProps = (async context => {
	const session = await getSession(context)

	if (!session) {
		return {
			redirect: {
				destination: `/connect`,
				permanent: false
			}
		}
	}

	return {
		props: {
			session
		}
	}
}) satisfies GetServerSideProps<{ session: Session | null }>

Page.getLayout = page => <TabsProvider>{page}</TabsProvider>

export default Page
