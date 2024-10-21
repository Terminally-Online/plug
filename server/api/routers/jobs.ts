// server/api/routers/jobs.ts
import { z } from "zod";
import { apiKeyProcedure, createTRPCRouter } from "../trpc";
import { TRPCError } from "@trpc/server";
import { Prisma } from "@prisma/client";

// Set the cleanup period to 7 days
const CLEANUP_OLDER_THAN_DAYS = 7;

export const jobs = createTRPCRouter({
  testAuth: apiKeyProcedure.query(() => {
    return {
      message: "API key authentication successful!",
      timestamp: new Date().toISOString(),
    };
  }),

  cleanupAnonymousUsers: apiKeyProcedure.mutation(async ({ ctx }) => {
    const { db } = ctx;

    // Use a transaction to ensure data integrity
    return await db.$transaction(async (prisma) => {
      try {
        const cutoffDate = new Date();
        cutoffDate.setDate(cutoffDate.getDate() - CLEANUP_OLDER_THAN_DAYS);

        console.log(`Deleting anonymous users created before: ${cutoffDate.toISOString()}`);

        // Find UserSockets to delete
        const userSocketsToDelete = await prisma.userSocket.findMany({
          where: {
            id: {
              startsWith: 'anonymous-'
            },
            createdAt: {
              lt: cutoffDate
            }
          },
          select: {
            id: true
          }
        });

        const userSocketIds = userSocketsToDelete.map(u => u.id);

        console.log(`Found ${userSocketIds.length} users matching the criteria`);

        // Delete related Collectibles
        const deletedCollectibles = await prisma.collectible.deleteMany({
          where: {
            cache: {
              socketId: {
                in: userSocketIds
              }
            }
          }
        });

        console.log(`Deleted ${deletedCollectibles.count} related Collectibles`);

        // Delete CollectibleCaches
        const deletedCaches = await prisma.collectibleCache.deleteMany({
          where: {
            socketId: {
              in: userSocketIds
            }
          }
        });

        console.log(`Deleted ${deletedCaches.count} CollectibleCaches`);

        // Delete UserSockets
        const deletedUsers = await prisma.userSocket.deleteMany({
          where: {
            id: {
              in: userSocketIds
            }
          }
        });

        console.log(`Deleted ${deletedUsers.count} UserSockets`);

        return {
          success: true,
          deletedCollectibles: deletedCollectibles.count,
          deletedCaches: deletedCaches.count,
          deletedUsers: deletedUsers.count,
          message: `Successfully deleted ${deletedUsers.count} anonymous users and related data older than ${CLEANUP_OLDER_THAN_DAYS} days.`
        };
      } catch (error) {
        console.error("Error in cleanupAnonymousUsers:", error);
        if (error instanceof Prisma.PrismaClientKnownRequestError) {
          throw new TRPCError({
            code: "INTERNAL_SERVER_ERROR",
            message: `Database error: ${error.message}`,
            cause: error
          });
        } else if (error instanceof Error) {
          throw new TRPCError({
            code: "INTERNAL_SERVER_ERROR",
            message: `An error occurred: ${error.message}`,
            cause: error
          });
        } else {
          throw new TRPCError({
            code: "INTERNAL_SERVER_ERROR",
            message: "An unknown error occurred."
          });
        }
      }
    });
  }),
});