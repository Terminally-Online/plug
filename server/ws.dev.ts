import { createContext } from "./context";
import { appRouter } from "./routers/app";
import { applyWSSHandler } from "@trpc/server/adapters/ws";

import ws from 'ws';

const wss = new ws.Server({ port: 3001 });
const handler = applyWSSHandler({
	wss, router: appRouter, createContext
});

wss.on('connection', (ws) => {
	console.log(`WebSocket client connected. Total: ${wss.clients.size}`);

	ws.on('close', () => {
		console.log(`WebSocket client disconnected. Total: ${wss.clients.size}`);
	})
})

console.log('✔︎ WebSocket server listening on port 3001.');

process.on('SIGTERM', () => {
	console.log('SIGTERM received. Shutting down WebSocket server.');

	handler.broadcastReconnectNotification();

	// * Close the socket and terminate the process.
	wss.close(() => {
		process.exit(0);
	});
});
