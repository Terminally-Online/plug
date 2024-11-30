/*
  Warnings:

  - Added the required column `userAddress` to the `Workflow` table without a default value. This is not possible if the table is not empty.

*/
-- AlterTable
ALTER TABLE "Workflow" ADD COLUMN     "userAddress" TEXT NOT NULL;
