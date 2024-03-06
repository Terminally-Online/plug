import { GetServerSideProps, InferGetServerSidePropsType } from "next"

import { getSession } from "next-auth/react"

import { Viewport } from "@/components/viewport/viewport"
import { TabsProvider } from "@/contexts/TabsProvider"
import { NextPageWithLayout } from "@/lib/types"

type PageProps = {
	id: string
}

const Page: NextPageWithLayout<
	InferGetServerSidePropsType<typeof getServerSideProps>
> = ({ id }) => {
	return <Viewport id={id} />
}

export const getServerSideProps = (async context => {
	const session = await getSession(context)

	if (!session) {
		return {
			redirect: {
				destination: `/?connect=true`,
				permanent: false
			}
		}
	}

	const { id } = context.query

	if (!id || Array.isArray(id)) throw new Error("Single id required.")

	const props = { id }

	return {
		props
	}
}) satisfies GetServerSideProps<PageProps>

Page.getLayout = page => <TabsProvider>{page}</TabsProvider>

export default Page
