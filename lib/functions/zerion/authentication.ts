import { env } from "@/env"

export const getZerionApiKey = () => {
	return `Basic ${env.ZERION_KEY}`
}
