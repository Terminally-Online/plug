/*
  Warnings:

  - The primary key for the `ENS` table will be changed. If it partially fails, the table could be left without primary key constraint.
  - You are about to drop the column `ensName` on the `SocketIdentity` table. All the data in the column will be lost.
  - Added the required column `socketId` to the `ENS` table without a default value. This is not possible if the table is not empty.

*/
-- DropForeignKey
ALTER TABLE "SocketIdentity" DROP CONSTRAINT "SocketIdentity_ensName_fkey";

-- DropIndex
DROP INDEX "ENS_name_key";

-- DropIndex
DROP INDEX "SocketIdentity_ensName_key";

-- AlterTable
ALTER TABLE "ENS" DROP CONSTRAINT "ENS_pkey",
ADD COLUMN     "socketId" TEXT NOT NULL,
ADD CONSTRAINT "ENS_pkey" PRIMARY KEY ("socketId");

-- AlterTable
ALTER TABLE "SocketIdentity" DROP COLUMN "ensName";

-- CreateIndex
CREATE INDEX "ENS_name_idx" ON "ENS"("name");

-- AddForeignKey
ALTER TABLE "ENS" ADD CONSTRAINT "ENS_socketId_fkey" FOREIGN KEY ("socketId") REFERENCES "SocketIdentity"("socketId") ON DELETE CASCADE ON UPDATE CASCADE;
