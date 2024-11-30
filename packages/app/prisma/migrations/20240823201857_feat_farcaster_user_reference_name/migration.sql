/*
  Warnings:

  - You are about to drop the column `farcasterUserId` on the `SocketIdentity` table. All the data in the column will be lost.
  - A unique constraint covering the columns `[farcasterId]` on the table `SocketIdentity` will be added. If there are existing duplicate values, this will fail.

*/
-- DropForeignKey
ALTER TABLE "SocketIdentity" DROP CONSTRAINT "SocketIdentity_farcasterUserId_fkey";

-- DropIndex
DROP INDEX "SocketIdentity_farcasterUserId_key";

-- AlterTable
ALTER TABLE "SocketIdentity" DROP COLUMN "farcasterUserId",
ADD COLUMN     "farcasterId" TEXT;

-- CreateIndex
CREATE UNIQUE INDEX "SocketIdentity_farcasterId_key" ON "SocketIdentity"("farcasterId");

-- AddForeignKey
ALTER TABLE "SocketIdentity" ADD CONSTRAINT "SocketIdentity_farcasterId_fkey" FOREIGN KEY ("farcasterId") REFERENCES "FarcasterUser"("id") ON DELETE SET NULL ON UPDATE CASCADE;
