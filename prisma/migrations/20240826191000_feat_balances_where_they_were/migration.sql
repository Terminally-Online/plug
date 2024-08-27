/*
  Warnings:

  - The primary key for the `ImplementationBalance` table will be changed. If it partially fails, the table could be left without primary key constraint.
  - You are about to drop the column `cacheId` on the `ImplementationBalance` table. All the data in the column will be lost.
  - Added the required column `socketId` to the `ImplementationBalance` table without a default value. This is not possible if the table is not empty.

*/
-- DropForeignKey
ALTER TABLE "ImplementationBalance" DROP CONSTRAINT "ImplementationBalance_cacheId_fkey";

-- AlterTable
ALTER TABLE "ImplementationBalance" DROP CONSTRAINT "ImplementationBalance_pkey",
DROP COLUMN "cacheId",
ADD COLUMN     "socketId" TEXT NOT NULL,
ADD CONSTRAINT "ImplementationBalance_pkey" PRIMARY KEY ("socketId", "implementationChain", "implementationContract");

-- AddForeignKey
ALTER TABLE "ImplementationBalance" ADD CONSTRAINT "ImplementationBalance_socketId_fkey" FOREIGN KEY ("socketId") REFERENCES "UserSocket"("id") ON DELETE RESTRICT ON UPDATE CASCADE;
