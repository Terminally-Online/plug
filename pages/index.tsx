import Link from 'next/link'

import { useSession } from 'next-auth/react'

import Siwe from '@/components/auth/siwe'

export default function IndexPage() {
	const { data: session } = useSession()

	return (
		<>
			<h1>Hello again</h1>

			<pre>{JSON.stringify(session, null, 2)}</pre>

			<Siwe />

			<Link href="/canvas">Enter App</Link>
		</>
	)
}
