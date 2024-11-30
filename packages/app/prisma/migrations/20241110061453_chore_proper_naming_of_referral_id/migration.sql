/*
  Warnings:

  - You are about to drop the column `referralId` on the `SocketIdentity` table. All the data in the column will be lost.

*/
-- DropForeignKey
ALTER TABLE "SocketIdentity" DROP CONSTRAINT "SocketIdentity_referralId_fkey";

-- AlterTable
ALTER TABLE "SocketIdentity" DROP COLUMN "referralId",
ADD COLUMN     "referrerId" TEXT;

-- AddForeignKey
ALTER TABLE "SocketIdentity" ADD CONSTRAINT "SocketIdentity_referrerId_fkey" FOREIGN KEY ("referrerId") REFERENCES "UserSocket"("id") ON DELETE SET NULL ON UPDATE CASCADE;
