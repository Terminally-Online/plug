import { NextApiRequest, NextApiResponse } from "next"

import { renderTrpcPanel } from "trpc-panel"

import { appRouter } from "@/server/api/root"

export default async function handler(_: NextApiRequest, res: NextApiResponse) {
	res.setHeader("Content-Type", "text/html")
	res.status(200).send(
		renderTrpcPanel(appRouter, {
			url: `${process.env.NEXT_PUBLIC_APP_URL || "http://localhost:3000"}/api/trpc`,
			transformer: "superjson"
		})
	)
}
