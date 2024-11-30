/*
  Warnings:

  - You are about to drop the column `referredBy` on the `SocketIdentity` table. All the data in the column will be lost.

*/
-- AlterTable
ALTER TABLE "SocketIdentity" DROP COLUMN "referredBy",
ADD COLUMN     "referralCode" TEXT,
ADD COLUMN     "referralId" TEXT;

-- AddForeignKey
ALTER TABLE "SocketIdentity" ADD CONSTRAINT "SocketIdentity_referralId_fkey" FOREIGN KEY ("referralId") REFERENCES "UserSocket"("id") ON DELETE SET NULL ON UPDATE CASCADE;
