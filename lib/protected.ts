import { getSession } from "next-auth/react"

export const isConnected = async (ctx: any, callback: () => any) => {
	if (!(await getSession(ctx))) {
		return {
			redirect: {
				destination: `?connect=true`,
				permanent: false
			}
		}
	}

	return callback()
}
