/*
  Warnings:

  - You are about to alter the column `balance` on the `TokenBalance` table. The data in that column could be lost. The data in that column will be cast from `BigInt` to `Integer`.
  - You are about to alter the column `rawBalance` on the `TokenBalance` table. The data in that column could be lost. The data in that column will be cast from `BigInt` to `Integer`.

*/
-- AlterTable
ALTER TABLE "TokenBalance" ALTER COLUMN "balance" SET DATA TYPE INTEGER,
ALTER COLUMN "rawBalance" SET DATA TYPE INTEGER;
