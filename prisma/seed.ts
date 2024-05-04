import { PrismaClient } from "@prisma/client"

const prisma = new PrismaClient()

const main = async () => {}

main()
	.catch(e => {
		throw e
	})
	.finally(async () => {
		await prisma.$disconnect()
	})
