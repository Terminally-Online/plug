import ws from 'ws'

import { appRouter } from './api/root'
import { createInnerTRPCContext } from './api/trpc'
import { applyWSSHandler } from '@trpc/server/adapters/ws'
import { getSession } from 'next-auth/react'

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

wss.on('connection', ws => {
	console.log(`WebSocket client connected. Total: ${wss.clients.size}`)

	ws.on('close', () => {
		console.log(`WebSocket client disconnected. Total: ${wss.clients.size}`)
	})
})

console.log('✔︎ WebSocket server listening on port 3001.')

process.on('SIGTERM', () => {
	console.log('SIGTERM received. Shutting down WebSocket server.')

	handler.broadcastReconnectNotification()

	// * Close the socket and terminate the process.
	wss.close(() => {
		process.exit(0)
	})
})
