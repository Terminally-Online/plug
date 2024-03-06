import { useEffect, useState } from "react"

import type { GetServerSideProps } from "next"

import { useRouter } from "next/router"

import { getSession } from "next-auth/react"

import Fireworks from "@/components/canvas/fireworks"
import { TabsProvider } from "@/contexts/TabsProvider"
import { api } from "@/lib/api"
import { NextPageWithLayout } from "@/lib/types"

const loadingTips = [
	"Spinning up a new Plug Canvas.",
	"Reserving your indexers.",
	"Reserving your indexers..",
	"Reserving your indexers.",
	"Authenticating your peer.",
	"Handshake confirmed.",
	"Waiting for indexer pulse.",
	"Waiting for indexer pulse..",
	"Have fun!"
]

const Page: NextPageWithLayout = () => {
	const router = useRouter()

	const [tip, setTip] = useState(loadingTips[0])
	const [redirecting, setRedirecting] = useState(false)

	const createCanvas = api.canvas.add.useMutation({
		onSuccess: data => {
			router.push(`/canvas/${data.id}`)
		}
	})

	useEffect(() => {
		let index = loadingTips.indexOf(tip)
		const interval = setInterval(() => {
			setTip(loadingTips[index])

			if (index === loadingTips.length - 1) {
				setRedirecting(true)
				clearInterval(interval)
			}

			index++
		}, 1000)

		return () => clearInterval(interval)
	}, [tip])

	useEffect(() => {
		if (!redirecting) return

		const timeout = setTimeout(() => {
			createCanvas.mutate({
				name: "Untitled Canvas",
				public: false,
				color: `#${Math.floor(Math.random() * 16777215).toString(16)}`
			})
		}, 2500)

		return () => clearTimeout(timeout)
	}, [createCanvas, redirecting])

	return (
		<>
			<Fireworks enabled={redirecting} />

			<div className="flex h-screen w-screen items-center justify-center gap-4 bg-stone-900 text-white">
				<div role="status">
					<svg
						aria-hidden="true"
						className="inline h-4 w-4 animate-spin fill-gray-600 text-gray-200 dark:fill-gray-300 dark:text-gray-600"
						viewBox="0 0 100 101"
						fill="none"
						xmlns="http://www.w3.org/2000/svg"
					>
						<path
							d="M100 50.5908C100 78.2051 77.6142 100.591 50 100.591C22.3858 100.591 0 78.2051 0 50.5908C0 22.9766 22.3858 0.59082 50 0.59082C77.6142 0.59082 100 22.9766 100 50.5908ZM9.08144 50.5908C9.08144 73.1895 27.4013 91.5094 50 91.5094C72.5987 91.5094 90.9186 73.1895 90.9186 50.5908C90.9186 27.9921 72.5987 9.67226 50 9.67226C27.4013 9.67226 9.08144 27.9921 9.08144 50.5908Z"
							fill="currentColor"
						/>
						<path
							d="M93.9676 39.0409C96.393 38.4038 97.8624 35.9116 97.0079 33.5539C95.2932 28.8227 92.871 24.3692 89.8167 20.348C85.8452 15.1192 80.8826 10.7238 75.2124 7.41289C69.5422 4.10194 63.2754 1.94025 56.7698 1.05124C51.7666 0.367541 46.6976 0.446843 41.7345 1.27873C39.2613 1.69328 37.813 4.19778 38.4501 6.62326C39.0873 9.04874 41.5694 10.4717 44.0505 10.1071C47.8511 9.54855 51.7191 9.52689 55.5402 10.0491C60.8642 10.7766 65.9928 12.5457 70.6331 15.2552C75.2735 17.9648 79.3347 21.5619 82.5849 25.841C84.9175 28.9121 86.7997 32.2913 88.1811 35.8758C89.083 38.2158 91.5421 39.6781 93.9676 39.0409Z"
							fill="currentFill"
						/>
					</svg>
					<span className="sr-only">Loading...</span>
				</div>

				<h1 className="tabular-nums">{tip}</h1>
			</div>
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
		props: {}
	}
}) satisfies GetServerSideProps

Page.getLayout = page => <TabsProvider>{page}</TabsProvider>

export default Page
