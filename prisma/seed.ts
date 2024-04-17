import { PrismaClient } from '@prisma/client'
import { chains } from "../lib/blockchain"

const prisma = new PrismaClient()

// NOTE: Create all of the chains that can be utilized in the database.
const seedChains = async () => { 
    const upsertPromises = chains.map(async chain => {
        await prisma.chain.upsert({
            where: {
                id: chain.id
            },
            update: {
                name: chain.name
            },
            create: {
                id: chain.id,
                name: chain.name
            }
        })  
    })

    await Promise.all(upsertPromises)
}

const main = async () => {
    await seedChains()
}

main()
    .catch(e => {
        throw e
    })
    .finally(async () => {
        await prisma.$disconnect()
    })