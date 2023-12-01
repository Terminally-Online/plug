import { Session } from 'next-auth'
import { getSession } from 'next-auth/react'

import { IncomingMessage } from 'http'
import ws from 'ws'

import * as trpcNext from '@trpc/server/adapters/next'
import { NodeHTTPCreateContextFnOptions } from '@trpc/server/adapters/node-http'

import { db } from '@/server/db'

import { getServerAuthSession } from './auth'

export const createContext = async (
	opts:
		| NodeHTTPCreateContextFnOptions<IncomingMessage, ws>
		| trpcNext.CreateNextContextOptions
) => {
	let session: Session | null = null

	const { req, res } = opts
	// session = await getServerAuthSession({ req, res })
	session = await getSession(opts)

	return {
		session,
		db
	}
}

export type Context = Awaited<ReturnType<typeof createContext>>
