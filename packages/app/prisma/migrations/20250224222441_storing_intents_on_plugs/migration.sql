/*
  Warnings:

  - You are about to drop the `Intent` table. If the table is not empty, all the data it contains will be lost.
  - You are about to drop the `Run` table. If the table is not empty, all the data it contains will be lost.

*/
-- DropForeignKey
ALTER TABLE "Intent" DROP CONSTRAINT "Intent_plugId_fkey";

-- DropForeignKey
ALTER TABLE "Run" DROP CONSTRAINT "Run_intentId_fkey";

-- DropTable
DROP TABLE "Intent";

-- DropTable
DROP TABLE "Run";
