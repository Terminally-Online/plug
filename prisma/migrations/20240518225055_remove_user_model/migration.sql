/*
  Warnings:

  - You are about to drop the column `userAddress` on the `Workflow` table. All the data in the column will be lost.
  - You are about to drop the `User` table. If the table is not empty, all the data it contains will be lost.

*/
-- DropForeignKey
ALTER TABLE "UserSocket" DROP CONSTRAINT "UserSocket_userAddress_fkey";

-- DropForeignKey
ALTER TABLE "Workflow" DROP CONSTRAINT "Workflow_userAddress_fkey";

-- AlterTable
ALTER TABLE "Workflow" DROP COLUMN "userAddress";

-- DropTable
DROP TABLE "User";
