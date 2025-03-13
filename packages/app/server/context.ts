import { getSession } from "next-auth/react"

import { CreateNextContextOptions } from "@trpc/server/adapters/next"

import { db } from "@/server/db"

export const createContext = async (opts: CreateNextContextOptions) => ({
	session: await getSession(opts),
	db
})

export type Context = Awaited<ReturnType<typeof createContext>>
