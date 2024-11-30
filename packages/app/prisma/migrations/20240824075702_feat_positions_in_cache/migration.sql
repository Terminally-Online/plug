/*
  Warnings:

  - Added the required column `cacheId` to the `Position` table without a default value. This is not possible if the table is not empty.

*/
-- AlterTable
ALTER TABLE "Position" ADD COLUMN     "cacheId" TEXT NOT NULL,
ALTER COLUMN "balance" DROP NOT NULL;

-- AddForeignKey
ALTER TABLE "Position" ADD CONSTRAINT "Position_cacheId_fkey" FOREIGN KEY ("cacheId") REFERENCES "PositionCache"("socketId") ON DELETE CASCADE ON UPDATE CASCADE;
