/*
  Warnings:

  - Made the column `cacheSocketId` on table `TokenBalance` required. This step will fail if there are existing NULL values in that column.
  - Made the column `cacheChain` on table `TokenBalance` required. This step will fail if there are existing NULL values in that column.

*/
-- DropForeignKey
ALTER TABLE "TokenBalance" DROP CONSTRAINT "TokenBalance_cacheSocketId_cacheChain_fkey";

-- AlterTable
ALTER TABLE "TokenBalance" ALTER COLUMN "cacheSocketId" SET NOT NULL,
ALTER COLUMN "cacheChain" SET NOT NULL;

-- AddForeignKey
ALTER TABLE "TokenBalance" ADD CONSTRAINT "TokenBalance_cacheSocketId_cacheChain_fkey" FOREIGN KEY ("cacheSocketId", "cacheChain") REFERENCES "TokenBalanceCache"("socketId", "chain") ON DELETE RESTRICT ON UPDATE CASCADE;
