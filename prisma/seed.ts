import { PrismaClient } from "@prisma/client"

const prisma = new PrismaClient()

const main = async () => {
	// await seedTags()
	// await seedActionCategories()
	// await seedActions()
}

main()
	.catch(e => {
		throw e
	})
	.finally(async () => {
		await prisma.$disconnect()
	})
