import { useSearchParams } from "next/navigation"
import { useRouter } from "next/router"

export const useNavigation = () => {
	const router = useRouter()
	const id = router.query.id as string
	const searchParams = useSearchParams()
	const from = searchParams.get("from") || undefined

	return { id, from }
}
