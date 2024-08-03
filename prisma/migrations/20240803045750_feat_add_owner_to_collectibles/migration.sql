/*
  Warnings:

  - Added the required column `owner` to the `OpenseaCollectible` table without a default value. This is not possible if the table is not empty.

*/
-- AlterTable
ALTER TABLE "OpenseaCollectible" ADD COLUMN     "owner" TEXT NOT NULL;
