/*
  Warnings:

  - You are about to drop the column `teamId` on the `Canvas` table. All the data in the column will be lost.
  - You are about to drop the column `position` on the `Component` table. All the data in the column will be lost.
  - You are about to drop the column `size` on the `Component` table. All the data in the column will be lost.
  - You are about to drop the column `type` on the `Component` table. All the data in the column will be lost.
  - You are about to drop the `Team` table. If the table is not empty, all the data it contains will be lost.
  - You are about to drop the `UserOnTeam` table. If the table is not empty, all the data it contains will be lost.
  - Added the required column `userId` to the `Canvas` table without a default value. This is not possible if the table is not empty.

*/
-- DropForeignKey
ALTER TABLE "Canvas" DROP CONSTRAINT "Canvas_teamId_fkey";

-- DropForeignKey
ALTER TABLE "UserOnTeam" DROP CONSTRAINT "UserOnTeam_teamId_fkey";

-- DropForeignKey
ALTER TABLE "UserOnTeam" DROP CONSTRAINT "UserOnTeam_userId_fkey";

-- AlterTable
ALTER TABLE "Canvas" DROP COLUMN "teamId",
ADD COLUMN     "userId" TEXT NOT NULL;

-- AlterTable
ALTER TABLE "Component" DROP COLUMN "position",
DROP COLUMN "size",
DROP COLUMN "type";

-- DropTable
DROP TABLE "Team";

-- DropTable
DROP TABLE "UserOnTeam";

-- DropEnum
DROP TYPE "ComponentType";

-- AddForeignKey
ALTER TABLE "Canvas" ADD CONSTRAINT "Canvas_userId_fkey" FOREIGN KEY ("userId") REFERENCES "User"("id") ON DELETE RESTRICT ON UPDATE CASCADE;
