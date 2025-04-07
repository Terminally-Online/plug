/*
  Warnings:

  - You are about to drop the `Companion` table. If the table is not empty, all the data it contains will be lost.

*/
-- DropForeignKey
ALTER TABLE "Companion" DROP CONSTRAINT "Companion_socketId_fkey";

-- DropTable
DROP TABLE "Companion";
