/*
  Warnings:

  - The `balance` column on the `TokenBalance` table would be dropped and recreated. This will lead to data loss if there is data in the column.
  - The `rawBalance` column on the `TokenBalance` table would be dropped and recreated. This will lead to data loss if there is data in the column.

*/
-- AlterTable
ALTER TABLE "TokenBalance" DROP COLUMN "balance",
ADD COLUMN     "balance" BIGINT,
DROP COLUMN "rawBalance",
ADD COLUMN     "rawBalance" BIGINT;
