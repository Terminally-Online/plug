import { PrismaClient } from '@prisma/client'

const PrismaClientSingleton = () => {
	return new PrismaClient()
}

type PrismaClientSingleton = ReturnType<typeof PrismaClientSingleton>

const globalForPrisma = globalThis as unknown as {
	p: PrismaClientSingleton | undefined
}

export const p = globalForPrisma.p ?? PrismaClientSingleton()

if (process.env.VERCEL_ENV !== 'production') globalForPrisma.p = p
