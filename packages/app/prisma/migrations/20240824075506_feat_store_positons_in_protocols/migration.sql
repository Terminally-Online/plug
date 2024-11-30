/*
  Warnings:

  - You are about to drop the column `cacheId` on the `Position` table. All the data in the column will be lost.
  - Made the column `balance` on table `Position` required. This step will fail if there are existing NULL values in that column.

*/
-- DropForeignKey
ALTER TABLE "Position" DROP CONSTRAINT "Position_cacheId_fkey";

-- AlterTable
ALTER TABLE "Position" DROP COLUMN "cacheId",
ALTER COLUMN "balance" SET NOT NULL;
