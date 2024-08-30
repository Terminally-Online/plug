import { PrismaClient } from "@prisma/client"

const prisma = new PrismaClient()

const seedSockets = async () => {
	const DEFAULT_SOCKETS = [
		"0x0Bb5d848487B10F8CFBa21493c8f6D47e8a8B17E",
		"0x62180042606624f02d8a130da8a3171e9b33894d",
		"0x581BEf12967f06f2eBfcabb7504fA61f0326CD9A",
		"0xda70761A63d5D0DdE3bdE3b179126127Cccb44b3"
	]

	await prisma.userSocket.createMany({
		data: DEFAULT_SOCKETS.map(id => ({
			id,
			socketAddress: id
		}))
	})
}

const main = async () => {
	await seedSockets()
}

main()
	.catch(e => {
		throw e
	})
	.finally(async () => {
		await prisma.$disconnect()
	})
