import { useEffect } from "react"

import { useRouter, useSearchParams } from "next/navigation"

import { api } from "@/server/client"

export const Page = () => {
	const searchParams = useSearchParams()
	const from = searchParams.get("from")

	const router = useRouter()

	const handleAdd = api.plug.add.useMutation({
		onSuccess: data =>
			router.push(`/app/plugs/${data.plug.id}?from=${from}`)
	})

	useEffect(() => handleAdd.mutate(), [handleAdd])
}

export default Page
