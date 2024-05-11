import { getSession } from "next-auth/react"

import ws from "ws"

import { applyWSSHandler } from "@trpc/server/adapters/ws"

import { version } from "../package.json"
import { appRouter } from "./api/root"
import { createInnerTRPCContext } from "./api/trpc"

const wss = new ws.Server({ port: 3001 })
const handler = applyWSSHandler({
	wss,
	router: appRouter,
	createContext: async opts => {
		const session = await getSession(opts)

		return createInnerTRPCContext({
			session
		})
	}
})

wss.on("connection", ws => {
	console.log(`WebSocket client connected. Total: ${wss.clients.size}`)

	ws.on("close", () => {
		console.log(`WebSocket client disconnected. Total: ${wss.clients.size}`)
	})
})

process.on("SIGTERM", () => {
	console.log("SIGTERM received. Shutting down WebSocket server.")

	handler.broadcastReconnectNotification()

	// * Close the socket and terminate the process.
	wss.close(() => {
		process.exit(0)
	})
})

console.log(`   🔌 Plug Websockets ${version} 
- Local: ws://localhost:3001
- Environments: .env
`)