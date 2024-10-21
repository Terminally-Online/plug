// server/api/routers/jobs.ts
import { TRPCError } from "@trpc/server"
import { z } from "zod"
import { Prisma } from "@prisma/client"
import { apiKeyProcedure, createTRPCRouter } from "../trpc"

// Set the cleanup period to 7 days
const CLEANUP_OLDER_THAN_DAYS = 7

export const jobs = createTRPCRouter({
    cleanupAnonymousUsers: apiKeyProcedure.mutation(async ({ ctx }) => {
        const { db } = ctx

        try {
            const cutoffDate = new Date()
            cutoffDate.setDate(cutoffDate.getDate() - CLEANUP_OLDER_THAN_DAYS)

            const deletedUsers = await db.userSocket.deleteMany({
                where: {
                    id: {
                        startsWith: "anonymous-"
                    },
                    createdAt: {
                        lt: cutoffDate
                    }
                }
            })

            return {
                success: true,
                deletedUsers: deletedUsers.count,
                message: `Successfully deleted ${deletedUsers.count} anonymous users older than ${CLEANUP_OLDER_THAN_DAYS} days.`
            }
        } catch (error) {
            console.error("Error in cleanupAnonymousUsers:", error)
            if (error instanceof Prisma.PrismaClientKnownRequestError) {
                throw new TRPCError({
                    code: "INTERNAL_SERVER_ERROR",
                    message: `Database error: ${error.message}`,
                    cause: error
                })
            } else if (error instanceof Error) {
                throw new TRPCError({
                    code: "INTERNAL_SERVER_ERROR",
                    message: `An error occurred: ${error.message}`,
                    cause: error
                })
            } else {
                throw new TRPCError({
                    code: "INTERNAL_SERVER_ERROR",
                    message: "An unknown error occurred."
                })
            }
        }
    })
})