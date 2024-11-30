/*
  Warnings:

  - Added the required column `chain` to the `OpenseaCollection` table without a default value. This is not possible if the table is not empty.

*/
-- AlterTable
ALTER TABLE "OpenseaCollection" ADD COLUMN     "chain" TEXT NOT NULL;
