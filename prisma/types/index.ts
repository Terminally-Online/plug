import { Prisma } from "@prisma/client"

const consoleColumnModel = Prisma.validator<Prisma.ConsoleColumnDefaultArgs>()(
	{}
)
export type ConsoleColumnModel = Prisma.ConsoleColumnGetPayload<
	typeof consoleColumnModel
>

const userSocketModel = Prisma.validator<Prisma.UserSocketDefaultArgs>()({
	include: { columns: true }
})
export type UserSocketModel = Prisma.UserSocketGetPayload<
	typeof userSocketModel
>
