import { PrismaClient } from "@prisma/client"

import { actionCategories, actions, tags } from "@/lib/constants"

const prisma = new PrismaClient()

const seedTags = async () => {
	Promise.all(
		tags.map(async tag => {
			await prisma.tag.upsert({
				where: { id: tag },
				create: { id: tag },
				update: {}
			})
		})
	)
}

const seedActionCategories = async () => {
	Promise.all(
		Object.keys(actionCategories).map(async categoryName => {
			const categoryFields =
				actionCategories[categoryName as keyof typeof actionCategories]

			await prisma.actionCategory.upsert({
				where: { id: categoryName },
				create: {
					id: categoryName,
					...categoryFields
				},
				update: {
					...categoryFields
				}
			})
		})
	)
}

const seedActions = async () => {
	Promise.all(
		Object.keys(actionCategories).map(async categoryName => {
			const categoryFields = actions[categoryName as keyof typeof actions]

			await Promise.all(
				Object.keys(categoryFields).map(async actionName => {
					const actionFields =
						categoryFields[
							actionName as keyof typeof categoryFields
						]

					await prisma.action.upsert({
						where: { id: actionName },
						create: {
							id: actionName,
							actionCategoryId: categoryName,
							...actionFields
						},
						update: {
							actionCategoryId: categoryName,
							...actionFields
						}
					})
				})
			)
		})
	)
}

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
