import { createNextApiHandler } from "@trpc/server/adapters/next"

import { appRouter as router } from "@/server/api/root"
import { createTRPCContext as createContext } from "@/server/api/trpc"

export default createNextApiHandler({
	router,
	createContext,
	responseMeta(opts) {
		const { ctx, paths, errors, type } = opts

		const services = paths && paths.every(path => path.startsWith("service"))
		const allOk = errors.length === 0
		const isQuery = type === "query"

		if (ctx?.res && allOk && isQuery && services) {
			const ONE_DAY_IN_SECONDS = 60 * 60 * 24
			return {
				headers: new Headers({
					"Cache-Control": `s-maxage=${ONE_DAY_IN_SECONDS}, stale-while-revalidate=${ONE_DAY_IN_SECONDS}`
				})
			}
		}

		return {}
	},
	onError:
		process.env.NODE_ENV === "development"
			? ({ path, error }) => {
					console.error(`❌ tRPC failed on ${path ?? "<no-path>"}: ${error.message}`)
				}
			: undefined
})
