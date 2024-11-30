/*
  Warnings:

  - You are about to drop the column `rawBalance` on the `TokenBalance` table. All the data in the column will be lost.

*/
-- AlterTable
ALTER TABLE "TokenBalance" DROP COLUMN "rawBalance",
ALTER COLUMN "balance" SET DATA TYPE DOUBLE PRECISION;
