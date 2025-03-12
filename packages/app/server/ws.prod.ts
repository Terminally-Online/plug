import next from "next"
import { getSession } from "next-auth/react"
import { applyWSSHandler } from "@trpc/server/adapters/ws"
import { createServer } from "node:http"
import { parse } from "node:url"
import { WebSocketServer } from "ws"
import { version } from "../package.json" // Added version import for logging
import { appRouter } from "./api/root"
import { createInnerTRPCContext } from "./api/trpc"

// Get port from environment variable or use default
const PORT = parseInt(process.env.PORT || "3000", 10)
const dev = process.env.NODE_ENV !== "production"

const app = next({ dev })
const handle = app.getRequestHandler()

void app.prepare().then(() => {
  const server = createServer(async (req, res) => {
    if (!req.url) return
    const parsedUrl = parse(req.url, true)
    await handle(req, res, parsedUrl)
  })

  const wss = new WebSocketServer({ noServer: true })
  const handler = applyWSSHandler({
    wss,
    router: appRouter,
    createContext: async (opts) => {
      const session = await getSession(opts)
      return createInnerTRPCContext({
        session
      })
    }
  })

  // Track and log WebSocket connections
  wss.on("connection", (ws) => {
    console.log(`WebSocket client connected. Total: ${wss.clients.size}`)
    ws.on("close", () => {
      console.log(`WebSocket client disconnected. Total: ${wss.clients.size}`)
    })
  })

  // Handle WebSocket upgrade
  server.on("upgrade", (req, socket, head) => {
    const pathname = parse(req.url || '').pathname || '';
    
    // You might want to check the path to ensure only valid WebSocket connections
    // are upgraded - for example if your WS endpoint is at /api/trpc/ws
    if (pathname.startsWith('/api/trpc')) {
      wss.handleUpgrade(req, socket, head, (ws) => {
        wss.emit("connection", ws, req)
      })
    } else {
      // For any other WebSocket paths, close the connection
      socket.destroy()
    }
  })

  // Graceful shutdown handling
  process.on("SIGTERM", () => {
    console.log("SIGTERM received. Shutting down WebSocket server.")
    handler.broadcastReconnectNotification()
    
    // Close the server and WebSocket server properly
    server.close(() => {
      wss.close(() => {
        console.log("Server and WebSocket server closed successfully")
        process.exit(0)
      })
    })
  })

  // Handle other common signals
  process.on("SIGINT", () => {
    console.log("SIGINT received. Shutting down.")
    handler.broadcastReconnectNotification()
    server.close(() => process.exit(0))
  })

  server.listen(PORT, () => {
    console.log(`
   🚀 Server started successfully
   🔌 Plug Websockets ${version}
   🌐 Mode: ${dev ? "development" : "production"}
   🌎 Listening: http://localhost:${PORT}
    `)
  })
})
