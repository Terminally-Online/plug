import { useState } from 'react'

import { api } from '@/lib/api'

export default function ApiExamplePage() {
	const publicQuery = api.public.useQuery()
	const protectedQuery = api.protected.useQuery()

	const [randomNumber, setRandomNumber] = useState(0)
	api.randomNumber.useSubscription(undefined, {
		onData: data => {
			setRandomNumber(data)
		}
	})

	const [counter, setCounter] = useState(0)
	const increment = api.increment.useMutation()
	api.onIncremenet.useSubscription(undefined, {
		onData: data => {
			console.log(data)
			setCounter(data)
		}
	})

	return (
		<>
			<h2>Public</h2>
			<p>{publicQuery.data ? publicQuery.data : 'Loading...'}</p>

			<h2>Protected</h2>
			<p>{protectedQuery.data ? protectedQuery.data : 'Loading...'}</p>

			<h2>Random Number Subscription</h2>
			<p>{randomNumber}</p>

			<h2>Increment Subscription</h2>
			<p>{counter}</p>
			<button onClick={() => increment.mutate()}>Increment</button>
		</>
	)
}
