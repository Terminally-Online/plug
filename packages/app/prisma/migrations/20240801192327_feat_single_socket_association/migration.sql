/*
  Warnings:

  - You are about to drop the column `socketAddress` on the `ConsoleColumn` table. All the data in the column will be lost.
  - You are about to drop the column `userAddress` on the `ConsoleColumn` table. All the data in the column will be lost.
  - The primary key for the `UserSocket` table will be changed. If it partially fails, the table could be left without primary key constraint.
  - You are about to drop the column `userAddress` on the `UserSocket` table. All the data in the column will be lost.
  - Added the required column `socketId` to the `ConsoleColumn` table without a default value. This is not possible if the table is not empty.
  - The required column `id` was added to the `UserSocket` table with a prisma-level default value. This is not possible if the table is not empty. Please add this column as optional, then populate it before making it required.

*/
-- DropForeignKey
ALTER TABLE "ConsoleColumn" DROP CONSTRAINT "ConsoleColumn_userAddress_socketAddress_fkey";

-- AlterTable
ALTER TABLE "ConsoleColumn" DROP COLUMN "socketAddress",
DROP COLUMN "userAddress",
ADD COLUMN     "socketId" TEXT NOT NULL;

-- AlterTable
ALTER TABLE "UserSocket" DROP CONSTRAINT "UserSocket_pkey",
DROP COLUMN "userAddress",
ADD COLUMN     "id" TEXT NOT NULL,
ADD CONSTRAINT "UserSocket_pkey" PRIMARY KEY ("id");

-- AddForeignKey
ALTER TABLE "ConsoleColumn" ADD CONSTRAINT "ConsoleColumn_socketId_fkey" FOREIGN KEY ("socketId") REFERENCES "UserSocket"("id") ON DELETE RESTRICT ON UPDATE CASCADE;
