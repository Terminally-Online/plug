import { getSession } from "next-auth/react"

import { db } from "@/server/db"
import { CreateNextContextOptions } from "@trpc/server/adapters/next"
import { NodeHTTPCreateContextFnOptions } from "@trpc/server/adapters/node-http"

import { IncomingMessage } from "http"
import ws from "ws"

export const createContext = async (
	opts: NodeHTTPCreateContextFnOptions<IncomingMessage, ws> | CreateNextContextOptions
) => ({
	session: await getSession(opts),
	db
})

export type Context = Awaited<ReturnType<typeof createContext>>
