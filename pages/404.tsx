import { useEffect, useState } from "react"

import Link from "next/link"
import { useRouter } from "next/router"

import { Home } from "lucide-react"

const NotFound = () => {
	const { asPath } = useRouter()

	const [origin, setOrigin] = useState("")

	useEffect(() => {
		setOrigin(
			typeof window !== "undefined" && window.location.origin
				? window.location.origin
				: ""
		)
	}, [])

	return (
		<div className="flex h-screen flex-col items-center justify-center gap-2 bg-stone-900 text-white">
			<div className="flex w-full flex-col items-center justify-center">
				<div className="max-w-[480px]">
					<p className="border-[1px] border-t-[0px] border-stone-950 bg-red-900/20 p-4 text-center text-red-700">
						Sorry! The URL <b>{asPath}</b> does not exist on{" "}
						<b>{origin}</b> and has no cache that can be served.
						Just in case, our team has been notified of this error.
						If you believe this happened by accident, please wait a
						moment and try again.
					</p>

					<Link
						href="/"
						className="flex w-full flex-row items-center justify-center gap-4 border-[1px] border-t-[0px] border-stone-950 p-4 px-8 text-center text-white"
					>
						<Home className="opacity-40" size={18} />
						RETURN BACK HOME
					</Link>
				</div>
			</div>
		</div>
	)
}

export default NotFound
