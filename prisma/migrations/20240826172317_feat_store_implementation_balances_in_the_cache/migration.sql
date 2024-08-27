/*
  Warnings:

  - The primary key for the `ImplementationBalance` table will be changed. If it partially fails, the table could be left without primary key constraint.
  - You are about to drop the column `socketId` on the `ImplementationBalance` table. All the data in the column will be lost.
  - Added the required column `cacheId` to the `ImplementationBalance` table without a default value. This is not possible if the table is not empty.

*/
-- DropForeignKey
ALTER TABLE "ImplementationBalance" DROP CONSTRAINT "ImplementationBalance_socketId_fkey";

-- AlterTable
ALTER TABLE "ImplementationBalance" DROP CONSTRAINT "ImplementationBalance_pkey",
DROP COLUMN "socketId",
ADD COLUMN     "cacheId" TEXT NOT NULL,
ADD CONSTRAINT "ImplementationBalance_pkey" PRIMARY KEY ("cacheId", "implementationChain", "implementationContract");

-- AddForeignKey
ALTER TABLE "ImplementationBalance" ADD CONSTRAINT "ImplementationBalance_cacheId_fkey" FOREIGN KEY ("cacheId") REFERENCES "PositionCache"("socketId") ON DELETE RESTRICT ON UPDATE CASCADE;
