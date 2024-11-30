/*
  Warnings:

  - Added the required column `rawBalance` to the `TokenBalance` table without a default value. This is not possible if the table is not empty.

*/
-- AlterTable
ALTER TABLE "TokenBalance" ADD COLUMN     "rawBalance" TEXT NOT NULL;
