/*
  Warnings:

  - Made the column `lastFeedAt` on table `Companion` required. This step will fail if there are existing NULL values in that column.

*/
-- AlterTable
ALTER TABLE "Companion" ALTER COLUMN "lastFeedAt" SET NOT NULL,
ALTER COLUMN "lastFeedAt" SET DEFAULT CURRENT_TIMESTAMP;
