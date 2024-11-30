/*
  Warnings:

  - You are about to drop the column `userAddress` on the `Workflow` table. All the data in the column will be lost.

*/
-- AlterTable
ALTER TABLE "Workflow" DROP COLUMN "userAddress",
ADD COLUMN     "socketId" TEXT;

-- AddForeignKey
ALTER TABLE "Workflow" ADD CONSTRAINT "Workflow_socketId_fkey" FOREIGN KEY ("socketId") REFERENCES "UserSocket"("id") ON DELETE SET NULL ON UPDATE CASCADE;
