/*
  Warnings:

  - You are about to drop the column `socketId` on the `Execution` table. All the data in the column will be lost.

*/
-- DropForeignKey
ALTER TABLE "Execution" DROP CONSTRAINT "Execution_socketId_fkey";

-- AlterTable
ALTER TABLE "Execution" DROP COLUMN "socketId";
